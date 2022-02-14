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

	http.Handle("/", accessControl(router))

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

//CORS
func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token"
		allowedMethods := "GET, POST, PUT, DELETE"

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		w.Header().Set("Access-Control-Allow-Methods", allowedMethods)

		h.ServeHTTP(w, r)
	})
}

//PORT

func httpPort() string {
	port := "80"
	if os.Getenv("HTTP_PORT") != "" {
		port = os.Getenv("HTTP_PORT")
	}

	return fmt.Sprintf(":%s", port)
}
