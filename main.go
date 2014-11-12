package main

import (
    "log"
    "net/http"
    "os/exec"

    "github.com/Unknwon/macaron"
)

func main() {
    m := macaron.Classic()
    m.Get("/auto", autoHandler)

    log.Println("Server is running...")
    log.Println(http.ListenAndServe("0.0.0.0:4000", m))
}

func autoHandler(ctx *macaron.Context) string {
    //output, _ := exec.Command("cmd.exe", "/c", "ping", "127.0.0.1").Output()
    output, _ := exec.Command("/bin/sh", "-c", "ping 127.0.0.1").Output()
    log.Println(string(output))

    return "the request path is: " + ctx.Req.RequestURI
}