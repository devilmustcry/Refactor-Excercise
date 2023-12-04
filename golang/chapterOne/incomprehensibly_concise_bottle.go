package chapterone

import (
	"fmt"
	"strings"
)

type IncomprehensiblyConcise struct {
}

func (IncomprehensiblyConcise) Verse(number int) string {
	capBottle := fmt.Sprintf("%d", number)
	nonCapBottle := fmt.Sprintf("%d", number)
	pluralized := "s"
	take := ""
	left := fmt.Sprintf("%d", number-1)
	leftPluralized := "s"
	if number == 0 {
		capBottle = "No more"
		nonCapBottle = "no more"
	}
	if number == 1 {
		pluralized = ""
	}
	if number > 0 {
		start := "Take"
		middle := "it"
		if number > 1 {
			middle = "one"
		}
		end := "down and pass it around"
		take = fmt.Sprintf("%s %s %s", start, middle, end)
	} else {
		take = "Go to the store and buy some more"
	}
	if number-1 == 1 {
		leftPluralized = ""
	}
	if number-1 < 0 {
		left = "99"
	} else if number-1 == 0 {
		left = "no more"
	}
	return fmt.Sprintf(
		"%s bottle%s of beer on the wall, %s bottle%s of beer.\n%s, %s bottle%s of beer on the wall.\n",
		capBottle,
		pluralized,
		nonCapBottle,
		pluralized,
		take,
		left,
		leftPluralized,
	)
}

func (s IncomprehensiblyConcise) Verses(start int, end int) string {
	verses := make([]string, 0)

	for i := start; i >= end; i-- {
		verses = append(verses, s.Verse(i))
	}
	return strings.Join(verses, "\n")
}

func (s IncomprehensiblyConcise) Song() string {
	return s.Verses(99, 0)
}
