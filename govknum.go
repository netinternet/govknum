package govknum

import (
	"math"
	"strconv"
)

func Vknum(n string) bool {
	if len(n) != 10 {
		return false
	}

	sum := 0
	for i := 0; i < 9; i++ {
		strToint, _ := strconv.Atoi(string(n[i]))
		i1 := (strToint + (9 - i)) % 10
		i2 := (i1 * int(math.Pow(2, (9-float64(i))))) % 9
		if i1 != 0 && i2 == 0 {
			i2 = 9
		}
		sum += i2
	}

	lastdigit := 0
	if (sum % 10) != 0 {
		lastdigit = (10 - (sum % 10))
	}
	last := strconv.Itoa(lastdigit)

	if string(n[9]) != last {
		return false
	}
	return true
}
