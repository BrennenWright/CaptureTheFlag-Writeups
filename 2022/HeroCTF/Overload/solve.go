package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

func main() {

    if len(os.Args) < 2 {
        fmt.Println("usage: solve.go overload.txt")
        return
    }

    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    text, err := ioutil.ReadAll(file)
    if err != nil {
        log.Fatal(err)
    }

    for _, char := range text {
        if char > 96 {
            fmt.Printf("%c", char)
        }

    }

}
