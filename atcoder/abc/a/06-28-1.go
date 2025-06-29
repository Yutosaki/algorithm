package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func readInt() int {
  scanner.Scan()
  i, _ := strconv.Atoi(scanner.Text())
  return i
}

func main(){
  scanner.Split(bufio.ScanWords) 

  n := readInt() 

  counter := 0
  for i := 0; i< n; i++ {
    a := readInt()
    b := readInt()

    if a < b {
      counter++
    }
  }
  fmt.Println(counter)
}
