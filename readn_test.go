package readn_test

import (
	"math"
	"readn"
	"testing"
)

var text = `To be, or not to be, that is the question:
Whether 'tis nobler in the mind to suffer
The slings and arrows of outrageous fortune,
Or to take arms against a sea of troubles
And by opposing end them. To die—to sleep,
No more; and by a sleep to say we end
The heart-ache and the thousand natural shocks
That flesh is heir to: 'tis a consummation
Devoutly to be wish'd. To die, to sleep;
To sleep, perchance to dream—ay, there's the rub:
For in that sleep of death what dreams may come,
When we have shuffled off this mortal coil,
Must give us pause—there's the respect
That makes calamity of so long life.
For who would bear the whips and scorns of time,
Th'oppressor's wrong, the proud man's contumely,
The pangs of dispriz'd love, the law's delay,
The insolence of office, and the spurns
That patient merit of th'unworthy takes,
When he himself might his quietus make
With a bare bodkin? Who would fardels bear,
To grunt and sweat under a weary life,
But that the dread of something after death,
The undiscovere'd country, from whose bourn
No traveller returns, puzzles the will,
And makes us rather bear those ills we have
Than fly to others that we know not of?
Thus conscience doth make cowards of us all,
And thus the native hue of resolution
Is sicklied o'er with the pale cast of thought,
And enterprises of great pith and moment
With this regard their currents turn awry
And lose the name of action.`

func TestFleschKincaid(t *testing.T) {
	res := readn.FleschKincaid(text)

	if math.Ceil(float64(res.Ease)) != 53 {
		t.Fatalf("FleschKincaid gave an incorrect answer %.2f rounded up does not equal 53", res.Ease)
	}
	if math.Ceil(float64(res.Level)) != 17 {
		t.Fatalf("FleschKincaid gave an incorrect answer %.2f rounded up does not equal 17", res.Level)
	}
}

func TestSMOG(t *testing.T) {
	res, err := readn.SimpleMeasureOfGobbledygook(text)

	if err == nil || res != 0 {
		t.Fatalf("SimpleMeasureOfGobbledygook() failed, should have errored because text is too small")
	}

	// TODO: get valid SMOG text

}

func TestARI(t *testing.T) {
	res := readn.AutomatedReadabilityIndex(text)
	if res != 21 {
		t.Fatalf("AutomatedReadabilityIndex() got an incorrect value for text %d should be 21", res)
	}
}

func TestInfoParsing(t *testing.T) {
	res := readn.GetTextData(text)

	if res.Words != 256 {
		t.Fatalf("GetTextData() gave an incorrect number of words %.f does not equal 256", res.Words)
	}

	if res.Sentences != 6 {
		t.Fatalf("GetTextData() gave an incorrect number of sentences %.f does not equal 6", res.Sentences)
	}

	if res.PolySylabicWords != 14 {
		t.Fatalf("GetTextData() gave an incorrect number of poly-syllabic words %.f does not equal 14", res.PolySylabicWords)
	}

	if res.Characters != 1099 {
		t.Fatalf("GetTextData() gave an incorrect number of characters %.f does not equal 1150", res.Characters)
	}

	if res.Syllables != 335 {
		t.Fatalf("GetTextData() gave an incorrect number of characters %.f does not equal 335", res.Syllables)
	}

}
