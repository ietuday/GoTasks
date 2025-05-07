package models

import "testing"

func TestValidateTask(t *testing.T) {
	cases := []struct {
		name     string
		task     Task
		expected bool
	}{
		{
			name:     "Valid Task",
			task:     Task{Title: "Learn Docker"},
			expected: true,
		},
		{
			name:     "Empty Title",
			task:     Task{Title: " "},
			expected: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := c.task.Validate()
			if result != c.expected {
				t.Errorf("Expected %v, got %v", c.expected, result)
			}
		})
	}
}
