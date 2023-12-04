module.exports = class Bottle {

  constructor() {

  }

  pluralize(n) {
    return n === 1 ? 'bottle' : 'bottles'
  }

  verse(n) {
    return `${n === 0 ? 'No more': n} bottle${n!==1 ? 's' : ''} of beer on the wall, ${n === 0 ? 'no more' : n} bottle${n !== 1 ? 's' : ''} of beer.\n${n > 0 ? `Take ${n > 1 ? 'one' : 'it'} down and pass it around` : 'Go to the store and buy some more'}, ${n-1 < 0 ? 99 : n-1 === 0 ? 'no more' : n-1} bottle${n-1 !== 1 ? 's' : ''} of beer on the wall.\n`
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