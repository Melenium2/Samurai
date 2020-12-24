package executor

import (
	"Samurai/config"
	"Samurai/internal/pkg/api"
	"Samurai/internal/pkg/api/inhuman"
	"Samurai/internal/pkg/api/mobilerpc"
	"Samurai/internal/pkg/api/models"
	"Samurai/internal/pkg/db"
	"Samurai/internal/pkg/logus"
	"context"
	"fmt"
	"log"
	"strings"
	"time"
)

var categories = []string{
	"apps_topselling_free",
	"apps_topgrossing",
	"apps_movers_shakers",
	"apps_topselling_paid",
}

type Worker interface {
	Work() error
	Done()
}

type Samurai struct {
	Config config.AppConfig

	TaskId    int
	isWorking bool
	logger    logus.Logus
	ctx       context.Context
	api       api.Requester
	db        db.Tracking
}

func (w *Samurai) Work() error {
	p := w.Config.Period
	for w.isWorking && p > 0 {
		// Why? Because DB clear ctx after transaction
		ctxWithTimeout, _ := context.WithTimeout(w.ctx, time.Minute*6)

		if err := w.Tick(ctxWithTimeout); err != nil {
			return err
		}

		p--
		w.logger.LogMany(logus.NewLUnit("Work()", "process"), logus.NewLUnit(p, "times left"))
		time.Sleep(w.Config.Intensity)
	}
	w.Done()

	return nil
}

func (w *Samurai) Tick(ctx context.Context) error {
	app, err := w.api.App(w.Config.Bundle)

	if err != nil {
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

	if err := w.UpdateMeta(ctx, app); err != nil {
		return err
	}

	for _, k := range w.Config.Keywords {
		keys, err := w.api.Flow(k)
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

	for _, subCat := range categories {
		cat := models.NewCategory(app.Categories, subCat)
		chart, err := w.api.Charts(ctx, cat)
		if err != nil {
			return err
		}
		pos := w.position(w.Config.Bundle, chart)
		if err = w.UpdateTrack(ctx, pos, string(cat)); err != nil {
			return err
		}
	}

	w.logger.Log("Tick()", "Tick completed")
	return nil
}

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

func (w *Samurai) UpdateMeta(ctx context.Context, app models.App) error {
	_, err := w.db.Insert(ctx, db.Meta{
		BundleId:         w.TaskId,
		Title:            app.Title,
		Price:            app.Price,
		Picture:          app.Picture,
		Screenshots:      app.Screenshots,
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

func (w *Samurai) UpdateTrack(ctx context.Context, pos int, t string) error {
	_, err := w.db.Insert(ctx, db.Track{
		BundleId: w.TaskId,
		Type:     t,
		Date:     time.Now(),
		Place:    int32(pos) + 1,
	})
	return err
}

func (w *Samurai) Done() {
	w.isWorking = false
	log.Print("Shutdown...")
}

func (w *Samurai) bundles(apps []models.App) []string {
	r := make([]string, len(apps))
	for i := 0; i < len(apps); i++ {
		r[i] = apps[i].Bundle
	}

	return r
}

func (w *Samurai) position(find string, values []string) int {
	lfind := strings.ToLower(find)
	for i, k := range values {
		if strings.ToLower(k) == lfind {
			return i
		}
	}

	return -1
}

func New(config config.AppConfig, logger logus.Logus, api api.Requester, repo db.Tracking) *Samurai {
	return &Samurai{
		ctx:       context.Background(),
		Config:    config,
		logger:    logger,
		isWorking: true,
		api:       api,
		db:        repo,
	}
}

func NewDefault(config config.Config, logger logus.Logus) *Samurai {
	requester := api.New(
		mobilerpc.New(mobilerpc.FromConfig(config)),
		inhuman.NewApiPlay(inhuman.FromConfig(config)),
	)
	repo := db.NewWithConfig(config.Database)

	return New(config.App, logger, requester, repo)
}
