// If you uncomment the commented code,
// and follow the instructions before the defaul case
// to make the channel pass in the `select`

package main

import "fmt"
// import "time"

func main() {
  messages := make(chan string)
  signals := make(chan bool)

  select {
  case msg := <-messages:
    fmt.Println("received message", msg)
  default:
    fmt.Println("no message received")
  }

  // go func() {
  //   time.Sleep(1 * time.Second)
  //   messages <- "hi"
  // }()

  msg := "hi"
  select {
  case messages <- msg :
    fmt.Println("sent message", msg)
  // comment here to make the channel pass
  default:
    fmt.Println("no message sent")
  }

  select {
  case msg := <-messages:
    fmt.Println("received message", msg)
  case sig := <-signals:
    fmt.Println("received signal", sig)
  default:
    fmt.Println("no activity")
  }

}
