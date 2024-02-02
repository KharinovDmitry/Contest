package app

import (
	"contest/internal/compiler"
	"contest/internal/compiler/linux"
	"contest/internal/config"
	"contest/internal/server/handlers"
	"contest/internal/server/middleware"
	"contest/internal/services"
	"contest/internal/storage"
	"contest/internal/storage/postgres"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"os"
)

type App struct {
	port        int
	router      *mux.Router
	store       *storage.Storage
	logger      *slog.Logger
	compiler    compiler.Compiler
	testService services.ITestService
}

func New(cfg *config.Config) *App {
	logger, err := setupLogger(cfg.Env)
	if err != nil {
		panic(err)
	}
	logger.Info("Логер запущен")

	router := mux.NewRouter()
	router.Use(middleware.AuthMiddleware(cfg.ApiKey))

	store, err := postgres.NewStorage(cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
	if err != nil {
		panic(err)
	}

	linuxCompiler := linux.NewLinuxCompiler()
	testService := services.NewTestService(linuxCompiler, store.TestRepository)

	app := &App{
		port:        cfg.Port,
		router:      router,
		store:       store,
		logger:      logger,
		compiler:    linuxCompiler,
		testService: testService,
	}
	app.setupRouter()

	logger.Info("Приложение собрано")
	return app
}

func (a *App) setupRouter() {
	a.router.HandleFunc("/test", handlers.RunTest(a.testService, a.logger)).Methods("GET")

	crudSubrouter := a.router.PathPrefix("/crud").Subrouter()
	crudSubrouter.HandleFunc("/test", handlers.AddTest(a.testService, a.logger)).Methods("PUT")
	crudSubrouter.HandleFunc("/test/{id}", handlers.DeleteTest(a.testService, a.logger)).Methods("DELETE")
	crudSubrouter.HandleFunc("/test/{id}", handlers.UpdateTest(a.testService, a.logger)).Methods("PATCH")
	crudSubrouter.HandleFunc("/test/{id}", handlers.GetTest(a.testService, a.logger)).Methods("GET")
	crudSubrouter.HandleFunc("/tests", handlers.GetTests(a.testService, a.logger)).Methods("GET")
	crudSubrouter.HandleFunc("/tests/{task_id}", handlers.GetTestsByTaskID(a.testService, a.logger)).Methods("GET")
}

func setupLogger(env string) (*slog.Logger, error) {
	switch env {
	case "local":
		return slog.New(slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			})), nil
	case "dev":
		file, err := os.OpenFile("app_logs", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			return nil, err
		}
		return slog.New(slog.NewTextHandler(
			file,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			})), nil
	case "prod":
		file, err := os.OpenFile("app_logs", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			return nil, err
		}
		return slog.New(slog.NewTextHandler(
			file,
			&slog.HandlerOptions{
				Level: slog.LevelInfo,
			})), nil
	default:
		return nil, errors.New("Unknown ENV: " + env)
	}
}

func (a *App) MustRun() {
	if err := http.ListenAndServe(fmt.Sprintf(":%d", a.port), a.router); err != nil {
		panic(err.Error())
	}
}
