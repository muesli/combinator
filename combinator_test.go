package combinator

import (
	"testing"
)

func TestCombinator(t *testing.T) {
	type Data struct {
		Color   string
		Number  int
		Enabled bool

		// DataTests ignore this field
		Untouched bool
	}

	type DataTests struct {
		Color   []string
		Number  []int
		Enabled []bool

		// DataTests ignore this field
		Untouched []bool
	}

	td := DataTests{
		Color:   []string{"red", "green", "blue"},
		Number:  []int{0, 1},
		Enabled: []bool{false, true},

		// DataTests ignore this field
		Untouched: []bool{},
	}
	tdl := len(td.Color) * len(td.Number) * len(td.Enabled)

	// initialize check matrix
	matrix := map[string]map[int]map[bool]bool{}
	for ci := 0; ci < len(td.Color); ci++ {
		matrix[td.Color[ci]] = map[int]map[bool]bool{}

		for ni := 0; ni < len(td.Number); ni++ {
			matrix[td.Color[ci]][td.Number[ni]] = map[bool]bool{}

			for bi := 0; bi < len(td.Enabled); bi++ {
				matrix[td.Color[ci]][td.Number[ni]][td.Enabled[bi]] = false
			}
		}
	}

	var data []Data
	err := Generate(&data, td)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(data) != tdl {
		t.Errorf("expected %d permutations, got %d", tdl, len(data))
	}
	for _, v := range data {
		// fmt.Println(i, v)

		cmatrix, ok := matrix[v.Color]
		if !ok {
			t.Errorf("unexpected value %s", v.Color)
			continue
		}
		nmatrix, ok := cmatrix[v.Number]
		if !ok {
			t.Errorf("unexpected value %s %d", v.Color, v.Number)
			continue
		}
		_, ok = nmatrix[v.Enabled]
		if !ok {
			t.Errorf("unexpected value %s %d %v", v.Color, v.Number, v.Enabled)
			continue
		}

		// flag combination as found
		matrix[v.Color][v.Number][v.Enabled] = true
	}

	// check all combinations have been found
	for ck, cv := range matrix {
		for nk, nv := range cv {
			for bk, bv := range nv {
				if !bv {
					t.Errorf("combination not found: %s %d %v", ck, nk, bk)
				}
			}
		}
	}
}

func TestConstValue(t *testing.T) {
	type Data struct {
		Number  int
		Enabled bool
	}

	type DataTests struct {
		Number  []int
		Enabled bool
	}

	td := DataTests{
		Number:  []int{0, 1},
		Enabled: true,
	}

	var data []Data
	err := Generate(&data, td)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for _, v := range data {
		if !v.Enabled {
			t.Fatalf("expected enabled to be set to true")
		}
	}
}

func TestUnmatchedField(t *testing.T) {
	type Data struct {
		Number int
	}

	type DataTests struct {
		Number []int

		// Data does not actually contain this field
		Unmatched []bool
	}

	td := DataTests{
		Number: []int{0, 1},

		// Data does not actually contain this field
		Unmatched: []bool{false},
	}

	var data []Data
	err := Generate(&data, td)
	if err == nil {
		t.Errorf("expected error for unmatched fields in data, got nil")
	}
}

func TestInvalidType(t *testing.T) {
	td := struct {
		Number []int
	}{}

	var s string
	err := Generate(s, td)
	if err == nil {
		t.Errorf("expected error for non-pointer type, got nil")
	}

	err = Generate(&s, td)
	if err == nil {
		t.Errorf("expected error for invalid type, got nil")
	}

	err = Generate(&[]string{}, "")
	if err == nil {
		t.Errorf("expected error for invalid type, got nil")
	}
}
