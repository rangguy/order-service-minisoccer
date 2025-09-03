package clients

type MidtransResponse struct {
	Code   int          `json:"code"`
	Status string       `json:"status"`
	Data   MidtransData `json:"data"`
}

type MidtransData struct {
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
}
