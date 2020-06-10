package iptc

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"testing"

	"github.com/dsoprea/go-logging"
)

func TestTest_String(t *testing.T) {
	tag := Tag{
		recordNumber:  11,
		datasetNumber: 22,
		dataSize:      33,
	}

	s := tag.String()

	if s != "Tag<DATASET=(11:22) DATA-SIZE=(33)>" {
		t.Fatalf("String not correct: [%s]", s)
	}
}

func TestDecodeTag(t *testing.T) {
	filepath := GetTestDataFilepath()
	f, err := os.Open(filepath)
	log.PanicIf(err)

	defer f.Close()

	actual, err := DecodeTag(f)
	log.PanicIf(err)

	expected := Tag{
		recordNumber:  2,
		datasetNumber: 0,
		dataSize:      2,
	}

	if reflect.DeepEqual(actual, expected) != true {
		t.Fatalf("Tag not correct: (%d) (%d) (%d)", actual.recordNumber, actual.datasetNumber, actual.dataSize)
	}
}

func TestStreamTagKey_String(t *testing.T) {
	stk := StreamTagKey{
		RecordNumber:  11,
		DatasetNumber: 22,
	}

	s := stk.String()

	if s != "11:22" {
		t.Fatalf("String not correct: [%s]", s)
	}
}

func TestTagData_IsPrintable_Hit(t *testing.T) {
	tg := TagData([]byte("abc"))
	if tg.IsPrintable() != true {
		t.Fatalf("Expected printable.")
	}
}

func TestTagData_IsPrintable_Miss(t *testing.T) {
	tg := TagData([]byte{'a', 'b', 3})
	if tg.IsPrintable() != false {
		t.Fatalf("Expected non-printable.")
	}
}

func TestTagData_String_Printable(t *testing.T) {
	tg := TagData([]byte("abc"))
	if tg.String() != "abc" {
		t.Fatalf("String not correct: [%s]", tg)
	}
}

func TestTagData_String_Nonprintable(t *testing.T) {
	tg := TagData([]byte{'a', 'b', 3})
	if tg.String() != "BINARY<(3) bytes>" {
		t.Fatalf("String not correct: [%s]", tg)
	}
}

func TestParseStream(t *testing.T) {
	filepath := GetTestDataFilepath()
	f, err := os.Open(filepath)
	log.PanicIf(err)

	defer f.Close()

	tags, err := ParseStream(f)
	log.PanicIf(err)

	if len(tags) != 16 {
		t.Fatalf("Number of tags not correct: (%d)", len(tags))
	}

	stk := StreamTagKey{2, 105}
	matchingTags := tags[stk]
	data := matchingTags[0]
	actual := string(data)

	if actual != "Emma Thompson" {
		t.Fatalf("Tag value not correct: [%s]", actual)
	}
}

func TestGetSimpleDictionaryFromParsedTags(t *testing.T) {
	filepath := GetTestDataFilepath()
	f, err := os.Open(filepath)
	log.PanicIf(err)

	defer f.Close()

	pt, err := ParseStream(f)
	log.PanicIf(err)

	actual := GetSimpleDictionaryFromParsedTags(pt)

	expected := map[string]string{
		"By-line":                         "Joel Ryan",
		"By-line Title":                   "INV",
		"Caption/Abstract":                "Actress Emma Thompson arrives for the British Independent Film Awards at Old Billingsgate Market in central London, Sunday, Dec. 7, 2014. (Photo by Joel Ryan/Invision/AP)",
		"City":                            "London",
		"Country/Primary Location Name":   "GBR",
		"Credit":                          "Joel Ryan/Invision/AP",
		"Date Created":                    "20141207",
		"Headline":                        "Emma Thompson",
		"Object Name":                     "Britain Independent Film Awards",
		"Original Transmission Reference": "LENT108",
		"Source":                          "Invision",
		"Special Instructions":            "12071419889",
		"Supplemental Category":           "ENT",
		"Time Created":                    "200406+0000",
		"Urgency":                         "5 ",
	}

	if reflect.DeepEqual(actual, expected) != true {
		// Print actual.

		fmt.Printf("\n")
		fmt.Printf("ACTUAL:\n")
		fmt.Printf("\n")

		keys := make([]string, len(actual))

		i := 0
		for k := range actual {
			keys[i] = k
			i++
		}

		sort.Strings(keys)

		for _, k := range keys {
			fmt.Printf("%s: [%s]\n", k, actual[k])
		}

		// Print expected.

		keys = make([]string, len(expected))

		fmt.Printf("\n")
		fmt.Printf("EXPECTED:\n")
		fmt.Printf("\n")

		i = 0
		for k := range expected {
			keys[i] = k
			i++
		}

		sort.Strings(keys)

		for _, k := range keys {
			fmt.Printf("%s: [%s]\n", k, expected[k])
		}

		t.Fatalf("Dictionary not correct.")
	}
}

func TestGetDictionaryFromParsedTags(t *testing.T) {
	filepath := GetTestDataFilepath()
	f, err := os.Open(filepath)
	log.PanicIf(err)

	defer f.Close()

	pt, err := ParseStream(f)
	log.PanicIf(err)

	actual := GetDictionaryFromParsedTags(pt)

	expected := map[string]string{
		"By-line":       "Joel Ryan",
		"By-line Title": "INV",
		"Caption/Abstract": `Actress Emma Thompson arrives for the British Independent Film Awards at Old Billingsgate Market in central London, Sunday, Dec. 7, 2014. (Photo by Joel Ryan/Invision/AP)
`,
		"City":                            "London",
		"Country/Primary Location Name":   "GBR",
		"Credit":                          "Joel Ryan/Invision/AP",
		"Date Created":                    "20141207",
		"Headline":                        "Emma Thompson",
		"Object Name":                     "Britain Independent Film Awards",
		"Original Transmission Reference": "LENT108",
		"Source":                          "Invision",
		"Special Instructions":            "12071419889",
		"Supplemental Category":           "ENT",
		"Time Created":                    "200406+0000",
		"Urgency":                         "5 ",
		"Record Version":                  "[BINARY] 00 01",
	}

	if reflect.DeepEqual(actual, expected) != true {
		// Print actual.

		fmt.Printf("\n")
		fmt.Printf("ACTUAL:\n")
		fmt.Printf("\n")

		keys := make([]string, len(actual))

		i := 0
		for k := range actual {
			keys[i] = k
			i++
		}

		sort.Strings(keys)

		for _, k := range keys {
			fmt.Printf("%s: [%s]\n", k, actual[k])
		}

		// Print expected.

		keys = make([]string, len(expected))

		fmt.Printf("\n")
		fmt.Printf("EXPECTED:\n")
		fmt.Printf("\n")

		i = 0
		for k := range expected {
			keys[i] = k
			i++
		}

		sort.Strings(keys)

		for _, k := range keys {
			fmt.Printf("%s: [%s]\n", k, expected[k])
		}

		t.Fatalf("Dictionary not correct.")
	}
}
