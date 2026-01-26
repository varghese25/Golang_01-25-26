# Golang_01-25-26

-- Practical experience working with various programming languages (GOlang is preferred) Company Name SwiftConnect learn this language 
Learned with in Hour completed basic & excerice in W3school Really,

Go is fun and easy to learn

# How to Run Code
-- Folder Name -> PS E:\Golang_01-25-26\Loop>
-- Enter the Folder -> Execute code PS E:\Golang_01-25-26\Loop> go run lopInc.go
-- Exectuing file inside it must be main()


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


# GoLang has Package Main Each

# Golden rule (remember this forever)

One folder = one executable = one main()

<p>
Why the conflict happens

In Go:

All .go files in the same folder are compiled together

So if your folder has:
E:\Golang_01-25-26
├── lopInc.go
├── test.go


And both files contain:

package main

BUT:

One file has func main()

Another file has no func main()
OR

Both files have func main()

👉 Go gets confused and throws errors like:

function main is undeclared

main redeclared in this block

The CORRECT way to practice Go (very important)
🟢 Option 1: ONE program per folder (BEST for beginners)

This is the recommended approach while learning.


Golang_Practice
├── loop_increment
│   └── main.go
├── addition
│   └── main.go
├── condition
│   └── main.go


Each folder:

Has only one main.go

Has one package main

Has one func main()

Run inside that folder:

go run .

✅ No conflicts
✅ Clean learning

Option 2: Keep helper functions, ONE main()

If you want multiple files in one folder:


// loop.go
package main

func loopIncrement() {
    // logic
}

// main.go
package main

func main() {
    loopIncrement()
}

Same package main
✔ Only one func main()
</p>