package apiserver

import "fmt"

type APIServer struct {
	port string
}

func NewAPIServer() APIServer {
	fmt.Println("New API server")
	return APIServer{
		port: "8080",
	}
}
