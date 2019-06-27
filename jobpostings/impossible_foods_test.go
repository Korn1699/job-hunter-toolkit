package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetImpossibleFoodsJobPostings(t *testing.T) {
	jobPostings, err := GetImpossibleFoodsJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}