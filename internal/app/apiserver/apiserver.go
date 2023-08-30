package apiserver

import (
	"avitotest/internal/app/store"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(conf *Config) *APIServer {
	return &APIServer{
		config: conf,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.ConfigureLogger(); err != nil {
		return err
	}
	s.configureRouter()
	if err := s.configureStore(); err != nil {
		return err
	}
	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) ConfigureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/users/all", s.HandleGetAllUsers()).Methods("GET")
	s.router.HandleFunc("/users/create/{name}", s.HandleUsersCreate()).Methods("GET")
	s.router.HandleFunc("/users/delete/{id:[0-9]+}", s.HandleUsersDelete()).Methods("GET")
	s.router.HandleFunc("/users/{id:[0-9]+}", s.HandleGetUserById()).Methods("GET")
	s.router.HandleFunc("/users/{id:[0-9]+}/change", s.HandleChangeUserSegments()).Methods("POST")
	s.router.HandleFunc("/segments/all", s.HandleGetAllSegments()).Methods("GET")
	s.router.HandleFunc("/segments/create/{name}", s.HandleSegmentsCreate()).Methods("GET")
	s.router.HandleFunc("/segments/delete/{name}", s.HandleSegmentsDelete()).Methods("GET")
	s.router.HandleFunc("/segments/{name}", s.HandleGetSegmentByName()).Methods("GET")
	s.router.HandleFunc("/segments/{name}/change", s.HandleChangeSegmentUsers()).Methods("POST")
	s.router.HandleFunc("/hello", s.HandleHello())
}
