package main

import "fmt"
import "time"

// broke the exercise into separate functions
func main() {
  thisNum(1);
  thisNum(2);
  thisNum(3);

  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend")
  default:
    fmt.Println("It's the weekday")
  }

  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("It's before noon")
  default:
    fmt.Println("It's after noon")
  }

  whatAmI(true)
  whatAmI(1)
  whatAmI("hey")
}

func whatAmI(i interface{}) {
  switch t := i.(type) {
  case bool:
    fmt.Println("I'm a bool")
  case int:
    fmt.Println("I'm an int")
  default:
    fmt.Printf("Don't know type %T\n", t)
  }
}

func thisNum(n int) {
  fmt.Println("Write ", n, " as ")
  switch n {
  case 1:
    fmt.Println("one")

  case 2:
    fmt.Println("two")

  case 3:
    fmt.Println("three")
  }
}
