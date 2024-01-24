package api

func (rt *_router) Close() error {
    if rt.db != nil {
        if err := rt.db.Close(); err != nil {
            return err // return error if db.Close() fails
        }
    }

    return nil
}
