package database

func (db *appdbimpl) SetName(name string, token string, oldname string) error {
	_, err := db.c.Exec("UPDATE follow SET username = ? WHERE username = ?", name, oldname)
	if err != nil {
		return err
	}
	_, err = db.c.Exec("UPDATE bans SET username = ? WHERE username = ?", name, oldname)
	if err != nil {
		return err
	}
	_, err = db.c.Exec("UPDATE comments SET username = ? WHERE username = ?", name, oldname)
	if err != nil {
		return err
	}
	_, err = db.c.Exec("UPDATE likes SET username = ? WHERE username = ?", name, oldname)
	if err != nil {
		return err
	}
	_, err = db.c.Exec("UPDATE photos SET username = ? WHERE username = ?", name, oldname)
	if err != nil {
		return err
	}
	_, err = db.c.Exec("UPDATE users SET username = ? WHERE user_id = ?", name, token)
	return err
}
