package v1

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Surname          string `json:"surname"`
	MiddleName       string `json:"middle_name"`
	PhoneNumber      string `json:"phone_number"`
	AccessToken      string `json:"access_token"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}
