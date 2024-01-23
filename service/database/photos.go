package database

func (db *appdbimpl) AddPhoto(photo Photo) error {
	// SQL statement to insert a new photo
	stmt := `INSERT INTO photos (username, image_data, timestamp) VALUES (?, ?, ?)`

	// Executing the SQL statement
	_, err := db.c.Exec(stmt, photo.Username, photo.ImageData, photo.Timestamp)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) DeletePhoto(photo_id int) error {
	stmt := `DELETE FROM photos WHERE photo_id = ?`

	_, err := db.c.Exec(stmt, photo_id)
	if err != nil {
		return err
	}

	return nil
}
