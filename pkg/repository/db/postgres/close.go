package postgres

func (p Postgres) Close() {
	p.crud.Release()
}
