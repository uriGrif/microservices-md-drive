package internal

import permissions "files-service/protogen/golang"

type File struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

var FileUpdatePermissionLevels []permissions.PermissionLevel = []permissions.PermissionLevel{
	permissions.PermissionLevel_OWNER,
	permissions.PermissionLevel_EDITOR,
}

var FileDeletePermissionLevels []permissions.PermissionLevel = []permissions.PermissionLevel{
	permissions.PermissionLevel_OWNER,
}
