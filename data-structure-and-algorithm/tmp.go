package main

import (
    "log"
    "time"
)

func main() {
    ch := make(chan string)
    go func() {
        ch <- "hello"
        time.Sleep(10 * time.Second)
        ch <- "hello"
        close(ch)
    } ()
    for d := range ch {
        log.Printf("d:%+v", d)
    }
    log.Printf("outputs:+%v", "aaa")
}
