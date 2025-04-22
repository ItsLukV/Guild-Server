package app

import (
	"fmt"
	"log"
	"os"

	"github.com/ItsLukV/Guild-Server/src/model"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

func SyncDatabase(users *[]model.User) (*xorm.Engine, error) {

	// Initialize connection string
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	// Initialize XORM engine
	engine, err := xorm.NewEngine("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
		return nil, err
	}

	// Use snake_case mapper for table and column names
	engine.SetMapper(names.SnakeMapper{})

	// Sync the table structure
	if err := engine.Sync(
		new(model.APIToken),
		new(model.User),
		new(model.DianaData),
		new(model.DungeonsData),
		new(model.GuildEvent),
		new(model.MiningData),
	); err != nil {
		log.Fatalf("Failed to sync database schema: %v", err)
		return nil, err
	}

	// Fetch users from the database
	if err := engine.Find(users); err != nil {
		log.Fatalf("Failed to fetch users: %v", err)
		return nil, err
	}

	return engine, nil
}
