package util

import (
	"testing"
)

func TestGuessMimeType(t *testing.T) {
	expected := "image/png"
	result := GuessMimeType("examples/test.png")

	if expected != result {
		t.Error(FormatTest("GuessMimeType", result, expected))
	}

	expected = "image/png"
	result = GuessMimeType(".png")

	if expected != result {
		t.Error(FormatTest("GuessMimeType", result, expected))
	}

	expected = "application/octet-stream"
	result = GuessMimeType("png")

	if expected != result {
		t.Error(FormatTest("GuessMimeType", result, expected))
	}

	expected = "application/octet-stream"
	result = GuessMimeType("examples/test")

	if expected != result {
		t.Error(FormatTest("GuessMimeType", result, expected))
	}
}
