package main

import (
	"math/rand"
	"strings"
)

type CharLine struct {
	len  int
	arr  []string
	ept  map[int]bool
	eptn int
}

// GenRandomChar returns a string randomly, with length 1 as a single character, in ascii [33, 127)
func GenRandomChar() (char string) {
	return string(33 + rand.Intn(127-33))
}

// GenCharLine returns an instance of CharLine
func GenCharLine(len int) (lin CharLine) {
	cl := new(CharLine)
	cl.len = len
	cl.ept = make(map[int]bool)
	for i := 0; i < cl.len; i++ {
		cl.ept[i] = true
	}
	cl.arr = make([]string, len)
	for i := range cl.arr {
		cl.arr[i] = GenRandomChar()
	}
	cl.eptn = len
	return *cl
}

// GenRandCharLine returns a random instance, with occupation rate as (1 / er)
func GenRandCharLine(len, er int) (lin CharLine) {
	lin = GenCharLine(len)
	for i := 0; i < lin.len; i++ {
		if rand.Intn(er) == 0 {
			lin.ept[i] = false
			lin.eptn--
		}
	}
	lin.ResetChars()
	return lin
}

// ToString converts a CharLine instance into string
func (cl CharLine) ToString() string {
	str := make([]string, cl.len)
	ind := 0
	for i := range str {
		if cl.ept[i] {
			str[i] = " "
		} else {
			str[i] = cl.arr[ind]
			ind += 1
		}
	}
	return strings.Join(str, "")
}

// ResetChars resets the character set of cl
func (cl *CharLine) ResetChars() {
	ind := 0
	for _, isEmpty := range cl.ept {
		if !isEmpty {
			cl.arr[ind] = GenRandomChar()
			ind += 1
		}
	}
}

// slotInfo returns the occupation/empty index list
func (cl CharLine) slotInfo() (occup, empty []int) {
	empty = make([]int, cl.eptn)
	occup = make([]int, cl.len-cl.eptn)

	indEmpty := 0
	indOccup := 0
	for i, isEmpty := range cl.ept {
		if isEmpty {
			empty[indEmpty] = i
			indEmpty += 1
		} else {
			occup[indOccup] = i
			indOccup += 1
		}
	}
	return occup, empty
}

// Transi changes the cl with different row (by addn), addn < 0 will end a row
func (cl *CharLine) Transi(addn int) {
	occup, empty := cl.slotInfo()
	swplen := min3(len(occup), len(empty), intabs(addn))
	if addn < 0 {
		for i := 0; i < swplen; i++ {
			ind := rand.Intn(len(occup))
			empty = append(empty, occup[ind])
			occup = append(occup[:ind], occup[ind+1:]...)
		}
	} else {
		for i := 0; i < swplen; i++ {
			ind := rand.Intn(len(empty))
			occup = append(occup, empty[ind])
			empty = append(empty[:ind], empty[ind+1:]...)
		}
	}
	cl.ResetSlot(occup, empty)
}

// ResetSlot reset the empty infomation by occupation/empty index lists
func (cl *CharLine) ResetSlot(occup, empty []int) {
	for i := range occup {
		cl.ept[occup[i]] = false
	}
	for i := range empty {
		cl.ept[empty[i]] = true
	}
	cl.eptn = len(empty)
}

// min3 returns the minimul integer value of given three parameters
func min3(a, b, c int) int {
	z := a
	v := []int{a, b, c}
	for i := range v {
		if v[i] <= z {
			z = v[i]
		}
	}
	return z
}

// intabs returns the absolute value of the integer
func intabs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
