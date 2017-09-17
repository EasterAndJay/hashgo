package hash

import (
	"fmt"
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

func TestSetAndGetNew(t *testing.T) {
	tables := []struct {
		input, expected int32
	}{
		{0, 0},
		{1, 1},
		{-100, -100},
	}
	h, _ := New(uint32(1000))
	for i, pair := range tables {
		h.Set(fmt.Sprintf("key%d", i), pair.input)
		if h.Get(fmt.Sprintf("key%d", i)) != pair.expected {
			t.Errorf("Set and Get returned incorrect value")
		}
	}
}

func TestSetAndGetUpdate(t *testing.T) {
	tables := []struct {
		input, expected int32
	}{
		{0, 0},
		{1, 1},
		{-100, -100},
	}
	h, _ := New(uint32(1))
	for _, pair := range tables {
		set := h.Set("key", pair.input)
    if !set {
      t.Errorf("Set should return true when updating")
    }
		if h.Get("key") != pair.expected {
			t.Errorf("Set and Get returned incorrect value")
		}
	}
}

func TestSetNewFull(t *testing.T) {
	tables := []struct {
		input, expected int32
	}{
		{0, 0},
		{1, 1},
		{-100, -100},
	}
	h, _ := New(uint32(1))
  h.Set("key", tables[0].input)
	for i, pair := range tables {
		set := h.Set(fmt.Sprintf("key%d", i), pair.input)
		if set {
			t.Errorf("Set should return false when full")
		}
		if h.Get("key") != tables[0].expected {
			t.Errorf("Set and Get returned incorrect value")
		}
	}
}

func TestSetAndGetUpdateFull(t *testing.T) {
  tables := []struct {
    input, expected int32
  }{
    {0, 0},
    {1, 1},
    {-100, -100},
  }
  h, _ := New(uint32(1))
  h.Set("key", tables[0].input)
  for _, pair := range tables {
    set := h.Set("key", pair.input)
    if !set {
      t.Errorf("Set should return true when updating")
    }
    if h.Get("key") != pair.expected {
      t.Errorf("Set and Get returned incorrect value")
    }
  }
}

func TestGetNotExists(t *testing.T) {
  h, _ := New(uint32(1))
  val := h.Get("key")
  if val != nil {
    t.Errorf("Get empty returned non-nil value")
  }
}

func TestDeleteExists(t *testing.T) {

}

func TestDeleteNotExists(t *testing.T) {

}

func TestLoad(t *testing.T) {

}
