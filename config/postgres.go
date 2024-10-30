package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresConfig struct {
	Host              string `mapstructure:"host"`
	Port              uint   `mapstructure:"password"`
	DBName            string `mapstructure:"dbname"`
	Username          string `mapstructure:"username"`
	Password          string `mapstructure:"password"`
	MaxIdleConnection int    `mapstructure:"maxidleconnection"`
	MaxOpenConnection int    `mapstructure:"maxopenconnection"`
	MaxIdleTime       int    `mapstructure:"maxidletime"`
}

func NewPostgresGormConnection() (db *gorm.DB) {
	//connect DB
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), //io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(
		postgres.Open(
			postgresDSN(),
		),
		&gorm.Config{
			Logger: newLogger,
		},
	)

	if err != nil {
		panic(err)
	}

	dbSql, err := db.DB()
	if err != nil {
		panic(err)
	}

	postgresPoolConf(dbSql)
	if err := dbSql.Ping(); err != nil {
		panic(err)
	}
	log.Println("successfully connect to Postgres")
	return db
}

func postgresPoolConf(dbSQL *sql.DB) {
	// set extended config
	dbSQL.SetMaxIdleConns(Load.DataSource.Postgres.Master.MaxIdleConnection)
	dbSQL.SetMaxOpenConns(Load.DataSource.Postgres.Master.MaxOpenConnection)
	dbSQL.SetConnMaxIdleTime(time.Duration(Load.DataSource.Postgres.Master.MaxIdleTime))
}

func postgresDSN() string {
	return fmt.Sprintf(`
		host=%v
		port=%v
		user=%v
		password=%v
		dbname=%v
		sslmode=disable`, Load.DataSource.Postgres.Master.Host, Load.DataSource.Postgres.Master.Port, Load.DataSource.Postgres.Master.Username, Load.DataSource.Postgres.Master.Password, Load.DataSource.Postgres.Master.DBName)
}
