package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/t6x-io/go/mods/loggingFormatter"
)

var appPort = 8753
var endpoint = "/mail/delivery"
var TZ = "America/New_York"

func main() {
	log.SetFormatter(&loggingFormatter.JSONFormatter{
		Timezone: TZ,
		Prefix:   "smpl-l-",
	})
	log.SetOutput(os.Stdout)
	Handler()
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"status\":\"OK\",\"message\":\"OK\"}}")
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Error in reading request body", err)
	}
	log.SetFormatter(&loggingFormatter.JSONFormatter{
		Timezone:       TZ,
		Prefix:         "smpl-l-",
		MessageNoQuote: true,
	})
	log.Info(string(requestBody))
	log.SetFormatter(&loggingFormatter.JSONFormatter{
		Timezone:       TZ,
		Prefix:         "smpl-l-",
		MessageNoQuote: false,
	})
}

func Handler() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc(endpoint, HomePage)

	log.Info("Application listening on port: " + fmt.Sprint(appPort))
	log.Info("Using protocol: HTTP")
	err := http.ListenAndServe(":"+fmt.Sprint(appPort), myRouter)
	if err != nil {
		log.Error("Error in listening and serving", err)
	}

}
