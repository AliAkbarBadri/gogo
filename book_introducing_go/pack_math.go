package main

import (
	"fmt"
	myMath "gogo/book_introducing_go/math"
)

func main() {
	fmt.Println(myMath.Average([]float64{1, 2, 3, 4}))
	fmt.Println(myMath.Average([]float64{}))
}
