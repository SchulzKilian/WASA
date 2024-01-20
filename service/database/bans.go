package database

// AddBan adds a ban to the database
func (db *appdbimpl) AddBan(banner, banned string) error {
    _, err := db.c.Exec("INSERT INTO bans (banner, banned) VALUES (?, ?)", banner, banned)
    return err
}

// DeleteBan removes a ban from the database
func (db *appdbimpl) DeleteBan(banner, banned string) error {
    _, err := db.c.Exec("DELETE FROM bans WHERE banner = ? AND banned = ?", banner, banned)
    return err
}
