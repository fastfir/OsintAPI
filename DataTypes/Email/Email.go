package OsintAPI

type Email struct {
	username, domain string
}

func (e Email) FullEmail() string {
	return e.username + "@" + e.domain
}
