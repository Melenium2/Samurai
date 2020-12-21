package main

import (
	"Samurai/config"
	"Samurai/internal/pkg/api"
	"Samurai/internal/pkg/api/mobilerpc"
	"Samurai/internal/pkg/db"
	"Samurai/internal/pkg/executor"
	"Samurai/internal/pkg/logus"
	"context"
	"flag"
	"fmt"
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

// 1) Полностью (почти) придется переписать/дописать конфигурацию для appstore
// 2) Манипуляции с бд
//		* Будем хранить весь список прил в одной табличке. Добавим поле чтоб
//		понимать google или appstore
//		* Все ключи и категории будем складировать в те же таблицы что и сейчас
//		* Добавить новую таблицу для хранения мета ифны для appstore (поля уточнить)
// 3) Api
//		* Написать api с 0 для получения данных с апи для appstore
// 			- Дописать методы для получения developerId из appstore
// 4) Executor
//		* Продумать как заменить некоторые модели данных на универсальные
// 5) Kafka для того чтобы отправлять информацию в амплитуду


// - Нужно дописать тесты для inhuman_ios_store
// - Разобраться какие параметры нужны для api store
// - Нужно проверить как преобразовываются структуры в тип APP
// - Проверить екзекутор на моменты для изменений
// - Переписать конфиг в целом


var ErrNotConfigured = func(value interface{}) error {
	return fmt.Errorf("error app not configured %v", value)
}

var PRODUCTION bool
var bundle string
var locale string
var period int
var intensity time.Duration
var email string
var password string
var proxy string
var token string
var gsfid int
var device string
var keywords string
var keyFile string
var force bool

func main() {

	p := os.Getenv("production")
	if p != "" {
		PRODUCTION = true
	}

	{
		flag.StringVar(&bundle, "bundle", "", "bundle of tracking application")
		flag.StringVar(&locale, "locale", "ru_RU", "lang of tracking application")
		flag.IntVar(&period, "period", 30, "period of tracking")
		flag.DurationVar(&intensity, "intensity", time.Hour*24, "time period for a new snapshot of information")
		flag.StringVar(&email, "email", "", "email of account")
		flag.StringVar(&password, "pass", "", "password for account")
		flag.StringVar(&proxy, "proxy", "", "proxy for use")
		flag.StringVar(&token, "token", "", "token instead email (must be with gsfid)")
		flag.IntVar(&gsfid, "gsfid", 0, "gsfid instead password (must be with token)")
		flag.StringVar(&device, "device", "whyred", "name of device")
		flag.StringVar(&keywords, "keys", "", "keywords separated by commas")
		flag.StringVar(&keyFile, "file", "", "file with keywords separated by '\\n'")
		flag.BoolVar(&force, "force", false, "force instance create new id for tracking")

		flag.Parse()
	}

	configPath := "./config/dev.yml"
	if PRODUCTION {
		configPath = "./config/prod.yml"
	}
	c := config.New(configPath)
	c.Database.Schema = "./config/schema.sql"

	{
		var grpc_address = os.Getenv("grpc_address")
		var grpc_port = os.Getenv("grpc_port")
		if grpc_address == "" {
			grpc_address = "localhost"
			log.Print(ErrNotConfigured("grpc address is empty, default localhost"))
		}
		c.Api.GrpcAddress = grpc_address

		if grpc_port == "" {
			grpc_port = "1000"
			log.Print(ErrNotConfigured("grpc port is empty, default 1000"))
		}
		c.Api.GrpcPort = grpc_port
	}

	{
		if bundle == "" {
			log.Fatal(ErrNotConfigured("empty bundle"))
		} else {
			c.App.Bundle = bundle
		}

		if email == "" && password == "" {
			if token == "" && gsfid == 0 {
				log.Fatal(ErrNotConfigured("need to provide email/password or token/gsfid"))
			}
		} else {
			c.Api.GrpcAccount.Login = email
			c.Api.GrpcAccount.Password = password
			c.Api.GrpcAccount.Token = token
			c.Api.GrpcAccount.GsfId = gsfid
		}

		if keyFile == "" && keywords == "" {
			log.Fatal(ErrNotConfigured("empty keywords"))
		} else {
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
			c.Api.GrpcAccount.Proxy = mobilerpc.NewProxy(proxy)
		}
	}

	{
		c.App.Lang = locale
		c.App.Period = period
		c.App.Intensity = intensity
		c.Api.GrpcAccount.Device = device
		c.Api.GrpcAccount.Locale = locale
	}

	conn, err := connection(c.Database)
	if err != nil {
		log.Fatal(err)
	}
	repository := db.NewWithConnection(conn)

	logger := logus.New(configureLogger())

	ex := executor.New(
		c.App,
		logger,
		api.New(c.Api, c.App.Lang),
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

	if err := ex.Work(); err != nil {
		log.Fatal(err)
	}

	log.Print("Off")
}

func configureLogger() murlog.Logger {
	c := murlog.NewConfig()
	c.TimePref(time.RFC822)
	c.CallerPref()

	return murlog.NewLogger(c)
}

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
