package models

type UserPrimaryKey struct {
	Id string `json:"id"`
}

type User struct {
	Id        string `json:"id"`
	FullName  string `json:"full_name"`
	NickName  string `json:"nick_name"`
	Photo     string `json:"photo"`
	Birthday  string `json:"birthday"`
	Location  string `json:"location"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateUser struct {
	FullName string `json:"full_name"`
	NickName string `json:"nick_name"`
	Photo    string `json:"photo"`
	Birthday string `json:"birthday"`
	Location string `json:"location"`
}

type UpdateUser struct {
	Id       string `json:"id"`
	FullName string `json:"full_name"`
	NickName string `json:"nick_name"`
	Photo    string `json:"photo"`
	Birthday string `json:"birthday"`
	Location string `json:"location"`
}

type GetListUserRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Fields string `json:"fields"`
	Sort   string `json:"sort"`
}

type GetListUserResponse struct {
	Count int64   `json:"count"`
	Users []*User `json:"users"`
}
