package main

import (
	"errors"
	"fmt"
	_ "github.com/harrydb/go/img/pnm"
	"image"
	"image/color"
	"log"
	"os"
)

type Question struct {
	Question string
	Answers  []Bubble
}

type Bubble struct {
	Answer string
	Rect   image.Rectangle
}

func main() {
	infile := os.Args[1]

	file, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	blackPixels := func(r image.Rectangle) int {
		b := 0
		for y := r.Min.Y; y <= r.Max.Y; y++ {
			for x := r.Min.X; x <= r.Max.X; x++ {
				if img.At(x, y).(color.RGBA).R <= 0x7f {
					b++
				}
			}
		}
		return b
	}

	chosen := func(q Question) (Bubble, error) {
		var filled Bubble
		numFilled := 0

		for _, b := range q.Answers {
			bubblePx := blackPixels(b.Rect)
			if bubblePx >= 125 {
				numFilled++
				filled = b
			}
		}

		if numFilled == 0 {
			return Bubble{Answer: "N/A"}, nil
		}

		if numFilled > 1 {
			return Bubble{}, errors.New("Multiple filled")
		}

		return filled, nil
	}

	Scantron := []Question{
		Question{
			"Presenter made educational goals clear at the start",
			[]Bubble{
				{"Excellent", image.Rectangle{image.Point{165, 706}, image.Point{205, 728}}},
				{"Good", image.Rectangle{image.Point{241, 706}, image.Point{281, 728}}},
				{"Average", image.Rectangle{image.Point{317, 706}, image.Point{357, 728}}},
				{"Fair", image.Rectangle{image.Point{391, 706}, image.Point{431, 728}}},
				{"Poor", image.Rectangle{image.Point{467, 706}, image.Point{507, 728}}},
			},
		},

		Question{
			"Content was applicable to my practice/ministry",
			[]Bubble{
				{"Excellent", image.Rectangle{image.Point{165, 756}, image.Point{205, 778}}},
				{"Good", image.Rectangle{image.Point{241, 756}, image.Point{281, 778}}},
				{"Average", image.Rectangle{image.Point{317, 756}, image.Point{357, 778}}},
				{"Fair", image.Rectangle{image.Point{391, 756}, image.Point{431, 778}}},
				{"Poor", image.Rectangle{image.Point{467, 756}, image.Point{507, 778}}},
			},
		},

		Question{
			"Presenter is knowledgable in his/her area",
			[]Bubble{
				{"Excellent", image.Rectangle{image.Point{165, 806}, image.Point{205, 828}}},
				{"Good", image.Rectangle{image.Point{241, 806}, image.Point{281, 828}}},
				{"Average", image.Rectangle{image.Point{317, 806}, image.Point{357, 828}}},
				{"Fair", image.Rectangle{image.Point{391, 806}, image.Point{431, 828}}},
				{"Poor", image.Rectangle{image.Point{467, 806}, image.Point{507, 828}}},
			},
		},

		Question{
			"Length of session was adequate",
			[]Bubble{
				{"Excellent", image.Rectangle{image.Point{165, 856}, image.Point{205, 878}}},
				{"Good", image.Rectangle{image.Point{241, 856}, image.Point{281, 878}}},
				{"Average", image.Rectangle{image.Point{317, 856}, image.Point{357, 878}}},
				{"Fair", image.Rectangle{image.Point{391, 856}, image.Point{431, 878}}},
				{"Poor", image.Rectangle{image.Point{467, 856}, image.Point{507, 878}}},
			},
		},

		Question{
			"Presenter made efficient use of time",
			[]Bubble{
				{"Excellent", image.Rectangle{image.Point{165, 906}, image.Point{205, 928}}},
				{"Good", image.Rectangle{image.Point{241, 906}, image.Point{281, 928}}},
				{"Average", image.Rectangle{image.Point{317, 906}, image.Point{357, 928}}},
				{"Fair", image.Rectangle{image.Point{391, 906}, image.Point{431, 928}}},
				{"Poor", image.Rectangle{image.Point{467, 906}, image.Point{507, 928}}},
			},
		},

		Question{
			"Presenter knew related psychological data",
			[]Bubble{
				{"Excellent", image.Rectangle{image.Point{165, 956}, image.Point{205, 978}}},
				{"Good", image.Rectangle{image.Point{241, 956}, image.Point{281, 978}}},
				{"Average", image.Rectangle{image.Point{317, 956}, image.Point{357, 978}}},
				{"Fair", image.Rectangle{image.Point{391, 956}, image.Point{431, 978}}},
				{"Poor", image.Rectangle{image.Point{467, 956}, image.Point{507, 978}}},
			},
		},

		Question{
			"Presentation was professional and thorough",
			[]Bubble{
				{"Excellent", image.Rectangle{image.Point{165, 1006}, image.Point{205, 1028}}},
				{"Good", image.Rectangle{image.Point{241, 1006}, image.Point{281, 1028}}},
				{"Average", image.Rectangle{image.Point{317, 1006}, image.Point{357, 1028}}},
				{"Fair", image.Rectangle{image.Point{391, 1006}, image.Point{431, 1028}}},
				{"Poor", image.Rectangle{image.Point{467, 1006}, image.Point{507, 1028}}},
			},
		},

		Question{
			"Content matched program advertising",
			[]Bubble{
				{"Excellent", image.Rectangle{image.Point{165, 1056}, image.Point{205, 1078}}},
				{"Good", image.Rectangle{image.Point{241, 1056}, image.Point{281, 1078}}},
				{"Average", image.Rectangle{image.Point{317, 1056}, image.Point{357, 1078}}},
				{"Fair", image.Rectangle{image.Point{391, 1056}, image.Point{431, 1078}}},
				{"Poor", image.Rectangle{image.Point{467, 1056}, image.Point{507, 1078}}},
			},
		},

		Question{
			"Presenter was responsive to questions/comments",
			[]Bubble{
				{"Excellent", image.Rectangle{image.Point{165, 1107}, image.Point{205, 1129}}},
				{"Good", image.Rectangle{image.Point{241, 1107}, image.Point{281, 1129}}},
				{"Average", image.Rectangle{image.Point{317, 1107}, image.Point{357, 1129}}},
				{"Fair", image.Rectangle{image.Point{391, 1107}, image.Point{431, 1129}}},
				{"Poor", image.Rectangle{image.Point{467, 1107}, image.Point{507, 1129}}},
			},
		},

		Question{
			"Presenter integrated content with spiritual principles",
			[]Bubble{
				{"Excellent", image.Rectangle{image.Point{165, 1157}, image.Point{205, 1179}}},
				{"Good", image.Rectangle{image.Point{241, 1157}, image.Point{281, 1179}}},
				{"Average", image.Rectangle{image.Point{317, 1157}, image.Point{357, 1179}}},
				{"Fair", image.Rectangle{image.Point{391, 1157}, image.Point{431, 1179}}},
				{"Poor", image.Rectangle{image.Point{467, 1157}, image.Point{507, 1179}}},
			},
		},

		Question{
			"Teaching aids/audiovisuals were used effectively",
			[]Bubble{
				{"Excellent", image.Rectangle{image.Point{165, 1207}, image.Point{205, 1229}}},
				{"Good", image.Rectangle{image.Point{241, 1207}, image.Point{281, 1229}}},
				{"Average", image.Rectangle{image.Point{317, 1207}, image.Point{357, 1229}}},
				{"Fair", image.Rectangle{image.Point{391, 1207}, image.Point{431, 1229}}},
				{"Poor", image.Rectangle{image.Point{467, 1207}, image.Point{507, 1229}}},
			},
		},

		Question{
			"Content was presented clearly and effectively",
			[]Bubble{
				{"Excellent", image.Rectangle{image.Point{165, 1257}, image.Point{205, 1279}}},
				{"Good", image.Rectangle{image.Point{241, 1257}, image.Point{281, 1279}}},
				{"Average", image.Rectangle{image.Point{317, 1257}, image.Point{357, 1279}}},
				{"Fair", image.Rectangle{image.Point{391, 1257}, image.Point{431, 1279}}},
				{"Poor", image.Rectangle{image.Point{467, 1257}, image.Point{507, 1279}}},
			},
		},

		Question{
			"Presenter adequately met learning objective #1",
			[]Bubble{
				{"Excellent", image.Rectangle{image.Point{165, 1307}, image.Point{205, 1329}}},
				{"Good", image.Rectangle{image.Point{241, 1307}, image.Point{281, 1329}}},
				{"Average", image.Rectangle{image.Point{317, 1307}, image.Point{357, 1329}}},
				{"Fair", image.Rectangle{image.Point{391, 1307}, image.Point{431, 1329}}},
				{"Poor", image.Rectangle{image.Point{467, 1307}, image.Point{507, 1329}}},
			},
		},

		Question{
			"Presenter adequately met learning objective #2",
			[]Bubble{
				{"Excellent", image.Rectangle{image.Point{165, 1357}, image.Point{205, 1379}}},
				{"Good", image.Rectangle{image.Point{241, 1357}, image.Point{281, 1379}}},
				{"Average", image.Rectangle{image.Point{317, 1357}, image.Point{357, 1379}}},
				{"Fair", image.Rectangle{image.Point{391, 1357}, image.Point{431, 1379}}},
				{"Poor", image.Rectangle{image.Point{467, 1357}, image.Point{507, 1379}}},
			},
		},

		Question{
			"Presenter adequately met learning objective #3",
			[]Bubble{
				{"Excellent", image.Rectangle{image.Point{165, 1407}, image.Point{205, 1429}}},
				{"Good", image.Rectangle{image.Point{241, 1407}, image.Point{281, 1429}}},
				{"Average", image.Rectangle{image.Point{317, 1407}, image.Point{357, 1429}}},
				{"Fair", image.Rectangle{image.Point{391, 1407}, image.Point{431, 1429}}},
				{"Poor", image.Rectangle{image.Point{467, 1407}, image.Point{507, 1429}}},
			},
		},

		Question{
			"I recommend presenter for future events",
			[]Bubble{
				{"Excellent", image.Rectangle{image.Point{165, 1459}, image.Point{205, 1481}}},
				{"Good", image.Rectangle{image.Point{241, 1459}, image.Point{281, 1481}}},
				{"Average", image.Rectangle{image.Point{317, 1459}, image.Point{357, 1481}}},
				{"Fair", image.Rectangle{image.Point{391, 1459}, image.Point{431, 1481}}},
				{"Poor", image.Rectangle{image.Point{467, 1459}, image.Point{507, 1481}}},
			},
		},
	}

	for _, q := range Scantron {
		bubble, err := chosen(q)
		if err != nil {
			fmt.Printf("\"%s\": \x1b[91;1m%s\x1b[0m\n", q.Question, err)
		} else {
			fmt.Printf("\"%s\": \x1b[92;1m%s\x1b[0m\n", q.Question, bubble.Answer)
		}
	}
}
