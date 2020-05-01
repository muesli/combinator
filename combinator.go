package combinator

import (
	"fmt"
	"reflect"
)

// Generate returns a slice of all possible value combinations for any given
// struct and a set of its potential member values.
func Generate(v interface{}, ov interface{}) error {
	vType := reflect.TypeOf(v).Elem().Elem()
	vPtr := reflect.ValueOf(v)
	value := vPtr.Elem()

	ovType := reflect.TypeOf(ov)
	ovValue := reflect.ValueOf(ov)

	// calculate combinations
	combinations := 1
	members := ovType.NumField()
	for i := 0; i < members; i++ {
		if ovValue.Field(i).Len() == 0 {
			// ignore empty option values
			continue
		}

		fname := ovType.Field(i).Name
		if _, ok := vType.FieldByName(fname); !ok {
			return fmt.Errorf("can't access struct field %s", fname)
		}

		combinations *= ovValue.Field(i).Len()
	}

	for i := 0; i < combinations; i++ {
		vi := reflect.Indirect(reflect.New(vType))

		offset := 1
		for j := 0; j < members; j++ {
			ovf := ovValue.Field(j)
			if ovf.Len() == 0 {
				// ignore empty option values
				continue
			}

			fname := ovType.Field(j).Name
			fvalue := ovf.Index((i / offset) % ovf.Len())
			if vi.FieldByName(fname).CanSet() {
				vi.FieldByName(fname).Set(fvalue)
			}

			// fmt.Println(fname, fvalue, offset)
			offset *= ovf.Len()
		}

		// append item to original slice
		value.Set(reflect.Append(value, vi))
	}

	return nil
}
