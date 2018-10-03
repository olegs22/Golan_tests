package main
import "fmt"

func add(a float32, b float32) float32{
  return a+b
}

func main(){
  fmt.Println("The sum is" , add(5.5,18.4))
}
