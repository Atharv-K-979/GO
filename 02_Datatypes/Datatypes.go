// ┌─────────────── Go Type System ────────────────┐
// │                                               │
// │  ┌──────────┐  ┌──────────┐  ┌──────────────┐ │
// │  │  Basic   │  │Composite │  │  Reference   │ │
// │  ├──────────┤  ├──────────┤  ├──────────────┤ │
// │  │ bool     │  │ array    │  │ slice        │ │
// │  │ int/uint │  │ struct   │  │ map          │ │
// │  │ float    │  │          │  │ channel      │ │
// │  │ complex  │  │          │  │ pointer      │ │
// │  │ string   │  │          │  │ function     │ │
// │  │ byte     │  │          │  │ interface    │ │
// │  │ rune     │  │          │  │              │ │
// │  └──────────┘  └──────────┘  └──────────────┘ │
// └───────────────────────────────────────────────┘

// Analogy: Every variable in Go gets a default gift — its zero value. No uninitialized garbage!
// Key Insight: Go has NO null — it has nil, and nil only applies to pointers, slices, maps, channels, functions, and interfaces.
// Key Insight: Go has NO ternary operator (? :). Use if-else instead. &^ is the bit clear (AND NOT) operator — unique to Go.





package main

import "fmt"

func main() {
    // 1) Integer types
    var a int = 10          // default integer type
    var b int8 = 100        // 8-bit integer
    var c int16 = 2000      // 16-bit integer
    var d int32 = 30000     // 32-bit integer
    var e int64 = 400000    // 64-bit integer
    fmt.Println("integers:", a, b, c, d, e)

    //2) Unsigned integers (only positive numbers)
    var u uint = 20
    var u8 uint8 = 255
    fmt.Println("unsigned:", u, u8)

    //3) Floating point numbers
    var f1 float32 = 3.14
    var f2 float64 = 9.87654321
    fmt.Println("floats:", f1, f2)

    //4) Boolean type
    var isOP bool = true
    fmt.Println("bool:", isOP)

    //5) String type
    var name string = "Atharv"
    fmt.Println("String:", name)

    //6) Byte (alias for uint8)
    var letter byte = 'A'
    fmt.Println("Byte:", letter)

    //7) Rune (Unicode char smae like ASCI valus)
    var r rune = 'अ'
    fmt.Println("Rune numeric:", r)        
    fmt.Printf("Rune character: %c\n", r)  

    //8) Array 
    var arr [3]int = [3]int{1, 2, 3}
    fmt.Println("Array:", arr)

}