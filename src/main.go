package main

import "syscall/js"

func main() {

    document := js.Global().Get("document")
    p := document.Call("createElement", "p")
    p.Set("innerHTML", "Hello World!")
    document.Get("body").Call("appendChild", p)
    
}
