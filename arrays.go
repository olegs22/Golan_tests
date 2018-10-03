package main
import "fmt"

func main(){
  var a [2]string
  a[0] = "hello"
  a[1] = "world!"
  fmt.Println(a[0],a[1])
  fmt.Println(a)

  var primes = [6]int{2,3,5,7,11,13} /*:= difines a variable*/
  fmt.Println(primes)

  var s_primes []int = primes[1:4]
  fmt.Println(s_primes)

//var s_looped []int: this could be used as a conteiner to append.
  s_looped := make([]int,10,10)
  for i := 0; i<10; i++ {
    s_looped[i] = i*i
  }
  fmt.Println(s_looped)
}
