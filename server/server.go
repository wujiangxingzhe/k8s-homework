package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/wujiangxingzhe/k8s-homework/middleware"
)

type Server struct {
	srv *http.Server
}

func NewServer(addr string) *Server {
	srv := &Server{
		srv: &http.Server{
			Addr: addr,
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/", srv.EchoHeader)
	router.HandleFunc("/env", srv.Env)
	router.HandleFunc("/healthz", srv.Healthz)

	srv.srv.Handler = middleware.Logging(router)

	return srv
}

func (s *Server) ListenAndServe() {
	go func() {
		s.srv.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGUSR1)
	select {
	case <-signalChan:
		fmt.Printf("gracefull shutdown, service will be closed after 5s")
		time.Sleep(time.Second * 5)
		s.Close()
	}
}

func (s *Server) Close() {
	err := s.srv.Close()
	if err != nil {
		panic(err)
	}
}

func (s *Server) EchoHeader(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("remote addr: %s, resp code:\n", strings.Split(r.RemoteAddr, ":")[0])
	for k, v := range r.Header {
		w.Header().Set(k, strings.Join(v, ";"))
		// io.WriteString(w, fmt.Sprintf("%s: %s\n", k, strings.Join(v, ";")))
		// _, err := io.WriteString(w, fmt.Sprintf("%s: %s\n", k, strings.Join(v, ";")))
		// if err != nil {
		// 	fmt.Println(k, v)
		// }
	}
	w.Write([]byte("ok"))
}

func (s *Server) Env(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("VERSION")
	if version == "" {
		version = "nil"
	}
	w.Header().Set("VERSION", version)
	_, err := w.Write([]byte(fmt.Sprintf("VERSION: %s", version)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *Server) Healthz(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "200")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
