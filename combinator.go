package combinator

import (
	"reflect"
)

// Generate returns a slice of all possible value combinations for any given
// struct and a set of its potential member values.
func Generate(v interface{}, ov interface{}) []interface{} {
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
		if _, ok := reflect.TypeOf(v).FieldByName(fname); !ok {
			// fmt.Println("can't access field", fname)
			continue
		}

		combinations *= ovValue.Field(i).Len()
	}

	var r []interface{}
	for i := 0; i < combinations; i++ {
		vi := reflect.Indirect(reflect.New(reflect.TypeOf(v)))

		offset := 1
		for j := 0; j < members; j++ {
			if ovValue.Field(j).Len() == 0 {
				// ignore empty option values
				continue
			}

			fname := ovType.Field(j).Name
			fvalue := ovValue.Field(j).Index((i / offset) % ovValue.Field(j).Len())
			if vi.FieldByName(fname).CanSet() {
				vi.FieldByName(fname).Set(fvalue)
			}

			// fmt.Println(fname, fvalue, offset)
			offset *= ovValue.Field(j).Len()
		}

		r = append(r, vi.Interface())
	}

	return r
}
