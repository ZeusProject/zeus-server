package database

import "math/rand"

type AuthCookieKind int

//go:generate stringer -type=AuthCookieKind
const (
	_ AuthCookieKind = iota

	AccountAuthCookie
	InterAuthCookie
)

const AuthCookieKindMask = 1 << 63

type AuthCookie struct {
	Cookie    uint64
	AccountID uint32
}

type AuthCookieStore interface {
	Create(cookie AuthCookie) error
	VerifyAndConsume(cookie AuthCookie) (bool, error)
}

func GenerateAuthCookie(accountID uint32, kind AuthCookieKind) AuthCookie {
	cookie := uint64(rand.Int63())

	if kind == AccountAuthCookie {
		cookie |= AuthCookieKindMask
	} else {
		cookie &^= AuthCookieKindMask
	}

	return AuthCookie{
		Cookie:    cookie,
		AccountID: accountID,
	}
}

func (ac AuthCookie) Kind() AuthCookieKind {
	if ac.Cookie&AuthCookieKindMask != 0 {
		return AccountAuthCookie
	} else {
		return InterAuthCookie
	}
}
