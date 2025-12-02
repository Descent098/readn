# Readn

*A very **quick and dirty** text reading difficulty analysis library*

This library provides a few different algorithms for determining the difficulty to read of a certain text. It is not accurate enough to be used in research, it's a good **estimation** tool, nothing more. For basically every index/level that exists you want to aim for 7-10 in any non-academic work intended for adults, though this is not a hard rule. It is also **single-threaded**, meaning it's quite slow on large text. This will be addressed at some point when I get time.

## Warning

The effectiveness of this library is massively limited by it's accuracy in:

- counting sentences
- counting sylables

The error percentages go up the more the text is not just in simple plaintext format with standard sentence formatting (i.e. ending sentences with proper punctuation). Likewise there are major limitations in what types of text are possible to analyze:

- Unicode characters are completely ignored
- Non-english text with english characters will be incorrectly parsed
- Numbers are ignored
- Non-sentence text data will cause errors (i.e. math formulas in markdown text)

## Automated Readability Index (ARI)

This is the algorithm I recomend for most people. The index is basically `index-1` years of education to read. So

| Score | Age | Grade Level |
|-------|-----|-------------|
| 1 | 5-6 | Kindergarten |
| 2 | 6-7 | First Grade |
| 3 | 7-8 | Second Grade |
| 4 | 8-9 | Third Grade |
| 5 | 9-10 | Fourth Grade |
| 6 | 10-11 | Fifth Grade |
| 7 | 11-12 | Sixth Grade |
| 8 | 12-13 | Seventh Grade |
| 9 | 13-14 | Eighth Grade |
| 10 | 14-15 | Ninth Grade |
| 11 | 15-16 | Tenth Grade |
| 12 | 16-17 | Eleventh Grade |
| 13 | 17-18 | Twelfth Grade |
| 14+ | 18-22 | College/University |


ARI is the most accurate in this package because it does not rely on sylables, opting for character count, which is trivial to do. This means the error % is just dependent on sentence count accuracy. It is also designed specifically for technical documents, and works best in this context, though it also works for more general use cases.

### Usage

```go
package main

import "https://github.com/Descent098/readn"

func main(){
    text := `some text here`

    index := readn.AutomatedReadabilityIndex(text)

    fmt.Printf("Your ARI is ~%.2f years education", index)
}
```
### Formula

$4.71(\frac{\textbf{characters}}{\textbf{words}})+0.5(\frac{\textbf{words}}{sentences})-21.43$


## Simple Measure Of Gobbledygook (SMOG)

SMOG tends to be the second most accurate of the algorithms. It is designed primarily for Medical writing, but it does work outside this context. However it's important to note it **does not work on text with less than 30 sentences**. The number given is essentially the number of years of education you should have to be able to read the text.

### Usage

```go
package main

import "https://github.com/Descent098/readn"

func main(){
    text := `some text here`

    val, err := readn.SimpleMeasureOfGobbledygook(text)

    if err !=nil{
        fmt.Fatal("Text was too short for SMOG analysis")
    }

    fmt.Printf("Your SMOG score is ~%.2f years education", index)
}
```

### Formula

$grade=1.0430 \sqrt{\textbf{number of polysyllabic words} \times \frac{30}{\textbf{number of sentences}}} +3.1291$

*A polysyllabic word is defined as any word with 3 or more syllables*

## Flesch-Kincaid

This method is typically the least accurate of the three. It tends to UNDERESTIMATE (lower level, higher ease). The method returns 2 values the `FleschKincaidResult.Ease` and the `FleschKincaidResult.Level`. The `Level` is essentially the numebr of years of education to understand, the ease is the opposite, it's a score where the higher the ease, the easier the text is to read. Ease is approximately:

| Score | School Level (US) | Notes |
|-------|-------------------|-------|
| 100-90 | 5Th Grade | Very easy to read. Easily understood by an average 11-year-old student |
| 90.0-80.0 | 	6th grade	| Easy to read. Conversational English for consumers. |
| 80.0-70.0 | 	7th grade	| Fairly easy to read. |
| 70.0-60.0 | 	8th & 9th grade	| Plain English. Easily understood by 13- to 15-year-old students. |
| 60.0-50.0 | 	10th to 12th grade	| Fairly difficult to read. |
| 50.0-30.0 | 	College	| Difficult to read. |
| 30.0-10.0 | 	College graduate	| Very difficult to read. Best understood by university graduates. |
| 10.0-0.0 | 	Professional	| Extremely difficult to read. Best understood by university graduates. |

### Usage

```go
package main

import "https://github.com/Descent098/readn"

func main(){
    text := `some text here`

    res:= readn.FleschKincaid(text)

    fmt.Printf("Your Kincaid ease is ~%.2f your education index it ~%.2f years education", res.Ease, res.Level)
}
```

### Formula

$ease=206.835-1.015(\frac{\textbf{total words}}{\textbf{total sentences}})-84.6(\frac{\textbf{total syllables}}{\textbf{total words}})$

$level=0.39(\frac{\textbf{total words}}{\textbf{total sentences}})+11.8(\frac{\textbf{total syllables}}{\textbf{total words}})-15.59$

## References

- https://en.wikipedia.org/wiki/SMOG
- https://en.wikipedia.org/wiki/Automated_readability_index
- https://en.wikipedia.org/wiki/Flesch%E2%80%93Kincaid_readability_tests
