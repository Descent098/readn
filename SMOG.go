package readn

import (
	"fmt"
	"math"
)

func SimpleMeasureOfGobbledygook(text string) (float64, error) {
	textInfo := GetTextData(text)
	if textInfo.Sentences < 30 {
		return 0, fmt.Errorf("Cannot calculate SMOG result on text with less than 30 sentences")
	}

	return (1.043 * math.Sqrt(float64(textInfo.PolySylabicWords)*float64(30/textInfo.Sentences))) + 3.1291, nil
}
