package main

import (
    "fmt"

    "encoding/hex"
    "crypto/sha256"
)

func toHash(password string) string {
    converted := sha256.Sum256([]byte(password))
    return hex.EncodeToString(converted[:])
}

func main() {
    str := "test"
    fmt.Printf("string: %v\n", str)
    fmt.Printf("hash   : %v\n", toHash(str))
    fmt.Printf("hash re: %v\n", toHash(str))
    str = "test2"
    fmt.Printf("string: %v\n", str)
    fmt.Printf("hash   : %v\n", toHash(str))
    fmt.Printf("hash re: %v\n", toHash(str))
}
