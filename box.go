package box

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Box struct {
	Appname     string
	Version     string
	port        int
	corsOptions *cors.Options
	server      *http.Server
	router      *mux.Router
}

func NewBox(appname string, version string) *Box {
	k := &http.Server{}
	r := mux.NewRouter()
	c := cors.Options{
		AllowedOrigins: []string{"*"},
	}
	port := 8080
	s := &Box{
		appname,
		version,
		port,
		&c,
		k,
		r,
	}
	s.init()
	return s
}
func (s *Box) init() {
	s.initRouter()
}

func (s *Box) CorsOptions(c *cors.Options) *Box {
	s.corsOptions = c
	return s
}
func (s *Box) Port(p int) *Box {
	s.port = p
	return s
}

func (s *Box) Router(r *mux.Router) *Box {
	s.router = r
	s.initRouter()
	return s
}
func (s *Box) GetRouter() *mux.Router {
	return s.router
}

func (s *Box) initRouter() {
	s.router.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		a, err := json.Marshal(s)
		if err == nil {
			w.Write(a)
		} else {
			w.Write([]byte("{'error':'not found'}"))
		}
	})
}

func (s *Box) Start(mode string) {
	c := cors.New(*s.corsOptions)
	handler := c.Handler(s.router)
	if mode == "local" {
		s.server = &http.Server{Addr: "127.0.0.1:" + strconv.Itoa(s.port), Handler: handler}
		log.Println("Data service is on 127.0.0.1:" + strconv.Itoa(s.port))
	} else {
		s.server = &http.Server{Addr: ":" + strconv.Itoa(s.port), Handler: handler}
		log.Println("Data service is on port " + strconv.Itoa(s.port))
	}
	err := s.server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
func (s *Box) Stop() error {
	log.Println("Shutdown data service...")
	err := s.server.Shutdown(context.TODO())
	return err
}
