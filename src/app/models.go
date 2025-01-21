package app

import (
	"fmt"
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
	Id                string `xorm:"varchar(255) pk notnull" json:"id"`
	ActiveProfileUUID string `xorm:"varchar(255) notnull active_profile_UUID" json:"active_profile_UUID"`
	DiscordSnowflake  string `xorm:"varchar(255) notnull" json:"discord_snowflake"`
	FetchData         bool   `xorm:"notnull" json:"fetch_data"`
}

// Use a different table name for this struct
func (User) TableName() string {
	return "users"
}

type GuildEventData interface {
	GetUserID() string
	TableName() string
	Subtract(other GuildEventData) (GuildEventData, error)
}

type DianaData struct {
	UserId          string    `xorm:"index notnull" json:"id"`
	FetchTime       time.Time `xorm:"notnull" json:"fetch_time"`
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

func (d DianaData) GetUserID() string {
	return d.UserId
}

func (d DianaData) Subtract(other GuildEventData) (GuildEventData, error) {
	otherData, ok := other.(DianaData)
	if !ok {
		return nil, fmt.Errorf("cannot subtract different types")
	}
	d.BurrowsTreasure -= otherData.BurrowsTreasure
	d.BurrowsCombat -= otherData.BurrowsCombat
	d.GaiaConstruct -= otherData.GaiaConstruct
	d.MinosChampion -= otherData.MinosChampion
	d.MinosHunter -= otherData.MinosHunter
	d.MinosInquisitor -= otherData.MinosInquisitor
	d.Minotaur -= otherData.Minotaur
	d.SiameseLynx -= otherData.SiameseLynx
	return d, nil
}

type DungeonsData struct {
	UserId            string             `xorm:"index notnull" json:"id"`
	FetchTime         time.Time          `xorm:"notnull" json:"fetch_time"`
	Experience        float64            `xorm:"DOUBLE notnull" json:"experience"`
	Completions       map[string]float32 `xorm:"json notnull" json:"completions"`
	MasterCompletions map[string]float32 `xorm:"json notnull" json:"master_completions"`
	ClassXp           map[string]float64 `xorm:"json notnull" json:"class_xp"`
	Secrets           int                `xorm:"INT notnull" json:"secrets"`
}

func (DungeonsData) TableName() string {
	return "dungeons_data"
}

func (d DungeonsData) GetUserID() string {
	return d.UserId
}

func (d DungeonsData) Subtract(other GuildEventData) (GuildEventData, error) {
	otherData, ok := other.(DungeonsData)
	if !ok {
		return nil, fmt.Errorf("cannot subtract different types")
	}
	d.Experience -= otherData.Experience
	d.Secrets -= otherData.Secrets

	for key, value := range otherData.Completions {
		d.Completions[key] -= value
	}

	for key, value := range otherData.MasterCompletions {
		d.MasterCompletions[key] -= value
	}

	for key, value := range otherData.ClassXp {
		d.ClassXp[key] -= value
	}

	return d, nil
}

type GuildEvent struct {
	Id        string         `xorm:"varchar(22) pk notnull" json:"id"`
	Users     []string       `xorm:"text[] user_ids" json:"users"`
	StartTime time.Time      `xorm:"notnull" json:"start_time"`
	Duration  int            `xorm:"INT notnull" json:"duration"`
	Type      GuildEventType `xorm:"varchar(255) not null" json:"type"`
	IsHidden  bool           `xorm:"notnull" json:"is_hidden"`
}

func (GuildEvent) TableName() string {
	return "guild_event"
}

type GuildEventType string

const (
	Dungeons GuildEventType = "dungeons"
	Diana    GuildEventType = "diana"
)
