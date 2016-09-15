// Units converts its numeric argument to various related units
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

func main() {
	if len(os.Args) == 1 {
		// no arguments; process lines of stdin instead
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() && input.Text() != "" {
			showConversions(input.Text())
		}
	}

	for _, arg := range os.Args[1:] {
		showConversions(arg)
	}
}

func showConversions(arg string) {
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		return
	}

	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))

	ft := Feet(t)
	m := Meters(t)
	fmt.Printf("%s = %s, %s = %s\n", ft, FeetToMeters(ft), m, MetersToFeet(m))
}
