package user

var (
	createUserQuery  = "INSERT INTO user(username, password) VALUES (?,?)"
	getUserQuery     = "SELECT id, username, password, role FROM user WHERE username=?"
	getUserByIDQuery = "SELECT id, username, password, role FROM user WHERE id=?"
)
