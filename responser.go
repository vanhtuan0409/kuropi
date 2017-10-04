package kuropi

import (
	"net/http"
)

type Responser interface {
	Handle(rw http.ResponseWriter, result interface{}, err error)
}
