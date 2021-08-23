package esslib

import (
	"fmt"
	"math/rand"
	"math"
	"time"
)

func Randint(arr[2]int)int{
	rand.Seed(time.Now().UnixNano())
	randint := rand.Intn(arr[1] - arr[0]) + arr[0]
	return randint
}

func RandItem(arr1 []string)string{
	rand.Seed(time.Now().UnixNano())
	var item = arr1[rand.Intn(len(arr1))]
	return item
}

func StrReverse(s string) string {
    rns := []rune(s)
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
        rns[i], rns[j] = rns[j], rns[i]
    }
    return string(rns)
}


func RemoveIndex(s []int, index int) []int {
    return append(s[:index], s[index+1:]...)
}

func Hypotenuse(h float64, b float64)float64{
	hsq := h*h
	bsq := b*b
	ade := bsq+hsq
	var hyp float64 = math.Pow(ade, 0.5)
	return hyp
}

func main(){
	// rant := []int{16,34,23,18,18}
	vt := Hypotenuse(3, 5)
	fmt.Println("%v", vt)
}