package models

import "testing"

// TestValidateTask tests the Validate method of the Task struct
func TestValidateTask(t *testing.T) {
	// Define test cases
	cases := []struct {
		name     string // The name of the test case
		task     Task   // The Task object to test
		expected bool   // The expected result of the Validate method (true or false)
	}{
		{
			name:     "Valid Task",                // Test case name
			task:     Task{Title: "Learn Docker"}, // A valid task with a non-empty title
			expected: true,                        // We expect the validation to return true
		},
		{
			name:     "Empty Title",    // Test case name
			task:     Task{Title: " "}, // A task with an empty title (only spaces)
			expected: false,            // We expect the validation to return false
		},
	}

	// Iterate over each test case
	for _, c := range cases {
		// Run each test case as a subtest
		t.Run(c.name, func(t *testing.T) {
			// Call the Validate method on the task and store the result
			result := c.task.Validate()

			// Compare the result with the expected value
			if result != c.expected {
				// If the result does not match the expected value, report an error
				t.Errorf("Expected %v, got %v", c.expected, result)
			}
		})
	}
}
