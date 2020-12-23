# Dice

The most simple RPG dice roller in go. As simple as fuck, change my mind

## Why?

Dude, why not?

## How to use?

```go
package main

func main() {
    println(dice.MustRoll("1d6"))
    println(dice.MustRoll("d6")) // like 1d6
    println(dice.MustRoll("d100"))
    println(dice.MustRoll("d%")) // like 1d100
    println(dice.MustRoll("1d1")) // always 1 you know
    println(dice.MustRoll("1d17")) // or any you want number
    println(dice.MustRoll("2d6")) // roll twice, will got [2,12] inclusive
}
```

## TODO maybe?

Honestly, i don't give a fuck to this package, but i'll approve any PR that you submit

My wishes:

- [ ] Support "1d6+2" evaluations
- [ ] best-of, worst-of (like "b3d6" and "w5d10" or SLT)
- [ ] ehmmm, multiple dices? like "d%-5d6"

## Contirbutions

Come on, create PR!

## License

You fucking kidding? This package is one structure, two funcs and 5 test cases! Don't mind, want patent it — go on, corpo freak