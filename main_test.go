package main

import (
	"baseify/fileop"
	"encoding/base64"
	"io/ioutil"
	"os"
	"testing"
)

func TestSimple(t *testing.T) {
	inputs := []string{
		"First",
		"Second",
		"Cool\t",
		"",
		"Wh\t\t\t\t\tat",
		`Wh   \t\t\t\
		tat`,
	}

	for _, input := range inputs {
		enc := base64.StdEncoding.EncodeToString([]byte(input))
		dec, err := base64.StdEncoding.DecodeString(enc)
		if err != nil {
			t.Errorf("expected decode to succeed: %s", err)
		}
		if input != string(dec) {
			t.Errorf("expected decoding to match input. input: %s != decoding: %s", input, string(dec))
		}
	}
}

func TestFile(t *testing.T) {
	defer func() {
		os.Remove("resources/tmp/test")
		os.Remove("resources/tmp/test_decoding")
	}()

	enc, err := fileop.Encode("resources/sample.txt")
	if err != nil {
		t.Errorf("unable to encode sample file: %s", err)
	}

	f, err := os.Create("resources/tmp/test")
	if err != nil {
		t.Errorf("Error creating file %s: %s", "resources/tmp/test", err)
	}
	_, err = f.WriteString(enc)
	if err != nil {
		t.Errorf("Error writing to file %s: %s", "resources/tmp/test", err)
	}
	err = f.Close()
	if err != nil {
		t.Errorf("Error closing file %s: %s", "resources/tmp/test", err)
	}

	b, err := ioutil.ReadFile("resources/tmp/test")
	if err != nil {
		t.Errorf("Error reading file %s: %s", "resources/tmp/test", err)
	}

	err = fileop.Decode(string(b), "resources/tmp/test_decoding")
	if err != nil {
		t.Errorf("Error decoding file %s: %s", "resources/tmp/test_decoding", err)
	}
}
