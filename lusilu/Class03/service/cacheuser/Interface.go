package cacheuser

type Interface interface {
	Register(username string, password string) (code int, msg string)
	Login(username string, password string) (token string, err error)
	Logout(token string) (code int, msg string)
	Delete(token string) (code int, msg string)
}
