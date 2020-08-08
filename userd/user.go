package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/spf13/viper"
	api "gitlab.com/renodesper/gokit-microservices/endpoint"
	"gitlab.com/renodesper/gokit-microservices/service"
	httptransport "gitlab.com/renodesper/gokit-microservices/transport/http"
)

var (
	host *string
	port *int
)

func init() {
	port = flag.Int("port", 8000, "port")
	host = flag.String("host", "0.0.0.0", "host")
	configFile := flag.String("config", "config/env/development.toml", "configuration path")
	flag.Parse()

	viper.SetConfigFile(*configFile)
	viper.BindEnv("app.env", "ENV")
	viper.BindEnv("log.level", "LOG_LEVEL")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if p := viper.GetInt("app.port"); p != 0 {
		port = &p
	}

	if viper.GetString("app.env") != "development" {
		panic("Do something when env is not development")
	}
}

func main() {
	env := viper.GetString("app.env")
	level := viper.GetString("log.level")

	fmt.Println(fmt.Sprintf("Enviroment: %s", env))
	fmt.Println(fmt.Sprintf("HTTP url: http://%s:%d", *host, *port))
	fmt.Println(fmt.Sprintf("Log level: %s", level))

	svc := service.New()
	// logger := ...
	// svc = loggingMiddleware{logger, svc}
	// svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	endpoint := api.New(svc, env)

	handler := httptransport.MakeHTTPHandler(endpoint)
	handler = cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowCredentials: true,
	}).Handler(handler)

	server := &http.Server{Addr: fmt.Sprintf("%s:%d", *host, *port), Handler: handler}
	log.Fatal(server.ListenAndServe())
}
