package kuropi

type Middleware func(HandlerFunc) HandlerFunc
