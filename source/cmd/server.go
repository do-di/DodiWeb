package main

import (
	"auth_server/application/usecase"
	"auth_server/infrastructure/repository"
	"auth_server/web/controller"
	"auth_server/web/routes"
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// load env
	LoadEnv()
	// create connection to db
	conn := connectDB()
	defer conn.Close()
	// create repository
	repository := repository.InitRepository(conn)
	// create usecase
	usecase := usecase.InitUseCase(*repository)

	// create controller
	controller := controller.InitController(usecase)

	// create router
	router := gin.Default()
	routes.InitRoutes(router, controller)
	router.Run(":8080")
}

// Load the env
func LoadEnv() {
	// default get value from enviroment
	if len(os.Args) < 2 {
		return
	}

	// Gets the file enviroment.
	environment := os.Args[1]
	filename := fmt.Sprintf(".env_%s", environment)
	err := godotenv.Load(filename)
	panic(err)
}

func connectDB() *pgxpool.Pool {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	// default using portgres.
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbName)
	conn, err := pgxpool.Connect(context.Background(), connString)
	if err == nil {
		panic(err)
	}
	return conn
}
