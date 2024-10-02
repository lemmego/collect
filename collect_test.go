package collect

import (
	"reflect"
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

func TestSliceCollectionCount(t *testing.T) {
	sc := NewSlice([]int{1, 2, 3, 4, 5})

	count := sc.Count(func(x int, _ int) bool {
		return x%2 == 0
	})

	if count != 2 {
		t.Errorf("Expected count to be 2, but got %d", count)
	}

	count = sc.Count(func(x int, _ int) bool {
		return x > 10
	})

	if count != 0 {
		t.Errorf("Expected count to be 0, but got %d", count)
	}
}

func TestMapCollectionCount(t *testing.T) {
	testCases := []struct {
		name      string
		input     map[string]int
		predicate func(int, string) bool
		expected  int
	}{
		{
			name:      "Count even values",
			input:     map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			predicate: func(v int, _ string) bool { return v%2 == 0 },
			expected:  2,
		},
		{
			name:      "Count keys with length > 1",
			input:     map[string]int{"a": 1, "bb": 2, "ccc": 3, "d": 4},
			predicate: func(_ int, k string) bool { return len(k) > 1 },
			expected:  2,
		},
		{
			name:      "Count all",
			input:     map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(_ int, _ string) bool { return true },
			expected:  3,
		},
		{
			name:      "Count none",
			input:     map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(_ int, _ string) bool { return false },
			expected:  0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mc := NewMap(tc.input)
			result := mc.Count(tc.predicate)
			if result != tc.expected {
				t.Errorf("Expected count %d, but got %d", tc.expected, result)
			}
		})
	}
}

func TestSliceCollectionSome(t *testing.T) {
	sc := NewSlice([]int{1, 2, 3, 4, 5})

	// Test when condition is met
	result := sc.Some(func(x int, _ int) bool {
		return x > 3
	})
	if !result {
		t.Errorf("Expected Some to return true, but got false")
	}

	// Test when condition is not met
	result = sc.Some(func(x int, _ int) bool {
		return x > 10
	})
	if result {
		t.Errorf("Expected Some to return false, but got true")
	}

	// Test with empty slice
	emptySlice := NewSlice([]int{})
	result = emptySlice.Some(func(x int, _ int) bool {
		return x > 0
	})
	if result {
		t.Errorf("Expected Some on empty slice to return false, but got true")
	}

	// Test using index in the predicate function
	result = sc.Some(func(x int, i int) bool {
		return i == 2 && x == 3
	})
	if !result {
		t.Errorf("Expected Some to return true for index-based condition, but got false")
	}
}

func TestMapCollectionSome(t *testing.T) {
	testCases := []struct {
		name      string
		input     map[string]int
		predicate func(int, string) bool
		expected  bool
	}{
		{
			name:      "Some elements satisfy condition",
			input:     map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(v int, k string) bool { return v > 2 },
			expected:  true,
		},
		{
			name:      "No elements satisfy condition",
			input:     map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(v int, k string) bool { return v > 5 },
			expected:  false,
		},
		{
			name:      "Empty map",
			input:     map[string]int{},
			predicate: func(v int, k string) bool { return true },
			expected:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mc := NewMap(tc.input)
			result := mc.Some(tc.predicate)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestSliceCollectionEvery(t *testing.T) {
	sc := NewSlice([]int{1, 2, 3, 4, 5})

	// Test when not all elements satisfy the condition
	allEven := sc.Every(func(x int, _ int) bool {
		return x%2 == 0
	})
	if allEven {
		t.Errorf("Expected not all elements to be even, but got true")
	}

	// Test when all elements satisfy the condition
	allPositive := sc.Every(func(x int, _ int) bool {
		return x > 0
	})
	if !allPositive {
		t.Errorf("Expected all elements to be positive, but got false")
	}

	// Test with empty slice
	emptySlice := NewSlice([]int{})
	allZero := emptySlice.Every(func(x int, _ int) bool {
		return x == 0
	})
	if !allZero {
		t.Errorf("Expected true for empty slice, but got false")
	}
}

func TestMapCollectionEvery(t *testing.T) {
	testCases := []struct {
		name      string
		input     map[string]int
		predicate func(int, string) bool
		expected  bool
	}{
		{
			name:      "All elements satisfy predicate",
			input:     map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(v int, k string) bool { return v > 0 },
			expected:  true,
		},
		{
			name:      "Some elements don't satisfy predicate",
			input:     map[string]int{"a": 1, "b": -2, "c": 3},
			predicate: func(v int, k string) bool { return v > 0 },
			expected:  false,
		},
		{
			name:      "Empty map",
			input:     map[string]int{},
			predicate: func(v int, k string) bool { return v > 0 },
			expected:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mc := NewMap(tc.input)
			result := mc.Every(tc.predicate)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestSliceCollectionNone(t *testing.T) {
	sc := NewSlice([]int{1, 2, 3, 4, 5})

	result := sc.None(func(x int, _ int) bool {
		return x > 10
	})
	if !result {
		t.Errorf("Expected None to return true, got false")
	}

	result = sc.None(func(x int, _ int) bool {
		return x == 3
	})
	if result {
		t.Errorf("Expected None to return false, got true")
	}
}

func TestMapCollectionNone(t *testing.T) {
	mc := NewMap(map[string]int{"a": 1, "b": 2, "c": 3})

	result := mc.None(func(v int, _ string) bool {
		return v > 10
	})
	if !result {
		t.Errorf("Expected None to return true, got false")
	}

	result = mc.None(func(v int, _ string) bool {
		return v == 2
	})
	if result {
		t.Errorf("Expected None to return false, got true")
	}
}

func TestSliceCollectionConcat(t *testing.T) {
	sc := NewSlice([]int{1, 2, 3})
	result := sc.Concat([]int{4, 5, 6})
	expected := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test with empty slice
	emptyResult := sc.Concat([]int{})
	if !reflect.DeepEqual(emptyResult, sc.Items()) {
		t.Errorf("Expected %v, but got %v", sc.Items(), emptyResult)
	}

	// Test with nil slice
	nilResult := sc.Concat(nil)
	if !reflect.DeepEqual(nilResult, sc.Items()) {
		t.Errorf("Expected %v, but got %v", sc.Items(), nilResult)
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
