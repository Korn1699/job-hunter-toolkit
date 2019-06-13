package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetMongoDBJobPostings(t *testing.T) {
	jobPostings, err := GetMongoDBJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}
