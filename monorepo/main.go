package main

import (
	"fmt"
	"log"
	"net/http"
	"v1/monorepo/handlers"
	"v1/monorepo/users/repository"
	"v1/monorepo/users/service"
	dbconfig "v1/monorepo/util/db_config"
	config "v1/monorepo/util/route_config"
	"v1/monorepo/wb_router/routes"

	"github.com/gorilla/mux"
)

// USED TO GENERATE THE KEY - very simple
// func init() {
// 	key := make([]byte, 64)

// 	if _, err := rand.Read(key); err != nil {
// 		log.Fatal(err)
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString(key)
// 	fmt.Println(stringBase64)
// }

func main() {
	config.Load()

	db, err := dbconfig.Connect()
	if err != nil {
		log.Fatal("DB connection error:", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r := mux.NewRouter()
	r = routes.Config(r, userHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
