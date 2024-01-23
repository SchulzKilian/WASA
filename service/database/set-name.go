package database

func (db *appdbimpl) SetName(name string, token string) error {
	_, err := db.c.Exec("UPDATE users SET username = ? WHERE user_id = ?", name, token)
	return err
}
