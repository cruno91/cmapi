package storage

import (
	"CMAPI/internal/entity"
	"CMAPI/internal/entity/model"
	"context"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"strings"
)

type PostgresRepository struct {
	db *sql.DB
}

// CreateEntityType inserts a new entity type.
func (r *PostgresRepository) CreateEntityType(ctx context.Context, et *model.Type) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO entity_types (id, name, description)
		VALUES ($1, $2, $3)
	`, et.ID, et.Name, et.Description)
	return err
}

// CreateBundle inserts metadata and creates the physical table.
func (r *PostgresRepository) CreateBundle(ctx context.Context, b *model.Bundle, fields []model.Field) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	tableName := b.EntityTypeID + "__" + b.ID

	_, err = tx.ExecContext(ctx, `
		INSERT INTO bundles (id, entity_type_id, name, label, table_name)
		VALUES ($1, $2, $3, $4, $5)
	`, b.ID, b.EntityTypeID, b.Name, b.Label, tableName)
	if err != nil {
		return err
	}

	// Insert field definitions
	for _, f := range fields {
		_, err := tx.ExecContext(ctx, `
			INSERT INTO field_definitions (id, bundle_id, name, type, required, settings)
			VALUES ($1, $2, $3, $4, $5, $6)
		`, f.ID, b.ID, f.Name, f.Type, f.Required, f.Settings)
		if err != nil {
			return err
		}
	}

	// Create table
	createSQL := fmt.Sprintf(`
		CREATE TABLE %s (
			id UUID PRIMARY KEY,
			%s,
			created_at TIMESTAMPTZ DEFAULT now(),
			updated_at TIMESTAMPTZ DEFAULT now()
		)
	`, pq.QuoteIdentifier(tableName), buildFieldSQL(fields))

	_, err = tx.ExecContext(ctx, createSQL)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// InsertEntity inserts a new entity.
func (r *PostgresRepository) InsertEntity(ctx context.Context, tableName string, data map[string]any) error {
	columns := []string{}
	values := []any{}
	placeholders := []string{}

	i := 1
	for k, v := range data {
		columns = append(columns, pq.QuoteIdentifier(k))
		values = append(values, v)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
		i++
	}

	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`,
		pq.QuoteIdentifier(tableName),
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "))

	_, err := r.db.ExecContext(ctx, query, values...)
	return err
}

func buildFieldSQL(fields []model.Field) string {
	var columns []string
	for _, f := range fields {
		fieldType, ok := entity.GetFieldType(f.Type)
		if !ok {
			// You can panic here or return an error instead
			panic(fmt.Sprintf("unknown field type: %s", f.Type))
		}

		sqlType := fieldType.SQLType(f.Settings)
		nullability := "NULL"
		if f.Required {
			nullability = "NOT NULL"
		}

		col := fmt.Sprintf("%s %s %s", pq.QuoteIdentifier(f.Name), sqlType, nullability)
		columns = append(columns, col)
	}
	return strings.Join(columns, ",\n")
}
