package raindrops

import "fmt"

const testVersion = 3

func Convert(num int) string {
	sounds := ""

	if num%3 == 0 {
		sounds += "Pling"
	}

	if num%5 == 0 {
		sounds += "Plang"
	}

	if num%7 == 0 {
		sounds += "Plong"
	}

	if len(sounds) == 0 {
		sounds = fmt.Sprintf("%d", num)
	}
	return sounds
}
