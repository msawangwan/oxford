package oxford_test

import (
	"testing"

	"github.com/msawangwan/oxford"
)

const configPath = "../config.json"

func TestDoesWordExist(t *testing.T) {
	t.Run(
		"[some]",
		testDoesWordExist(
			[]string{"some"},
			[]bool{true}))
	t.Run(
		"[some, asddg]",
		testDoesWordExist(
			[]string{"some", "asddg"},
			[]bool{true, false}))
	t.Run(
		"[some, more, abcde, bible, constant]",
		testDoesWordExist(
			[]string{"some", "more", "abcde", "bible", "constant"},
			[]bool{true, true, false, true, true}))
}

func testDoesWordExist(cases []string, expected []bool) func(t *testing.T) {
	return func(t *testing.T) {
		ox, err := oxford.New(configPath, oxford.EN)
		if err != nil {
			t.Error(err)
		}

		for i, c := range cases {
			exists, err := ox.Exists(c)
			if err != nil {
				if err == oxford.ErrInvalidWord {
					t.Logf("word is invalid: %s", c)
					continue
				}

				t.Error(err)
				continue
			}

			if exists != expected[i] {
				t.Errorf("test should be %t, got %t: %s", expected[i], exists, c)
				continue
			}

			t.Logf("word is valid and exists in dict: %s", c)
		}
	}
}
