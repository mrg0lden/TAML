//Package taml implements a Space-indentation free YAML

// because Tabs are objectively better.

package taml

import (
	"bytes"
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
	type Job struct {
		Name   string
		Salary int
	}
	type Person struct {
		Name string
		Job  Job
	}

	tests := []struct {
		name    string
		in      interface{}
		want    []byte
		wantErr bool
	}{
		{
			name: "A simple struct",
			in: Person{
				Name: "Someone",
				Job: Job{
					Name:   "Doctor",
					Salary: 100000,
				},
			},
			want: []byte("name: Someone\njob:\n\tname: Doctor\n\tsalary: 100000\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			got, err := Marshal(tt.in)
			if tt.wantErr {
				assert.NotNil(err)
				return
			}

			assert.Equal(tt.want, got)
		})
	}
}

func TestUnmarshal(t *testing.T) {
	type Job struct {
		Name   string
		Salary int
	}
	type Person struct {
		Name string
		Job  Job
	}

	type args struct {
		in  []byte
		out interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid TAML",
			args: args{[]byte("name: Someone\njob:\n\tname: Doctor\n\tsalary: 100000\n"), &Person{}},
		},
		{
			name:    "Invalid TAML",
			args:    args{[]byte("name: Someone\njob:\n    name: Doctor\n    salary: 100000\n"), &Person{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			err := Unmarshal(tt.args.in, tt.args.out)
			if tt.wantErr {
				assert.NotNil(err)
			}
		})
	}
}

func TestDecoder_Decode(t *testing.T) {
	type Job struct {
		Name   string
		Salary int
	}
	type Person struct {
		Name string
		Job  Job
	}
	tests := []struct {
		name    string
		v       interface{}
		dec     *Decoder
		wantErr bool
	}{
		{
			name: "Valid TAML",
			dec:  NewDecoder(bytes.NewBuffer([]byte("name: Someone\njob:\n\tname: Doctor\n\tsalary: 100000\n"))),
			v:    &Person{},
		},
		{
			name:    "Invalid TAML",
			dec:     NewDecoder(bytes.NewBuffer([]byte("name: Someone\njob:\n    name: Doctor\n    salary: 100000\n"))),
			v:       &Person{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			err := tt.dec.Decode(tt.v)
			if tt.wantErr {
				assert.NotNil(err)
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
