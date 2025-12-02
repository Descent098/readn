package readn

type FleschKincaidResult struct {
	Ease  float32
	Level float32
}

// Calculates the Flesch Kindcade information of text
//
// # Parameters
//   - text(string): The text to analyze
//
// # Returns
//   - *FleschKincaidResult: the analytics of the text
//
// # Notes
//   - Accuracy is limited by syllable estimation and sentence count estimation
//   - Typically this function will UNDERESTIMATE difficulty (lower level, higher ease)
//
// # Reference
//   - https://en.wikipedia.org/wiki/Flesch%E2%80%93Kincaid_readability_tests
//   - https://github.com/mholtzscher/syllapy Syllable Heuristic
func FleschKincaid(text string) *FleschKincaidResult {
	textInfo := GetTextData(text)

	return &FleschKincaidResult{
		Ease:  206.835 - (1.015 * (textInfo.Words / textInfo.Sentences)) - (84.6 * (textInfo.Syllables / textInfo.Words)),
		Level: 0.39*(textInfo.Words/textInfo.Sentences) + 11.8*(textInfo.Syllables/textInfo.Words) - 15.59,
	}

}
