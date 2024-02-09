package main

import (
    "fmt"
    "crypto/rand"
)

func main() {
    var numGUIDs int

    // Prompt the user for the number of GUIDs
    fmt.Println("Welcome to GUID Gen by Xanny.")
    fmt.Println("How many GUIDs do you need today? ")
    for {
        _, err := fmt.Scanln(&numGUIDs)
        if err != nil || numGUIDs <= 0 {
            fmt.Println("Invalid input. Please enter a number above zero.")
            continue
        }
        break
    }

    // Generate and print GUIDs
    fmt.Println("Generated GUIDs:")
    for i := 0; i < numGUIDs; i++ {
        guid, err := generateGUID()
        if err != nil {
            fmt.Println("Error generating GUID:", err)
            continue
        }
        fmt.Println(guid)
    }
}

// Generate a random GUID (version 4)
func generateGUID() (string, error) {
    b := make([]byte, 16)
    _, err := rand.Read(b)
    if err != nil {
        return "", err
    }

    // Set version bits (4th most significant bit to 1)
    b[4] = (b[4] | 0x80) & 0xaf

    // Set variant bits (most significant 2 bits to 10)
    b[8] = (b[8] | 0x80) & 0xbf

    return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}
