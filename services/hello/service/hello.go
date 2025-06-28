package service

import (
	"fmt"
)

type HelloService struct{}

func NewHelloService() *HelloService {
	return &HelloService{}
}

func (s *HelloService) SayHello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
