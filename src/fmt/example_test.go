// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fmt_test

import (
	"fmt"
	"os"
)

// The Errorf function lets us use formatting features
// to create descriptive error messages.
func ExampleErrorf() {
	const name, id = "bueller", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err.Error())
	// Output: user "bueller" (id 17) not found
}

func ExampleSprintf() {
	i := 30
	s := "Aug"
	sf := fmt.Sprintf("Today is %d %s", i, s)
	fmt.Println(sf)
	fmt.Println(len(sf))
	// Output:
	// Today is 30 Aug
	// 15
}

func ExamplePrintln() {
	n, err := fmt.Println("there", "are", 99, "gophers")
	if err != nil {
		panic(err)
	}
	fmt.Print(n)
	// Output:
	// there are 99 gophers
	// 21
}

func ExampleSprintln() {
	s := "Aug"
	sl := fmt.Sprintln("Today is 30", s)
	fmt.Printf("%q", sl)
	// Output:
	// "Today is 30 Aug\n"
}

func ExampleFprintln() {
	n, err := fmt.Fprintln(os.Stdout, "there", "are", 99, "gophers")
	if err != nil {
		panic(err)
	}
	fmt.Print(n)
	// Output:
	// there are 99 gophers
	// 21
}
