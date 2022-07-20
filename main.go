package main

import "fmt"

func main() {

	t := TSlice{[]interface{}{}, 0}

	t.push(123)
	t.push(456)

	fmt.Println(t.top())

	t.pop()

	fmt.Println(t.top())
	fmt.Println(t.length())

}
