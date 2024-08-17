package di

import (
	"context"
	"log"
	"transportadora/controller"
	"transportadora/infra/pgstore"
	"transportadora/tests"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go"
	"go.uber.org/dig"
)

type DI struct {
	Dig *dig.Container
}

func New() *DI {
	return &DI{
		Dig: dig.New(),
	}
}

func (d *DI) Close(test bool) {
	if test {
		d.Dig.Invoke(func(db pgstore.DBTX) {
			pool, ok := db.(*pgxpool.Pool)
			if !ok {
				panic("db is not a *pgxpool.Pool")
			}
			pool.Close()
		})

		d.Dig.Invoke(func(c testcontainers.Container) {
			cxt := context.Background()
			if err := c.Terminate(cxt); err != nil {
				log.Fatalf("Failed to terminate container: %v", err)
			}
		})
		return
	}

	err := d.Dig.Invoke(func(db pgstore.DBTX) {
		pool, ok := db.(*pgxpool.Pool)
		if !ok {
			panic("db is not a *pgxpool.Pool")
		}
		pool.Close()
	})

	if err != nil {
		panic(err)
	}
}

func (d *DI) Inject(test bool) {

	if test {
		d.Dig.Provide(tests.NewContainer)
		d.Dig.Provide(tests.SetupTestDB, dig.As(new(pgstore.DBTX)))
	} else {
		d.Dig.Provide(pgstore.NewPool, dig.As(new(pgstore.DBTX)))
	}

	d.Dig.Provide(pgstore.New)
	d.Dig.Provide(controller.NewAPI)

	d.Dig.Invoke(func(db pgstore.DBTX) {

		pool, ok := db.(*pgxpool.Pool)
		if !ok {
			log.Fatalf("db is not a *pgxpool.Pool")
		}

		err := pool.Ping(context.Background())
		if err != nil {
			log.Fatalf("error ping database: %v", err)
		}

		var path string
		if test {
			path = "file://../infra/pgstore/migrations"
		} else {
			path = "file://infra/pgstore/migrations"
		}

		m, err := migrate.New(
			path,
			pool.Config().ConnString(),
		)
		if err != nil {
			log.Fatalf("error create migration: %v", err)
		}

		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("error execute migrations: %v", err)
		}
	})

	d.injectRepos()
	d.injectServices()
	d.injectControllers()
}
