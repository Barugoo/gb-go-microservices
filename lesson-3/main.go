package main

const (
	USER_STATUS_NOT_SET = iota
	USER_STATUS_ACTIVE
	USER_STATUS_BLOCKED
)

type User struct {
	ID     int32
	Name   string
	Status int32
}

func main() {
	statusFilter := USER_STATUS_BLOCKED

	sqlReq := "SELECT * FROM users"

	switch statusFilter {
	case USER_STATUS_NOT_SET:
		sqlReq += " WHERE status = 0"
	case USER_STATUS_ACTIVE:
		sqlReq += " WHERE status = 1"
	case USER_STATUS_BLOCKED:
		sqlReq += " WHERE status = 2"
	}
}
