package utils

import (
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type SQLClient struct {
	DB *gorm.DB

	// dbHost     string
	// dbUser     string
	// dbPort     string
	// dbPassword string
	// dbName     string
}

func (c *SQLClient) RegisterFlags(cmd *cobra.Command) {
	// cmd.PersistentFlags().StringVarP(&c.dbHost, "db-host", "", "", "")
	// cmd.PersistentFlags().StringVarP(&c.dbUser, "db-user", "", "postgres", "")
	// cmd.PersistentFlags().StringVarP(&c.dbPort, "db-port", "", "5432", "")
	// cmd.PersistentFlags().StringVarP(&c.dbPassword, "db-password", "", "", "")
	// cmd.PersistentFlags().StringVarP(&c.dbName, "db-name", "", "postgres", "")

	// cmd.MarkPersistentFlagRequired("db-host")
	// cmd.MarkPersistentFlagRequired("db-password")
}

func (c *SQLClient) InitializeClient(l logger.LogLevel) {
	c.Connect(l)
}

func (c *SQLClient) Connect(l logger.LogLevel) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: logger.Default.LogMode(l),
	})
	Check(err)

	c.DB = db
}

func SelectAll[T any](c *SQLClient, columns []string) []T {
	var output []T
	result := c.DB.Select(columns).Find(&output)
	Check(result.Error)

	return output
}

func SelectWhere[T any](c *SQLClient, columns []string, whereQuery interface{}, whereArgs ...interface{}) []T {
	var output []T
	result := c.DB.Where(whereQuery, whereArgs).Select(columns).Find(&output)
	Check(result.Error)

	return output
}

func Create[T any](c *SQLClient, objs []T) {
	result := c.DB.Create(objs)
	Check(result.Error)
}

func CreateOrOverwrite[T any](c *SQLClient, objs []T) {
	result := c.DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(objs)
	Check(result.Error)
}
