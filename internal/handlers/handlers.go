package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Pavel-Sergeev-ekb/first_http_server/internal/service"
)

func BaseHandle(h http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(h, "method not supporting", http.StatusInternalServerError)
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(h, r)
		return
	}

	file, err := os.Open("index.html")
	if err != nil {
		http.Error(h, "file not found", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	h.Header().Set("Content-Type", "text/html")

	http.ServeContent(h, r, "index.html", time.Time{}, file)

	//	if err := recover(); err != nil {
	//	http.Error(h, "internal server error", http.StatusInternalServerError)
	//}
}

func UploadHandle(h http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/upload" {
		http.NotFound(h, r)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(h, "method not supported", http.StatusInternalServerError)
		return
	}
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(h, "error parsing form", http.StatusInternalServerError)
		return
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(h, "file not found", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(h, "error reading file", http.StatusInternalServerError)
		return
	}
	convertedString, err := service.Convert(string(data))
	if err != nil {
		http.Error(h, "failed to convert message", http.StatusInternalServerError)
		return
	}

	timeStamp := time.Now().UTC().String()
	extension := filepath.Ext(handler.Filename)
	newFileName := fmt.Sprintf("%s%s", timeStamp, extension)

	fileExt, err := os.Create(newFileName)
	if err != nil {
		http.Error(h, "error create file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = io.WriteString(fileExt, convertedString)
	if err != nil {
		http.Error(h, "error write to file", http.StatusInternalServerError)
		return
	}

	h.Write([]byte(convertedString))

}
