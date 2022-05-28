package main

import (
	"fmt"
	"golangchallenge/internal/core/domain"
	"golangchallenge/internal/infrastructure/configuration"
	"golangchallenge/internal/infrastructure/configuration/authentication"
	"golangchallenge/internal/utils"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	if err := utils.Init(); err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %s", err.Error()))
	}

	utils.Logger.Debug("Logging to custom file")

	//Instantiating router
	router := chi.NewRouter()

	//Initializing Firebase auth client and setting up the middleware to authenticate requests
	firebaseAuth := authentication.FirebaseInit()

	//Instantiate middlewares
	authenticationMiddleware := authentication.NewAuthenticationMiddleware(firebaseAuth)

	db := initializeMySqlDatabase()

	//Instantiate controllers, services and repositories and set up routes
	srv := configuration.NewServer(router)
	srv.Initialize(firebaseAuth, authenticationMiddleware, db)

	utils.Logger.Info("Server is running!")
	if err := http.ListenAndServe(":4000", router); err != nil {
		utils.Logger.Panic(fmt.Sprintf("Error running application: %s", err.Error()))
		panic(err)
	}
}

func initializeMySqlDatabase() *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_SERVICE_NAME"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"))

	count := 0
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		for {
			if err == nil {
				break
			}

			time.Sleep(time.Second)
			count++
			if count > 10 {
				utils.Logger.Error(fmt.Sprintf("Failed to connect to the database; attempt number %d: %s", count, err.Error()))
				panic(err)
			}
			db, err = gorm.Open("mysql", connectionString)
		}
	}

	utils.Logger.Info("MySQL database instantiated successfully, proceeding to ping the database")

	db.AutoMigrate(&domain.Course{})
	db.AutoMigrate(&domain.AuthenticationData{})

	db.DB().SetConnMaxLifetime(time.Minute * 3)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)

	return db
}
