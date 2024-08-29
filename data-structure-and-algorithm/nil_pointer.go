package main

import (
    "fmt"
    "os"
)

func Foo() error {
    var err *os.PathError = nil
    return err
}

func main() {
    err := Foo()
    fmt.Println(err == nil) // false
    if err != nil {
        fmt.Printf("nil != nil")
    }

    fmt.Printf("%+v", err)        // <nil>
}
