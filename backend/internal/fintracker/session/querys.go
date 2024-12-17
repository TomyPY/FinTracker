package session

var (
	getSessionQuery    = "SELECT token FROM session WHERE token = ?"
	createSessionQuery = "INSERT INTO session(token,user_id) VALUES (?,?)"
)
