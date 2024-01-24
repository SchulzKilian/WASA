package database

import (
	"errors"
	"strconv"
)

// AddComment adds a comment to the database
func (db *appdbimpl) AddComment(comment Comment) error {
	if comment.Content == "" || comment.Commenter == "" || comment.PhotoID == "" {
		return errors.New("comment, commenter, and photoID must be provided")
	}

	_, err := db.c.Exec("INSERT INTO comments (content, commenter, photo_id) VALUES (?, ?, ?)",
		comment.Content, comment.Commenter, comment.PhotoID)
	return err
}

// DeleteComment removes a comment from the database
// Assuming you have a unique identifier for comments (like commentID)
func (db *appdbimpl) DeleteComment(commentID, commenter string) error {
	var existingCommenter string
	id, err := strconv.Atoi(commentID)
	if err != nil {
		return errors.New("invalid comment ID")
	}

	err = db.c.QueryRow("SELECT commenter FROM comments WHERE comment_id = ?", id).Scan(&existingCommenter)
	if err != nil {
		return errors.New("Comment does not exist")

	}

	// Check if the commenter matches
	if existingCommenter != commenter {
		return errors.New("unauthorized: commenter does not match")
	}

	_, err = db.c.Exec("DELETE FROM comments WHERE comment_id = ?", commentID)
	return err
}

func (db *appdbimpl) GetComments(photoid string) ([]Comment, error) {

	var comments []Comment
	query := "SELECT content, comment_id, commenter, photo_id FROM comments WHERE photo_id = ?"
	rows, err := db.c.Query(query, photoid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.Content, &comment.CommentID, &comment.Commenter, &comment.PhotoID); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}
