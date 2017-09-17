package hash

import(
  "testing"
)

func TestNew(t *testing.T) {
	tables := []struct {
		input, expected uint32
	}{
		{0, 0},
		{1, 1},
		{100, 100},
	}
	for _, pair := range tables {
		h, err := New(pair.input)
		if pair.input <= 0 {
      if err == nil {
        t.Errorf("hash.New Expected error with 0 size")
      }
		} else {
      if err != nil {
        t.Errorf("hash.New unexpected error")
        continue
      }
      if h.Size() != pair.expected {
        t.Errorf("hash.New incorrect size expected: %d, actual: %d",
          pair.expected, pair.input)
      }
      if h.Count() != 0 {
        t.Errorf("hash.New incorrect count")
      }
    }
	}
}

func TestSetNew(t *testing.T) {

}

func TestSetAndUpdate(t *testing.T) {

}

func TestSetNewFull(t *testing.T) {

}

func TestSetAndUpdateFull(t *testing.T) {

}

func TestGetExists(t *testing.T) {

}

func TestGetNotExists(t *testing.T) {

}

func TestDeleteExists(t *testing.T) {

}

func TestDeleteNotExists(t *testing.T) {

}

func TestLoad(t *testing.T) {

}
