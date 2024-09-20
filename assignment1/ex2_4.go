package main

import (
    "encoding/json"
    "fmt"
)


type Product struct {
    Name string  `json:"name"`
    Price float64 `json:"price"`
    Quantity int  `json:"quantity"`
}


func productToJSON(p Product) (string, error) {
    jsonData, err := json.Marshal(p)
    if err != nil {
        return "", err
    }
    return string(jsonData), nil
}


func jsonToProduct(jsonStr string) (Product, error) {
    var p Product
    err := json.Unmarshal([]byte(jsonStr), &p)
    if err != nil {
        return p, err
    }
    return p, nil
}

func main() {

    product := Product{Name: "phonecase", Price: 21.99, Quantity: 20}

    jsonStr, err := productToJSON(product)
    if err != nil {
        fmt.Println("Error encoding to JSON:", err)
    } else {
        fmt.Println("Product as JSON:", jsonStr)
    }
    

    decodedProduct, err := jsonToProduct(jsonStr)
    if err != nil {
        fmt.Println("Error decoding JSON:", err)
    } else {
        fmt.Println("Decoded Product:", decodedProduct)
    }
}