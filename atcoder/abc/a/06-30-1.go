package main

import(
  "fmt"
  "os"
  "bufio"
  "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
  sc.Scan()
  i,_ := strconv.Atoi(sc.Text())
  return i
}

func main(){
  sc.Split(bufio.ScanWords) 
  P := ""
  fmt.Scan(&P)

  max := readInt()

  if len(P) >= max{
    fmt.Println("Yes")
  }else{
    fmt.Println("No")
  }
}
