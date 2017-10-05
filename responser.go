package kuropi

import (
	"net/http"
)

type Responser interface {
	Handle(rw http.ResponseWriter, result interface{}, err error)
}

type ResponserMap map[string]Responser

func (rm ResponserMap) Clone() ResponserMap {
	responsers := map[string]Responser{}
	for name, responser := range rm {
		responsers[name] = responser
	}
	return responsers
}
