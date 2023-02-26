package v1

type CreateUser struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	MiddleName string `json:"middle_name"`
	Email      string `json:"email"`
	Login      string `json:"login"`
	Password   string `json:"password"`
}

type GetUsersById struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	MiddleName string `json:"middle_name"`
	Email      string `json:"email"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
type GetUsersByLogin struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	MiddleName string `json:"middle_name"`
	Email      string `json:"email"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Login      string `json:"login"`
	Password   string `json:"password"`
}

type GetUserListResponse struct {
	Count int64          `json:"count"`
	Users []GetUsersById `json:"users"`
}
