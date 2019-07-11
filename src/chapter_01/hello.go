package main 

import "fmt"

const englishhelloPrexif = "Hello, "

func Hello(name string) string {
    if name == "" {
        name = "world"
    }

    return englishhelloPrexif + name

}

func main() {
    fmt.Print(Hello(""))
}
