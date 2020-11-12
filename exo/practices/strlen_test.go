package main

import "testing"

func TestStrlen(t* testing.T) {
    if strlen("strlen") != len("strlen") {
        t.Error("NOT EQUAL")
    }
}
