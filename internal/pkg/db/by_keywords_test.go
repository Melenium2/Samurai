package db_test

import (
	"Samurai/internal/pkg/db"
	"time"
)

func NewTrack() db.Track {
	return db.Track{
		Bundle: "com.bundle.go",
		Type:   "game_free_how",
		Date:   time.Now(),
		Place:  43,
	}
}

