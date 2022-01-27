package main

/*
	1- the routes required controllers
	2- the app.go required gorm
	3- the models required config
*/

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pluralsight/webservice/pkg/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
	fmt.Println("Running")
}
