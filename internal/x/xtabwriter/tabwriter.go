package xtabwriter

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"
)

const structFieldTagKey = "tab"

// NewWriter chooses reasonable defaults for a human readable output of tabular data.
func NewWriter() *tabwriter.Writer {
	return tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
}

// StructValues tab delimits the values of a given struct.
//
// Tag a field `tab:"-"` to hide it from output.
func StructValues(data interface{}) string {
	v := reflect.ValueOf(data)
	s := &strings.Builder{}
	for i := 0; i < v.NumField(); i++ {
		if shouldHideField(v.Type().Field(i)) {
			continue
		}
		s.WriteString(fmt.Sprintf("%v\t", v.Field(i).Interface()))
	}
	return s.String()
}

// StructFieldNames tab delimits the field names of a given struct.
//
// Tag a field `tab:"-"` to hide it from output.
func StructFieldNames(data interface{}) string {
	v := reflect.ValueOf(data)
	s := &strings.Builder{}
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		if shouldHideField(field) {
			continue
		}
		s.WriteString(fmt.Sprintf("%s\t", field.Name))
	}
	return s.String()
}

// WriteTable writes the given list elements to stdout in a human readable
// tabular format. Headers abide by the `tab` struct tag.
//
// `tab:"-"` omits the field and no tag defaults to the Go identifier.
func WriteTable(length int, each func(i int) interface{}) error {
	if length < 1 {
		return nil
	}
	w := NewWriter()
	defer w.Flush()
	for ix := 0; ix < length; ix++ {
		item := each(ix)
		if ix == 0 {
			_, err := fmt.Fprintln(w, StructFieldNames(item))
			if err != nil {
				return err
			}
		}
		_, err := fmt.Fprintln(w, StructValues(item))
		if err != nil {
			return err
		}
	}
	return nil
}

func shouldHideField(f reflect.StructField) bool {
	return f.Tag.Get(structFieldTagKey) == "-"
}
