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

// A map is an unordered and changeable collection that does not allow duplicates.
/*Each element in a map is a key:value pair.

A map is an unordered and changeable collection that does not allow duplicates.

The length of a map is the number of its elements. You can find it using the len() function.

The default value of a map is nil.

Maps hold references to an underlying hash table.

Go has multiple ways for creating maps.*/


/*Var a = with infer can used Insid & outside Function
b := Only used inside Function 

In Go, := is used only inside functions, it declares and initializes variables, and the type is inferred automatically. Outside functions, you must use var or const.*/