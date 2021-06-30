package isogram

// Source: exercism/problem-specifications
// Commit: 74869e8 isogram: Add test case for dupe after non-letter (fixes #1390)
// Problem Specifications Version: 1.7.0

var testCases = []struct {
	description string
	input       string
	expected    bool
}{
	{
		description: "empty string",
		input:       "",
		expected:    true,
	},
	{
		description: "isogram with only lower case characters",
		input:       "isogram",
		expected:    true,
	},
	{
		description: "word with one duplicated character",
		input:       "eleven",
		expected:    false,
	},
	{
		description: "word with one duplicated character from the end of the alphabet",
		input:       "zzyzx",
		expected:    false,
	},
	{
		description: "longest reported english isogram",
		input:       "subdermatoglyphic",
		expected:    true,
	},
	{
		description: "word with duplicated character in mixed case",
		input:       "Alphabet",
		expected:    false,
	},
	{
		description: "word with duplicated character in mixed case, lowercase first",
		input:       "alphAbet",
		expected:    false,
	},
	{
		description: "hypothetical isogrammic word with hyphen",
		input:       "thumbscrew-japingly",
		expected:    true,
	},
	{
		description: "hypothetical word with duplicated character following hyphen",
		input:       "thumbscrew-jappingly",
		expected:    false,
	},
	{
		description: "isogram with duplicated hyphen",
		input:       "six-year-old",
		expected:    true,
	},
	{
		description: "made-up name that is an isogram",
		input:       "Emily Jung Schwartzkopf",
		expected:    true,
	},
	{
		description: "duplicated character in the middle",
		input:       "accentor",
		expected:    false,
	},
	{
		description: "same first and last characters",
		input:       "angola",
		expected:    false,
	},
	{
		description: "all numbers",
		input:       "0123456789",
		expected:    true,
	},
	{
		description: "one repeating number",
		input:       "01234567890",
		expected:    false,
	},
	{
		description: "with unique unicode characters",
		input:       "àèîõü",
		expected:    true,
	},
	{
		description: "with repeating unicode characters",
		input:       "àèîõüàèîõü",
		expected:    false,
	},
	{
		description: "with unique unicode and numbers",
		input:       "0à1è2î3õ4ü5",
		expected:    true,
	},
	{
		description: "with duplicate unicode in unicode + numbers",
		input:       "0à1è2î3õ4ü5è",
		expected:    false,
	},
	{
		description: "with duplicate numbers in unicode + numbers",
		input:       "0à1è2î3õ4ü2",
		expected:    false,
	},
	{
		description: "with unique unicode in a sentence",
		input:       "ہیلو جناب",
		expected:    true,
	},
	{
		description: "with duplicate unicode in a sentence",
		input:       "ہیلو دنیا",
		expected:    false,
	},
}
