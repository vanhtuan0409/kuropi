package kuropi

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)

type Route struct {
	Method      Method
	Middlewares []Middleware
	Path        string
	HandlerFunc HandlerFunc
}
