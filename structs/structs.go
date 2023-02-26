package structs

type ById struct {
	Id string `json:"id"`
}

type ByIds struct {
	Id []string `json:"id"`
}

type ListRequest struct {
	Id    string `json:"id"`
	Limit int32  `json:"limit"`
	Page  int32  `json:"page"`
}

type ShortUrl struct {
	ShortUrl string `json:"short_url"`
}
