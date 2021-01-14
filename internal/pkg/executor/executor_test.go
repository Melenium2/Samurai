package executor_test

import (
	"Samurai/config"
	"Samurai/internal/pkg/api"
	"Samurai/internal/pkg/api/imgprocess"
	"Samurai/internal/pkg/api/inhuman"
	"Samurai/internal/pkg/api/mobilerpc"
	"Samurai/internal/pkg/api/models"
	"Samurai/internal/pkg/db"
	"Samurai/internal/pkg/executor"
	"Samurai/internal/pkg/logus"
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

func (m mock_api) Charts(ctx context.Context, chart models.Category) ([]string, error) {
	return []string{"bundle1", "com.app", "bundle2"}, nil
}

func (m mock_api) App(bundle string) (models.App, error) {
	return models.App{
		Bundle:      "com.app",
		Description: "123",
		Categories:  "GAME_RACING",
	}, nil
}

func (m mock_api) Flow(key string) ([]models.App, error) {
	return []models.App{{Bundle: "com.app"}, {Bundle: "bundle1"}, {Bundle: "bundle2"}}, nil
}

func TestSamurai_NewApp_ShouldInsertNewAppToDb_NoError(t *testing.T) {
	c := config.AppConfig{
		Period: 30,
		Lang:   "ru_RU",
	}
	ex := executor.New(c, logus.NewEmptyLogger(), mock_api{}, nil, &mock_repo{appdb: make(map[int]db.App)})

	id, err := ex.NewApp(context.Background(), models.App{Bundle: "123"})
	assert.NoError(t, err)
	assert.Equal(t, 1, id)
}

func TestSamurai_UpdateMeta_ShouldInsertNewMetaInfoToDb_NoError(t *testing.T) {
	c := config.AppConfig{
		Period: 30,
		Lang:   "ru_RU",
	}
	repo := &mock_repo{metadb: make(map[int]db.Meta)}
	ex := executor.New(c, logus.NewEmptyLogger(), mock_api{}, nil, repo)
	ex.TaskId = 1

	err := ex.UpdateMeta(context.Background(), models.App{Description: "123"})
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
	ex := executor.New(c, logus.NewEmptyLogger(), mock_api{}, nil, repo)
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
	c.Categories = models.CategoriesGoogle
	repo := &mock_repo{
		appdb:      make(map[int]db.App),
		metadb:     make(map[int]db.Meta),
		categorydb: make(map[int]db.Track),
		keydb:      make(map[int]db.Track),
	}
	ex := executor.New(c, logus.NewEmptyLogger(), mock_api{}, nil, repo)

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
		Bundle:    "com.app",
		Period:    3,
		Intensity: time.Second * 1,
		Lang:      "ru_RU",
		Keywords:  []string{"1", "2", "3"},
		Categories: models.CategoriesGoogle,
	}
	repo := &mock_repo{
		appdb:      make(map[int]db.App),
		metadb:     make(map[int]db.Meta),
		categorydb: make(map[int]db.Track),
		keydb:      make(map[int]db.Track),
	}
	ex := executor.New(c, logus.NewEmptyLogger(), mock_api{}, nil, repo)

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

func TestSamurai_Done_ShouldStopAfterFirstIteration_NoError(t *testing.T) {
	c := config.AppConfig{
		Bundle:    "com.app",
		Period:    5,
		Intensity: time.Second * 1,
		Lang:      "ru_RU",
		Keywords:  []string{"1", "2", "3"},
		Categories: models.CategoriesGoogle,
	}
	repo := &mock_repo{
		appdb:      make(map[int]db.App),
		metadb:     make(map[int]db.Meta),
		categorydb: make(map[int]db.Track),
		keydb:      make(map[int]db.Track),
	}
	ex := executor.New(c, logus.NewEmptyLogger(), mock_api{}, nil, repo)

	go func() {
		time.Sleep(time.Second * 1)
		ex.Done()
	}()

	err := ex.Work()
	assert.NoError(t, err)
}

func DatabaseConnection(config config.DBConfig) *pgx.Conn {
	url, err := db.ConnectionUrl(config)
	if err != nil {
		panic(err)
	}
	conn, err := db.Connect(url)
	if err != nil {
		panic(err)
	}

	return conn
}

func TestSamurai_NewApp_ShouldInsertNewRowToDb_NoError(t *testing.T) {
	c := config.New("../../../config/dev.yml")
	c.App.Lang = "ru_RU"
	c.App.Period = 30

	ex := executor.NewDefault(c, logus.NewEmptyLogger(), nil)

	id, err := ex.NewApp(context.Background(), models.App{
		Bundle: "com.com.com",
	})

	assert.NoError(t, err)
	assert.Greater(t, id, 0)

	conn := DatabaseConnection(c.Database)

	row := conn.QueryRow(context.Background(), "select bundle from app_tracking where id = $1", id)
	var bundle string
	assert.NoError(t, row.Scan(&bundle))

	assert.Equal(t, "com.com.com", bundle)

	_, err = conn.Exec(context.Background(), "truncate table app_tracking cascade")
	if err != nil {
		panic(err)
	}
}

func TestSamurai_NewApp_ShouldntInsertRowCozContextEmpty_NoError(t *testing.T) {
	c := config.New("../../../config/dev.yml")
	c.App.Lang = "ru_RU"
	c.App.Period = 30

	ex := executor.New(c.App, logus.NewEmptyLogger(), mock_api{}, nil, db.NewAppTracking(DatabaseConnection(c.Database)))

	assert.Panics(t, func() {
		ex.NewApp(nil, models.App{
			Bundle: "com.com.com",
		})
	})
}

func TestSamurai_UpdateMeta_ShouldInsertNewRow_NoError(t *testing.T) {
	c := config.New("../../../config/dev.yml")
	c.App.Lang = "ru_RU"
	c.App.Period = 30

	conn := DatabaseConnection(c.Database)
	tr := db.NewWithConnection(conn)
	ex := executor.New(c.App, logus.NewEmptyLogger(), mock_api{}, nil, tr)

	app := models.App{
		Bundle:      "com.com.com",
		Description: "123",
	}
	ctx := context.Background()

	id, err := ex.NewApp(ctx, app)
	assert.NoError(t, err)
	ex.TaskId = id

	assert.NoError(t, ex.UpdateMeta(ctx, app))

	row := conn.QueryRow(ctx, "select id from meta_tracking where bundleid = $1", id)
	var newid int
	assert.NoError(t, row.Scan(&newid))
	assert.Greater(t, newid, 0)

	_, err = conn.Exec(context.Background(), "truncate table app_tracking, meta_tracking cascade")
	if err != nil {
		panic(err)
	}
}

func TestSamurai_UpdateTrack_ShouldInsertNewKeywordsAndCategories(t *testing.T) {
	c := config.New("../../../config/dev.yml")
	c.App.Lang = "ru_RU"
	c.App.Period = 30

	conn := DatabaseConnection(c.Database)
	tr := db.NewWithConnection(conn)
	ex := executor.New(c.App, logus.NewEmptyLogger(), mock_api{}, nil, tr)

	app := models.App{
		Bundle:      "com.com.com",
		Description: "123",
	}
	ctx := context.Background()

	id, err := ex.NewApp(ctx, app)
	assert.NoError(t, err)
	ex.TaskId = id

	assert.NoError(t, ex.UpdateTrack(ctx, 50, "key"))
	assert.NoError(t, ex.UpdateTrack(ctx, 50, "finance|apps_top_selling_hello"))

	row := conn.QueryRow(ctx, "select type from keyword_tracking where bundleid = $1", id)
	var key string
	assert.NoError(t, row.Scan(&key))
	assert.Equal(t, "key", key)

	row = conn.QueryRow(ctx, "select type from category_tracking where bundleid = $1", id)
	var cat string
	assert.NoError(t, row.Scan(&cat))
	assert.Equal(t, "finance|apps_top_selling_hello", cat)

	_, err = conn.Exec(context.Background(), "truncate table app_tracking, keyword_tracking, category_tracking cascade")
	if err != nil {
		panic(err)
	}
}

func TestSamurai_Tick_ShouldDoneOnTickOnlyDb(t *testing.T) {
	c := config.New("../../../config/dev.yml")
	c.App.Lang = "ru_RU"
	c.App.Period = 30
	c.App.Bundle = "com.app"
	c.App.Keywords = []string{"key1", "key2", "key3"}

	conn := DatabaseConnection(c.Database)
	tr := db.NewWithConnection(conn)
	ex := executor.New(c.App, logus.NewEmptyLogger(), mock_api{}, nil, tr)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*120)

	assert.NoError(t, ex.Tick(ctx))
	assert.Greater(t, ex.TaskId, 0)

	row := conn.QueryRow(context.Background(), "select bundle from app_tracking where id = $1", ex.TaskId)
	var bundle string
	assert.NoError(t, row.Scan(&bundle))
	assert.Equal(t, "com.app", bundle)

	row = conn.QueryRow(context.Background(), "select description from meta_tracking where bundleid = $1", ex.TaskId)
	var desc string
	assert.NoError(t, row.Scan(&desc))
	assert.Equal(t, "123", desc)

	rows, err := conn.Query(context.Background(), "select type from keyword_tracking where bundleid = $1", ex.TaskId)
	assert.NoError(t, err)

	var keywords []string
	for rows.Next() {
		var k string
		assert.NoError(t, rows.Scan(&k))
		keywords = append(keywords, k)
	}
	rows.Close()
	assert.Equal(t, 3, len(keywords))

	rows, err = conn.Query(context.Background(), "select type from category_tracking where bundleid = $1", ex.TaskId)
	assert.NoError(t, err)

	var cats []string
	for rows.Next() {
		var c string
		assert.NoError(t, rows.Scan(&c))
		cats = append(cats, c)
	}
	rows.Close()
	assert.Equal(t, 4, len(cats))

	_, err = conn.Exec(context.Background(), "truncate table app_tracking, keyword_tracking, meta_tracking, category_tracking cascade")
	if err != nil {
		panic(err)
	}
}

func TestSamurai_Tick_ShouldDoneOnTick(t *testing.T) {
	c := config.New("../../../config/dev.yml")
	c.App.Lang = "ru_RU"
	c.App.Intensity = time.Hour
	c.App.Period = 1
	c.App.Bundle = "com.duolingo"
	c.App.Keywords = []string{"translate", "lingo", "english"}
	c.App.ItemsCount = 250
	c.App.OnlyMeta = true

	conn := DatabaseConnection(c.Database)
	tr := db.NewWithConnection(conn)

	c.Api.GrpcAddress = "localhost"
	c.Api.GrpcPort = "1000"
	c.Api.GrpcAccount = config.Account{
		Login:    "markovskiiikhiura@gmail.com",
		Password: "k4kdffz9m",
		Locale:   "ru_RU",
		Proxy: &config.Proxy{
			Http:  "http://5AHKey:mPNZg8@45.11.127.53:8000",
			Https: "https://5AHKey:mPNZg8@45.11.127.53:8000",
		},
		Device: "whyred",
	}
	request := api.New(
		mobilerpc.New(mobilerpc.FromConfig(c)),
		inhuman.NewApiPlay(inhuman.FromConfig(c)),
	)

	ex := executor.New(c.App, logus.NewEmptyLogger(), request, imgprocess.New(c.Api.ImageProcessing), tr)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*120)

	assert.NoError(t, ex.Tick(ctx))
	assert.Greater(t, ex.TaskId, 0)

	row := conn.QueryRow(context.Background(), "select bundle from app_tracking where id = $1", ex.TaskId)
	var bundle string
	assert.NoError(t, row.Scan(&bundle))
	assert.Equal(t, c.App.Bundle, bundle)

	row = conn.QueryRow(context.Background(), "select description from meta_tracking where bundleid = $1", ex.TaskId)
	var desc string
	assert.NoError(t, row.Scan(&desc))
	assert.NotEmpty(t, desc)

	if !c.App.OnlyMeta {
		rows, err := conn.Query(context.Background(), "select type from keyword_tracking where bundleid = $1", ex.TaskId)
		assert.NoError(t, err)

		var keywords []string
		for rows.Next() {
			var k string
			assert.NoError(t, rows.Scan(&k))
			keywords = append(keywords, k)
		}
		rows.Close()
		assert.Equal(t, 3, len(keywords))

		rows, err = conn.Query(context.Background(), "select type from category_tracking where bundleid = $1", ex.TaskId)
		assert.NoError(t, err)

		var cats []string
		for rows.Next() {
			var c string
			assert.NoError(t, rows.Scan(&c))
			cats = append(cats, c)
		}
		rows.Close()
		assert.Equal(t, 4, len(cats))
	}

	_, err := conn.Exec(context.Background(), "truncate table app_tracking, keyword_tracking, meta_tracking, category_tracking cascade")
	if err != nil {
		panic(err)
	}
}

func TestSamurai_NewApp_ShouldInsertNewIosAppToDb(t *testing.T) {
	c := config.New("../../../config/dev.yml")
	k := []string{"game", "sub", "way", "сабвей", "метро"}
	c.App.Lang = "ru_RU"
	c.App.Intensity = time.Hour
	c.App.Period = 1
	c.App.Bundle = "512939461"
	c.App.Keywords = k
	c.App.ItemsCount = 200
	c.App.Categories = models.CategoriesIos

	conn := DatabaseConnection(c.Database)
	tr := db.NewWithConnection(conn)

	req := api.NewRequester(
		inhuman.NewApiStore(inhuman.FromConfig(c)),
	)

	ex := executor.New(c.App, logus.NewEmptyLogger(), req, imgprocess.New(c.Api.ImageProcessing), tr)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*120)

	assert.NoError(t, ex.Tick(ctx))
	assert.Greater(t, ex.TaskId, 0)

	row := conn.QueryRow(context.Background(), "select bundle from app_tracking where id = $1", ex.TaskId)
	var bundle string
	assert.NoError(t, row.Scan(&bundle))
	assert.Equal(t, c.App.Bundle, bundle)

	row = conn.QueryRow(context.Background(), "select description from meta_tracking where bundleid = $1", ex.TaskId)
	var desc string
	assert.NoError(t, row.Scan(&desc))
	assert.NotEmpty(t, desc)

	rows, err := conn.Query(context.Background(), "select type from keyword_tracking where bundleid = $1", ex.TaskId)
	assert.NoError(t, err)

	var keywords []string
	for rows.Next() {
		var k string
		assert.NoError(t, rows.Scan(&k))
		keywords = append(keywords, k)
	}
	rows.Close()
	assert.Equal(t, len(k), len(keywords))

	rows, err = conn.Query(context.Background(), "select type from category_tracking where bundleid = $1", ex.TaskId)
	assert.NoError(t, err)

	var cats []string
	for rows.Next() {
		var c string
		assert.NoError(t, rows.Scan(&c))
		cats = append(cats, c)
	}
	rows.Close()
	assert.Equal(t, len(c.App.Categories.Get())*4, len(cats))

	_, err = conn.Exec(context.Background(), "truncate table app_tracking, keyword_tracking, meta_tracking, category_tracking cascade")
	if err != nil {
		panic(err)
	}
}
