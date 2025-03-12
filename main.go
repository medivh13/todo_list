package main

import (
	"context"
	"database/sql"
	"log"

	usecases "todo_list/src/app/usecases"

	"todo_list/src/infra/config"

	postgres "todo_list/src/infra/persistence/postgres"

	userRepo "todo_list/src/app/repositories/user"

	"todo_list/src/interface/rest"

	ms_log "todo_list/src/infra/log"

	userUC "todo_list/src/app/usecases/user"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"

	"github.com/sirupsen/logrus"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// init context
	ctx := context.Background()

	// read the server environment variables
	conf := config.Make()

	// check is in production mode
	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	// logger setup
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(conf.Log.Name),
		ms_log.IsProduction(isProd),
		ms_log.LogAdditionalFields(m))

	postgresdb, err := postgres.New(conf.SqlDb, logger)
	if err != nil {
		logger.Fatalf("Failed to initialize Postgres: %s", err)
	}

	defer func(l *logrus.Logger, sqlDB *sql.DB, dbName string) {
		err := sqlDB.Close()
		if err != nil {
			l.Errorf("error closing sql database %s: %s", dbName, err)
		} else {
			l.Printf("sql database %s successfuly closed.", dbName)
		}
	}(logger, postgresdb.Conn.DB, postgresdb.Conn.DriverName())

	userRepository := userRepo.NewUserRepository(postgresdb.Conn)

	httpServer, err := rest.New(
		conf.Http,
		isProd,
		logger,
		usecases.AllUseCases{

			UserUC: userUC.NewUserUseCase(userRepository),
		},
	)
	if err != nil {
		panic(err)
	}

	httpServer.Start(ctx)

}
