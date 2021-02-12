package executor

import (
	"Samurai/config"
	"Samurai/internal/pkg/api"
	"Samurai/internal/pkg/api/inhuman"
	"Samurai/internal/pkg/api/mobilerpc"
	"Samurai/internal/pkg/api/models"
	"Samurai/internal/pkg/db"
	"Samurai/internal/pkg/logus"
	"Samurai/internal/pkg/retry"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

// Samurai implementing Worker
type Samurai struct {
	Config config.AppConfig
	TaskId int

	isWorking     bool
	logger        logus.Logus
	ctx           context.Context
	api           api.Requester
	imgProcessing api.ImageProcessingApi
	db            db.Tracking
}

// Work start tracking application for a given period
// Every Intensity start new cycle
func (w *Samurai) Work() error {
	p := w.Config.Period

	w.externalLog(fmt.Sprintf("noerror. Starting working on bundle = %s lang = %s", w.Config.Bundle, w.Config.Lang))

	for w.isWorking && p > 0 {
		ctx := context.Background()
		if err := w.Tick(ctx); err != nil {
			w.externalLog(fmt.Sprintf("Error inside working loop %s", err.Error()))

			return err
		}

		p--
		w.logger.LogMany(logus.NewLUnit("Work()", "process"), logus.NewLUnit(p, "times left"))
		time.Sleep(w.Config.Intensity)
	}

	w.Done()

	return nil
}

// One cycle of tracking
// Creates new app for tracking if taskId not defined
// Collect metadata of application and collect information about
// positions of our app by keywords and categories
func (w *Samurai) Tick(ctx context.Context) error {
	roptions := []retry.Option{
		retry.WithContext(ctx),
		retry.WithFactor(1.3),
		retry.WithMaxAttempts(10),
		retry.WithMaxRetryTime(time.Minute * 2),
	}

	var app models.App
	var err error
	err = retry.Go(func() error {
		app, err = w.api.App(w.Config.Bundle)
		return err
	}, roptions...)

	if err != nil {
		w.logger.Log("errors in app: ", err)
		return err
	}

	if w.TaskId == 0 {
		id, err := w.NewApp(ctx, app)
		if err != nil {
			return err
		}
		w.logger.Log("Tick()", "Create new app for tracking")
		w.TaskId = id
	}

	if w.imgProcessing != nil {
		if err = w.replaceImages(ctx, &app, roptions...); err != nil {
			w.logger.Log("Tick() replace image", err)
			return err
		}
	}

	if err := w.UpdateMeta(ctx, app); err != nil {
		return err
	}

	if !w.Config.OnlyMeta {
		for _, k := range w.Config.Keywords {
			var keys []models.App
			err = retry.Go(func() error {
				keys, err = w.api.Flow(k)
				return err
			}, roptions...)

			if err != nil {
				w.logger.Log("error in flow", fmt.Sprintf("keyword '%s' response with: %s", k, err))
				continue
			}

			bundles := w.bundles(keys)
			pos := w.position(w.Config.Bundle, bundles)
			if err = w.UpdateTrack(ctx, pos, k); err != nil {
				return err
			}
		}

		appCategories := strings.Split(app.Categories, ", ")
		for _, subCat := range w.Config.Categories.Get() {
			for _, category := range appCategories {
				cat := models.NewCategory(category, subCat)
				var chart []string
				err = retry.Go(func() error {
					chart, err = w.api.Charts(ctx, cat)
					return err
				}, roptions...)

				if err != nil {
					w.logger.Log("errors in category: %s: ", fmt.Sprintf("errors in category: %s, error: %s", cat, err))
					return err
				}
				pos := w.position(w.Config.Bundle, chart)
				if err = w.UpdateTrack(ctx, pos, string(cat)); err != nil {
					return err
				}
			}
		}
	}

	w.logger.Log("Tick()", "Tick completed")
	return nil
}

// Insert NewApp to database
func (w *Samurai) NewApp(ctx context.Context, app models.App) (int, error) {
	return w.db.Insert(ctx, db.App{
		Bundle:      app.Bundle,
		Category:    app.Categories,
		DeveloperId: app.DeveloperId,
		Developer:   app.Developer,
		Geo:         w.Config.Lang,
		StartAt:     time.Now(),
		Period:      uint32(w.Config.Period),
	})
}

// Insert new metadata to dataabase
func (w *Samurai) UpdateMeta(ctx context.Context, app models.App) error {
	screenshots := make([]string, len(app.Screenshots))
	i := 0
	for _, v := range app.Screenshots {
		b, _ := json.Marshal(v)
		screenshots[i] = string(b)
		i++
	}

	_, err := w.db.Insert(ctx, db.Meta{
		BundleId:         w.TaskId,
		Title:            app.Title,
		Price:            app.Price,
		Picture:          app.Picture,
		Screenshots:      screenshots,
		Rating:           app.Rating,
		ReviewCount:      app.ReviewCount,
		RatingHistogram:  app.RatingHistogram,
		Description:      app.Description,
		ShortDescription: app.ShortDescription,
		RecentChanges:    app.RecentChanges,
		ReleaseDate:      app.ReleaseDate,
		LastUpdateDate:   app.LastUpdateDate,
		AppSize:          app.AppSize,
		Installs:         app.Installs,
		Version:          app.Version,
		AndroidVersion:   app.AndroidVersion,
		ContentRating:    app.ContentRating,
		DeveloperContacts: db.DeveloperContacts{
			Email:    app.DeveloperContacts.Email,
			Contacts: app.DeveloperContacts.Contacts,
		},
		PrivacyPolicy: app.PrivacyPolicy,
		Date:          time.Now(),
	})

	return err
}

// Insert new position of application to database
func (w *Samurai) UpdateTrack(ctx context.Context, pos int, t string) error {
	_, err := w.db.Insert(ctx, db.Track{
		BundleId: w.TaskId,
		Type:     t,
		Date:     time.Now(),
		Place:    int32(pos) + 1,
	})
	return err
}

// Close Worker
func (w *Samurai) Done() {
	w.isWorking = false

	w.externalLog(fmt.Sprintf("noerror. Ending  working on bundle = %s lang = %s", w.Config.Bundle, w.Config.Lang))

	log.Print("Shutdown...")
}

// bundles return only bundleid of given []models.App
func (w *Samurai) bundles(apps []models.App) []string {
	r := make([]string, len(apps))
	for i := 0; i < len(apps); i++ {
		r[i] = apps[i].Bundle
	}

	return r
}

// position return position of string in given slice
func (w *Samurai) position(find string, values []string) int {
	lfind := strings.ToLower(find)
	for i, k := range values {
		if strings.ToLower(k) == lfind {
			return i
		}
	}

	return -1
}

// send log to external system for tg notification
func (w *Samurai) externalLog(log string) {
	if w.Config.ExternalLog != "" {
		w.logger.LogExternal(w.Config.ExternalLog, logus.Log{
			Type:    "error",
			Module:  "samurai",
			Message: log,
		})
	}
}

// replaceImages replaces images in the app bundle with a url of the remote resource
func (w *Samurai) replaceImages(ctx context.Context, app *models.App, roptions ...retry.Option) error {
	var err error
	// Replace logo
	{
		var images []string
		err = retry.Go(func() error {
			images, err = w.imgProcessing.Process(ctx, []string{app.Picture})
			return err
		}, roptions...)
		if err != nil {
			w.logger.Log("errors in replace image - logo: ", err)
			return err
		}
		app.Picture = images[0]
	}
	// Replace screenshots
	{
		var preparedScreenshots []string
		for _, v := range app.Screenshots {
			preparedScreenshots = append(preparedScreenshots, v.Screens...)
		}
		err = retry.Go(func() error {
			preparedScreenshots, err = w.imgProcessing.Process(ctx, preparedScreenshots)
			return err
		}, roptions...)
		if err != nil {
			w.logger.Log("errors in replace image - screenshots: ", err)
			return err
		}
		index := 0
		for _, v := range app.Screenshots {
			for i := range v.Screens {
				v.Screens[i] = preparedScreenshots[index]
				index++
			}
		}
	}
	return nil
}

func New(
	config config.AppConfig,
	logger logus.Logus,
	api api.Requester,
	imgprocess api.ImageProcessingApi,
	repo db.Tracking,
) *Samurai {
	return &Samurai{
		ctx:           context.Background(),
		Config:        config,
		logger:        logger,
		isWorking:     true,
		api:           api,
		imgProcessing: imgprocess,
		db:            repo,
	}
}

func NewDefault(config config.Config, logger logus.Logus, imgprocess api.ImageProcessingApi) *Samurai {
	requester := api.New(
		mobilerpc.New(mobilerpc.FromConfig(config)),
		inhuman.NewApiPlay(inhuman.FromConfig(config)),
	)
	repo := db.NewWithConfig(config.Database)

	return New(config.App, logger, requester, imgprocess, repo)
}
