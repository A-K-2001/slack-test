package store

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStorage struct {
	pgOnce      *sync.Once
	customTypes [][2]string
	Pool        *pgxpool.Pool
}

func NewPostgresStorage() *PostgresStorage {
	return &PostgresStorage{
		pgOnce:      &sync.Once{},
		customTypes: [][2]string{},
	}
}

func (ps *PostgresStorage) Connect(ctx context.Context, uri string) error {
	cfg, err := pgxpool.ParseConfig(uri)

	if err != nil {
		return fmt.Errorf("unable to parse database uri: %v", err)
	}

	cfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		return ps.registerCustomTypes(ctx, conn)
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)

	if err != nil {
		return fmt.Errorf("unable create connection pool: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("unable to ping database: %v", err)
	}

	ps.Pool = pool

	return nil
}

func (ps *PostgresStorage) Close() {
	if ps.Pool == nil {
		return
	}

	ps.Pool.Close()
}

func (ps *PostgresStorage) registerCustomTypes(ctx context.Context, conn *pgx.Conn) error {
	ps.pgOnce.Do(func() {
		rows, err := conn.Query(ctx, `SELECT DISTINCT pg_type.typname AS enum_type,
     concat('_', pg_type.typname) as array_type FROM pg_type JOIN pg_enum
     ON pg_enum.enumtypid = pg_type.oid;`)

		if err != nil {
			log.Fatalf("Unable to query custom types: %v", err)
		}

		defer rows.Close()

		fmt.Println("\nRetrieving custom postgres types:")

		ps.customTypes, err = pgx.CollectRows[[2]string](rows, (func(row pgx.CollectableRow) ([2]string, error) {
			var result [2]string

			if err := row.Scan(&result[0], &result[1]); err != nil {
				return result, err
			}

			fmt.Println("  type: ", result)

			return result, nil
		}))

		fmt.Printf("Finished retrieving custom postgres types\n\n")

		if err != nil {
			log.Fatalf("Unable to collect custom types: %v", err)
		}
	})

	for _, t := range ps.customTypes {
		fmt.Println("Loading ", t)
		loaded, err := conn.LoadType(ctx, fmt.Sprintf("\"%s\"",t[0]))

		if err != nil {
			return fmt.Errorf("unable to load type: %v", err)
		}

		conn.TypeMap().RegisterType(loaded)

		loaded, err = conn.LoadType(context.Background(), fmt.Sprintf("\"%s\"",t[1]))

		if err != nil {
			return fmt.Errorf("unable to load type: %v", err)
		}

		conn.TypeMap().RegisterType(loaded)
	}

	return nil
}
