package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sirclo/project-capstone/database"
	"sirclo/project-capstone/repository/dayRepository"
	"sirclo/project-capstone/repository/officeRepository"
	"sirclo/project-capstone/repository/userRepository"
	_routes "sirclo/project-capstone/router"
	"syscall"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot Read .env file")
	}

	var router = mux.NewRouter()
	var userRepo userRepository.UserRepoInterface
	var officeRepo officeRepository.OfficeRepoInterface
	var dayRepo dayRepository.DayRepoInterface

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
	officeRepo = officeRepository.NewMySQLOfficeRepository(dbMysql)
	dayRepo = dayRepository.NewMySQLDayRepository(dbMysql)

	router = _routes.Routes(
		userRepo,
		officeRepo,
		dayRepo,
	)

	// http.Handle("/", accessControl(router))
	credentials := handlers.AllowCredentials()
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port : ", httpPort())
		errs <- http.ListenAndServe(httpPort(), handlers.CORS(credentials, methods, origins)(router))
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
