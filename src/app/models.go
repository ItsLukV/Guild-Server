package app

import (
	"time"

	"xorm.io/xorm"
)

type App struct {
	Engine *xorm.Engine
	Users  []User
}

func (app *App) GetToken(service string) (string, error) {
	var apiToken APIToken

	// Query the database using XORM's Where method
	has, err := app.Engine.Where("service_name = ?", service).Get(&apiToken)
	if err != nil {
		return "", err // Return error if query fails
	}
	if !has {
		return "", nil // Return nil if no token found for the service
	}

	return apiToken.Token, nil
}

type APIToken struct {
	ID          int64     `xorm:"pk autoincr"`
	ServiceName string    `xorm:"unique notnull"`
	Token       string    `xorm:"varchar(512)"`
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAt   time.Time `xorm:"updated"`
}

type User struct {
	Id                int    `xorm:"INT pk notnull autoincr" json:"id"`
	Name              string `xorm:"varchar(255) notnull" json:"name"`
	ActiveProfileUUID string `xorm:"varchar(255) notnull active_profile_UUID" json:"active_profile_UUID"`
}

// Use a different table name for this struct
func (User) TableName() string {
	return "users"
}

type DianaData struct {
	Id              int       `xorm:"INT pk notnull autoincr" json:"id"`
	UserId          int       `xorm:"index notnull" json:"user"`
	FetchTime       time.Time `xorm:"created notnull" json:"fetch_time"`
	BurrowsTreasure float32   `xorm:"DOUBLE notnull" json:"burrows_treasure"`
	BurrowsCombat   float32   `xorm:"DOUBLE notnull" json:"burrows_combat"`
	GaiaConstruct   int       `xorm:"INT notnull" json:"gaia_construct"`
	MinosChampion   int       `xorm:"INT notnull" json:"minos_champion"`
	MinosHunter     int       `xorm:"INT notnull" json:"minos_hunter"`
	MinosInquisitor int       `xorm:"INT notnull" json:"minos_inquisitor"`
	Minotaur        int       `xorm:"INT notnull" json:"minotaur"`
	SiameseLynx     int       `xorm:"INT notnull" json:"siamese_lynx"`
}

func (DianaData) TableName() string {
	return "diana_data"
}

type DungeonsData struct {
	Id                int                `xorm:"INT pk notnull autoincr" json:"id"`
	UserId            int                `xorm:"index notnull" json:"user"`
	FetchTime         time.Time          `xorm:"created notnull" json:"fetch_time"`
	Experience        float64            `xorm:"DOUBLE notnull" json:"experience"`
	Completions       map[string]float32 `xorm:"json notnull" json:"completions"`
	MasterCompletions map[string]float32 `xorm:"json notnull" json:"master_completions"`
	ClassXp           map[string]float64 `xorm:"json notnull" json:"class_xp"`
	Secrets           int                `xorm:"INT notnull" json:"secrets"`
}

func (DungeonsData) TableName() string {
	return "dungeons_data"
}

type GuildEvent struct {
	Id        int       `xorm:"INT pk notnull autoincr" json:"id"`
	Users     []int     `xorm:"'user_ids' json" json:""`
	StartTime time.Time `xorm:"created notnull" json:"start_time"`
	Duration  int       `xorm:"INT notnull" json:"duration"`
	Type      string    `xorm:"varchar(255) not null" json:"type"`
}

func (GuildEvent) TableName() string {
	return "guild_event"
}
