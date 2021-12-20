package database

import (
	"errors"
	"fmt"

	"github.com/noploy/noploy-back/repositories/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connector struct {
	DB       *gorm.DB
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

var ConnectionNotOpened = errors.New("connection not opened")

func (c *Connector) Open() error {
	var connectionUrl = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.Host, c.User, c.Password, c.DBName, c.Port)
	db, err := gorm.Open(postgres.Open(connectionUrl), &gorm.Config{})
	if err != nil {
		return err
	}
	c.DB = db
	return nil
}

func (c *Connector) AutoMigrate() error {
	if c.DB == nil {
		return ConnectionNotOpened
	}
	err := c.DB.AutoMigrate(entities.Entities)
	if err != nil {
		return err
	}
	return nil
}
