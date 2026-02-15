package main

import (
	"fmt"
)

func main() {
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}  // var a = map[string]string{...}  // ✅ allowed globally

  b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4} // b := map[string]int{...}  // ❌ error if outside function


  fmt.Printf("a\t%v\n", a)
  fmt.Printf("b\t%v\n", b)
}

/*Var a = with infer can used Insid & outside Function
b := Only used inside Function 

In Go, := is used only inside functions, it declares and initializes variables, and the type is inferred automatically. Outside functions, you must use var or const.*/