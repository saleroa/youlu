package model

import "time"

type User struct {
	UID      int64     `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Avatar   string    `json:"avatar"`
	Email    string    `json:"email"`
	CreatAt  time.Time `json:"creat_at"`
	UpdateAt time.Time `json:"update_at"`
}

type UserInfo struct {
	UID      int64  `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}
