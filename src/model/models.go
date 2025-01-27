package model

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
	Id                string `xorm:"varchar(255) pk notnull" json:"id"`
	ActiveProfileUUID string `xorm:"varchar(255) notnull active_profile_UUID" json:"active_profile_UUID"`
	DiscordSnowflake  string `xorm:"varchar(255) notnull" json:"discord_snowflake"`
	FetchData         bool   `xorm:"notnull" json:"fetch_data"`
}

// Use a different table name for this struct
func (User) TableName() string {
	return "users"
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
	Mining   GuildEventType = "mining"
)
