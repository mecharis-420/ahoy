package parsing

import (
    "encoding/json"
    "fmt"
    "os"
)

type AhoyKeys struct {
    Name    string `json:"name"`
    Desc    string `json:"desc"`
    Version string `json:"ver"`
}

func parsing() {
    file, err := os.Open(".ahoyfile")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    var info AhoyKeys
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&info)
    if err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }

    fmt.Println(info)
}