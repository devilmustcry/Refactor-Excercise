class Verse {
  constructor(number, lyric) {
    this.number = number
    this.lyric = lyric
  }

  text() {
    return this.lyric(this)
  }
}

module.exports = class Bottle {

  constructor() {
  }

  nomore(verse) {
    return "No more bottles of beer on the wall, " + "no more bottles of beer.\n" +
    "Go to the store and buy some more, " + "99 bottles of beer on the wall.\n"
  }

  lastOne(verse) {
    return "1 bottle of beer on the wall, " +
    "1 bottle of beer.\n" +
    "Take it down and pass it around, " + "no more bottles of beer on the wall.\n"
  }

  penultimate(verse) {
    return "2 bottles of beer on the wall, " +
    "2 bottles of beer.\n" +
    "Take one down and pass it around, " + "1 bottle of beer on the wall.\n"
  }

  default(verse) {
    return `${verse.number} bottles of beer on the wall, ` + `${verse.number} bottles of beer.\n` +
    `Take one down and pass it around, ` + `${verse.number - 1} bottles of beer on the wall.\n`
  }

  verseFor(number) {
    switch(number) {
      case 0:
        return new Verse(number, this.nomore).text()
      case 1:
        return new Verse(number, this.lastOne).text()
      case 2:
        return new Verse(number, this.penultimate).text()
      default:
        return new Verse(number, this.default).text()
    }
  }

  verse(n) {
    return this.verseFor(n)
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