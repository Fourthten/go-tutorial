package pg

import "github.com/go-pg/pg/v11/internal/pool"

func (db *DB) Pool() pool.Pooler {
	return db.pool
}

func (ln *Listener) CurrentConn() *pool.Conn {
	return ln.cn
}
