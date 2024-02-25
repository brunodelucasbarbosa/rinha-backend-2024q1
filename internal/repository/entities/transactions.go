package entities

import "time"

type Transacoes struct {
	ClientId    int        `db:"cliente_id"`
	Value       int        `db:"valor"`
	Type        string     `db:"tipo"`
	Description string     `db:"descricao"`
	CreatedAt   *time.Time `db:"realizada_em"`
}
