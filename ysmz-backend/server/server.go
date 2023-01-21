package server

func Init() {
	r := newRouter()
	r.Run()
}
