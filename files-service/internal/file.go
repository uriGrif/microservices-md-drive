package internal

type File struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
