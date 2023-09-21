package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func downloadFile(w http.ResponseWriter, r *http.Request) {
	filename := r.FormValue("filename")
	if filename == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	filepath := filepath.Join("/Users/zhaoxin/Desktop/", filename)
	file, err := os.Open(filepath)
	if err != nil {
		http.Error(w, "File Not Found", http.StatusNotFound)
		return
	}
	defer file.Close()

	fileHeader := make([]byte, 512)
	file.Read(fileHeader)
	fileStat, _ := file.Stat()
	size := strconv.FormatInt(fileStat.Size(), 10)
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", size)
	file.Seek(0, 0)
	io.Copy(w, file)
}

func main() {
	http.HandleFunc("/download", downloadFile)

	fmt.Println("Server is running at http://localhost:8080/download")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
