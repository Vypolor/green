package handler

import (
	"encoding/json"
	"net/http"

	greenapi "green/internal/clients/green-api"
)

func (h *Handler) GetSettings(w http.ResponseWriter, r *http.Request) {
	var req instanceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if req.IDInstance == "" || req.APITokenInstance == "" {
		writeError(w, http.StatusBadRequest, "idInstance and apiTokenInstance are required")
		return
	}

	result, err := h.greenAPI.GetSettings(r.Context(), req.IDInstance, req.APITokenInstance)
	if err != nil {
		writeError(w, http.StatusBadGateway, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (h *Handler) GetStateInstance(w http.ResponseWriter, r *http.Request) {
	var req instanceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if req.IDInstance == "" || req.APITokenInstance == "" {
		writeError(w, http.StatusBadRequest, "idInstance and apiTokenInstance are required")
		return
	}

	result, err := h.greenAPI.GetStateInstance(r.Context(), req.IDInstance, req.APITokenInstance)
	if err != nil {
		writeError(w, http.StatusBadGateway, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (h *Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	var req sendMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if req.IDInstance == "" || req.APITokenInstance == "" {
		writeError(w, http.StatusBadRequest, "idInstance and apiTokenInstance are required")
		return
	}
	if req.ChatID == "" || req.Message == "" {
		writeError(w, http.StatusBadRequest, "chatId and message are required")
		return
	}

	result, err := h.greenAPI.SendMessage(r.Context(), req.IDInstance, req.APITokenInstance, greenapi.SendMessageRequest{
		ChatID:  req.ChatID,
		Message: req.Message,
	})
	if err != nil {
		writeError(w, http.StatusBadGateway, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (h *Handler) SendFileByUrl(w http.ResponseWriter, r *http.Request) {
	var req sendFileByUrlRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if req.IDInstance == "" || req.APITokenInstance == "" {
		writeError(w, http.StatusBadRequest, "idInstance and apiTokenInstance are required")
		return
	}
	if req.ChatID == "" || req.URLFile == "" {
		writeError(w, http.StatusBadRequest, "chatId and urlFile are required")
		return
	}

	fileName := req.FileName
	if fileName == "" {
		fileName = "file"
	}

	result, err := h.greenAPI.SendFileByUrl(r.Context(), req.IDInstance, req.APITokenInstance, greenapi.SendFileByUrlRequest{
		ChatID:   req.ChatID,
		URLFile:  req.URLFile,
		FileName: fileName,
	})
	if err != nil {
		writeError(w, http.StatusBadGateway, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}
