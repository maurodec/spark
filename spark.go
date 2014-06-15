// The MIT License (MIT)

// Copyright (c) 2014 Mauro de Carvalho

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"flag"
	"fmt"
	"github.com/maurodec/bars"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var alternateStyle bool
	flag.BoolVar(&alternateStyle, "alternate", false, "Displays the graph in an alternate style")
	flag.Parse()

	// Unfortunately we need to read all the input to be able to draw the graph.
	// We could read it line by line or in chunks too.
	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	args := strings.Fields(string(stdin))
	numbers := make([]float64, len(args))

	// Convert the given input to numbers that we can graph.
	for i := 0; i < len(args); i++ {
		parsed, err := strconv.ParseFloat(args[i], 64)
		if err != nil {
			log.Fatal(err)
		}

		numbers[i] = parsed
	}

	// Select style to display the graph.
	var style bars.BarSet
	if alternateStyle {
		style = bars.BraileBarSet
	} else {
		style = bars.NiceBarSet
	}

	// Create and display the graph.
	sparkline := bars.MakeBar(numbers, style)
	fmt.Println(string(sparkline))
}
