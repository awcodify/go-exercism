package raindrops

import "fmt"

const testVersion = 3

func Convert(num int) string {
	rain := make([]string, 10)
	rain[3] = "Pling"
	rain[5] = "Plang"
	rain[7] = "Plong"
	sounds := ""
	for i, sound := range rain {
		if 0 < i && num%i == 0 {
			sounds += sound
		}
	}
	if len(sounds) == 0 {
		sounds = fmt.Sprintf("%d", num)
	}
	return sounds
}
