package structs

type ById struct {
	Id string `json:"id"`
}

type ByIds struct {
	Id []string `json:"id"`
}

type ListRequest struct {
	Limit int32 `json:"limit"`
	Page  int32 `json:"page"`
}
