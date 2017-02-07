package main

import("fmt"
       "bufio"
       "github.com/atotto/clipboard"
       "os")

func intToHexString(num int32) string {
    hex := ""
    for (num > 0) {
        t := num % 16
        num = num / 16
        if(t < 10) {
            hex = string('0' + t) + hex
        } else if (t == 10) {
            hex = "a" + hex
        } else if (t == 11) {
            hex = "b" + hex
        } else if (t == 12) {
            hex = "c" + hex
        } else if (t == 13) {
            hex = "d" + hex
        } else if (t == 14) {
            hex = "e" + hex
        } else {
            hex = "f" + hex
        }
    }
    return hex
}

func toHex(str string) string {
    hex := ""
    for i := 0; i < len(str); i++ {
        chr := int32(str[i])
        hex += intToHexString(chr)
    }
    return hex
}

func main() {
    str := ""
    read := bufio.NewReader(os.Stdin)
    enoughArgs := (len(os.Args) > 1)
    if(enoughArgs) {
        str = os.Args[1]
    } else {
        buf, _, _ := read.ReadLine()
        str = string(buf)
    }
    hex := toHex(str)
    fmt.Println(hex)
    clipboard.WriteAll(hex)
    fmt.Println("Skopiowano do schowka")
    if(!enoughArgs) {
        fmt.Println("Wpisz cokolwiek")
        read.ReadLine()
    }
}