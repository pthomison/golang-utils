package utils

import (
	"fmt"

	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type DBClient struct {
	DB *gorm.DB

	dbEngine string

	sqliteFile string

	dbHost     string
	dbUser     string
	dbPort     string
	dbPassword string
	dbName     string
}

func (c *DBClient) RegisterFlags(cmd *cobra.Command) {

	cmd.PersistentFlags().StringVarP(&c.dbEngine, "db-engine", "", "sqlite", "")

	cmd.PersistentFlags().StringVarP(&c.sqliteFile, "sqlite-file", "", "gorm.db", "")

	cmd.PersistentFlags().StringVarP(&c.dbHost, "postgres-host", "", "", "")
	cmd.PersistentFlags().StringVarP(&c.dbUser, "postgres-user", "", "postgres", "")
	cmd.PersistentFlags().StringVarP(&c.dbPort, "postgres-port", "", "5432", "")
	cmd.PersistentFlags().StringVarP(&c.dbPassword, "postgres-password", "", "", "")
	cmd.PersistentFlags().StringVarP(&c.dbName, "postgres-name", "", "postgres", "")

}

func (c *DBClient) InitializeClient(l logger.LogLevel) {
	c.Connect(l)
}

func (c *DBClient) Connect(l logger.LogLevel) {

	var db *gorm.DB
	var err error

	if c.dbEngine == "postgres" {
		psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.dbHost, c.dbPort, c.dbUser, c.dbPassword, c.dbName)

		db, err = gorm.Open(postgres.Open(psqlconn), &gorm.Config{
			Logger: logger.Default.LogMode(l),
		})
		Check(err)
	} else if c.dbEngine == "sqlite" {
		db, err = gorm.Open(sqlite.Open(c.sqliteFile), &gorm.Config{})
		Check(err)
	}

	c.DB = db
}

func SelectAll[T any](c *DBClient, columns []string) []T {
	var output []T
	result := c.DB.Select(columns).Find(&output)
	Check(result.Error)

	return output
}

func SelectWhere[T any](c *DBClient, columns []string, whereQuery interface{}, whereArgs ...interface{}) []T {
	var output []T
	result := c.DB.Where(whereQuery, whereArgs).Select(columns).Find(&output)
	Check(result.Error)

	return output
}

func Create[T any](c *DBClient, objs []T) {
	result := c.DB.Create(objs)
	Check(result.Error)
}

func CreateOrOverwrite[T any](c *DBClient, objs []T) {
	result := c.DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(objs)
	Check(result.Error)
}
