package main

import "fmt"

func familyName(fname string) {
    fmt.Println("Hello", fname, "Refsnes")
}

func main() {
    familyName("Liam")
    familyName("Jenny")
    familyName("Anja")
}


/* When a parameter is passed to the function, it is called an argument. 
So, from the example above: fname is a parameter, while Liam, Jenny and Anja are arguments.*/