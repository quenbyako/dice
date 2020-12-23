package dice

import (
	"errors"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/xelaj/go-dry"
)

type Dice struct {
	Rolls    int
	DiceSize int
}

var isDiceRegexp = regexp.MustCompilePOSIX(`^([0-9]*)d([0-9]+|\%)$`)

func IsDice(in string) bool {
	return isDiceRegexp.MatchString(in)
}

func (d *Dice) Parse(in string) error {
	if !IsDice(in) {
		return errors.New("not a dice")
	}

	item := strings.IndexRune(in, 'd')
	rollsStr := in[:item]
	rolls := 1
	if len(rollsStr) != 0 {
		rollsTmp, err := strconv.ParseInt(rollsStr, 10, 0)
		dry.PanicIfErr(err)
		rolls = int(rollsTmp)
	}

	diceSizeStr := in[item+1:]
	diceSize := 100
	if diceSizeStr != "%" {
		diceSizeTmp, err := strconv.ParseInt(diceSizeStr, 10, 0)
		dry.PanicIfErr(err)
		diceSize = int(diceSizeTmp)
	}

	*d = Dice{
		Rolls:    rolls,
		DiceSize: diceSize,
	}

	return nil
}

func (d *Dice) Roll() int {
	res := 0

	for i := 0; i < d.Rolls; i++ {
		rand.Seed(time.Now().UnixNano())

		res += rand.Intn(d.DiceSize) + 1
	}

	return res
}

func Roll(in string) (int, error) {
	d := &Dice{}
	err := d.Parse(in)
	if err != nil {
		return 0, err
	}

	return d.Roll(), nil
}


func MustRoll(in string) int {
	i, err := Roll(in)
	dry.PanicIfErr(err)
	return i
}