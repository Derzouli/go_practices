package main

// string is in bytes
func strlen(str string) int {
    var len int
// runes in encoding with unicode
    runeStr :=  []rune(str)
    len = 0
// difference with C in Go no terminating string with "\0"
    for i := range runeStr {
        _ = i
        len += 1
    }
    return len
}

