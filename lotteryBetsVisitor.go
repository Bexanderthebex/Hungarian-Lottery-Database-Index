package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type LotteryBetsVisitor struct {
	bitmap    BitMapIndex
	separator string
}

func NewLotteryBetsVisitor(bitmap BitMapIndex, separator string) *LotteryBetsVisitor {
	return &LotteryBetsVisitor{
		bitmap:    bitmap,
		separator: separator,
	}
}

func (l *LotteryBetsVisitor) Visit(lottoBet string) {
	// TODO: made the cap similar to width
	lottoPicks := make([]string, 0, 5)
	for _, s := range strings.Split(lottoBet, l.separator) {
		lottoPicks = append(lottoPicks, s)
	}

	formattedLottoPicks := make([]uint, 0, 5)
	for _, lottoPick := range lottoPicks {
		validLottoPickFormat, _ := strconv.ParseUint(lottoPick, 10, 8)
		formattedLottoPicks = append(formattedLottoPicks, uint(validLottoPickFormat))
	}

	if picksValid(formattedLottoPicks) {
		recordId := l.bitmap.GetTotalRecords()
		for _, flt := range formattedLottoPicks {
			l.bitmap.SetValue(flt, recordId, true)
		}

		l.boolMap.IncrementTotalRecords()
	} else {
		fmt.Println("Detected invalid lotto picks. Discarding")
		fmt.Println(lottoPicks)
	}
}
