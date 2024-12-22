package session

var (
	getSessionQuery        = "SELECT user_id, token, is_valid FROM session WHERE token = ?"
	createSessionQuery     = "INSERT INTO session(token,user_id,is_valid) VALUES (?,?,?)"
	invalidateSessionQuery = "UPDATE session SET is_valid = false WHERE user_id = ?"
)
