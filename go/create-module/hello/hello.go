package main

import (
  "fmt"
  "log"
  "example.com/greetings"
)

func main() {
  log.SetPrefix("greetings: ")
  log.SetFlags(0)

  names:=[]string{"Mehmet", "Cihat", "Altuntaş"}
  messages, err := greetings.Hellos(names)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Test message")
  //message:=greetings.Hello("Cihat")
  fmt.Println(messages)
}
