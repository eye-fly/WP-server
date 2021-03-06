/*
 * Secret Server
 *
 * This is an API of a secret service. You can save your secret by using the API. You can restrict the access of a secret after the certen number of views or after a certen period of time.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"log"
	"net/http"

	"example.com/secrets"
	"example.com/swagger"
	"github.com/gorilla/mux"
)

type specificHandler struct {
	Thing string
}

func (h *specificHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.Thing))
}

// func main() {
// 	s := "Helloefdsghre world!"
// 	http.Handle("/something", &specificHandler{Thing: s})
// 	http.ListenAndServe(":8080", nil)
// }

// func main() {
// 	log.Printf("Server started")

// 	router := swagger.NewRouter()

// 	log.Fatal(http.ListenAndServe(":8080", router))
// }

func main() {
	var UploadedSecret = secrets.New()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/v1/secret/{hash}", swagger.GetSecretByHashfunc(UploadedSecret))
	router.HandleFunc("/v1/secret", swagger.AddSecret(UploadedSecret))
	log.Fatal(http.ListenAndServe(":8080", router))
}
