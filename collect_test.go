package collect

import (
	"testing"
)

func TestEach(t *testing.T) {
	arr := []int{1, 2, 3}
	Each(arr, func(x int, i int) {
		if x != i+1 {
			t.Errorf("Expected %d, got %d", i+1, x)
		}
	})
}

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3}
	result := Map(arr, func(x int, i int) int {
		return x + i
	})
	if result[0] != 1 {
		t.Errorf("Expected 1, got %d", result[0])
	}
	if result[1] != 3 {
		t.Errorf("Expected 3, got %d", result[1])
	}
	if result[2] != 5 {
		t.Errorf("Expected 5, got %d", result[2])
	}
}

func TestFilter(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	result := Filter(arr, func(x int, i int) bool {
		return x%2 == 0
	})
	if result[0] != 2 {
		t.Errorf("Expected 2, got %d", result[0])
	}
	if result[1] != 4 {
		t.Errorf("Expected 4, got %d", result[1])
	}
}

func TestCount(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	result := Count(arr, func(x int, i int) bool {
		return x%2 == 0
	})
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}

func TestFindIndex(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	result := FindIndex(arr, func(x int) bool {
		return x == 3
	})
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}

func TestFindLastIndex(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	result := FindLastIndex(arr, func(x int) bool {
		return x == 3
	})
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}

func TestFind(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	result, ok := Find(arr, func(x int) bool {
		return x == 3
	})
	if !ok || result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}

func TestFindLast(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	result, ok := FindLast(arr, func(x int) bool {
		return x == 3
	})
	if !ok || result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}

func TestSome(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	result := Some(arr, func(x int, i int) bool {
		return x == 3
	})
	if !result {
		t.Errorf("Expected true, got %t", result)
	}
}

func TestNone(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	result := None(arr, func(x int, i int) bool {
		return x == 6
	})
	if !result {
		t.Errorf("Expected true, got %t", result)
	}
}

func TestEvery(t *testing.T) {
	arr := []int{2, 4, 6, 8, 10}
	result := Every(arr, func(x int, i int) bool {
		return x%2 == 0
	})
	if !result {
		t.Errorf("Expected true, got %t", result)
	}
}

func TestConcat(t *testing.T) {
	arr := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	result := Concat(arr, arr2)
	if len(result) != 6 {
		t.Errorf("Expected 6, got %d", len(result))
	}

	for i, x := range result {
		if x != i+1 {
			t.Errorf("Expected %d, got %d", i+1, x)
		}
	}
}

func TestConcatMap(t *testing.T) {
	arr := []int{1, 2, 3}
	result := ConcatMap(arr, func(x int) []int {
		return []int{x, x * 2}
	})
	t.Log(result)
	if len(result) != 6 {
		t.Errorf("Expected 6, got %d", len(result))
	}
}

func TestReverse(t *testing.T) {
	arr := []int{1, 2, 3}
	result := Reverse(arr)
	if result[0] != 3 {
		t.Errorf("Expected 3, got %d", result[0])
	}
	if result[1] != 2 {
		t.Errorf("Expected 2, got %d", result[1])
	}
	if result[2] != 1 {
		t.Errorf("Expected 1, got %d", result[2])
	}
}

func TestUniq(t *testing.T) {
	arr := []int{1, 2, 3, 2, 1}
	result := Uniq(arr)
	if len(result) != 3 {
		t.Errorf("Expected 3, got %d", len(result))
	}
}

func TestUniqBy(t *testing.T) {
	arr := []int{1, 2, 3, 2, 1}
	result := UniqBy(arr, func(x int) int {
		return x
	})
	if len(result) != 3 {
		t.Errorf("Expected 3, got %d", len(result))
	}
}
