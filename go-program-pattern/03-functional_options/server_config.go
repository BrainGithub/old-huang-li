package main

import (
    "crypto/tls"
    "time"
)

type Server struct {
    Addr     string
    Port     int
    Protocol string
    Timeout  time.Duration
    MaxConns int
    TLS      tls.Config
}

func NewDefaultServer(addr string, port int) (*Server, error) {
    return &Server{
        addr,
        port,
        "tcp",
        30 * time.Second,
        1024,
        nil}, nil
}

func NewTLSServer(addr string, port int, tls tls.Config) (*Server, error) {
    return &Server{
        addr,
        port,
        "tcp",
        30 * time.Second,
        1024,
        tls}, nil
}

func NewServerWithTimeout(addr string, port int, timeout time.Duration) (*Server, error) {
    return &Server{
        addr,
        port,
        "tcp",
        timeout,
        1024,
        nil}, nil
}

func NewTLSServerWithMaxConnAndTimeout(addr string, port int, maxconns int, timeout time.Duration, tls tls.Config) (*Server, error) {
    return &Server{addr, port, "tcp", 30 * time.Second, maxconns, tls}, nil
}

type Config struct {
    Protocol string
    Timeout time.Duration
    Maxconns int
    TLS tls.Config
}

type Server2 struct {
    Addr string
    Port int
    Conf Config
}

func NewServer(addr string, port int, conf *Config) (*Server2, error) {
    //return &Server2{addr, port, *conf}, nil
    return nil, nil
}

//func main() {
//    svr1, _ := NewServer("localhost", 9000, nil)
//
//    conf := &Config {}
//    svr2, _ := NewServer("localhost", 9000, conf)
//
//    fmt.Printf("svr1:%+v, svr2:%+v", svr1, svr2)
//}
