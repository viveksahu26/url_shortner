package src

type URLResponse struct {
	OriginalURL string `json:"originalURL"`
	ShortURL    string `json:"shortURL"`
}

// builds shortURL as response
func BuildURLWithResponse(host string, shortURL string, originalURL string) *URLResponse {
	resp := &URLResponse{
		OriginalURL: originalURL,
		ShortURL:    "http://" + host + "/" + shortURL,
	}
	return resp
}
