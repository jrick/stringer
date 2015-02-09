// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains simple golden tests for various examples.
// Besides validating the results when the implementation changes,
// it provides a way to look at the generated code without having
// to execute the print statements in one's head.

package main

import (
	"strings"
	"testing"
)

// Golden represents a test case.
type Golden struct {
	name   string
	input  string // input; the package clause is provided when running the test.
	output string // exected output.
}

var golden = []Golden{
	{"day", day_in, day_out},
	{"offset", offset_in, offset_out},
	{"gap", gap_in, gap_out},
	{"num", num_in, num_out},
	{"unum", unum_in, unum_out},
	{"prime", prime_in, prime_out},
}

// Each example starts with "type XXX [u]int", with a single space separating them.

// Simple test: enumeration of type int starting at 0.
const day_in = `type Day int
const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)
`

const day_out = `
const stringernamesDay = "MondayTuesdayWednesdayThursdayFridaySaturdaySunday"

var stringerindexDay = [...]uint8{0, 6, 13, 22, 30, 36, 44, 50}

// String returns the Day in a human-readable form.
func (i Day) String() string {
	if i < 0 || i+1 >= Day(len(stringerindexDay)) {
		return fmt.Sprintf("Day(%d)", i)
	}
	return stringernamesDay[stringerindexDay[i]:stringerindexDay[i+1]]
}
`

// Enumeration with an offset.
// Also includes a duplicate.
const offset_in = `type Number int
const (
	_ Number = iota
	One
	Two
	Three
	AnotherOne = One  // Duplicate; note that AnotherOne doesn't appear below.
)
`

const offset_out = `
const stringernamesNumber = "OneTwoThree"

var stringerindexNumber = [...]uint8{0, 3, 6, 11}

// String returns the Number in a human-readable form.
func (i Number) String() string {
	i -= 1
	if i < 0 || i+1 >= Number(len(stringerindexNumber)) {
		return fmt.Sprintf("Number(%d)", i+1)
	}
	return stringernamesNumber[stringerindexNumber[i]:stringerindexNumber[i+1]]
}
`

// Gaps and an offset.
const gap_in = `type Gap int
const (
	Two Gap = 2
	Three Gap = 3
	Five Gap = 5
	Six Gap = 6
	Seven Gap = 7
	Eight Gap = 8
	Nine Gap = 9
	Eleven Gap = 11
)
`

const gap_out = `
const (
	stringernamesGap0 = "TwoThree"
	stringernamesGap1 = "FiveSixSevenEightNine"
	stringernamesGap2 = "Eleven"
)

var (
	stringerindexGap0 = [...]uint8{0, 3, 8}
	stringerindexGap1 = [...]uint8{0, 4, 7, 12, 17, 21}
	stringerindexGap2 = [...]uint8{0, 6}
)

// String returns the Gap in a human-readable form.
func (i Gap) String() string {
	switch {
	case 2 <= i && i <= 3:
		i -= 2
		return stringernamesGap0[stringerindexGap0[i]:stringerindexGap0[i+1]]
	case 5 <= i && i <= 9:
		i -= 5
		return stringernamesGap1[stringerindexGap1[i]:stringerindexGap1[i+1]]
	case i == 11:
		return stringernamesGap2
	default:
		return fmt.Sprintf("Gap(%d)", i)
	}
}
`

// Signed integers spanning zero.
const num_in = `type Num int
const (
	m_2 Num = -2 + iota
	m_1
	m0
	m1
	m2
)
`

const num_out = `
const stringernamesNum = "m_2m_1m0m1m2"

var stringerindexNum = [...]uint8{0, 3, 6, 8, 10, 12}

// String returns the Num in a human-readable form.
func (i Num) String() string {
	i -= -2
	if i < 0 || i+1 >= Num(len(stringerindexNum)) {
		return fmt.Sprintf("Num(%d)", i+-2)
	}
	return stringernamesNum[stringerindexNum[i]:stringerindexNum[i+1]]
}
`

// Unsigned integers spanning zero.
const unum_in = `type Unum uint
const (
	m_2 Unum = iota + 253
	m_1
)

const (
	m0 Unum = iota
	m1
	m2
)
`

const unum_out = `
const (
	stringernamesUnum0 = "m0m1m2"
	stringernamesUnum1 = "m_2m_1"
)

var (
	stringerindexUnum0 = [...]uint8{0, 2, 4, 6}
	stringerindexUnum1 = [...]uint8{0, 3, 6}
)

// String returns the Unum in a human-readable form.
func (i Unum) String() string {
	switch {
	case 0 <= i && i <= 2:
		return stringernamesUnum0[stringerindexUnum0[i]:stringerindexUnum0[i+1]]
	case 253 <= i && i <= 254:
		i -= 253
		return stringernamesUnum1[stringerindexUnum1[i]:stringerindexUnum1[i+1]]
	default:
		return fmt.Sprintf("Unum(%d)", i)
	}
}
`

// Enough gaps to trigger a map implementation of the method.
// Also includes a duplicate to test that it doesn't cause problems
const prime_in = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const prime_out = `
const stringernamesPrime = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var stringermapPrime = map[Prime]string{
	2:  stringernamesPrime[0:2],
	3:  stringernamesPrime[2:4],
	5:  stringernamesPrime[4:6],
	7:  stringernamesPrime[6:8],
	11: stringernamesPrime[8:11],
	13: stringernamesPrime[11:14],
	17: stringernamesPrime[14:17],
	19: stringernamesPrime[17:20],
	23: stringernamesPrime[20:23],
	29: stringernamesPrime[23:26],
	31: stringernamesPrime[26:29],
	41: stringernamesPrime[29:32],
	43: stringernamesPrime[32:35],
}

// String returns the Prime in a human-readable form.
func (i Prime) String() string {
	if str, ok := stringermapPrime[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}
`

func TestGolden(t *testing.T) {
	for _, test := range golden {
		var g Generator
		input := "package test\n" + test.input
		file := test.name + ".go"
		g.parsePackage(".", []string{file}, input)
		// Extract the name and type of the constant from the first line.
		tokens := strings.SplitN(test.input, " ", 3)
		if len(tokens) != 3 {
			t.Fatalf("%s: need type declaration on first line", test.name)
		}
		g.generate(tokens[1])
		got := string(g.format())
		if got != test.output {
			t.Errorf("%s: got\n====\n%s====\nexpected\n====%s", test.name, got, test.output)
		}
	}
}
