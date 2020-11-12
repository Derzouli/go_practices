package main

import (
    "testing"
)

func TestHelloWorld(t *testing.T) {
    if hello_world() != "Hello World" {
        t.Error("Hello World function is not working")
    }
}
