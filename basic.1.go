package main
import "fmt"
func main(){
  fmt.Println(max(3, 5))
}

func max(num1,num2 int) int {
  var result int
  if(num1 > num2){
    result = num1
  } else {
    result = num2
  }
  return result
}
