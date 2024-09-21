package collect

import (
	"testing"
)

func TestSliceCollectionReduce(t *testing.T) {
	sc := NewSlice([]int{1, 2, 3, 4, 5})

	sum := sc.Reduce(func(acc, curr, index int) int {
		return acc + curr
	}, 0)

	expected := 15
	if sum != expected {
		t.Errorf("Expected sum to be %d, but got %d", expected, sum)
	}

	product := sc.Reduce(func(acc, curr, index int) int {
		return acc * curr
	}, 1)

	expected = 120
	if product != expected {
		t.Errorf("Expected product to be %d, but got %d", expected, product)
	}
}

func TestMapCollectionReduce(t *testing.T) {
	mc := NewMap(map[string]int{"a": 1, "b": 2, "c": 3})

	result := mc.Reduce(func(acc int, val int, key string) int {
		return acc + val
	}, 0)

	expected := 6
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Test with string concatenation
	mcString := NewMap(map[int]string{1: "hello", 2: "world", 3: "!"})

	resultString := mcString.Reduce(func(acc string, val string, key int) string {
		return acc + val
	}, "")

	expectedString := "helloworld!"
	if resultString != expectedString {
		t.Errorf("Expected %s, but got %s", expectedString, resultString)
	}
}

func TestSliceCollectionFindIndex(t *testing.T) {
	sc := NewSlice([]int{1, 2, 3, 4, 5})

	index := sc.FindIndex(func(x int) bool {
		return x == 3
	})

	if index != 2 {
		t.Errorf("Expected index 2, got %d", index)
	}

	index = sc.FindIndex(func(x int) bool {
		return x == 6
	})

	if index != -1 {
		t.Errorf("Expected index -1, got %d", index)
	}
}

func TestSliceCollectionFindLast(t *testing.T) {
	sc := NewSlice([]int{1, 2, 3, 4, 5})

	result, found := sc.FindLast(func(x int) bool {
		return x%2 == 0
	})

	if !found {
		t.Error("Expected to find an even number")
	}

	if result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}

	_, notFound := sc.FindLast(func(x int) bool {
		return x > 10
	})

	if notFound {
		t.Error("Expected not to find a number greater than 10")
	}
}

func TestSliceCollectionFindLastIndex(t *testing.T) {
	sc := NewSlice([]int{1, 2, 3, 4, 5, 4, 3, 2, 1})

	index := sc.FindLastIndex(func(x int) bool {
		return x == 4
	})

	if index != 5 {
		t.Errorf("Expected index 5, got %d", index)
	}

	index = sc.FindLastIndex(func(x int) bool {
		return x == 6
	})

	if index != -1 {
		t.Errorf("Expected index -1, got %d", index)
	}
}

// ================== Base Functions ==================

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
