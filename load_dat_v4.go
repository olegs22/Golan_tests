package main
import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)

type Data struct{
  X,Y []float64
}

func Count_dim(file_name string) int{
  f,err := os.Open(file_name)
  if err != nil{
    panic(err)
  }

  scanner := bufio.NewScanner(f)
  count := 0
  for scanner.Scan(){
    line := scanner.Text()
    text := strings.TrimRight(line, " ")

    if string(text[0]) != "#"{
      for j:=0;j<len(string(text));j++{
        if string(text[j]) == " " && string(text[j+1]) != " " {
          count++
        }
      }
      break
    }else if string(text[0]) == "#"{
      continue
    }
  }
return count + 1
}
func loadtxt(file_name string) *Data{
  //var err2 error
  dim := Count_dim(file_name)

  data := Data{}

  f,err := os.Open(file_name)
  if err != nil{
    panic(err)
  }
  scanner := bufio.NewScanner(f)
  for scanner.Scan(){
    line := scanner.Text()
    text := strings.TrimRight(line, " ")

    if string(text[0]) == "#"{
      continue
    }
    count :=0
    for j:=0;j<len(string(text));j++{
      if string(text[j]) == " "{
        break
      }
      count++
    }
    if dim == 1{
      val1,err2 := strconv.ParseFloat(line[:count],64)
      if err2 != nil{
        fmt.Println(err2)
        break
        }
      data.X = append(data.X,val1)
      }
    if dim == 2{
      val1,err2 := strconv.ParseFloat(line[:count],64)
      if err2 != nil{
        fmt.Println(err2)
        break
        }
      data.X = append(data.X,val1)
      val2,err2 := strconv.ParseFloat(line[count+1:],64)
      if err2 != nil{
        fmt.Println(err2)
        break
        }
      data.Y = append(data.Y,val2)
      }
    }
  return &data
}

func main()  {
  file := "random_data.dat"
  N := Count_dim(file)
  loaded_data :=loadtxt(file)
  fmt.Println(N)
  fmt.Println(loaded_data.X)
  fmt.Println(loaded_data.Y)
}
