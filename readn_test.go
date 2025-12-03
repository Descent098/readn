package readn_test

import (
	"log"
	"math"
	"os"
	"testing"

	"github.com/descent098/readn"
)

func TestFleschKincaid(t *testing.T) {
	res := readn.FleschKincaid(text)

	if math.Ceil(float64(res.Ease)) != 57 {
		t.Fatalf("FleschKincaid gave an incorrect ease %.2f rounded up does not equal 57", res.Ease)
	}
	if math.Ceil(float64(res.Level)) != 14 {
		t.Fatalf("FleschKincaid gave an incorrect level %.2f rounded up does not equal 14", res.Level)
	}

	res2 := readn.FleschKincaid(paradiseLostBookOne)

	if math.Ceil(float64(res2.Ease)) != 69 {
		t.Fatalf("FleschKincaid gave an incorrect ease %.2f rounded up does not equal 69", res2.Ease)
	}
	if math.Ceil(float64(res2.Level)) != 10 {
		t.Fatalf("FleschKincaid gave an incorrect level %.2f rounded up does not equal 10", res2.Level)
	}

	res3 := readn.FleschKincaid(crimeAndPunishment)
	if math.Ceil(float64(res3.Ease)) != 75 {
		t.Fatalf("FleschKincaid gave an incorrect ease %.2f rounded up does not equal 75", res3.Ease)
	}
	if math.Ceil(float64(res3.Level)) != 7 {
		t.Fatalf("FleschKincaid gave an incorrect level %.2f rounded up does not equal 7", res3.Level)
	}
}

func TestSMOG(t *testing.T) {
	res, err := readn.SimpleMeasureOfGobbledygook(text)

	if err == nil || res != 0 {
		t.Fatalf("SimpleMeasureOfGobbledygook() failed, should have errored because text is too small")
	}

	res2, err := readn.SimpleMeasureOfGobbledygook(paradiseLostBookOne)

	if err != nil || res2 == 0 {
		t.Fatalf("SimpleMeasureOfGobbledygook() failed, text should be SMOG compatible")
	}

	if math.Ceil(float64(res2)) != 11 {
		t.Fatalf("SimpleMeasureOfGobbledygook() gave an incorrect answer %.2f rounded up does not equal 11", res2)
	}

}

func TestARI(t *testing.T) {
	res := readn.AutomatedReadabilityIndex(text)
	if res != 16 {
		t.Fatalf("AutomatedReadabilityIndex() got an incorrect value for text %d should be 16", res)
	}

	res2 := readn.AutomatedReadabilityIndex(paradiseLostBookOne)
	if res2 != 13 {
		t.Fatalf("AutomatedReadabilityIndex() got an incorrect value for text %d should be 13", res2)
	}

	res3 := readn.AutomatedReadabilityIndex(crimeAndPunishment)
	if res3 != 6 {
		t.Fatalf("AutomatedReadabilityIndex() got an incorrect value for text %d should be 6", res3)

	}
}

func TestInfoParsing(t *testing.T) {
	res := readn.GetTextData(text)

	if res.Words != 156 {
		t.Fatalf("GetTextData() gave an incorrect number of words %.f does not equal 156", res.Words)
	}

	if res.Sentences != 5 {
		t.Fatalf("GetTextData() gave an incorrect number of sentences %.f does not equal 5", res.Sentences)
	}

	if res.PolySylabicWords != 14 {
		t.Fatalf("GetTextData() gave an incorrect number of poly-syllabic words %.f does not equal 14", res.PolySylabicWords)
	}

	if res.Characters != 701 {
		t.Fatalf("GetTextData() gave an incorrect number of characters %.f does not equal 701", res.Characters)
	}

	if res.Syllables != 218 {
		t.Fatalf("GetTextData() gave an incorrect number of syllables %.f does not equal 218", res.Syllables)
	}

	res2 := readn.GetTextData(paradiseLostBookOne)

	if res2.Words != 5996 {
		t.Fatalf("GetTextData() gave an incorrect number of words %.f does not equal 256", res2.Words)
	}

	if res2.Sentences != 241 {
		t.Fatalf("GetTextData() gave an incorrect number of sentences %.f does not equal 241", res2.Sentences)
	}

	if res2.PolySylabicWords != 374 {
		t.Fatalf("GetTextData() gave an incorrect number of poly-syllabic words %.f does not equal 374", res2.PolySylabicWords)
	}

	if res2.Characters != 26991 {
		t.Fatalf("GetTextData() gave an incorrect number of characters %.f does not equal 26991", res2.Characters)
	}

	if res2.Syllables != 7997 {
		t.Fatalf("GetTextData() gave an incorrect number of syllables %.f does not equal 7997", res2.Syllables)
	}
	res3 := readn.GetTextData(crimeAndPunishment)
	if res3.Words != 203470 {
		t.Fatalf("GetTextData() gave an incorrect number of words %.f does not equal 203470", res3.Words)
	}

	if res3.Sentences != 14393 {
		t.Fatalf("GetTextData() gave an incorrect number of sentences %.f does not equal 14393", res3.Sentences)
	}

	if res3.PolySylabicWords != 18786 {
		t.Fatalf("GetTextData() gave an incorrect number of poly-syllabic words %.f does not equal 18786", res3.PolySylabicWords)
	}

	if res3.Characters != 876136 {
		t.Fatalf("GetTextData() gave an incorrect number of characters %.f does not equal 876136", res3.Characters)
	}

	if res3.Syllables != 282999 {
		t.Fatalf("GetTextData() gave an incorrect number of syllables %.f does not equal 282999", res3.Syllables)
	}

}

func TestInfoParsingShort(t *testing.T) {
	res3 := readn.GetTextData(sampleSentence1)

	if res3.Words != 6 {
		t.Fatalf("GetTextData() gave an incorrect number of words %.f does not equal 6", res3.Words)
	}

	if res3.Sentences != 1 {
		t.Fatalf("GetTextData() gave an incorrect number of sentences %.f does not equal 1", res3.Sentences)
	}

	if res3.PolySylabicWords != 0 {
		t.Fatalf("GetTextData() gave an incorrect number of poly-syllabic words %.f does not equal 0", res3.PolySylabicWords)
	}

	if res3.Characters != 17 {
		t.Fatalf("GetTextData() gave an incorrect number of characters %.f does not equal 17", res3.Characters)
	}

	if res3.Syllables != 6 {
		t.Fatalf("GetTextData() gave an incorrect number of syllables %.f does not equal 6", res3.Syllables)
	}

	res4 := readn.GetTextData(sampleSentence2)

	if res4.Words != 16 {
		t.Fatalf("GetTextData() gave an incorrect number of words %.f does not equal 16", res4.Words)
	}

	if res4.Sentences != 1 {
		t.Fatalf("GetTextData() gave an incorrect number of sentences %.f does not equal 1", res4.Sentences)
	}

	if res4.PolySylabicWords != 0 {
		t.Fatalf("GetTextData() gave an incorrect number of poly-syllabic words %.f does not equal 0", res3.PolySylabicWords)
	}

	if res4.Characters != 68 {
		t.Fatalf("GetTextData() gave an incorrect number of characters %.f does not equal 68", res4.Characters)
	}

	if res4.Syllables != 23 {
		t.Fatalf("GetTextData() gave an incorrect number of syllables %.f does not equal 23", res4.Syllables)
	}

	res5 := readn.GetTextData(sampleSentence3)

	if res5.Words != 13 {
		t.Fatalf("GetTextData() gave an incorrect number of words %.f does not equal 13", res5.Words)
	}

	if res5.Sentences != 1 {
		t.Fatalf("GetTextData() gave an incorrect number of sentences %.f does not equal 1", res5.Sentences)
	}

	if res5.PolySylabicWords != 2 {
		t.Fatalf("GetTextData() gave an incorrect number of poly-syllabic words %.f does not equal 2", res5.PolySylabicWords)
	}

	if res5.Characters != 68 {
		t.Fatalf("GetTextData() gave an incorrect number of characters %.f does not equal 68", res5.Characters)
	}

	// It actually has 25 syllables, but 21 is as close as I can get it
	if res5.Syllables != 21 {
		t.Fatalf("GetTextData() gave an incorrect number of syllables %.f does not equal 21", res5.Syllables)
	}
}

func init() {
	data1, err := os.ReadFile("testdata/crime_and_punishment.txt")
	if err != nil {
		log.Fatalf("failed to read input file: %v", err)
	}

	data2, err := os.ReadFile("testdata/paradise_lost_book_one.txt")
	if err != nil {
		log.Fatalf("failed to read input file: %v", err)
	}
	crimeAndPunishment = string(data1)
	paradiseLostBookOne = string(data2)
}

// From https://en.wikipedia.org/wiki/Flesch%E2%80%93Kincaid_readability_tests
var sampleSentence1 = `The cat sat on the mat.`
var sampleSentence2 = `This sentence, taken as a reading passage unto itself, is being used to prove a point.`
var sampleSentence3 = `The Australian platypus is seemingly a hybrid of a mammal and reptilian creature.`

// https://www.gutenberg.org/cache/epub/2600/pg2600.txt
/*
   This eBook is for the use of anyone anywhere in the United States and most
   other parts of the world at no cost and with almost no restrictions
   whatsoever. You may copy it, give it away or re-use it under the terms
   of the Project Gutenberg License included with this eBook or online
   at www.gutenberg.org. If you
   are not located in the United States, you will have to check the laws
   of the country where you are located before using this eBook.
*/
var text = `The princess smiled. She rose with the same unchanging smile with which
she had first entered the room—the smile of a perfectly beautiful
woman. With a slight rustle of her white dress trimmed with moss
and ivy, with a gleam of white shoulders, glossy hair, and sparkling
diamonds, she passed between the men who made way for her, not looking
at any of them but smiling on all, as if graciously allowing each the
privilege of admiring her beautiful figure and shapely shoulders,
back, and bosom—which in the fashion of those days were very much
exposed—and she seemed to bring the glamour of a ballroom with her as
she moved toward Anna Pávlovna. Hélène was so lovely that not only
did she not show any trace of coquetry, but on the contrary she even
appeared shy of her unquestionable and all too victorious beauty. She
seemed to wish, but to be unable, to diminish its effect.`

// https://www.gutenberg.org/cache/epub/20/pg20.txt
/*
   This eBook is for the use of anyone anywhere in the United States and most
   other parts of the world at no cost and with almost no restrictions
   whatsoever. You may copy it, give it away or re-use it under the terms
   of the Project Gutenberg License included with this eBook or online
   at www.gutenberg.org. If you
   are not located in the United States, you will have to check the laws
   of the country where you are located before using this eBook.
*/
var paradiseLostBookOne string

// https://www.gutenberg.org/cache/epub/2554/pg2554.txt
/*
   This eBook is for the use of anyone anywhere in the United States and most
   other parts of the world at no cost and with almost no restrictions
   whatsoever. You may copy it, give it away or re-use it under the terms
   of the Project Gutenberg License included with this eBook or online
   at www.gutenberg.org. If you
   are not located in the United States, you will have to check the laws
   of the country where you are located before using this eBook.
*/
var crimeAndPunishment string
