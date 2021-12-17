package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"

	"golang.org/x/term"
)

const UpdateEveryLineOverHeight float32 = 0.5
const UpdateTimeInMillisecond int = 20
const emptyRate int = 10
const addnMax int = 3

func main() {

	color.Set(color.FgGreen, color.Bold)
	defer color.Unset() // Use it in your function

	ind := 0

	cl, linupd := StateInit()

	for {
		ind++
		if ind%linupd == 0 {
			cl.Update(rand.Intn(addnMax), genDurs())
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
	cl = NewCharLine(wid, int(hei/2))
	cl.Update(int(wid/emptyRate), genDurs())
	return cl, linupd
}

func genDurs() []int {
	wid, hei, _ := term.GetSize(0)
	return GenDurList(int(hei/2), int(hei/2), wid)
}
