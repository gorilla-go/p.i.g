package Session

import "php-in-go/Include/Contracts/Cache"

type ISession interface {
	StartSessionManager(cache Cache.ICache)
	CloseSessionManager()
	GetSession(s string) interface{}
	SetSession(key string, v interface{})
	GetSessionList() map[string]interface{}
}
