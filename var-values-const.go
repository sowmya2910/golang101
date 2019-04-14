package main
import "fmt"
import "math"

const s string = "0_o"

func main(){
  x := "Vary"

  fmt.Println("go"+"lang", 1+1, 7.0/3.0, true||false,x,s)
  const n = 500000000
  const d = 3e20 / n
  fmt.Println(int64(d))
  fmt.Println(math.Sin(n))

}
