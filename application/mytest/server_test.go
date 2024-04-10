package mytest

import (
	"testing"
	"wailstemplate/application/services"
)

func TestServer(t *testing.T) {
	StaticServerInstance := services.NewStaticServer()
	StaticServerInstance.StartStaticServer("C:\\Users\\Administrator\\Pictures", "8080")
	StaticServerInstance.StartStaticServer("C:\\Users\\Administrator\\Pictures\\dlh", "8080")
	select {}
}
