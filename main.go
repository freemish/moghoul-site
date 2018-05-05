package main

import (
	"log"
	"net/http"
	"os"
	"path"
	"portfolio/pages"

	"github.com/gorilla/mux"
)

func FileServerWithCustom404(fs http.FileSystem) http.Handler {
	fsh := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fs.Open(path.Clean(r.URL.Path))
		if os.IsNotExist(err) {
			pages.NotFoundPageHandler(w, r)
			return
		}
		if err != nil {
			log.Println(err.Error() + " (in FileServerWithCustom404)")
		}
		fsh.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// Templated pages
	r.HandleFunc("/", pages.HomePageHandler)

	// Serve static pages
	s := FileServerWithCustom404(http.Dir("assets/"))
	r.PathPrefix("/").Handler(s)
	http.Handle("/", r)

	http.ListenAndServe(":8080", r)
}
