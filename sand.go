package box

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

/*Box : a communicated multi window and muiti panel system.
 */
type Box struct {
	Appname string
	Root    string
	Home    string
	Version string
	server  *http.Server
}

func modesText(d map[string]string) string {
	var buffer bytes.Buffer
	for k, v := range d {
		buffer.WriteString(fmt.Sprintf("%s=%s&", k, v))
	}
	s := buffer.String()
	if len(s) > 0 {
		s = strings.TrimRight(s, "&")
	}
	return s
}

func (s *Box) InitRouter(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/version", http.StatusTemporaryRedirect)
	})
	router.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		a, err := json.Marshal(s)
		if err == nil {
			w.Write(a)
		} else {
			w.Write([]byte("{'error':'not found'}"))
		}
	})
}
