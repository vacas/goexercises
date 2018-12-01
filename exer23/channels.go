package main

import "fmt"

// Made the exercise a bit more interesting
// by passing channel through a func as a variable
func getPing(msg <-chan string) string {
  result := <-msg

  return result
}

func main() {
  messages := make(chan string)

  go func() { messages <- "ping"}()

  fmt.Println(getPing(messages))
}
