package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
)

//----------------------------------config-----------------------
const target = "https://challenges.hackrocks.com/lechuck/user/lechuck" //chals target goes here
const flag = "{flag}"                                                  //events flag prefix goes here

func main() {

    fmt.Println("--------------------------------------------------------")
    fmt.Println("--------------Solve LeChuck is back!--------------------")
    fmt.Println("--------------by: Ir0nM4n-------------------------------")
    fmt.Println("--------------@: https://github.com/brennenwright-------")
    fmt.Println("\n-----Setup")

    //step one: attack the site looking for endpoints
    //https://challenges.hackrocks.com/lechuck/user/lechuck
    fmt.Println("\n-----Step One: check endpoints")

    var payload = "https://challenges.hackrocks.com/lechuck/user/lechuck"

    fmt.Println("trying payload:" + payload)

    resp, err := http.Get(target)
    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

    resptext, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    //parse for the flag
    fmt.Println("\n-----Flag: ")
    i := strings.Index(string(resptext), flag)
    fmt.Println(string(resptext)[i+len(flag):])

}
