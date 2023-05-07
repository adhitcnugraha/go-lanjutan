package unittest

import "testing"

func TestHomeworkBE5304103(t *testing.T) {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wants := 3.7
	result := HomeworkBE5304103(numbers)
	if result != wants {
		t.Errorf("Expected result to be %f, but got %f instead.", wants, result)
	}

	numbers1 := []int{}
	wants1 := 0.0
	result1 := HomeworkBE5304103(numbers1)
	if result1 != wants1 {
		t.Errorf("Expected result to be %f, but got %f", wants1, result1)
	}

	numbers2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	wants2 := -1.0
	result2 := HomeworkBE5304103(numbers2)
	if result2 != wants2 {
		t.Errorf("Expected result to be %f, but got %f", wants2, result2)
	}

	numbers3 := []int{3, 6, 9}
	wants3 := 0.0
	result3 := HomeworkBE5304103(numbers3)
	if result3 != wants3 {
		t.Errorf("Expected result to be %f, but got %f", wants3, result3)
	}
}
