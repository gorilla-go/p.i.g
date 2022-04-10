package Cache

import "time"

type Cache struct {
	Key    string
	Value  interface{}
	Expire time.Time
}
