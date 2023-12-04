package chapterone

import (
	"fmt"
	"strings"
)

type ShamelessGreenBottle struct {
}

func (ShamelessGreenBottle) Verse(number int) string {
	switch number {
	case 0:
		return "No more bottles of beer on the wall, " + "no more bottles of beer.\n" +
			"Go to the store and buy some more, " + "99 bottles of beer on the wall.\n"
	case 1:
		return "1 bottle of beer on the wall, " +
			"1 bottle of beer.\n" +
			"Take it down and pass it around, " + "no more bottles of beer on the wall.\n"
	case 2:
		return "2 bottles of beer on the wall, " +
			"2 bottles of beer.\n" +
			"Take one down and pass it around, " + "1 bottle of beer on the wall.\n"
	default:
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", number, number, number-1)
	}
}

func (s ShamelessGreenBottle) Verses(start int, end int) string {
	verses := make([]string, 0)

	for i := start; i >= end; i-- {
		verses = append(verses, s.Verse(i))
	}
	return strings.Join(verses, "\n")
}

func (s ShamelessGreenBottle) Song() string {
	return s.Verses(99, 0)
}
