package readn

import (
	"regexp"
	"slices"
	"strings"
)

type textInformation struct {
	Sentences        float32
	Words            float32
	Syllables        float32
	PolySylabicWords float32
	Characters       float32
}

var (
	vowels      = []rune{'a', 'e', 'i', 'o', 'u'}
	punctuation = []rune{'!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '.', '/', ':', ';', '<', '=', '>', '?', '@', '[', '\\', ']', '^', '_', '`', '{', '|', '}', '~'}
)

// Checks if a rune is a letter
//
// # Parameters
//   - character (rune): The character to check
//
// # Returns
//
//   - bool: If the character is a letter
//
// # Notes
//
//   - Using ascii only, no unicode
func IsLetter(character rune) bool {
	if (character < 'a' || character > 'z') && (character < 'A' || character > 'Z') {
		return false
	}
	return true
}

// Cleans a word (removes any non-letter characters except hyphens)
//
// # Parameters
//   - word (string): The word to clean
//
// # Returns
//   - string; the cleaned string
//   - bool; if the word was changed
func CleanWord(word string) (string, bool) {
	changed := false
	noPunctuation := ""

	for _, punctuationMark := range punctuation {
		noPunctuation = strings.ReplaceAll(word, string(punctuationMark), "")
	}
	if noPunctuation != word {
		changed = true
	}
	resultWord := ""

	for _, character := range noPunctuation {
		if !IsLetter(character) && character != '-' {
			changed = true
			continue
		}
		resultWord += string(character)
	}
	return resultWord, changed
}

// Calculates text information about input text
//
// # Parameters
//   - text(string): the text to analyze
//
// # Returns
//   - *textInformation: The information about the text
func GetTextData(text string) *textInformation {
	result := textInformation{}
	var (
		syllables        int
		wordCount        int
		sentenceCount    int
		polySylabicWords int
		characterTotal   int
	)
	re := regexp.MustCompile(`[.?!]\s+`)
	sentences := re.Split(text, -1)

	var cleanedSentences []string
	for _, s := range sentences {
		trimmed := strings.TrimSpace(s)
		if trimmed != "" {
			cleanedSentences = append(cleanedSentences, trimmed)
			sentenceCount += 1
			// fmt.Println("Cleaned sentence: ", trimmed)
		}
	}

	words := strings.FieldsSeq(strings.Join(cleanedSentences, " "))

	for word := range words {
		res, changed := CleanWord(word)
		if changed {
			if res == "" { // skip now empty words (were punctuation)
				continue
			}
		}
		wordCount += 1
		wordSyllables := EstimateWordSyllables(res)
		syllables += wordSyllables
		if wordSyllables >= 3 {
			polySylabicWords += 1
		}
		characterTotal += len(res)

	}

	result.Sentences = float32(sentenceCount)
	result.Words = float32(wordCount)
	result.Syllables = float32(syllables)
	result.PolySylabicWords = float32(polySylabicWords)
	result.Characters = float32(characterTotal)
	return &result
}

// Estimates the number of syllables in a word
//
// # Parameters
//
//   - word(string): the word to count the syllables of
//
// # Returns
//
//   - int: The estimated number of syllables
//
// # References
//   - https://github.com/mholtzscher/syllapy provided the heuristic I'm using
func EstimateWordSyllables(word string) int {
	var (
		result int
	)
	word = strings.ToLower(strings.TrimSpace(word))
	if strings.Contains(word, "-") {
		words := strings.Split(word, "-")
		for _, word := range words {
			result += EstimateWordSyllables(word)
		}
		return result

	}
	if word == "" {
		return 0
	}

	if slices.Contains(vowels, rune(word[0])) {
		result += 1
	}

	for index, letter := range word {
		if index == 0 { // Skip first letter
			continue
		}
		// If current letter is a vowel, and previous letter is not a vowel
		if slices.Contains(vowels, letter) && !slices.Contains(vowels, rune(word[index-1])) {
			result += 1
		}
	}

	if strings.HasSuffix(word, "e") {
		result -= 1
	}

	if len(word) > 2 {
		if strings.HasSuffix(word, "le") && !slices.Contains(vowels, rune(word[len(word)-3])) {
			result += 1
		}
	}

	if result == 0 {
		// Non-empty words cannot have no syllables
		result += 1
	}

	return result
}
