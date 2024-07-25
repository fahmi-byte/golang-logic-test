package test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type Comment struct {
	CommentId      int       `json:"commentId"`
	CommentContent string    `json:"commentContent"`
	Replies        []Comment `json:"replies,omitempty"` // omitempty skips this field if it's empty
}

func TestComments(t *testing.T) {
	jsonComments := `
	[
		{
			"commentId": 1,
			"commentContent": "Hai",
			"replies": [
				{
					"commentId": 11,
					"commentContent": "Hai Juga",
					"replies": [
						{
							"commentId": 111,
							"commentContent": "Haai juga hai jugaa"
						},
						{
							"commentId": 112,
							"commentContent": "Haai juga hai jugaa"
						}
					]
				},
				{
					"commentId": 12,
					"commentContent": "Hai Juga",
					"replies": [
						{
							"commentId": 121,
							"commentContent": "Haai juga hai jugaa"
						}
					]
				}
			]
		},
		{
			"commentId": 2,
			"commentContent": "Halooo"
		},
	]`

	var comments []Comment

	err := json.Unmarshal([]byte(jsonComments), &comments)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	totalComments := CountTotalComments(comments)

	fmt.Printf("Total Commentnya adalah: %d\n", totalComments)

}

func CountTotalComments(comments []Comment) int {
	total := len(comments) // Start by counting all top-level comments

	for _, comment := range comments {
		// Add the count of replies recursively
		total += CountTotalComments(comment.Replies)
	}

	return total
}
