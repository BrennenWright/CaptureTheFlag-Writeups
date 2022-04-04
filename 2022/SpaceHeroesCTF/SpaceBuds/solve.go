 package main

 import (
 	"fmt"
 	"net/http"
 	"os"
 	"log"
    "strings"
    "io/ioutil"
 )

 func main() {
 	username := os.Args[1]

 	
 	cookie := &http.Cookie{
        Name:   "userID",
        Value:  username,
        MaxAge: 300,
    }
    
    req, err := http.NewRequest("POST", "http://45.79.204.27/getcookie", strings.NewReader("nm="+username))
    if err != nil {
        log.Fatalf("Got error %s", err.Error())
    }
    req.Header.Set("Content-Type", "application/json; charset=UTF-8")
    req.AddCookie(cookie)
    
    for _, c := range req.Cookies() {
        fmt.Println(c)
    }
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalf("Error occured. Error is: %s", err.Error())
    }
    defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
 }
