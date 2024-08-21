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

type UploadRequestInput struct {
	File string `json:"file"`
}

type UploadRequest struct {
	Action map[string]interface{} `json:"action"`
	Input  struct {
		UserInput UploadRequestInput `json:"input"`
	} `json:"input"`
	RequestQuery     string                 `json:"request_query"`
	SessionVariables map[string]interface{} `json:"session_variables"`
}

type UploadResponse struct {
	ImageURL string `json:"imageUrl"`
}

var cld *cloudinary.Cloudinary

func init() {
	var err error
	cld, err = cloudinary.NewFromParams("debm4el9a", "753363681219597", "58A4AzLoPplWqfB3Tusx-N_Cfck")
	if err != nil {
		fmt.Println("Error initializing Cloudinary:", err)
	}
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Invalid request method"}`, http.StatusMethodNotAllowed)
		return
	}

	var req UploadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Failed to decode request"}`, http.StatusBadRequest)
		return
	}

	fileData := req.Input.UserInput.File
	if fileData == "" {
		http.Error(w, `{"error": "No file data provided"}`, http.StatusBadRequest)
		return
	}

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

	response := UploadResponse{
		ImageURL: resp.SecureURL,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
	}
}
