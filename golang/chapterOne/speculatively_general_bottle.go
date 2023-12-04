package chapterone

import (
	"fmt"
	"strings"
)

type Verse struct {
	number int
	lyric  func(v Verse) string
}

func (v Verse) text() string {
	return v.lyric(v)
}

type SpeculativelyGeneralBottle struct {
}

func (s SpeculativelyGeneralBottle) nomore(v Verse) string {
	return "No more bottles of beer on the wall, " + "no more bottles of beer.\n" +
		"Go to the store and buy some more, " + "99 bottles of beer on the wall.\n"
}

func (s SpeculativelyGeneralBottle) lastOne(v Verse) string {
	return "1 bottle of beer on the wall, " +
		"1 bottle of beer.\n" +
		"Take it down and pass it around, " + "no more bottles of beer on the wall.\n"
}

func (s SpeculativelyGeneralBottle) penultimate(v Verse) string {
	return "2 bottles of beer on the wall, " +
		"2 bottles of beer.\n" +
		"Take one down and pass it around, " + "1 bottle of beer on the wall.\n"
}

func (s SpeculativelyGeneralBottle) defaultVerse(v Verse) string {
	return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", v.number, v.number, v.number-1)
}

func (s SpeculativelyGeneralBottle) verseFor(number int) string {
	switch number {
	case 0:
		return Verse{number: number, lyric: s.nomore}.text()
	case 1:
		return Verse{number: number, lyric: s.lastOne}.text()
	case 2:
		return Verse{number: number, lyric: s.penultimate}.text()
	default:
		return Verse{number: number, lyric: s.defaultVerse}.text()
	}
}

func (s SpeculativelyGeneralBottle) Verse(number int) string {
	return s.verseFor(number)
}

func (s SpeculativelyGeneralBottle) Verses(start int, end int) string {
	verses := make([]string, 0)

	for i := start; i >= end; i-- {
		verses = append(verses, s.Verse(i))
	}
	return strings.Join(verses, "\n")
}

func (s SpeculativelyGeneralBottle) Song() string {
	return s.Verses(99, 0)
}
