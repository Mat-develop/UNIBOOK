package main

import (
	"fmt"
	"log"
	"net/http"
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
	r := mux.NewRouter()
	fmt.Println(config.SecretKey)
	r = routes.Config(r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
