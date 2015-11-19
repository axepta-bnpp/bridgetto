package main

import (
    "net"
    "io"
    "log"
    "os"
)

func proxy(local net.Conn, target string){
    remote, err := getRemote(target)
    if (err == nil) {
        go func() {
            _, err := io.Copy(local, remote)
            log.Println("remote connection closed")
            local.Close()
            if err != nil {
                log.Println("error is: %v", err)
            }
            defer remote.Close()
        }()
        go func() {
            _, err := io.Copy(remote, local)
            log.Println("incoming connection closed")
            if err != nil {
                log.Printf("error is: %v", err)
            }
            defer remote.Close()
        }()
    }
}

func getRemote(target string) (net.Conn, error) {
    remote, err := net.Dial("tcp", target)
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
    log.Println("the bridge - (c) 2015 RpG")
    log.Println("--------------------------------")
    if (len(os.Args[1:]) < 2) {
        log.Fatal("local:port and remote:port are required arguments")
        return
    }

    listenOn := os.Args[1]
    target := os.Args[2]

    server, err := net.Listen("tcp", listenOn)
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
        proxy(local, target)
    }
}
