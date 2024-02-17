# gojp 

simple json parser in golang 

## Quick start
``` Golang 
package main

import (
    "github.com/2asm/gojp"
    "fmt"
    "log"
)

func main() {
    sample_json := `
        {
            "key": "String",
            "Number": 1,
            "array": [1,2,3],	
            "nested": {
                "array2": [true, false, null, 234, "String"]
            }	
        }
    `
    l := gojp.Lexer{
        Str:  []rune(sample_json),
        Cur:  0,
        Size: len([]rune(sample_json)),
    }

    tokens, err := l.Lex()
    if err != nil {
        log.Fatal(err)
    }

    p := gojp.Parser{
        Tokens: tokens,
    }

    out, err := p.Parse()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(out)
}
```

Output
``` Console
map[Number:1 array:[1 2 3] key:String nested:map[array2:[true false <nil> 234 String]]]

[Process exited 0]
```
