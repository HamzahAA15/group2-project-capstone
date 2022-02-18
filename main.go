package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sirclo/project-capstone/database"
	"sirclo/project-capstone/repository/userRepository"
	_routes "sirclo/project-capstone/router"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot Read .env file")
	}

	var router = mux.NewRouter()
	var userRepo userRepository.UserRepoInterface

	dbMysql := database.MySQLConnection(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=%s&loc=%s",
		os.Getenv("mysqlUser"),
		os.Getenv("mysqlPassword"),
		os.Getenv("mysqlHost"),
		os.Getenv("mysqlPort"),
		os.Getenv("mysqlName"),
		os.Getenv("mysqlParsetime"),
		os.Getenv("mysqlTimeLocation"),
	))
	defer dbMysql.Close()

	userRepo = userRepository.NewMySQLUserRepository(dbMysql)

	router = _routes.Routes(
		userRepo,
	)

	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})

	router.Use(corsOptions.Handler)
	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port : ", httpPort())
		errs <- http.ListenAndServe(httpPort(), nil)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("terminated %s", <-errs)
}

//PORT

func httpPort() string {
	port := "80"
	if os.Getenv("HTTP_PORT") != "" {
		port = os.Getenv("HTTP_PORT")
	}

	return fmt.Sprintf(":%s", port)
}
