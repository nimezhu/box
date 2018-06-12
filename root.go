package box

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strconv"
	"syscall"

	"github.com/gorilla/mux"
)

func (s *Box) InitHome(root string) {
	path1 := s.Root //TODO
	if _, err := os.Stat(path1); os.IsNotExist(err) {
		os.Mkdir(path1, os.ModePerm)
	}
	path2 := path.Join(path1, "sessions")
	if _, err2 := os.Stat(path2); os.IsNotExist(err2) {
		os.Mkdir(path2, os.ModePerm)
	}

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for sig := range c {
			if sig == os.Interrupt || sig == syscall.SIGTERM {
				os.Exit(1)
			}
		}
	}()
}

/*InitIdxRoot : bigwig index local storage directory.
 *   default should be HOME/apphome/index
 */
func (s *Box) InitIdxRoot(root string) string {
	if root == "" {
		s.Root = path.Join(os.Getenv("HOME"), s.Home)
	} else {
		s.Root = path.Join(root, s.Home)
	}
	idxRoot := path.Join(s.Root, "index")
	if _, err := os.Stat(idxRoot); os.IsNotExist(err) {
		os.Mkdir(idxRoot, os.ModePerm)
	}
	return idxRoot
}

func (s *Box) Start(mode string, port int, router *mux.Router) {
	s._startApp(mode, port, router)
}
func (s *Box) _startApp(mode string, port int, router *mux.Router) {
	server := &http.Server{Addr: ":" + strconv.Itoa(port), Handler: router}
	//add router add header
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	log.Println("Please open http://127.0.0.1:" + strconv.Itoa(port))
}