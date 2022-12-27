package main

import (
    "fmt"
    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    var message string
    message = greetings.Hello("Gladys")
    fmt.Println(message)
}

