package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct{
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer{
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func  (a *APIServer) Start() error{
	if err := a.configureLogger(); err != nil{
		return err
	}
	a.configureRouter()

	a.logger.Info("starting api server")

	return http.ListenAndServe(a.config.BindAddr, a.router)
}

func (a *APIServer) configureLogger() error{
	level, err := logrus.ParseLevel(a.config.LogLevel)
	if err != nil{
		return err
	}

	a.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter(){
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}