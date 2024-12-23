package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ItsLukV/Guild-Server/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/exp/rand"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var engine *xorm.Engine

// Test users
var users []User = []User{
	{
		Name: "LukV",
	},
	{
		Name: "flyingshepfan_69",
	},
	{
		Name: "22um",
	},
	{
		Name: "stepjeppe",
	},
	{
		Name: "emilmz",
	},
}

func main() {

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

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
	var err error
	engine, err = xorm.NewEngine("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer engine.Close()

	// Use snake_case mapper for table and column names
	engine.SetMapper(names.SnakeMapper{})

	// Sync the `test` struct with the database schema
	err = engine.Sync(new(User))
	if err != nil {
		log.Fatalf("Failed to sync database schema: %v", err)
	}
	err = engine.Sync(new(DianaData))
	if err != nil {
		log.Fatalf("Failed to sync database schema: %v", err)
	}

	userdata := fetchData()
	_, err = engine.Insert(userdata)
	if err != nil {
		log.Fatalf("Failed to insert into database: %v", err)
	}
	log.Println("Record inserted successfully!")

	// Start a background Goroutine to fetch data every hour
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				log.Println("Fetching data")
				userdata = fetchData()
				_, err = engine.Insert(userdata)
				if err != nil {
					log.Fatalf("Failed to insert into database: %v", err)
				}
				log.Println("Record inserted successfully!")

			}
		}
	}()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	words := []string{"It's me mario", "Meow", "LukV", "LukV", "LukV", "LukV", "LukV"}

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": words[rand.Intn(len(words))],
		})
	})

	router.Run(":8080")

}

type User struct {
	Name string `xorm:"varchar(255) pk notnull" json:"name"`
}

type DianaData struct {
	id              string    `xorm:"varchar(255) pk notnull" json:"id"`
	User            string    `xorm:"index notnull" json:"user"`
	FetchTime       time.Time `xorm:"created notnull" json:"fetch_time"`
	BurrowsTreasure float32   `xorm:"INT notnull" json:"burrows_treasure"`
	BurrowsCombat   float32   `xorm:"INT notnull" json:"burrows_combat"`
	GaiaConstruct   int       `xorm:"INT notnull" json:"gaia_construct"`
	MinosChampion   int       `xorm:"INT notnull" json:"minos_champion"`
	MinosHunter     int       `xorm:"INT notnull" json:"minos_hunter"`
	MinosInquisitor int       `xorm:"INT notnull" json:"minos_inquisitor"`
	Minotaur        int       `xorm:"INT notnull" json:"minotaur"`
	SiameseLynx     int       `xorm:"INT notnull" json:"siamese_lynx"`
}

func fetchData() []DianaData {
	out := make([]DianaData, 0)
	for _, user := range users {
		uuid, err := utils.GetMCUUID(user.Name)
		if err != nil {
			log.Println("Failed to fetch api: ", err)
		}

		profile, err := utils.FetchActivePlayerProfile(uuid.Id)
		if err != nil {
			log.Println("Failed to fetch api: ", err)
		}

		data, err := utils.FetchPlayerData(uuid.Id, profile)
		if err != nil {
			log.Println("Failed to fetch api: ", err)
		}
		dianaData := DianaData{
			User:            user.Name,
			BurrowsTreasure: 0,
			BurrowsCombat:   0,
			GaiaConstruct:   0,
			MinosChampion:   0,
			MinosHunter:     0,
			MinosInquisitor: 0,
			Minotaur:        0,
			SiameseLynx:     0,
		}
		BestiaryKillsToDianaData(&dianaData, data.Profile.Members[uuid.Id].Bestiary.Kills)
		addMythData(&dianaData, data.Profile.Members[uuid.Id].PlayerStats.Mythos)
		out = append(out, dianaData)
	}
	return out
}

func BestiaryKillsToDianaData(dianaData *DianaData, kills utils.BestiaryKills) {

	dianaData.GaiaConstruct = kills.GaiaConstruct
	dianaData.MinosChampion = kills.MinosChampion
	dianaData.MinosHunter = kills.MinosHunter
	dianaData.MinosInquisitor = kills.MinosInquisitor
	dianaData.Minotaur = kills.Minotaur
	dianaData.SiameseLynx = kills.SiameseLynx
}

func addMythData(dianaData *DianaData, data utils.SkyblockMythos) {
	dianaData.BurrowsCombat = data.BurrowsDugCombat.Legendary
	dianaData.BurrowsTreasure = data.BurrowsDugTreasure.Legendary
}
