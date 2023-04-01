package main

import (
    "fmt"
    "github.com/krisfollegg/funtemps/conv"
)

func main() {
	f := conv.Fahrenheit(68)
	c := conv.Celsius(20)
	fmt.Printf("%s = %s, %s = %s\n", f, conv.FToC(f), c, conv.CToF(c))
}


}
