package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"embed-file-spa/web"
)

func main() {
	// http.HandleFunc("/hello.json", handleHello)
	http.HandleFunc("/", handleSPA)
	log.Println("the server is listening to port 5050")
	log.Fatal(http.ListenAndServe(":5050", nil))
}

// func handleHello(w http.ResponseWriter, r *http.Request) {
// 	res, err := json.Marshal(map[string]string{
// 		"message": "hello from the server",
// 	})
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(res)
// }

func handleSPA(w http.ResponseWriter, r *http.Request) {
	buildPath := "dist"
	f, err := web.BuildFs.Open(filepath.Join(buildPath, r.URL.Path))
	if os.IsNotExist(err) {
		index, err := web.BuildFs.ReadFile(filepath.Join(buildPath, "index.html"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write(index)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	http.FileServer(web.BuildHTTPFS()).ServeHTTP(w, r)
}
