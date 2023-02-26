package v1

type CreateUser struct {
	Name             string `json:"name"`
	Surname          string `json:"surname"`
	MiddleName       string `json:"middle_name"`
	TelegramUsername string `json:"telegram_username"`
	PhoneNumber      string `json:"phone_number"`
}

type GetUsersById struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Surname          string `json:"surname"`
	MiddleName       string `json:"middle_name"`
	TelegramUsername string `json:"telegram_username"`
	PhoneNumber      string `json:"phone_number"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type GetUserListResponse struct {
	Count int64          `json:"count"`
	Users []GetUsersById `json:"users"`
}

type UpdateUserToken struct {
	Id           string `json:"id"`
}

type CreateMessageSupportRequst struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

type CreateMessageUserRequst struct {
	Id      string   `json:"id"`
	Message string   `json:"message"`
	Image   string   `json:"image"`
	Links   []string `json:"links"`
}

type CreateMessageRequst struct {
	Id        string   `json:"id"`
	Message   string   `json:"message"`
	CreatedBy string   `json:"created_by"`
	Image     string   `json:"image"`
	Links     []string `json:"links"`
	IsSeen    bool     `json:"is_seen"`
}

type CreateMessageAllUsersRequest struct {
	Message string   `json:"message"`
	Image   string   `json:"image"`
	Links   []string `json:"links"`
}

type GetMessage struct {
	Id        int64    `json:"id"`
	Message   string   `json:"message"`
	Send      bool     `json:"send"`
	CreatedAt string   `json:"created_at"`
	Image     string   `json:"image"`
	Links     []string `json:"links"`
}

type GetMessageResponse struct {
	Count    int64        `json:"count"`
	Messages []GetMessage `json:"messages"`
}
