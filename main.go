package cardboardcrusaders

import (
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	})
}
