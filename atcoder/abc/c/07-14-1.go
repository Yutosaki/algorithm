package main

import(
  "fmt"
)

func main(){
  n, q := 0, 0
  fmt.Scan(&n, &q)

  a:=make([]int, n)
  for i:=0;i<n;i++{
    a[i] = i
  }

  for i:=0;i<q;i++{
    type:=0
    fmt.Scan(&type)
    if type == 1{
      p,x:=0,0
      fmt.Scan(&p, &x)
      a[p] = x
    }else if type == 2{
      p:=0
      fmt.Scan(&p)
      fmt.Println(a[p])
    }else{
      for j:=0;j<n;j++{
        tmp :=make([]int,n)
        tmp[j] = a[n-k]
      }
    }
  }
}
