package chapterone

type BottleInterface interface {
	Verse(number int) string
	Verses(start, end int) string
	Song() string
}

type Bottle struct {
}
