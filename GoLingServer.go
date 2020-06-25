package main

import "net"

const (
	HOST_NAME = "GoLing Server"
	HOST_PORT = "2839"
)

type Server struct {
	port string
	host string
}

func StartServer() {
	connection, connectionError := net.Listen(HOST_NAME, HOST_PORT)
}
