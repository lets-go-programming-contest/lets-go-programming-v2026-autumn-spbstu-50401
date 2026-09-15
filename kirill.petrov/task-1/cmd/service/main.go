package main

import "fmt"

func main() {
  var d string = "horosho"

  var a, b int
  _, err := fmt.Scan(&a)
  if (err != nil) {
    d = "ploho"
  } else {
    _, err = fmt.Scan(&b)
    if (err != nil) {
      d = "ujasno"
    }
  }
  var s string
  fmt.Scan(&s)
  if (d == "horosho") {
    if (s == "+") {
      fmt.Println(a + b)
    } else if (s == "-") {
      fmt.Println(a - b)
    } else if (s == "*") {
      fmt.Println(a * b)
    } else if (s == "/") {
      if (b == 0) {
        fmt.Println("Division by zero")
      } else {
        fmt.Println(a / b)
      }
    } else {
      fmt.Println("Invalid operation")
    }
  } else if (d == "ploho"){
    fmt.Println("Invalid first operand")
  } else {
    fmt.Println("Invalid second operand")
  }
}
