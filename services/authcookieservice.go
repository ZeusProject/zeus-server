package services

import "github.com/zeusproject/zeus-server/model"

type AuthCookieService interface {
	Create(accountID uint32) (model.AuthCookie, error)
	VerifyAndConsume(cookie model.AuthCookie) (bool, error)
}
