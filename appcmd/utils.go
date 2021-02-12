package appcmd

import (
	"Samurai/config"
	"Samurai/internal/pkg/api"
	"Samurai/internal/pkg/api/imgprocess"
	"Samurai/internal/pkg/api/models"
	"Samurai/internal/pkg/db"
	"Samurai/internal/pkg/executor"
	"Samurai/internal/pkg/logus"
	"context"
	murlog "github.com/Melenium2/Murlog"
	"github.com/jackc/pgx/v4"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"
)

// Create config for tracking tools
// Config same with both gp and store
// Just need to provide config.TrackingType. Because you need to set
// different categories for tracking, store and gp have different categories
func loadConfig(trackingType config.TrackingType) config.Config {
	configPath := "./config/dev.yml"
	if PRODUCTION {
		configPath = "./config/prod.yml"
	}
	c := config.New(configPath)

	switch trackingType {
	case config.AppStore:
		c.App.Categories = models.CategoriesIos
	case config.GooglePlay:
		c.App.Categories = models.CategoriesGoogle
		if (email == "" && password == "") && (token == "" && gsfid == 0) {
			log.Fatal(ErrNotConfigured("need to provide email/password or token/gsfid"))
		} else {
			c.Api.GrpcAccount.Login = email
			c.Api.GrpcAccount.Password = password
			c.Api.GrpcAccount.Token = token
			c.Api.GrpcAccount.GsfId = gsfid
		}
	}

	c.Database.Schema = "./config/schema.sql"
	c.App.ItemsCount = itemsCount
	c.App.OnlyMeta = onlyMeta

	{
		if c.Api.GrpcAddress == "" || strings.ContainsAny(c.Api.GrpcAddress, "<>") {
			c.Api.GrpcAddress = "localhost"
			log.Print(ErrNotConfigured("grpc address is empty, default localhost"))
		}

		if c.Api.GrpcPort == "" || strings.ContainsAny(c.Api.GrpcPort, "<>") {
			c.Api.GrpcPort = "1000"
			log.Print(ErrNotConfigured("grpc port is empty, default 1000"))
		}
	}

	{
		if bundle == "" {
			log.Fatal(ErrNotConfigured("empty bundle"))
		} else {
			c.App.Bundle = bundle
		}
	}

	if keyFile != "" || keywords != "" {
		var splited []string
		if keyFile != "" {
			b, err := ioutil.ReadFile(keyFile)
			if err != nil {
				log.Fatal(ErrNotConfigured(err))
			}
			sep := "\n"
			if runtime.GOOS == "windows" {
				sep = "\r\n"
			}
			keywords = string(b)
			splited = strings.Split(keywords, sep)

		} else {
			splited = strings.Split(keywords, ", ")
		}
		c.App.Keywords = splited
	}

	if proxy != "" {
		c.Api.GrpcAccount.Proxy = config.NewProxy(proxy)
	}

	{
		c.App.Lang = locale
		c.App.Period = period
		c.App.Intensity = intensity
		c.Api.GrpcAccount.Device = device
		c.Api.GrpcAccount.Locale = locale
	}

	return c
}

// Method the same for gp and store trackers.
// loadTracker configure db with config.Config, then creating logger and executor.Worker.
// method return configured instance of executor.Worker
func loadTracker(c config.Config, a api.Requester) executor.Worker {
	conn, err := connection(c.Database)
	if err != nil {
		log.Fatal(err)
	}
	repository := db.NewWithConnection(conn)

	logger := logus.New(configureLogger())

	var imgproc api.ImageProcessingApi
	if imgProcessing {
		imgproc = imgprocess.New(c.Api.ImageProcessing)
	}

	ex := executor.New(
		c.App,
		logger,
		a,
		imgproc,
		repository,
	)

	if !force {
		id, err := loadTrackId(conn, bundle, locale, period)
		if err != nil {
			log.Fatal(err)
		}
		if id != 0 {
			ex.TaskId = id
		}
	}

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
		<-sig
		ex.Done()
		os.Exit(1)
	}()

	c.View()

	return ex
}

// connection create *pgx.Conn by config.DBConfig
// Method also create schema of database which is contained in
// the config.Schema file
func connection(config config.DBConfig) (*pgx.Conn, error) {
	url, err := db.ConnectionUrl(config)
	if err != nil {
		return nil, err
	}

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}

	if err := db.InitSchema(conn, config.Schema); err != nil {
		return nil, err
	}

	return conn, nil
}

// configureLogger return murlog.Logger configured by default
func configureLogger() murlog.Logger {
	c := murlog.NewConfig()
	c.TimePref(time.RFC822)
	c.CallerPref()

	return murlog.NewLogger(c)
}

// loadTrackId search in database app with same params bundle, locale, period
// if this app exists then return id of this app, otherwise return 0
func loadTrackId(conn *pgx.Conn, bundle, locale string, period int) (int, error) {
	row := conn.QueryRow(
		context.Background(),
		"select id from app_tracking where bundle = $1 and geo = $2 and period = $3 order by id desc",
		bundle, locale, period,
	)
	var id int
	if err := row.Scan(&id); err != nil {
		if err == pgx.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	return id, nil
}
