package main

import (
  "fmt"
)

func f(bolshemenshe string, now int, add int) int {
  if (string(bolshemenshe[0]) == ">" && string(bolshemenshe[1] == "=")) {
    if (now < add) {
      return add
    } else {
      return now
    }
  } else {
    if (now < add) {
      return now
    } else {
      return add
    }
  }
}

func main() {
  var N int
  var Kolichestvo int

}
