package gpx

import (
	"bytes"
	"encoding/xml"
	"testing"
)

// TestCDATA tests the CDATA type and its XML marshaling functionality
type testStruct struct {
	XMLName xml.Name `xml:"test"`
	Text    CDATA    `xml:"text"`
}

func TestCDATA(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple text",
			input:    "Hello World",
			expected: "<test>\n  <text><![CDATA[Hello World]]></text>\n</test>",
		},
		{
			name:     "Text with newline",
			input:    "First line\nSecond line",
			expected: "<test>\n  <text><![CDATA[First line\nSecond line]]></text>\n</test>",
		},
		{
			name:     "Text with special XML characters",
			input:    "<tag> & \"",
			expected: "<test>\n  <text><![CDATA[<tag> & \"]]></text>\n</test>",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "<test>\n  <text></text>\n</test>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			encoder := xml.NewEncoder(&buf)
			encoder.Indent("", "  ")

			test := testStruct{
				Text: CDATA(tt.input),
			}

			if err := encoder.Encode(test); err != nil {
				t.Errorf("Failed to encode: %v", err)
			}

			if err := encoder.Flush(); err != nil {
				t.Errorf("Failed to flush encoder: %v", err)
			}

			actual := buf.String()
			if actual != tt.expected {
				t.Errorf("Unexpected output\nGot:\n%s\nExpected:\n%s", actual, tt.expected)
			}
		})
	}
}

// TestCDATAUnmarshal tests that CDATA can be unmarshaled correctly
func TestCDATAUnmarshal(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple CDATA",
			input:    "<test>\n  <text><![CDATA[Hello World]]></text>\n</test>",
			expected: "Hello World",
		},
		{
			name:     "CDATA with newlines",
			input:    "<test>\n  <text><![CDATA[First line\nSecond line]]></text>\n</test>",
			expected: "First line\nSecond line",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var test testStruct
			if err := xml.Unmarshal([]byte(tt.input), &test); err != nil {
				t.Errorf("Failed to unmarshal: %v", err)
			}

			if string(test.Text) != tt.expected {
				t.Errorf("Unexpected value: got %q, want %q", test.Text, tt.expected)
			}
		})
	}
}
