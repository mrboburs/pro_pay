package main

import (
	"pro_pay/config"
	"pro_pay/pkg/handler"
	"pro_pay/pkg/repository"
	"pro_pay/pkg/service"
	"pro_pay/pkg/store"

	"context"
	"os"
	"os/signal"
	"pro_pay/server"
	"pro_pay/tools/logger"
	"syscall"
)

// @title pro_pay
// @version 1.0
// @description API Server for pro_pay Application
// @termsOfService gitlab.com
// @host gitlab.com
// @BasePath
// @contact.name   Bakhodir Yashin Mansur
// @contact.email  phapp0224mb@gmail.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	loggers := logger.GetLogger()
	cfg := config.Config()
	// Migration Up
	err := repository.MigratePsql(cfg.Postgres, loggers, cfg.Server.MigrationUp)
	if err != nil {
		loggers.Error("error while migrate up", err)
	}
	db, err := repository.NewPostgresDB(&cfg.Postgres, loggers)
	if err != nil {
		loggers.Fatalf("failed to initialize db: %s", err.Error())
	}
	minio, err := store.MinioConnection(&cfg.Minio, loggers)
	if err != nil {
		loggers.Fatal("error while connect to minio server", err)
	}
	// rbac.NewRBAC(db, loggers)
	repos := repository.NewRepository(db, loggers)
	newStore := store.NewStore(minio, &cfg.Minio, loggers)
	newService := service.NewService(repos, newStore, cfg, loggers)
	handlers := handler.NewHandler(newService, loggers)
	// go cronjobs.NewCronJobs(db, minio, loggers)
	srv := new(server.Server)
	go func() {
		if err := srv.Run(cfg.Server.ServerPort, handlers.InitRoutes()); err != nil {
			loggers.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err := srv.Shutdown(context.Background()); err != nil {
		loggers.Errorf("error occurred on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		loggers.Errorf("error occurred on db connection close: %s", err.Error())
	}

}
