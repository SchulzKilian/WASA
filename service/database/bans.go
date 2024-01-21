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


func (db *appdbimpl) AmIBanned(banner, banned string) (bool, error) {
    var exists bool
    query := "SELECT EXISTS(SELECT 1 FROM bans WHERE banner = ? AND banned = ?)"
    err := db.c.QueryRow(query, banner, banned).Scan(&exists)
    if err != nil {
        return false, err
    }
    return exists, nil
}