package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/neabparinya11/Golang-Project/config"
	middlewarehandler "github.com/neabparinya11/Golang-Project/modules/middleware/middlewareHandler"
	middlewarerepository "github.com/neabparinya11/Golang-Project/modules/middleware/middlewareRepository"
	middlewareusecase "github.com/neabparinya11/Golang-Project/modules/middleware/middlewareUsecase"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Server struct {
		app        *echo.Echo
		db         *mongo.Client
		cfg        *config.Config
		middleware middlewarehandler.MiddlewareHandlerService
	}
)

//Create new middleware via configuration.
func NewMiddleware(cfg *config.Config) middlewarehandler.MiddlewareHandlerService{
	repositpry := middlewarerepository.NewMiddlewareRepository()
	usecase := middlewareusecase.NewMiddlewareUsecase(repositpry)
	return middlewarehandler.NewMiddlewareHandler(cfg, usecase)
}

func (s *Server) GracefulShutdown(pctx context.Context, quit <-chan os.Signal){
	log.Printf("Start service: %s", s.cfg.App.Name)
	<-quit
	log.Printf("Shutting down server: %s ...", s.cfg.App.Name)

	ctx , cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	if err := s.app.Shutdown(ctx); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func (s *Server) HttpListening(){
	if err := s.app.Start(s.cfg.App.Url); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error: %v", err)
	}
}

//Function to start server.
//Input parameter: cfg configuration, db database mongo client.
func Start(pctx context.Context, cfg *config.Config, db *mongo.Client){
	s := &Server{
		app: echo.New(),
		db: db,
		cfg: cfg,
		middleware: NewMiddleware(cfg),
	}

	// Request Timeout
	s.app.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper: middleware.DefaultSkipper,
		ErrorMessage: "Error: Request Timeout",
		Timeout: 30*time.Second,
	}))

	//Cores
	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: middleware.DefaultSkipper,
		AllowOrigins: []string{ "*" },
		AllowMethods: []string{ echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
	}))

	//Body Limit
	s.app.Use(middleware.BodyLimit("10M"))

	//Custom middleware
	switch s.cfg.App.Name{
	case "auth":
		s.AuthService()
	case "inventory":
		s.InventoryService()
	case "item":
		s.ItemService()
	case "payment":
		s.PaymentService()
	case "player":
		s.PlayerService()
	}

	//Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	s.app.Use(middleware.Logger())
	go s.GracefulShutdown(pctx, quit)
	//Listening
	s.HttpListening()
}