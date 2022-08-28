/*
- Qual o resultado de fmt.Println...
    - true && true
    - true && false
    - true || true
    - true || false
    - !true
*/

package main

import "fmt"

func main() {
	a := true && true  // true
	b := true && false // false
	c := true || true  // true
	d := true || false // true
	e := !true         // false

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}
