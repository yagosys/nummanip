package calc
import (
	"errors"
)

func Add(numbers ...int) (error, int) {
	sum := 0
	if len(numbers) < 2 {
		return errors.New("need more than 2 variable"), sum
	} else {
		for _, num := range numbers {
			sum = sum + num
		}
		return nil, sum
	}
}
