package internal

type CreateFileRequest struct {
	Name string `json:"name"`
}

type ListFile struct {
	Id              string `json:"id"`
	OwnerId         string `json:"owner_id"`
	Name            string `json:"name"`
	CreatedAt       string `json:"created_at"`
	PermissionLevel string `json:"permission_level"`
}
