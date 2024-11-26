package service

type HelloService struct{}

func NewHelloService() *HelloService {
	return &HelloService{}
}

func (h *HelloService) SayHi(name string) string {
	return "hello: " + name
}
