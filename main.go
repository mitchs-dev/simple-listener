package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mitchs-dev/library-go/loggingFormatter"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var configFile string
var appPort int
var endpoint string
var TZ string

type Config struct {
	AppPort  int    `yaml:"appPort"`
	EndPoint string `yaml:"endpoint"`
	TimeZone string `yaml:"timeZone"`
}

var c Config

func (configItem *Config) GetConfig(configFile string) *Config {
	configData, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal("Error in reading config file", err)
	}
	err = yaml.Unmarshal(configData, configItem)
	if err != nil {
		log.Fatal("Error in unmarshalling config file", err)
	}
	return configItem
}

func main() {
	flag.StringVar(&configFile, "config", "config.yaml", "config file")
	flag.Parse()
	c.GetConfig(configFile)
	appPort = c.AppPort
	endpoint = c.EndPoint
	TZ = c.TimeZone
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
	log.Info("Application endpoint: " + endpoint)
	log.Info("Using protocol: HTTP")
	err := http.ListenAndServe(":"+fmt.Sprint(appPort), myRouter)
	if err != nil {
		log.Error("Error in listening and serving", err)
	}

}
