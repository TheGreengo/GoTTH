package main

import (
    "net/http"
    "github.com/a-h/templ"
)

func main() {
    comp := hello("TheGreengo")

    http.Handle("/", templ.Handler(comp)) 
    
    http.ListenAndServe(":4321", nil)
}
