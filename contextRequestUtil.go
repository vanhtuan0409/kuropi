package kuropi

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var (
	formDecoder = schema.NewDecoder()
)

func (c *context) Query(name string) string {
	return c.request.URL.Query().Get(name)
}

func (c *context) Var(name string) string {
	return mux.Vars(c.request)[name]
}

func (c *context) ParseForm(dst interface{}) error {
	if err := c.request.ParseForm(); err != nil {
		return err
	}
	return formDecoder.Decode(dst, c.request.PostForm)
}

func (c *context) ParseJson(dst interface{}) error {
	decoder := json.NewDecoder(c.request.Body)
	return decoder.Decode(dst)
}
