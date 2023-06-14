package database

func WithUsername(username string) func(*Database) {
	return func(db *Database) {
		db.username = username
	}
}

func WithPassword(password string) func(*Database) {
	return func(db *Database) {
		db.password = password
	}
}
func WithHost(host string) func(*Database) {
	return func(db *Database) {
		db.host = host
	}
}
func WithName(name string) func(*Database) {
	return func(db *Database) {
		db.name = name
	}
}
