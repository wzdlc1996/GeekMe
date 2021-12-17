package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

type CharLine struct {
	len int          // len is the length of one CharLine, should be the same as the window width
	arr []string     // arr is the body of CharLine, stores all available characters
	ept map[int]bool // ept is the state of CharLine, denotes whether a slot is empty
	dur []int        // dur is the time of rows of CharLine, if it gets to zero one row ends
	mid int
}

func NewCharLine(len, mid int) CharLine {
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
	cl.dur = make([]int, cl.len)
	cl.mid = mid
	return *cl
}

func (ch *CharLine) Update(addn int, dur []int) {
	emptySlot, eptn := ch.genEmptySlot()
	samp := RandSample(emptySlot, intmin(eptn, addn))
	ch.addAt(samp, dur)
	for i := 0; i < ch.len; i++ {
		ch.updateAt(i)
	}
}

// ToString converts a CharLine instance into string
func (cl CharLine) ToString() string {
	str := make([]string, cl.len)
	ind := 0
	for i := range str {
		if cl.ept[i] {
			str[i] = " "
		} else {
			str[i] = cl.rendChar(ind)
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

func (cl CharLine) rendChar(ind int) string {

	return cl.arr[ind]
}

func (ch *CharLine) addAt(slots, dur []int) {
	for i := range slots {
		slt := slots[i]
		ch.ept[slt] = false
		ch.dur[slt] = dur[i]
	}
}

func (ch *CharLine) updateAt(pos int) {
	if ch.dur[pos] <= 0 {
		ch.ept[pos] = true
		ch.dur[pos] = 0
		return
	}
	ch.dur[pos]--
}

func (ch CharLine) genEmptySlot() (emptySlot []int, eptn int) {
	emptySlot = make([]int, ch.len)
	eptn = 0
	for i, empty := range ch.ept {
		if empty {
			emptySlot[eptn] = i
			eptn++
		}
	}
	emptySlot = emptySlot[:eptn]
	return emptySlot, eptn
}

// RandSample returns a random sample from slice source with length n,
// n should be not greater than the length of source
func RandSample(source []int, n int) []int {
	if n > len(source) {
		fmt.Println("Error sampling setup")
		panic(n)
	}
	if n == len(source) {
		return source
	}
	if n == 0 {
		return []int{}
	}
	choice := rand.Intn(len(source))
	sub := append(source[:choice], source[choice+1:]...)
	return append(RandSample(sub, n-1), source[choice])
}

// GenRandomChar returns a string randomly, with length 1 as a single character, in ascii [33, 127)
func GenRandomChar() (char string) {
	return string(rune(33 + rand.Intn(127-33)))
}

func GenDur(min, rang int) int {
	return rand.Intn(rang) + min
}

func GenDurList(min, rang, len int) []int {
	res := make([]int, len)
	for i := range res {
		res[i] = GenDur(min, rang)
	}
	return res
}

func intmin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func IsContinue(arr []int) bool {
	res := true
	for i := range arr {
		res = res && ((arr[i] - arr[0]) == i)
	}
	return res
}

func GetAllContinuousSubList(arr []int, n int) (nlis []int) {
	nlis = make([]int, 0)
	for i := 0; i < len(arr)-n+1; i++ {
		if IsContinue(arr[i : i+n]) {
			nlis = append(nlis, i)
		}
	}
	return nlis
}

func GetRandomContinuousSubList(arr []int, n int) (indl, indr int, ok bool) {
	allind := GetAllContinuousSubList(arr, n)
	if len(allind) == 0 {
		return 0, 0, false
	}
	res := rand.Intn(len(allind))
	return allind[res], allind[res] + n - 1, true
}

func SwapSubList(sender, reciever []int, indl, indr int) (nsen, nrec []int) {
	nrec = append(reciever, sender[indl:indr+1]...)
	nsen = append(sender[:indl], sender[indr+1:]...)
	sort.Ints(nsen)
	sort.Ints(nrec)
	return nsen, nrec
}
