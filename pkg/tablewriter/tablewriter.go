package tablewriter

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"text/tabwriter"
)

const structFieldTagKey = "table"

// StructValues tab delimits the values of a given struct.
//
// Tag a field `table:"-"` to hide it from output.
// Tag a field `table:"_"` to flatten its subfields.
func StructValues(data interface{}) string {
	v := reflect.ValueOf(data)
	s := &strings.Builder{}
	for i := 0; i < v.NumField(); i++ {
		if shouldHideField(v.Type().Field(i)) {
			continue
		}
		if shouldFlatten(v.Type().Field(i)) {
			fmt.Fprintf(s, "%v", StructValues(v.Field(i).Interface()))
			continue
		}
		fmt.Fprintf(s, "%v\t", v.Field(i).Interface())
	}
	return s.String()
}

// StructFieldNames tab delimits the field names of a given struct.
//
// Tag a field `table:"-"` to hide it from output.
// Tag a field `table:"_"` to flatten its subfields.
func StructFieldNames(data interface{}) string {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	s := &strings.Builder{}
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		if shouldHideField(field) {
			continue
		}
		if shouldFlatten(field) {
			fmt.Fprintf(s, "%s", StructFieldNames(reflect.New(field.Type).Interface()))
			continue
		}
		fmt.Fprintf(s, "%s\t", fieldName(field))
	}
	return s.String()
}

// WriteTable writes the given list elements to stdout in a human readable
// tabular format. Headers abide by the `table` struct tag.
//
// `table:"-"` omits the field and no tag defaults to the Go identifier.
// `table:"_"` flattens a fields subfields.
func WriteTable(writer io.Writer, length int, each func(i int) interface{}) error {
	if length < 1 {
		return nil
	}
	w := tabwriter.NewWriter(writer, 0, 0, 4, ' ', 0)
	defer func() { _ = w.Flush() }() // Best effort.
	for ix := 0; ix < length; ix++ {
		item := each(ix)
		if ix == 0 {
			if _, err := fmt.Fprintln(w, StructFieldNames(item)); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprintln(w, StructValues(item)); err != nil {
			return err
		}
	}
	return nil
}

func fieldName(f reflect.StructField) string {
	custom, ok := f.Tag.Lookup(structFieldTagKey)
	if ok {
		return custom
	}
	return f.Name
}

func shouldFlatten(f reflect.StructField) bool {
	return f.Tag.Get(structFieldTagKey) == "_"
}

func shouldHideField(f reflect.StructField) bool {
	return f.Tag.Get(structFieldTagKey) == "-"
}
