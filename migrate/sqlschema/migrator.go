package sqlschema

import (
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type MigratorDialect interface {
	schema.Dialect
	Migrator(*bun.DB) Migrator
}

type Migrator interface {
	AppendSQL(b []byte, operation interface{}) ([]byte, error)
}

// migrator is a dialect-agnostic wrapper for sqlschema.MigratorDialect.
type migrator struct {
	Migrator
}

func NewMigrator(db *bun.DB) (Migrator, error) {
	md, ok := db.Dialect().(MigratorDialect)
	if !ok {
		return nil, fmt.Errorf("%q dialect does not implement sqlschema.Migrator", db.Dialect().Name())
	}
	return &migrator{
		Migrator: md.Migrator(db),
	}, nil
}

// BaseMigrator can be embeded by dialect's Migrator implementations to re-use some of the existing bun queries.
type BaseMigrator struct {
	db *bun.DB
}

func NewBaseMigrator(db *bun.DB) *BaseMigrator {
	return &BaseMigrator{db: db}
}

func (m *BaseMigrator) AppendCreateTable(b []byte, model interface{}) ([]byte, error) {
	return m.db.NewCreateTable().Model(model).AppendQuery(m.db.Formatter(), b)
}

func (m *BaseMigrator) AppendDropTable(b []byte, fqn FQN) ([]byte, error) {
	return m.db.NewDropTable().TableExpr(fqn.String()).AppendQuery(m.db.Formatter(), b)
}
