package main

import (
   "io/ioutil"
   "fmt"
   "net/http"
)


func main() {

    for {
       resp, _ := http.Get("http://173.230.134.127/seq/0872ed61-66a6-4e45-8a12-fae6c98621f1")

        //We Read the response body on the line below.
       body, _ := ioutil.ReadAll(resp.Body)

        //Convert the body to type string
       sb := string(body)
       fmt.Printf(sb)
       }
}

