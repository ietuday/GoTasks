package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    jsonData := `{"name": "John", "age": 30, "isEmployed": true}`

    var data map[string]interface{}
    err := json.Unmarshal([]byte(jsonData), &data)
    if err != nil {
        fmt.Println("Error parsing JSON:", err)
        return
    }

    for key, value := range data {
        switch v := value.(type) {
        case string:
            fmt.Printf("%s is a string: %s\n", key, v)
        case float64:
            fmt.Printf("%s is a number: %f\n", key, v)
        case bool:
            fmt.Printf("%s is a boolean: %t\n", key, v)
        default:
            fmt.Printf("%s is of an unknown type: %T\n", key, v)
        }
    }
}