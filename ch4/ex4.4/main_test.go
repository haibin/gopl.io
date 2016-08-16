package ex4point4

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 2)
	wanted := []int{2, 3, 4, 5, 0, 1}
	if !reflect.DeepEqual(s, wanted) {
		t.Errorf("rotated %v, want %v", s, wanted)
	}
}
