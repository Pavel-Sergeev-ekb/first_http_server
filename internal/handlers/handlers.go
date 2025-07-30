package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Pavel-Sergeev-ekb/first_http_server/internal/service"
)

func BaseHandle(h http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(h, "method not supporting", http.StatusInternalServerError)
		return
	}
	http.ServeFile(h, r, "index.html")
}

func UploadHandle(h http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(h, "method not supported", http.StatusInternalServerError)
		return
	}

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(h, "error parsing form", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(h, "file not found", http.StatusInternalServerError)
		fmt.Println(err)
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
	timeStamp := strings.ReplaceAll(
		time.Now().UTC().Format("2006-01-02_15-04-05"),
		":", "-")
	extension := filepath.Ext(handler.Filename)
	newFileName := fmt.Sprintf("%s%s", timeStamp, extension)

	dir := "./uploads"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			http.Error(h, "error creating directory", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
	}

	fileExt, err := os.Create(filepath.Join(dir, newFileName))
	if err != nil {
		http.Error(h, "error create file", http.StatusInternalServerError)
		fmt.Println(err)
		return

	}
	defer fileExt.Close()

	_, err = io.WriteString(fileExt, convertedString)
	if err != nil {
		http.Error(h, "error write to file", http.StatusInternalServerError)
		return
	}

	if _, err := h.Write([]byte(convertedString)); err != nil {
		http.Error(h, "error writing response", http.StatusInternalServerError)
		return
	}
	h.WriteHeader(http.StatusOK)
}
