package service

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/cors"
	"github.com/spf13/viper"
	api "gitlab.com/renodesper/gokit-microservices/endpoint"
	"gitlab.com/renodesper/gokit-microservices/service"
	httptransport "gitlab.com/renodesper/gokit-microservices/transport/http"
	"gitlab.com/renodesper/gokit-microservices/util/logger"
	"gitlab.com/renodesper/gokit-microservices/util/logger/zap"
)

var (
	host *string
	port *int
)

// Run ...
func Run() {
	initConfig()

	env := viper.GetString("app.env")
	level := viper.GetString("log.level")

	logger, err := initLogger(env, level)
	if err != nil {
		return
	}

	logger.Infof("Enviroment: %s", env)
	logger.Infof("HTTP url: http://%s:%d", *host, *port)
	logger.Infof("Log level: %s", level)

	svc := service.New()
	// svc = instrumentingMiddleware{svc, requestCount, requestLatency, countResult}

	endpoint := api.New(svc, env)

	handler := httptransport.MakeHTTPHandler(endpoint)
	handler = cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowCredentials: true,
	}).Handler(handler)

	errChan := make(chan error)
	server := &http.Server{Addr: fmt.Sprintf("%s:%d", *host, *port), Handler: handler}

	go func() {
		logger.Info("Service started!")
		errChan <- server.ListenAndServe()
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	logger.Error(<-errChan)
}

func initConfig() {
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

	// if viper.GetString("app.env") != "development" {
	// 	fmt.Println("Do something when env is not development")
	// }
}

func initLogger(env, level string) (logger.Logger, error) {
	z, err := zap.CreateLogger(env, level)
	if err != nil {
		return nil, err
	}

	ls := logger.New(z)
	return ls, nil
}
