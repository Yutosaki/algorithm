package main

import "fmt"

func main(){
  n:=0
  fmt.Scan(&n)

  s:=""
  for i:=0;i<n;i++{
    c,l:="",0
    fmt.Scan(&c,&l)
    if len(s) + l > 100{
      fmt.Println("Too Long")
      return
    }
    for j:=0;j<l;j++{
      s+=c
    }
  }
  fmt.Println(s)
}
