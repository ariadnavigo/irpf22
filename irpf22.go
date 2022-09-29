/* See LICENSE file for copyright and license details. */

package main

import (
	"fmt"
	"os"
	"strconv"
)

func die(str string) {
	fmt.Fprintf(os.Stderr, "%v\n", str)
	os.Exit(1)
}

func irpf(sal float32) (tax float32) {
	if sal >= 300000 {
		tax += (sal - 300000) * 0.47
		sal = 300000
	} 

	if sal >= 60000 {
		tax += (sal - 60000) * 0.45
		sal = 60000
	} 

	if sal >= 35200 {
		tax += (sal - 35200) * 0.37
		sal = 35200
	} 

	if sal >= 20200 {
		tax += (sal - 20200) * 0.30
		sal = 20200
	} 

	if sal >= 12450 {
		tax += (sal - 12450) * 0.24
		sal = 12450
	} 
	
	tax += sal * 0.19

	return tax
}

func main() {
	if len(os.Args) != 2 {
		die("usage: irpf22 num")
	}

	if sal, err := strconv.ParseFloat(os.Args[1], 32); err != nil {
		die("error: NaN")
	} else {
		fmt.Println(irpf(float32(sal)))
	}
}
