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

	testCases := []struct {
		sub      string
		obj      string
		act      string
		expected bool
	}{
		{"student", "class", "attend", true},
		{"student", "face_id", "use", true},
		{"student", "class_schedule", "view", true},
		{"student", "class_schedule", "access", false},

		{"teacher", "class", "attend", true},
		{"teacher", "class_schedule", "access", true},
		{"teacher", "exams", "grade", true},
		{"teacher", "face_id", "use", true},
		{"teacher", "class_schedule", "view", false},

		{"admin", "users", "manage", true},
		{"admin", "roles", "manage", true},
		{"admin", "permissions", "manage", true},
		{"admin", "class_schedule", "access", true},
		{"admin", "face_id", "use", true},

		{"guest", "class_schedule", "view", true},
		{"guest", "class_schedule", "access", false},
	}

	for _, tc := range testCases {
		ok, err := e.Enforce(tc.sub, tc.obj, tc.act)
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

		fmt.Printf("Access %s for %s to %s %s (expected: %s)\n", result, tc.sub, tc.act, tc.obj, expected)
	}
}
