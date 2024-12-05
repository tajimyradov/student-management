package main

import (
	"context"
	"flag"
	"github.com/chapsuk/grace"
	"go.elastic.co/apm/module/apmzap"
	"go.uber.org/zap"
	"log"
	"math/rand"
	"runtime"
	"student-management/internal/app"
	"student-management/internal/config"
	"student-management/internal/repositories"
	"student-management/pkg/logger"
	"student-management/pkg/postgres"
	"time"
)

const Name = "student-management"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.NewSource(time.Now().UnixNano())

	ctx := grace.ShutdownContext(context.Background())

	var version, environment, logLevel string
	flag.StringVar(&version, "v", "", "version")
	flag.StringVar(&environment, "e", "", "environment")
	flag.StringVar(&logLevel, "ll", "info", "logging level")
	//flag.StringVar(&test, "-gcflags", "-m", "")
	flag.Parse()

	appConfig, err := config.NewAppConfig("configs/" + environment + ".yml")
	if err != nil {
		log.Fatal("error while read config", err.Error())
	}

	ll, err := logger.NewWithSampler(
		Name,
		version,
		environment,
		logLevel,
		1,
		500,
		zap.WrapCore((&apmzap.Core{}).WrapCore),
	)

	if err != nil {
		log.Fatal("error while init logger", err.Error())
	}

	ll.Info(
		"flags",
		zap.String("version", version),
		zap.String("environment", environment),
		zap.String("log_level", logLevel),
	)

	studentDB, err := postgres.NewPostgresDB(appConfig.StudentDB)
	if err != nil {
		log.Fatal("error while init postgres db", err.Error())
		return
	}

	repos := repositories.NewRepository(studentDB, appConfig)
	application := app.New(appConfig, ll, repos)
	application.Start(ctx)
	application.Shutdown()
}
