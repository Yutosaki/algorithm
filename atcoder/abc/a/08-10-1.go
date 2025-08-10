package main

import "fmt"

func main(){
  N:=0
  fmt.Scan(&N)
  S:=""
  fmt.Scan(&S)

  if N < 3{
    fmt.Println("No")
    return
  }
  if S[N-1]=='a' && S[N-2]=='e' && S[N-3]=='t'{
    fmt.Println("Yes")
  }else {
    fmt.Println("No")
  }
}
