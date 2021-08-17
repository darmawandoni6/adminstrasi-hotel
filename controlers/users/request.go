package users

import "administrasi-hotel/busieness/users"

type ReqUsers struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *ReqUsers) ToDomain() *users.Domain {
	return &users.Domain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}
