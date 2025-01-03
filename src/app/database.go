package app

import (
	"fmt"
	"log"
	"os"

	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

func SyncDatabase(users *[]User) (*xorm.Engine, error) {

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
	if err := engine.Sync(new(User), new(DianaData), new(DungeonsData), new(GuildEvent)); err != nil {
		log.Fatalf("Failed to sync database schema: %v", err)
		return nil, err
	}

	// Add foreign key only if it doesn't already exist
	constraintCheckQuery := `
	SELECT constraint_name
	FROM information_schema.table_constraints
	WHERE table_name = 'diana_data' AND constraint_name = 'fk_user';`
	existingConstraint := make([]string, 0)

	if err := engine.SQL(constraintCheckQuery).Find(&existingConstraint); err != nil {
		log.Printf("Failed to check existing constraints: %v", err)
		return nil, err
	}

	if len(existingConstraint) == 0 {
		_, err = engine.Exec(`
		ALTER TABLE diana_data
		ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
		ALTER TABLE dungeons_data
		ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
		`)
		if err != nil {
			log.Printf("Failed to add foreign key constraint: %v", err)
			return nil, err
		}
	}

	// Fetch users from the database
	if err := engine.Find(users); err != nil {
		log.Fatalf("Failed to fetch users: %v", err)
		return nil, err
	}

	return engine, nil
}
