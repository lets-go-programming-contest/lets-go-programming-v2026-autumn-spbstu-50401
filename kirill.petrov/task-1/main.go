package main

import "fmt"

func main() {
  var a, b int
  fmt.Scan(&a, &b)
  var s string
  fmt.Scan(&s)
  if (s == "+") {
    fmt.Println(a + b)
  } else if (s == "-") {
    fmt.Println(a - b)
  } else if (s == "*") {
    fmt.Println(a * b)
  } else {
    if (b == 0) {
      fmt.Println("Division by zero")
    } else {
      fmt.Println(a / b)
    }
  }
}
