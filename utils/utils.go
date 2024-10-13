package utils

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
)

func Bar(iv, total int64) {
	mark := "█"
	bar := ""
	if total == 0 {
		return
	}
	percent := float32(iv*10000/total) / 100
	for i := 0; i < int(percent); i += 2 {
		bar += mark
	}
	fmt.Printf("\r[%-50s]%.2f%%  %8d/%d %s/%s", bar, percent, iv, total, humanize.Bytes(uint64(iv)), humanize.Bytes(uint64(total)))
}

func CloseBar() {
	fmt.Printf("\n")
}

var (
	Blue   = color.New(color.FgHiWhite, color.BgHiBlue).SprintFunc()
	Green  = color.New(color.FgHiWhite, color.BgHiGreen).SprintFunc()
	Red    = color.New(color.FgHiWhite, color.BgHiRed).SprintFunc()
	Yellow = color.New(color.FgHiWhite, color.BgHiYellow).SprintFunc()
)
