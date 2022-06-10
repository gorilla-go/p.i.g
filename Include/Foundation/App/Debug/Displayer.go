package Debug

import (
	"fmt"
	"strings"
)

type Console struct {
}

func (c Console) Log(s ...string) {
	fmt.Println(strings.Join(s, ","))
}
