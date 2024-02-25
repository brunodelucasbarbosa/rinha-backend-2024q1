package entities

type Client struct {
	Id    int    `db:"id"`
	Name  string `db:"nome"`
	Limit int    `db:"limite"`
}

func (c Client) Exists() bool {
	return c.Id != 0
}