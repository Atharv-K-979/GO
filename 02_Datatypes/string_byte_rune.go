s := "Hello, 世界"

// Iterate by BYTES
for i := 0; i < len(s); i++ {
    fmt.Printf("%x ", s[i])
}

// Iterate by RUNES (correct for Unicode)
for i, r := range s {
    fmt.Printf("%d:%c ", i, r)
}

// Conversions
bytes := []byte(s)        // string → byte slice
runes := []rune(s)        // string → rune slice
back  := string(runes)    // rune slice → string

// String builder (efficient concatenation)
var sb strings.Builder
sb.WriteString("Hello")
sb.WriteString(" World")
result := sb.String()