class Round {
  constructor(bottles) {
    this.bottles = bottles
    this.onWall = 'on the wall'
    this.beer = 'beer'
    this.buyNewBeer = 'Go to the store and buy some more'
  }

  pluralizedBottleForm() {
    return this.lastBeer() ? 'bottle' : 'bottles'
  }

  allOut() {
    return this.bottles === 0
  }

  lastBeer() {
    return this.bottles === 1
  }

  anglicizedBottleCount() {
    return {
      capitalized: this.allOut() ? 'No more' : this.bottles,
      nonCapitalized: this.allOut() ? 'no more': this.bottles
    }
  }

  bottleOfBeers() {
    return {
      capitalized: `${this.anglicizedBottleCount().capitalized} ${this.pluralizedBottleForm()} of ${this.beer}`,
      nonCapitalized: `${this.anglicizedBottleCount().nonCapitalized} ${this.pluralizedBottleForm()} of ${this.beer}`
    }
  }

  goToStoreOrTakeOneDown() {
    if (this.allOut()) {
      this.bottles = 99
      return this.buyNewBeer
    } else {
      let lyric = this.drinkBeer()
      this.bottles--
      return lyric
    }
  }

  drinkBeer() {
    return `Take ${this.itOrOne()} down and pass it around`
  }

  itOrOne() {
    return this.lastBeer() ? 'it' : 'one'
  }

  toString() {
    return this.challenge() + this.response()
  }
  challenge() {
    return this.bottleOfBeers().capitalized + ' ' + this.onWall + ', ' + this.bottleOfBeers().nonCapitalized + '.\n'
  }
  response() {
    return this.goToStoreOrTakeOneDown() + ', ' + this.bottleOfBeers().nonCapitalized + ' ' + this.onWall + '.\n'
  }
}







module.exports = class Bottle {

  constructor() {
  }

  pluralize(n) {
    return n === 1 ? 'bottle' : 'bottles'
  }

  verse(n) {
    this.bottles = new Round(n)
    return this.bottles.toString()
  }

  verses(start, end) {
    const verses = []
    for (let i = start; i >= end; i--) {
      verses.push(this.verse(i))
    }
    return verses.join('\n')
  }

  song() {
    return this.verses(99, 0)
  }
}