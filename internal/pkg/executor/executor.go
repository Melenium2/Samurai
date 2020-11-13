package executor

import (
	"Samurai/config"
	"Samurai/internal/pkg/api"
	"Samurai/internal/pkg/api/inhuman"
	"Samurai/internal/pkg/api/mobilerpc"
	"Samurai/internal/pkg/db"
	"context"
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
	ctx       context.Context
	api       api.Requester
	db        db.Tracking
}

func (w *Samurai) Work() error {
	p := w.Config.Period
	for w.isWorking && p > 0 {
		// Why? Because DB clear ctx after transaction
		ctxWithTimeout, _ := context.WithTimeout(w.ctx, time.Second*120)
		// Перед тем как добовлять новую запись, нуобходимо проверить присутствует
		// 		ли приложение с заданными параметрами. Если необходимо сделать новый слепок
		// 		то можно в начальных параметрах указывать --force флаг
		//		Для этого необходимо поменять get метод к бд

		if err := w.Tick(ctxWithTimeout); err != nil {
			return err
		}

		p--
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
		w.TaskId = id
	}

	if err := w.UpdateMeta(ctx, app); err != nil {
		return err
	}

	for _, k := range w.Config.Keywords {
		keys, err := w.api.Flow(k)
		if err != nil {
			return err
		}
		bundles := w.bundles(keys)
		pos := w.position(w.Config.Bundle, bundles)
		if err = w.UpdateTrack(ctx, pos, k); err != nil {
			return err
		}
	}

	for _, subCat := range categories {
		cat := mobilerpc.NewCategory(app.Categories, subCat)
		chart, err := w.api.Charts(ctx, cat)
		if err != nil {
			return err
		}
		pos := w.position(w.Config.Bundle, chart)
		if err = w.UpdateTrack(ctx, pos, string(cat)); err != nil {
			return err
		}
	}

	return nil
}

func (w *Samurai) NewApp(ctx context.Context, app *inhuman.App) (int, error) {
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

func (w *Samurai) UpdateMeta(ctx context.Context, app *inhuman.App) error {
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

func (w *Samurai) bundles(apps []inhuman.App) []string {
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

func New(config config.AppConfig, api api.Requester, repo db.Tracking) *Samurai {
	return &Samurai{
		ctx:       context.Background(),
		Config:    config,
		isWorking: true,
		api:       api,
		db:        repo,
	}
}

func NewDefault(config config.Config) *Samurai {
	requester := api.New(config.Api, config.App.Lang)
	repo := db.NewWithConfig(config.Database)

	return New(config.App, requester, repo)
}
