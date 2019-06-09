package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetAirbnbJobPostings(t *testing.T) {
	jobPostings, err := GetAirbnbJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}