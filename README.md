# Golang_01-25-26

Go is fun and easy to learn

# Type Inference in Variable Declarations

-- The most common use of type inference in Go is with the short variable declaration operator :=

# Explicit declaration:

var i int = 42 <br>
var f float64 = 3.14<br>
var s string = "hello"<br>

# Inferred declaration (using :=):

i := 42 // The compiler infers 'i' is an int<br>
f := 3.14 // The compiler infers 'f' is a float64<br>
s := "hello" // The compiler infers 's' is a <br>

<p>The compiler looks at the expression on the right side of the := and uses its type for the new variable on the left. The var name = expression syntax also works similarly for type inference.</p>
