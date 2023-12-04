package chapterone

import (
	"fmt"
	"strings"
)

type Round struct {
	bottles    int
	onWall     string
	beer       string
	buyNewBeer string
}

func (r Round) pluralizedBottleForm() string {
	if r.lastBeer() {
		return "bottle"
	}
	return "bottles"
}

func (r Round) allOut() bool {
	return r.bottles == 0
}

func (r Round) lastBeer() bool {
	return r.bottles == 1
}

func (r Round) anglicizedBottleCount() map[string]string {
	capitalized := fmt.Sprintf("%d", r.bottles)
	nonCapitalized := fmt.Sprintf("%d", r.bottles)
	if r.allOut() {
		capitalized = "No more"
		nonCapitalized = "no more"
	}
	return map[string]string{
		"capitalized":    capitalized,
		"nonCapitalized": nonCapitalized,
	}
}

func (r Round) bottleOfBeers() map[string]string {
	return map[string]string{
		"capitalized":    fmt.Sprintf("%s %s of %s", r.anglicizedBottleCount()["capitalized"], r.pluralizedBottleForm(), r.beer),
		"nonCapitalized": fmt.Sprintf("%s %s of %s", r.anglicizedBottleCount()["nonCapitalized"], r.pluralizedBottleForm(), r.beer),
	}
}

func (r *Round) goToStoreOrTakeOneDown() string {
	if r.allOut() {
		r.bottles = 99
		return r.buyNewBeer
	} else {
		lyric := r.drinkBeer()
		r.bottles = r.bottles - 1
		return lyric
	}
}

func (r Round) drinkBeer() string {
	return fmt.Sprintf(`Take %s down and pass it around`, r.itOrOne())
}

func (r Round) itOrOne() string {
	if r.lastBeer() {
		return "it"
	} else {
		return "one"
	}
}

func (r Round) toString() string {
	return r.challenge() + r.response()
}

func (r Round) challenge() string {
	return fmt.Sprintf("%s %s, %s.\n", r.bottleOfBeers()["capitalized"], r.onWall, r.bottleOfBeers()["nonCapitalized"])
}

func (r Round) response() string {
	return fmt.Sprintf("%s, %s %s.\n", r.goToStoreOrTakeOneDown(), r.bottleOfBeers()["nonCapitalized"], r.onWall)
}

type ConcretelyAbstractBottle struct {
}

func (ConcretelyAbstractBottle) Verse(n int) string {
	r := Round{
		bottles:    n,
		onWall:     "on the wall",
		beer:       "beer",
		buyNewBeer: "Go to the store and buy some more",
	}
	return r.toString()
}

func (s ConcretelyAbstractBottle) Verses(start int, end int) string {
	verses := make([]string, 0)

	for i := start; i >= end; i-- {
		verses = append(verses, s.Verse(i))
	}
	return strings.Join(verses, "\n")
}

func (s ConcretelyAbstractBottle) Song() string {
	return s.Verses(99, 0)
}
