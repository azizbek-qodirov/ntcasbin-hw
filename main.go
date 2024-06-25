package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

type Student struct {
	Name                        string
	IsPaymentCompleted          bool
	IsLessonHour                bool
	IsPaymentCompletedNextMonth bool
	Is20thLesson                bool
	UncompletedHomeworks        int
}

func main() {
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		log.Fatalf("Failed to create enforcer: %v", err)
	}

	alice := Student{
		Name:                        "alice",
		IsPaymentCompleted:          true,
		IsLessonHour:                true,
		IsPaymentCompletedNextMonth: true,
		Is20thLesson:                true,
		UncompletedHomeworks:        2,
	}

	// Check permissions
	ok, err := e.Enforce(alice.Name, "faceID", "use")
	if err != nil {
		log.Fatalf("Failed to enforce policy: %v", err)
	}
	if ok {
		fmt.Println("Alice is allowed to use faceID.")
	} else {
		fmt.Println("Alice is not allowed to use faceID.")
	}

	ok, err = e.Enforce(alice.IsPaymentCompleted, "lesson", "participate")
	if err != nil {
		log.Fatalf("Failed to enforce policy: %v", err)
	}
	if ok {
		fmt.Println("Alice is allowed to participate in lesson.")
	} else {
		fmt.Println("Alice is not allowed to participate in lesson.")
	}

	ok, err = e.Enforce(alice.Name, "exam", "participate")
	if err != nil {
		log.Fatalf("Failed to enforce policy: %v", err)
	}
	if ok {
		fmt.Println("Alice is allowed to participate in exam.")
	} else {
		fmt.Println("Alice is not allowed to participate in exam.")
	}
}
