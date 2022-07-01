package main

import (
  "fmt"
  "time"
)

func hello(done chan bool) {
  fmt.Println("hello go routine is going to sleep")
  time.Sleep(4 * time.Second)
  fmt.Println("hello go routine awake and going to write to done")
  done <- true
}

func main() {
  fmt.Println("Main going to call hello go routine")
  done := make(chan bool)
  go hello(done)
  <-done
  fmt.Println("Main received data")
}
