package main

import (
    "net"
    "io"
    "log"
    "os"
)

func proxy(local net.Conn, targetProtocol string, target string){
    remote, err := getRemote(targetProtocol, target)
    if (err == nil) {
        go func() {
            _, err := io.Copy(local, remote)
            log.Println("remote connection closed")
            local.Close()
            remote.Close()
            if err != nil {
//                log.Println("remote error is:", err)
            }
            //defer remote.Close()
        }()
        go func() {
            _, err := io.Copy(remote, local)
            log.Println("incoming connection closed")
            local.Close()
            remote.Close()
            if err != nil {
//                log.Printf("incoming error is:", err)
            }
            //defer remote.Close()
        }()
    }
}

func getRemote(targetProtocol string, target string) (net.Conn, error) {
    remote, err := net.Dial(targetProtocol, target)
    log.Println("Connecting remote", target, "wait...")
    if err != nil {
        log.Printf("Unable to connect - %s", err)
    } else {
        log.Println("remote connection to", target, "opened")
    }
    return remote, err
}

func main() {

    log.Println("--------------------------------")
    log.Println("smallbridge - (c) 2015 RpG")
    log.Println("--------------------------------")
    if (len(os.Args[1:]) < 2) {
        log.Fatal("local:port and remote:port (or unix:/path/to/filename) are required arguments")
        return
    }
    
    listenOn := os.Args[1]
    localProtocol := "tcp"
    target := os.Args[2]
    targetProtocol := "tcp"
    if (target[:5] == "unix:") {
        targetProtocol = "unix"
        target = target[5:]
    } else if(target[:4] == "tcp:") {
        targetProtocol = "tcp"
        target = target[4:]
    }
    log.Println("bridge targeting:", targetProtocol, target)
    
    server, err := net.Listen(localProtocol, listenOn)
    if err != nil {
		log.Fatalf("Port/Address busy - %s", err)
	}
    defer server.Close()

    log.Println("bridge listening on", listenOn)
    for {
        local, err := server.Accept()
        if err != nil {
            log.Fatalf("listen Accept failed - %s", err)
        }
        log.Println("incoming connection accepted from:", local.RemoteAddr())
        proxy(local, targetProtocol, target)
    }
}
