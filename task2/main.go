package task2

import (
	"errors"
	"fmt"
	"strings"
)

/**
Распаковка строки
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:

* "a4bc2d5e" => "aaaabccddddde"
* "abcd" => "abcd"
* "45" => "" (некорректная строка)

Дополнительное задание: поддержка escape - последовательности
* `qwe\4\5` => `qwe45` (*)
* `qwe\\5` => `qwe\\\\\` (*)
*/

const (
	EscapeSym = 92
	ZeroSym   = 48
	FirstSym  = 49
	NineSym   = 57
)

func UnpackString(str string) (unpacked string, err error) {
	var multiplier int
	var prev rune
	var escaped bool

	stringBuilder := strings.Builder{}
	for byteNum, sym := range str {
		if sym == ZeroSym { //if it is zero
			return "", errors.New(fmt.Sprintf("Zero symbol at byte %d", byteNum))
		}
		//between 1 and 9
		if sym > FirstSym && sym <= NineSym && !escaped {
			if prev > FirstSym && prev <= NineSym {
				return "", errors.New(fmt.Sprintf("Two numbers in a row at byte %d", byteNum))
			}
			multiplier = int(sym) - ZeroSym
			for ; multiplier > 1; multiplier-- {
				stringBuilder.WriteRune(prev)
			}
		} else if !escaped && sym == EscapeSym {
			escaped = true
		} else {
			stringBuilder.WriteRune(sym)
			escaped = false
		}

		prev = sym
	}

	return stringBuilder.String(), nil
}
