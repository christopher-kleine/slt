package slt_test

import (
	"fmt"

	"github.com/christopher-kleine/slt"
)

func ExampleJoin() {
	type Student struct {
		ID        string
		Name      string
		ProjectID string
	}

	type Project struct {
		ID   string
		Name string
	}

	students := []Student{
		{ID: "AAA", Name: "Benjamin", ProjectID: "DDD"},
		{ID: "BBB", Name: "Karla", ProjectID: "EEE"},
		{ID: "CCC", Name: "Peter", ProjectID: "DDD"},
	}

	projects := []Project{
		{ID: "DDD", Name: "Prepare classroom"},
		{ID: "EEE", Name: "Invite parents"},
	}

	projectByStudentID := slt.Join(
		students,
		projects,
		func(s Student) string { return s.ProjectID },
		func(p Project) string { return p.ID },
		func(s Student) string { return s.ID },
	)

	fmt.Printf("%+v", projectByStudentID)

	// Output:
	// map[AAA:{ID:DDD Name:Prepare classroom} BBB:{ID:EEE Name:Invite parents} CCC:{ID:DDD Name:Prepare classroom}]
}
