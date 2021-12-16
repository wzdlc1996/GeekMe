package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"

	"golang.org/x/term"
)

const UpdateEveryLineOverHeight float32 = 0.5
const UpdateTimeInMillisecond int = 20
const emptyRate int = 10
const rowWid int = 2

func main() {

	color.Set(color.FgGreen, color.Bold)
	defer color.Unset() // Use it in your function

	ind := 0
	rel := []int{1, 3, 2, -2, -3, -1, 3, 4, -4, -2, -1}
	relind := 0

	cl, linupd := StateInit()

	for {
		ind++
		if ind%linupd == 0 {
			cl.NewTransi(rel[relind])
			relind = (relind + 1) % len(rel)
		}
		cl.ResetChars()
		fmt.Print(cl.ToString())
		time.Sleep(time.Millisecond * time.Duration(UpdateTimeInMillisecond))
	}
}

// StateInit returns the CharLine instance and linupd for update
func StateInit() (cl CharLine, linupd int) {
	wid, hei, _ := term.GetSize(0)
	linupd = int(float32(hei) * UpdateEveryLineOverHeight)
	cl = GenRandCharLine(wid, emptyRate, rowWid)
	return cl, linupd
}
