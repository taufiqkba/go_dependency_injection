package simple

type Database struct {
	Name string
}

type DatabasePostgresQL Database
type DatabaseMongoDB Database

func NewDatabasePostgresQL() *DatabasePostgresQL {
	return (*DatabasePostgresQL)(&Database{Name: "PostgresQL"})
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}

type DatabaseRepository struct {
	DatabasePostgresQL *DatabasePostgresQL
	DatabaseMongoDB    *DatabaseMongoDB
}

func NewDatabaseRepository(postgresQL *DatabasePostgresQL, mongodb *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgresQL: postgresQL, DatabaseMongoDB: mongodb}
}
