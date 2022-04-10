package Cookie

import "time"

type ICookie interface {
	GetCookie(key string) string
	SetCookie(key string, value string, expire time.Duration, path string)
	GetCookieList() map[string]string
}
