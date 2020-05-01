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

		// Data does not actually contain this field
		Unmatched []bool
	}

	td := DataTests{
		Color:   []string{"red", "green", "blue"},
		Number:  []int{0, 1},
		Enabled: []bool{false, true},

		// Data does not actually contain this field
		Unmatched: []bool{false},
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

	c := Generate(Data{}, td)
	if len(c) != tdl {
		t.Errorf("expected %d permutations, got %d", tdl, len(c))
	}
	for _, v := range c {
		// fmt.Println(i, v.(Data))
		vData := v.(Data)

		cmatrix, ok := matrix[vData.Color]
		if !ok {
			t.Errorf("unexpected value %s", vData.Color)
			continue
		}
		nmatrix, ok := cmatrix[vData.Number]
		if !ok {
			t.Errorf("unexpected value %s %d", vData.Color, vData.Number)
			continue
		}
		_, ok = nmatrix[vData.Enabled]
		if !ok {
			t.Errorf("unexpected value %s %d %v", vData.Color, vData.Number, vData.Enabled)
			continue
		}

		// flag combination as found
		matrix[vData.Color][vData.Number][vData.Enabled] = true
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
