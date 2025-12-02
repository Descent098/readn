package readn

import "math"

func AutomatedReadabilityIndex(text string) int {
	textInfo := GetTextData(text)

	return int(math.Ceil(float64(4.71*(textInfo.Characters/textInfo.Words) + 0.5*(textInfo.Words/textInfo.Sentences) - 21.43)))

}
