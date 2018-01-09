package parser_test

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/achew22/acceptance-testing-vision-upload-server/parser"
)

const zooPath = "../testdata"

func TestParseZoo(t *testing.T) {
	fi, err := ioutil.ReadDir(zooPath)
	if err != nil {
		t.Error(err)
	}

	for _, f := range fi {
		name := f.Name()
		if !strings.HasSuffix(name, ".in") {
			continue
		}
		prefix := name[:len(name)-3]

		t.Run(name, func(t *testing.T) {
			in, err := ioutil.ReadFile(filepath.Join(zooPath, name))
			if err != nil {
				t.Error(err)
			}

			var out, suffix string
			parsedOutput, err := parser.Parse(in)
			if err != nil {
				suffix = ".err"

				// If an error, stringify it and check that against the expected out.
				out = err.Error()
			} else {
				suffix = ".out"

				for _, v := range parsedOutput {
					out += v.String() + "\n"
				}
			}

			expectedOut, err := ioutil.ReadFile(filepath.Join(zooPath, prefix+suffix))
			if err != nil {
				t.Error(err)
			}

			if out != string(expectedOut) {
				t.Errorf("\nGot:  \"%s\"\nWant: \"%s\"", out, expectedOut)
			}
		})
	}
}
