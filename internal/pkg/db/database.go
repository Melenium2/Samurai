package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

// TODO **DATABASE**
//	Entities
//		app (bundle, category, developerId, developerName, geo for tracking, start_date, how much time track app (days))
//		meta(bundle, .....)
//		TrackByKeywords(bundle, keyword, date, place(1-250)/or none)
//		TrackByCategory(bundle, category, date, place(1-500)/or none)

var ErrWrongDataType = fmt.Errorf("wrong data type")

type Inserter interface {
	Insert(ctx context.Context, data interface{}) error
}

type Getter interface {
	Inserter
	Get(ctx context.Context, value interface{}) (interface{}, error)
}

type Tracking interface {
}

type TrackingDatabase struct {
	app      Getter
	meta     Inserter
	keywords Inserter
	category Inserter
}

func (t *TrackingDatabase) Insert(ctx context.Context, data interface{}) error {
	switch v := data.(type) {
	case App:
		return t.app.Insert(ctx, v)
	case Meta:
		return t.meta.Insert(ctx, v)
	case Track:
		splited := strings.Split(v.Type, "_")
		if len(splited) >= 2 {
			return t.category.Insert(ctx, v)
		}

		return t.keywords.Insert(ctx, v)
	}

	return ErrWrongDataType
}

func NewWithConnection(db *sql.DB) *TrackingDatabase {
	return &TrackingDatabase{
		app:      NewAppTracking(db),
		meta:     NewMetaTracking(db),
		keywords: NewKeywordsTracking(db),
		category: NewCategoryTracking(db),
	}
}

func New(app Getter, meta Inserter, keys Inserter, cats Inserter) *TrackingDatabase {
	return &TrackingDatabase{
		app, meta, keys, cats,
	}
}
