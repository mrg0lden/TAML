//Package taml implements a Space-indentation free YAML

// because Tabs are objectively better.

package taml

import (
	"bytes"
	"io"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func mustMarshal(v interface{}) []byte {
	b, err := yaml.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}

func Test_replaceSpacesWithTabs(t *testing.T) {

	type Job struct {
		Name   string
		Salary int
	}
	type Person struct {
		Name string
		Job  Job
	}

	tests := []struct {
		name string
		in   interface{}
		want []byte
	}{
		{
			name: "Valid YAML",
			in: Person{Name: "Someone",
				Job: Job{
					Name:   "name",
					Salary: 1000,
				},
			},
			want: []byte("name: Someone\njob:\n\tname: name\n\tsalary: 1000\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			got := replaceSpacesWithTabs(mustMarshal(tt.in))
			assert.Equal(tt.want, got)
		})
	}
}

func Test_replaceTabsWithSpaces(t *testing.T) {

	tests := []struct {
		name    string
		in      []byte
		want    []byte
		wantErr bool
	}{
		{
			name: "valid input",
			in:   []byte("name: Someone\njob:\n\tname: name\n\tsalary: 1000\n"),
			want: []byte("name: Someone\njob:\n    name: name\n    salary: 1000\n"),
		},
		{
			name:    "mixing spaces and tabs",
			in:      []byte("name: Someone\njob:\n\tname: name\n    salary: 1000\n"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			got, err := replaceTabsWithSpaces(tt.in)
			if tt.wantErr {
				assert.NotNil(err)
			}

			assert.Equal(tt.want, got)
		})
	}
}

func TestMarshal(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnmarshal(t *testing.T) {
	type args struct {
		in  []byte
		out interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Unmarshal(tt.args.in, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewDecoder(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want *Decoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDecoder(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDecoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_Decode(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		dec     *Decoder
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.dec.Decode(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Decoder.Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewEncoder(t *testing.T) {
	tests := []struct {
		name  string
		want  *Encoder
		wantW string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if got := NewEncoder(w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEncoder() = %v, want %v", got, tt.want)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("NewEncoder() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestEncoder_Close(t *testing.T) {
	tests := []struct {
		name    string
		e       *Encoder
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_Encode(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		e       *Encoder
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.Encode(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.Encode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
