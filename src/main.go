package main

import (
  "creditcard"
  "bufio"
  "fmt"
  "os"
)

func main() {
  input := bufio.NewReader(os.Stdin)

  for {
    fmt.Print("> ")

    line, err := input.ReadString('\n')
    if (err != nil) {
      break
    }

    if (len(line) > 0) {
      c := creditcard.NewCreditCard(line)
      var s string

      if c.Valid() {
        s = "✔"
      } else {
        s = "✘"
      }

      fmt.Printf("%s %s %s\n", c.Type(), c, s)
    }
  }
}