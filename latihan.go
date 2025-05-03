package main
import "fmt"

func main(){
	var x int
	fmt.Scan(&x)
	hasil := sumOfTheDigitsOfHarshadNumber(x)
	fmt.Print(hasil)
}
func sumOfTheDigitsOfHarshadNumber(x int) int {
    var n, i, bagi int
    n = x
    for n > 0{
        i = n%10
        bagi += i
		n = n/10
    }
    if x%bagi == 0{
        return bagi
    }else{
        return -1
    }
}
