package color

import (
	"strings"
	"testing"
)

func TestConvertStrToColor(t *testing.T) {
	type test struct {
		name  string
		param string
		res   string
	}

	tests := []test{
		{
			name:  "Convert an empty string to a color of the color wheel",
			param: "",
			res:   "red-orange",
		},
		{
			name:  "Convert a short string to a color of the color wheel",
			param: "kubernetes",
			res:   "red-violet",
		},
		{
			name:  "Convert a very long string with numbers and dashes to a color of the color wheel",
			param: "this-is-a-very-very-very-very-very-very-long-t3st",
			res:   "blue",
		},
	}

	err := 0
	for _, test := range tests {
		if result := ConvertStrToColor(test.param); result != test.res {
			t.Errorf("Result %v not equal to the expected result %v\nTest: %s\nParameter: %v",
				result, test.res, strings.ToLower(test.name), test.param)
			err++
		}
	}
	t.Logf("%d%% tests passed with success (%d/%d)", (len(tests)-err)*100/len(tests), len(tests)-err, len(tests))
}
