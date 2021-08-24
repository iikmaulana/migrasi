module github.com/iikmaulana/migrasi

go 1.15

require (
	github.com/gearintellix/u2 v1.0.9
	github.com/google/uuid v1.1.1
	github.com/iikmaulana/gateway v0.0.4
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.2.0
	github.com/opentracing/opentracing-go v1.1.0
	gopkg.in/rethinkdb/rethinkdb-go.v6 v6.2.1
)

//replace github.com/iikmaulana/gateway => ../gateway
