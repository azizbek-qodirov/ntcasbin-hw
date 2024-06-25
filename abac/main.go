package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func main() {
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Define test cases
	testCases := []struct {
		sub      string
		obj      string
		act      string
		time     string
		paid     string
		present  string
		expected bool
	}{
		{"student", "classes", "attend", "day_time", "any", "any", true},
		{"student", "classes", "attend", "night_time", "any", "any", false},

		{"student", "exams", "take", "any", "paid", "any", true},
		{"student", "exams", "take", "any", "unpaid", "any", false},

		{"student", "homework", "submit", "any", "any", "present", true},
		{"student", "homework", "submit", "any", "any", "absent", false},

		{"student", "classes", "attend", "any", "paid", "present", false},
		{"student", "exams", "take", "day_time", "paid", "any", true},
		{"student", "homework", "submit", "night_time", "any", "present", false},
	}

	for _, tc := range testCases {
		ok, err := e.Enforce(tc.sub, tc.obj, tc.act, tc.time, tc.paid, tc.present)
		if err != nil {
			fmt.Printf("Error enforcing policy: %v\n", err)
			continue
		}

		result := "denied"
		if ok {
			result = "granted"
		}

		expected := "denied"
		if tc.expected {
			expected = "granted"
		}

		fmt.Printf("Access %s for %s to %s %s (time: %s, paid: %s, present: %s) (expected: %s)\n", result, tc.sub, tc.act, tc.obj, tc.time, tc.paid, tc.present, expected)
	}

}
