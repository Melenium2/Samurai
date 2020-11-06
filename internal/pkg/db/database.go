package db

import (
	"context"
	"fmt"
	"strings"
)

// TODO **DATABASE**
//	Entities
//		App (bundle, category, developerId, developerName, geo for tracking, start_date, how much time track app (days))
//		Meta(bundle, .....)
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
	App      Getter
	Meta     Inserter
	Keywords Inserter
	Category Inserter
}

func (t *TrackingDatabase) Insert(ctx context.Context, data interface{}) error {
	switch v := data.(type) {
	case App:
		return t.App.Insert(ctx, v)
	case Meta:
		return t.Meta.Insert(ctx, v)
	case Track:
		splited := strings.Split(v.Type, "_")
		if len(splited) >= 2 {
			return t.Category.Insert(ctx, v)
		}

		return t.Keywords.Insert(ctx, v)
	}

	return ErrWrongDataType
}

func New(app Getter, meta Inserter, keys Inserter, cats Inserter) *TrackingDatabase {
	return &TrackingDatabase{
		app, meta, keys, cats,
	}
}
