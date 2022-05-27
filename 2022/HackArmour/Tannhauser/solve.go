package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    filename := "stripped.txt"

    filebuffer, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    inputdata := string(filebuffer)
    data := bufio.NewScanner(strings.NewReader(inputdata))
    data.Split(bufio.ScanRunes)

    var count = 1
    for data.Scan() {
        if data.Bytes()[0] == '\n' {
            //fmt.Print("")
        } else if data.Bytes()[0] == 32 {
            fmt.Print(".")
        } else if data.Bytes()[0] == 9 {
            fmt.Print("-")
        }

        //if count == 5 {
        //    fmt.Print(" ")
        //    count = 0
        //}
        count++
    }
}
