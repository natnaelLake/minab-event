package controller

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type UploadImagesInput struct {
	Files []string `json:"files"`
}

type UploadImagesRequest struct {
	Action          map[string]interface{} `json:"action"`
	Input           struct {
		UserInput UploadImagesInput `json:"input"`
	} `json:"input"`
	RequestQuery     string                 `json:"request_query"`
	SessionVariables map[string]interface{} `json:"session_variables"`
}

type UploadImagesResponse struct {
	ImageURLs []string `json:"imageUrls"`
}

var cld *cloudinary.Cloudinary

func init() {
	var err error
	cld, err = cloudinary.NewFromParams("debm4el9a", "753363681219597", "58A4AzLoPplWqfB3Tusx-N_Cfck")
	if err != nil {
		fmt.Println("Error initializing Cloudinary:", err)
	}
}

func UploadImagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Invalid request method"}`, http.StatusMethodNotAllowed)
		return
	}

	var req UploadImagesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Failed to decode request"}`, http.StatusBadRequest)
		return
	}

	fileDataList := req.Input.UserInput.Files
	if len(fileDataList) == 0 {
		http.Error(w, `{"error": "No file data provided"}`, http.StatusBadRequest)
		return
	}

	var imageUrls []string
	for _, fileData := range fileDataList {
		decodedFile, err := base64.StdEncoding.DecodeString(fileData)
		if err != nil {
			http.Error(w, `{"error": "Failed to decode file"}`, http.StatusBadRequest)
			return
		}

		resp, err := cld.Upload.Upload(r.Context(), bytes.NewReader(decodedFile), uploader.UploadParams{Folder: "tasks"})
		if err != nil {
			http.Error(w, `{"error": "Failed to upload to Cloudinary"}`, http.StatusInternalServerError)
			return
		}

		imageUrls = append(imageUrls, resp.SecureURL)
	}

	response := UploadImagesResponse{
		ImageURLs: imageUrls,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
	}
}
