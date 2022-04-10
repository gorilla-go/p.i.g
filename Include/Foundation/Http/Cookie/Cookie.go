package Cookie

import "time"

type Cookie struct {
}

func (c *Cookie) GetCookie(key string) string {
	//TODO implement me
	panic("implement me")
}

func (c *Cookie) SetCookie(key string, value string, expire time.Duration, path string) {
	//TODO implement me
	panic("implement me")
}

func (c *Cookie) GetCookieList() map[string]string {
	//TODO implement me
	panic("implement me")
}
