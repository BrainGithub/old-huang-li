package main

import (
    "fmt"
)

type Server3 struct {
    Addr     string
    Port     int
    Protocol string
    //Timeout  time.Duration
    MaxConns int
    //TLS      tls.Config
}

type ServerBuilder struct {
    Server3
}

func (sb *ServerBuilder) Create(addr string, port int) *ServerBuilder {
    sb.Addr = addr
    sb.Port = port
    return sb
}

func (sb *ServerBuilder) WithProtocol(protocol string) *ServerBuilder {
    sb.Protocol = protocol
    return sb
}

//func (sb *ServerBuilder) WithTLS(tls tls.Config) *ServerBuilder {
//    sb.TLS = tls
//    return sb
//}

func (sb *ServerBuilder) Build() Server3 {
    return sb.Server3
}

func main() {
    sb := ServerBuilder{}
    server := sb.Create("127.0.0.1", 8080).WithProtocol("udp").Build()
    fmt.Println(server)
}
