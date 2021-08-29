package reversed_words

import "testing"

type reversedWordsTests struct {
	input    string
	expected string
}

var testCases = []reversedWordsTests{
	{
		input:    "hello world!",
		expected: "world! hello",
	},
	{
		input:    "yoda doesn't speak like this",
		expected: "this like speak doesn't yoda",
	},
	{
		input:    "foobar",
		expected: "foobar",
	},
	{
		input:    "kata editor",
		expected: "editor kata",
	},
	{
		input:    "row row row your boat",
		expected: "boat your row row row",
	},
}

func TestReverseWords(t *testing.T) {
	for _, test := range testCases {
		actual := ReverseWords(test.input)
		if actual != test.expected {
			t.Errorf("FAIL - actual: %s ; expected %s",
				actual, test.expected)
		} else {
			t.Logf("FAIL - actual: %s ; expected %s",
				actual, test.expected)
		}
	}
}
