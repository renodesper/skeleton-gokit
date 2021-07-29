package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-pg/pg/v10"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	api "gitlab.com/renodesper/gokit-microservices/endpoint"

	"gitlab.com/renodesper/gokit-microservices/repository/postgre"
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
func main() {
	initConfig()

	env := viper.GetString("app.env")
	level := viper.GetString("log.level")

	log, err := initLogger(env, level)
	if err != nil {
		return
	}

	scheme := viper.GetString("app.scheme")

	log.Infof("Enviroment: %s", env)
	log.Infof("HTTP url: %s://%s:%d", scheme, *host, *port)
	log.Infof("Log level: %s", level)

	dbUsername := viper.GetString("db.username")
	dbPassword := viper.GetString("db.password")
	dbHost := viper.GetString("db.host")
	dbPort := viper.GetInt("db.port")
	dbName := viper.GetString("db.name")
	db := initDB(dbUsername, dbPassword, dbHost, dbPort, dbName)
	defer db.Close()

	healthSvc := service.NewHealthService()
	googleAuthSvc := service.NewGoogleOauthService(log, db)
	authSvc := service.NewOauthService(log, db)
	userSvc := service.NewUserService(log, db)
	verificationSvc := service.NewVerificationService(log, db)

	endpoint := api.New(env, healthSvc, googleAuthSvc, authSvc, userSvc, verificationSvc)
	handler := httptransport.NewHTTPHandler(endpoint, log)
	handler = cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowCredentials: true,
	}).Handler(handler)

	errChan := make(chan error)
	server := &http.Server{Addr: fmt.Sprintf("%s:%d", *host, *port), Handler: handler}

	go func() {
		log.Info("Service started!")
		errChan <- server.ListenAndServe()
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	log.Error(<-errChan)
}

func initConfig() {
	host = flag.String("host", "0.0.0.0", "host")
	port = flag.Int("port", 8000, "port")

	configFile := flag.String("config", "config/env/development.toml", "configuration path")
	flag.Parse()

	viper.SetConfigFile(*configFile)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if h := viper.GetString("app.host"); h != "" {
		host = &h
	}
	if p := viper.GetInt("app.port"); p != 0 {
		port = &p
	}

	if viper.GetString("app.env") != "production" {
		fmt.Println("\n> You're not in production mode")
		fmt.Println()
	}
}

func initLogger(env, level string) (logger.Logger, error) {
	z, err := zap.CreateLogger(env, level)
	if err != nil {
		return nil, err
	}

	ls := logger.New(z)
	return ls, nil
}

func initDB(username, password, host string, port int, dbName string) *pg.DB {
	db := postgre.NewPostgreClient(username, password, host, port, dbName)
	return db
}
