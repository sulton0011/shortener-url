package v1

type CreateUrlRequest struct {
	Title        string `json:"title"`
	LongUrl      string `json:"long_url"`
	ShortUrl     string `json:"short_url"`
	ExpiresAt    string `json:"expires_at"`
	ExpiresCount int64  `json:"expires_count"`
}

type CreateUrlResponse struct {
	Id       string `json:"id"`
	ShortUrl string `json:"short_url"`
}

type UpdateUrlRequest struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	ShortUrl     string `json:"short_url"`
	ExpiresAt    string `json:"expires_at"`
	ExpiresCount int64 `json:"expires_count"`
	UsedCount    int64  `json:"used_count"`
}

type GetUrlResponse struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	ShortUrl     string `json:"short_url"`
	LongUrl      string `json:"long_url"`
	ExpiresAt    string `json:"expires_at"`
	ExpiresCount int64  `json:"expires_count"`
	UsedCount    int64  `json:"used_count"`
	QrCode       []byte `json:"qr_code"`
	CreatedBy    string `json:"created_by"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type GetUrlListResponse struct {
	Count int64            `json:"count"`
	Urls  []GetUrlResponse `json:"urls"`
}

type Message struct {
	Message string `json:"message"`
}
