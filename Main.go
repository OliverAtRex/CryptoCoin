package main
import (
		"net/http"
		"google.golang.org/appengine"
)
func askHandler(w http.ResponseWriter, r *http.Request) {

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}
func main() {
	http.HandleFunc("/ask", askHandler)
	http.HandleFunc("/", indexHandler)
	appengine.Main()
}