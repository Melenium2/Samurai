package executor_test

import (
	"Samurai/config"
	"Samurai/internal/pkg/api/inhuman"
	"Samurai/internal/pkg/api/mobilerpc"
	"Samurai/internal/pkg/db"
	"Samurai/internal/pkg/executor"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

type mock_repo struct {
	appdb      map[int]db.App
	metadb     map[int]db.Meta
	categorydb map[int]db.Track
	keydb      map[int]db.Track
}

func (m *mock_repo) Insert(ctx context.Context, data interface{}) (int, error) {
	switch v := data.(type) {
	case db.App:
		id := len(m.appdb) + 1
		m.appdb[id] = v
		return id, nil
	case db.Meta:
		id := len(m.metadb) + 1
		m.metadb[id] = v
		return id, nil
	case db.Track:
		splited := strings.Split(v.Type, "_")
		if len(splited) >= 2 {
			id := len(m.categorydb) + 1
			m.categorydb[id] = v
			return id, nil
		}

		id := len(m.keydb) + 1
		m.keydb[id] = v
		return id, nil
	}

	return 0, nil
}

func (m mock_repo) InsertTx(tx pgx.Tx, ctx context.Context, data interface{}) (int, error) {
	return 0, nil
}

type mock_api struct {
}

func (m mock_api) Charts(ctx context.Context, chart mobilerpc.Category) ([]string, error) {
	return []string {"bundle1", "com.app", "bundle2"}, nil
}

func (m mock_api) App(bundle string) (*inhuman.App, error) {
	return &inhuman.App{
		Bundle: "com.app",
		Description: "123",
		Categories: "GAME_RACING",
	}, nil
}

func (m mock_api) Flow(key string) ([]inhuman.App, error) {
	return []inhuman.App { {Bundle: "com.app" }, {Bundle: "bundle1"}, {Bundle: "bundle2" } }, nil
}

func TestSamurai_NewApp_ShouldInsertNewAppToDb_NoError(t *testing.T) {
	c := config.AppConfig{
		Period: 30,
		Lang:   "ru_RU",
	}
	ex := executor.New(c, mock_api{}, &mock_repo{appdb: make(map[int]db.App)})

	id, err := ex.NewApp(context.Background(), &inhuman.App{Bundle: "123"})
	assert.NoError(t, err)
	assert.Equal(t, 1, id)
}

func TestSamurai_UpdateMeta_ShouldInsertNewMetaInfoToDb_NoError(t *testing.T) {
	c := config.AppConfig{
		Period: 30,
		Lang:   "ru_RU",
	}
	repo := &mock_repo{metadb: make(map[int]db.Meta)}
	ex := executor.New(c, mock_api{}, repo)
	ex.TaskId = 1

	err := ex.UpdateMeta(context.Background(), &inhuman.App{Description: "123"})
	assert.NoError(t, err)

	meta := repo.metadb[1]
	assert.Equal(t, 1, meta.BundleId)
}

func TestSamurai_UpdateTrack_ShouldInsertNewTrackInfoAsKeywordToDb_NoError(t *testing.T) {
	c := config.AppConfig{
		Period: 30,
		Lang:   "ru_RU",
	}
	repo := &mock_repo{keydb: make(map[int]db.Track)}
	ex := executor.New(c, mock_api{}, repo)
	ex.TaskId = 1

	err := ex.UpdateTrack(context.Background(), 50, "key")
	assert.NoError(t, err)

	track := repo.keydb[1]
	assert.Equal(t, 1, track.BundleId)
}

func TestSamurai_Tick_ShouldDoAllWorkOnce_NoError(t *testing.T) {
	c := config.AppConfig{
		Bundle:   "com.app",
		Period:   30,
		Lang:     "ru_RU",
		Keywords: []string{"1", "2", "3"},
	}
	repo := &mock_repo{
		appdb:      make(map[int]db.App),
		metadb:     make(map[int]db.Meta),
		categorydb: make(map[int]db.Track),
		keydb:      make(map[int]db.Track),
	}
	ex := executor.New(c, mock_api{}, repo)

	err := ex.Tick(context.Background())
	assert.NoError(t, err)

	app := repo.appdb[1]
	assert.Equal(t, "com.app", app.Bundle)
	assert.Equal(t, "ru_RU", app.Geo)

	assert.Equal(t, 1, ex.TaskId)

	meta := repo.metadb[1]
	assert.Equal(t, "123", meta.Description)

	assert.Equal(t, 3, len(repo.keydb))
	for _, v := range repo.keydb {
		assert.Equal(t, int32(1), v.Place)
	}

	assert.Equal(t, 4, len(repo.categorydb))
	for _, v := range repo.categorydb {
		assert.Equal(t, int32(2), v.Place)
	}
}

func TestSamurai_Work_ShouldDoAllWork3Times_NoError(t *testing.T) {
	c := config.AppConfig{
		Bundle:   "com.app",
		Period:   3,
		Intensity: time.Second * 1,
		Lang:     "ru_RU",
		Keywords: []string{"1", "2", "3"},
	}
	repo := &mock_repo{
		appdb:      make(map[int]db.App),
		metadb:     make(map[int]db.Meta),
		categorydb: make(map[int]db.Track),
		keydb:      make(map[int]db.Track),
	}
	ex := executor.New(c, mock_api{}, repo)

	err := ex.Work()
	assert.NoError(t, err)

	app := repo.appdb[1]
	assert.Equal(t, 1, len(repo.appdb))
	assert.Equal(t, "com.app", app.Bundle)
	assert.Equal(t, "ru_RU", app.Geo)

	assert.Equal(t, 1, ex.TaskId)

	assert.Equal(t, 3, len(repo.metadb))

	assert.Equal(t, 9, len(repo.keydb))

	for _, v := range repo.keydb {
		assert.Equal(t, int32(1), v.Place)
	}

	assert.Equal(t, 12, len(repo.categorydb))
	for _, v := range repo.categorydb {
		assert.Equal(t, int32(2), v.Place)
	}
}





