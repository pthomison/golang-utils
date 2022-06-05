package utils

import (
	"fmt"

	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type DBClient struct {
	DB *gorm.DB

	dbHost     string
	dbUser     string
	dbPort     string
	dbPassword string
	dbName     string
}

func (c *DBClient) RegisterFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&c.dbHost, "db-host", "", "", "")
	cmd.PersistentFlags().StringVarP(&c.dbUser, "db-user", "", "postgres", "")
	cmd.PersistentFlags().StringVarP(&c.dbPort, "db-port", "", "5432", "")
	cmd.PersistentFlags().StringVarP(&c.dbPassword, "db-password", "", "", "")
	cmd.PersistentFlags().StringVarP(&c.dbName, "db-name", "", "postgres", "")

	cmd.MarkPersistentFlagRequired("db-host")
	cmd.MarkPersistentFlagRequired("db-password")
}

func (c *DBClient) InitializeClient() {
	c.Connect()
}

func (c *DBClient) Connect() {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.dbHost, c.dbPort, c.dbUser, c.dbPassword, c.dbName)

	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	Check(err)

	c.DB = db
}

func SelectAll[T any](c *DBClient, columns []string) []T {
	var output []T
	result := c.DB.Select(columns).Find(&output)
	Check(result.Error)

	return output
}

func Create[T any](c *DBClient, objs []T) {
	result := dbc.DB.Create(objs)
	Check(result.Error)
}

func CreateOrOverwrite[T any](c *DBClient, objs []T) {
	result := dbc.DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(objs)
	Check(result.Error)
}
