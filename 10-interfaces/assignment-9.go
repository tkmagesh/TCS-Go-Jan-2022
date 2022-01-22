package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum(10, 20))                                                 //=> 30
	fmt.Println(sum(10, 20, 30, 40))                                         //=> 100
	fmt.Println(sum(10))                                                     //=> 10
	fmt.Println(sum())                                                       //=> 0
	fmt.Println(sum(10, "20", 30, "40"))                                     //=> 100
	fmt.Println(sum(10, "20", 30, "40", "abc"))                              //=> 100
	fmt.Println(sum(10, 20, []int{30, 40}))                                  //=> 100
	fmt.Println(sum(10, 20, []interface{}{30, 40, []int{10, 20}}))           //=> 130
	fmt.Println(sum(10, 20, []interface{}{30, 40, []interface{}{10, "20"}})) //=> 130
}

func sum(data ...interface{}) int {
	result := 0
	for _, v := range data {
		switch val := v.(type) {
		case int:
			result = result + val
		case string:
			if no, err := strconv.Atoi(val); err == nil {
				result = result + no
			}
		case []interface{}:
			result = result + sum(val...)
		case []int:
			for _, no := range val {
				result = result + no
			}
		}
	}
	return result
}
