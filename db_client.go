package utils

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func SelectAll_All[T any](c *Client) []T {
	var output []T
	result := c.DBClient.DB.Find(&output)
	Check(result.Error)

	return output
}

func SelectID_All[T any](c *Client) []T {
	var output []T
	result := c.DBClient.DB.Select("id").Find(&output)
	Check(result.Error)

	return output
}

func SelectAll_ByID[T any](c *Client, ids []uint64) []T {
	var output []T
	result := c.DBClient.DB.Where(ids).Find(&output)
	Check(result.Error)

	return output
}

func SelectAll_IDSpan[T any](c *Client, start uint64, end uint64) []T {
	var output []T
	result := c.DBClient.DB.Where("id >= ? AND id < ?", start, end).Find(&output)
	Check(result.Error)

	return output
}

func SelectAll_TimestampSpan[T any](c *Client, start time.Time, end time.Time) []T {
	var output []T
	result := c.DBClient.DB.Where("timestamp >= ? AND timestamp < ?", start, end).Order("timestamp asc").Find(&output)
	Check(result.Error)

	return output
}

func SelectID_IDSpan[T any](c *Client, start uint64, end uint64) []T {
	var output []T
	result := c.DBClient.DB.Where("id >= ? AND id < ?", start, end).Select("id").Find(&output)
	Check(result.Error)

	return output
}
