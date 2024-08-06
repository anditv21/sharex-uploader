package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

type TemplateData struct {
	Color            string
	OriginalName     string
	FileSize         string
	FileType         string
	UploadTime       string
	GeneratedSymbols string
	Author           string
	MediaFolder      string
	SiteName         string
}

func calculateFileSize(size int64) string {
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
	)

	switch {
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/float64(MB))
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/float64(KB))
	}
	return fmt.Sprintf("%d B", size)
}

func HandleUpload(config Configuration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate the token
		token := r.FormValue("key")
		if !isValidToken(token, config.Tokens) {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		err := r.ParseMultipartForm(10 << 20) // 10 MB
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		file, header, err := r.FormFile("sharex")
		if err != nil {
			http.Error(w, "Unable to get file from form", http.StatusBadRequest)
			return
		}
		defer file.Close()

		generatedSymbols := uuid.New().String()[:config.Length]
		directoryPath := filepath.Join(config.MediaFolder, generatedSymbols)
		err = os.MkdirAll(directoryPath, 0777)
		if err != nil {
			http.Error(w, "Unable to create directory", http.StatusInternalServerError)
			return
		}

		err = os.Chmod(config.MediaFolder, 0777)
		if err != nil {
			http.Error(w, "Unable to set media folder permissions", http.StatusInternalServerError)
			return
		}

		filePath := filepath.Join(directoryPath, "image"+filepath.Ext(header.Filename))
		out, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Unable to create file", http.StatusInternalServerError)
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			http.Error(w, "Unable to save file", http.StatusInternalServerError)
			return
		}

		err = os.Chmod(filePath, 0777)
		if err != nil {
			http.Error(w, "Unable to set file permissions", http.StatusInternalServerError)
			return
		}

		uploadTime := time.Now().UTC().Add(time.Hour).Format("Monday, January 2, 2006, 15:04")
		templateData := TemplateData{
			Color:            randColor(),
			OriginalName:     header.Filename,
			FileSize:         calculateFileSize(header.Size),
			FileType:         filepath.Ext(header.Filename)[1:],
			UploadTime:       uploadTime,
			GeneratedSymbols: generatedSymbols,
			Author:           config.Author,
			MediaFolder:      config.MediaFolder,
			SiteName:         config.SiteName,
		}

		templateFile := getTemplateFile(templateData.FileType)
		templateContent, err := ioutil.ReadFile(templateFile)
		if err != nil {
			http.Error(w, "Unable to read template file", http.StatusInternalServerError)
			return
		}

		for key, value := range map[string]string{
			"%color%":            templateData.Color,
			"%originalName%":     templateData.OriginalName,
			"%fileSize%":         templateData.FileSize,
			"%fileType%":         templateData.FileType,
			"%uploadTime%":       templateData.UploadTime,
			"%generatedSymbols%": templateData.GeneratedSymbols,
			"%author%":           templateData.Author,
			"%mediafolder%":      config.MediaFolder,
			"%sitename%":         config.SiteName,
		} {
			templateContent = bytes.ReplaceAll(templateContent, []byte(key), []byte(value))
		}

		err = ioutil.WriteFile(filepath.Join(directoryPath, "index.html"), templateContent, 0644)
		if err != nil {
			http.Error(w, "Unable to write HTML template", http.StatusInternalServerError)
			return
		}

		url := fmt.Sprintf("https://%s/%s/%s/", r.Host, config.MediaFolder, generatedSymbols)
		link := fmt.Sprintf(url)
		fmt.Fprintf(w, link)
	}
}

func isValidToken(token string, validTokens []string) bool {
	for _, validToken := range validTokens {
		if token == validToken {
			return true
		}
	}
	return false
}

func randColor() string {
	return fmt.Sprintf("#%06x", rand.Intn(0xffffff))
}

func getTemplateFile(fileType string) string {
	switch strings.ToLower(fileType) {
	case "mp4", "webm", "ogg", "mov", "wmv", "avi", "avchd", "mkv":
		return "templates/template_video.html"
	case "jpeg", "png", "gif", "tiff", "jpg", "jfif":
		return "templates/template_image.html"
	default:
		return "templates/template_default.html"
	}
}
