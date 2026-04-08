package handler

type instanceRequest struct {
	IDInstance       string `json:"idInstance"`
	APITokenInstance string `json:"apiTokenInstance"`
}

type sendMessageRequest struct {
	instanceRequest
	ChatID  string `json:"chatId"`
	Message string `json:"message"`
}

type sendFileByUrlRequest struct {
	instanceRequest
	ChatID   string `json:"chatId"`
	URLFile  string `json:"urlFile"`
	FileName string `json:"fileName"`
}
