package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  lekkewords := []string{
    "dog", "cat", "fish", "giraffe",
    "moo", "spider", "lion", "apple",
    "tree", "moon", "snake", "mountain lion",
    "trooper", "burger", "nasa", "yes",
  }

  rand.Seed(time.Now().UnixNano())
  var zelength int = len(lekkewords)
  var indexnum int = rand.Intn(zelength-1)
  word := lekkewords[indexnum]

  fmt.Println("Number of words:", zelength)
  fmt.Println("Selected index number:", indexnum)
  fmt.Println("Selected word is:", word)
}
