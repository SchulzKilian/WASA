package database

import (
    "errors"
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
func (db *appdbimpl) DeleteComment(commentID string) error {
    _, err := db.c.Exec("DELETE FROM comments WHERE comment_id = ?", commentID)
    return err
}