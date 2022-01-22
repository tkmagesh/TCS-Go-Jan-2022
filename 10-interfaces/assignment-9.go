package main

import (
	"fmt"
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

func sum( /*  */ ) int {

}
