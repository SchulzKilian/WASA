package database

func (db *appdbimpl) SetName(name string, token string, oldname string) error {
	_, err := db.c.Exec("UPDATE follow SET follower = ? WHERE follower = ?", name, oldname)
	if err != nil {
		return err
	}
	_, err = db.c.Exec("UPDATE follow SET followed = ? WHERE followed = ?", name, oldname)
	if err != nil {
		return err
	}
	_, err = db.c.Exec("UPDATE bans SET banner = ? WHERE banner = ?", name, oldname)
	if err != nil {
		return err
	}
	_, err = db.c.Exec("UPDATE bans SET banned = ? WHERE banned = ?", name, oldname)
	if err != nil {
		return err
	}
	_, err = db.c.Exec("UPDATE comments SET commenter = ? WHERE commenter = ?", name, oldname)
	if err != nil {
		return err
	}
	_, err = db.c.Exec("UPDATE likes SET liker = ? WHERE liker = ?", name, oldname)
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
