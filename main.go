package main

import (
	"context"
	"database/sql"
	"log"

	usecases "todo_list/src/app/usecases"

	"todo_list/src/infra/config"

	postgres "todo_list/src/infra/persistence/postgres"

	taskRepo "todo_list/src/app/repositories/task"
	userRepo "todo_list/src/app/repositories/user"

	"todo_list/src/interface/rest"

	ms_log "todo_list/src/infra/log"

	taskUC "todo_list/src/app/usecases/task"
	userUC "todo_list/src/app/usecases/user"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"

	"github.com/sirupsen/logrus"

	"todo_list/src/infra/broker/nats"
	natsPublisher "todo_list/src/infra/broker/nats/publisher"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize a new context
	ctx := context.Background()

	// Read configuration from environment variables
	conf := config.Make()

	// Check if the application is running in production mode
	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	// Setup logger with additional fields
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(conf.Log.Name),
		ms_log.IsProduction(isProd),
		ms_log.LogAdditionalFields(m))

	// Initialize PostgreSQL database connection
	postgresdb, err := postgres.New(conf.SqlDb, logger)
	if err != nil {
		logger.Fatalf("Failed to initialize Postgres: %s", err)
	}

	// Ensure the database connection is closed when the application exits
	defer func(l *logrus.Logger, sqlDB *sql.DB, dbName string) {
		err := sqlDB.Close()
		if err != nil {
			l.Errorf("error closing sql database %s: %s", dbName, err)
		} else {
			l.Printf("sql database %s successfully closed.", dbName)
		}
	}(logger, postgresdb.Conn.DB, postgresdb.Conn.DriverName())

	// Initialize repositories for user and task management
	userRepository := userRepo.NewUserRepository(postgresdb.Conn)
	taskRepository := taskRepo.NewTaskRepository(postgresdb.Conn)

	// Initialize NATS message broker
	Nats := nats.NewNats(conf.Nats, logger)
	// Initialize NATS publisher
	publisher := natsPublisher.NewPushWorker(Nats)

	// Initialize HTTP server with use cases
	httpServer, err := rest.New(
		conf.Http,
		isProd,
		logger,
		usecases.AllUseCases{
			UserUC: userUC.NewUserUseCase(userRepository), // User use case
			TaskUC: taskUC.NewTaskUseCase(publisher, taskRepository), // Task use case
		},
	)
	if err != nil {
		panic(err)
	}

	// Start the HTTP server
	httpServer.Start(ctx)
}
