package main

import (
    "fmt"
    "github.com/pkg/errors"
)

type authorizationError struct {
    operation string
    err error // original error
}

func (e *authorizationError) Error() string {
    return fmt.Sprintf("authorization failed during %s: %v", e.operation, e.err)
}

func main() {
    //错误包装
    err := errors.New("test error wrap1")
    err = errors.Wrap(err, "wraper1")
    err = errors.Wrap(err, "wraper2")

    t := errors.Cause(err)
    fmt.Printf("11111111err type:%+v", t)

    // Cause接口
    switch t := errors.Cause(err).(type) {
        //case *MyError: // handle specifically
        default: // unknown error
            fmt.Printf("2222err type:%+v", t)
    }
}
