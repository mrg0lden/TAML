//Package taml implements a Space-indentation free YAML
// because Tabs are objectively better.
package taml

import (
	"bytes"
	"errors"
	"io"

	"gopkg.in/yaml.v3"
)

var ErrMixingTabsAndSpaces = errors.New("mixing tabs and spaces is prohibited")

func replaceAtIndex(s, replace []byte, i, before, after int) []byte {
	before = i - before + 1
	after += i + 1
	return append(s[:before], append(replace, s[after:]...)...)
}

func replaceSpacesWithTabs(in []byte) []byte {
	op := 0
	const (
		lookForSpaces = iota
		lookForNewLine
	)

	spaces := 0 //expect 4

	for i, char := range in {
		if op == lookForNewLine {
			if char == '\n' {
				op = lookForSpaces
			}
			continue
		}

		if char != ' ' {
			op = lookForNewLine
			continue
		}

		spaces++
		if spaces == 4 {
			in = replaceAtIndex(in, []byte("\t"), i, 4, 0)
			op = lookForSpaces
			spaces = 0
		}

	}
	return in
}

func replaceTabsWithSpaces(in []byte) ([]byte, error) {
	op := 0
	const (
		lookForTabs = iota
		lookForNewLine
	)

	iDiff := 0 //diff between i and actual i (after modifying the slice)
	// since when a slice is modified, the Go runtime doesn't care, the range
	// will keep iterating over the original slice

	for i, char := range in {
		if op == lookForNewLine {
			if char == '\n' {
				op = lookForTabs
			}
			continue
		}

		if char == ' ' {
			return nil, ErrMixingTabsAndSpaces
		}

		if char == '\t' {
			in = replaceAtIndex(in, []byte("    "), i+iDiff, 1, 0)
			iDiff += 3
			continue
		}

		op = lookForNewLine
	}
	return in, nil
}

func Marshal(in interface{}) ([]byte, error) {
	out, err := yaml.Marshal(in)
	if err != nil {
		return nil, err
	}
	return replaceSpacesWithTabs(out), nil
}

func Unmarshal(in []byte, out interface{}) (err error) {
	in, err = replaceTabsWithSpaces(in)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(in, out)
}

type Decoder struct {
	*yaml.Decoder
	r   io.Reader
	buf *bytes.Buffer
}

func NewDecoder(r io.Reader) *Decoder {
	t := &bytes.Buffer{}
	return &Decoder{yaml.NewDecoder(t), r, t}
}

func (dec *Decoder) Decode(v interface{}) error {
	b, err := io.ReadAll(dec.r)
	if err != nil {
		return err
	}
	b, err = replaceTabsWithSpaces(b)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(b)
	*dec.buf = *buf
	return dec.Decode(v)
}

type Encoder struct {
	e   *yaml.Encoder
	buf *bytes.Buffer
	w   io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	b := &bytes.Buffer{}
	return &Encoder{yaml.NewEncoder(b), b, w}
}

func (e *Encoder) Close() error {
	return e.e.Close()
}

func (e *Encoder) Encode(v interface{}) error {
	err := e.e.Encode(v)
	if err != nil {
		return err
	}

	b := replaceSpacesWithTabs(e.buf.Bytes())
	_, err = e.w.Write(b)
	return err
}

type Marshaler interface {
	MarshalTAML() (interface{}, error)
}

type Unmarshaler interface {
	UnmarshaleTAML(*Node) error
}

type Node struct {
	*yaml.Node
}
