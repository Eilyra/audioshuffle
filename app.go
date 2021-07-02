package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/songs.json", func(rw http.ResponseWriter, r *http.Request) {
		var songs []string
		err := filepath.Walk(filepath.Join(filepath.Dir(ex), "music"), func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("error while walking (%q): %v", path, err)
				return err
			}
			if !info.IsDir() {
				songs = append(songs, info.Name())
			}
			return nil
		})
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("Something went wrong while walking files."))
			fmt.Printf("walk error: %v", err)
			return
		}
		json, err := json.Marshal(songs)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("Something went wrong while marshaling JSON"))
			fmt.Printf("marshal error: %v", err)
			return
		}
		rw.Write(json)
	})
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, filepath.Join(filepath.Dir(ex), "index.html"))
	})
	http.ListenAndServe(":9119", nil)
}
