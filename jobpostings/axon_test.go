package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetAxonJobPostings(t *testing.T) {
	jobPostings, err := GetAxonJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}