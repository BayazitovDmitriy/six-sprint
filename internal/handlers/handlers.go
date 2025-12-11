package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HtmlReturn(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "../index.html")
}

func HtmlParse(w http.ResponseWriter, req *http.Request) {

	file, handler, err := req.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	info, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	line, err := service.ConvertToMorse(string(info))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nameFile := time.Now().UTC().String() + filepath.Ext(handler.Filename)

	err = os.WriteFile(nameFile, []byte(info), 0755)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(line))
}
