package kuropi

func (a *app) Responser(name string, rs Responser) {
	a.responsers[name] = rs
}
