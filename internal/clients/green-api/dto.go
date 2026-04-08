package greenapi

type SendMessageRequest struct {
	ChatID  string `json:"chatId"`
	Message string `json:"message"`
}

type SendFileByUrlRequest struct {
	ChatID   string `json:"chatId"`
	URLFile  string `json:"urlFile"`
	FileName string `json:"fileName"`
	Caption  string `json:"caption,omitempty"`
}

type SendMessageResponse struct {
	IDMessage string `json:"idMessage"`
}

type SendFileByUrlResponse struct {
	IDMessage string `json:"idMessage"`
}
