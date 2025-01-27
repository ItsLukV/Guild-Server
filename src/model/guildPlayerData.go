package model

import (
	"fmt"
	"time"
)

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

type MiningData struct {
	UserId         string    `xorm:"index notnull" json:"id"`
	FetchTime      time.Time `xorm:"notnull" json:"fetch_time"`
	Mineshaft      int       `json:"mineshaft_count"`
	FossilDust     float32   `json:"fossil_dust"`
	TungstenCorpse int       `json:"tungsten_corpse"`
	UmberCorpse    int       `json:"umber_corpse"`
	LapisCorpse    int       `json:"lapis_corpse"`
	VanguardCorpse int       `json:"vanguard_corpse"`
	NucleusRuns    int       `json:"nucleus_runs"`
	MithrilPowder  int       `json:"mithril_powder"`
	GemstonePowder int       `json:"powder_gemstone"`
	GlacitePowder  int       `json:"glacite_powder"`
	Collections    struct {
		Mithril     int `json:"mithril"`
		Gemstone    int `json:"gemstone"`
		Gold_Ingot  int `json:"gold_Ingot"`
		Netherrack  int `json:"netherrack"`
		Diamond     int `json:"diamond"`
		Ice         int `json:"ice"`
		Redstone    int `json:"Redstone"`
		Lapis       int `json:"Lapis"`
		Sulphur     int `json:"sulphur"`
		Coal        int `json:"coal"`
		Emerald     int `json:"emerald"`
		EndStone    int `json:"end_stone"`
		Glowstone   int `json:"glowstone"`
		Gravel      int `json:"gravel"`
		IronIngot   int `json:"iron_ingot"`
		Mycelium    int `json:"mycelium"`
		Quartz      int `json:"quartz"`
		Obsidian    int `json:"Obsidian"`
		RedSand     int `json:"red_sand"`
		Sand        int `json:"sand"`
		Cobblestone int `json:"cobblestone"`
		HardStone   int `json:"hard_stone"`
		MetalHeart  int `json:"metal_heart"`
		Glacite     int `json:"glacite"`
		Umber       int `json:"umber"`
		Tungsten    int `json:"tungsten"`
	} `xorm:"json" json:"Collections"`
	ScathaKills int `json:"scatha_kills"`
	WormKills   int `json:"worm_kills"`
}

func (m MiningData) GetUserID() string {
	return m.UserId
}

func (m MiningData) TableName() string {
	return "mining_data"
}

func (m MiningData) Subtract(other GuildEventData) (GuildEventData, error) {
	otherData, ok := other.(MiningData)
	if !ok {
		return MiningData{}, fmt.Errorf("cannot subtract different types")
	}

	// Perform subtraction on all relevant fields
	result := MiningData{
		UserId:         m.UserId,
		Mineshaft:      m.Mineshaft - otherData.Mineshaft,
		FossilDust:     m.FossilDust - otherData.FossilDust,
		TungstenCorpse: m.TungstenCorpse - otherData.TungstenCorpse,
		UmberCorpse:    m.UmberCorpse - otherData.UmberCorpse,
		LapisCorpse:    m.LapisCorpse - otherData.LapisCorpse,
		VanguardCorpse: m.VanguardCorpse - otherData.VanguardCorpse,
		NucleusRuns:    m.NucleusRuns - otherData.NucleusRuns,
		MithrilPowder:  m.MithrilPowder - otherData.MithrilPowder,
		GemstonePowder: m.GemstonePowder - otherData.GemstonePowder,
		GlacitePowder:  m.GlacitePowder - otherData.GlacitePowder,
		ScathaKills:    m.ScathaKills - otherData.ScathaKills,
		WormKills:      m.WormKills - otherData.WormKills,
		Collections: struct {
			Mithril     int `json:"mithril"`
			Gemstone    int `json:"gemstone"`
			Gold_Ingot  int `json:"gold_Ingot"`
			Netherrack  int `json:"netherrack"`
			Diamond     int `json:"diamond"`
			Ice         int `json:"ice"`
			Redstone    int `json:"Redstone"`
			Lapis       int `json:"Lapis"`
			Sulphur     int `json:"sulphur"`
			Coal        int `json:"coal"`
			Emerald     int `json:"emerald"`
			EndStone    int `json:"end_stone"`
			Glowstone   int `json:"glowstone"`
			Gravel      int `json:"gravel"`
			IronIngot   int `json:"iron_ingot"`
			Mycelium    int `json:"mycelium"`
			Quartz      int `json:"quartz"`
			Obsidian    int `json:"Obsidian"`
			RedSand     int `json:"red_sand"`
			Sand        int `json:"sand"`
			Cobblestone int `json:"cobblestone"`
			HardStone   int `json:"hard_stone"`
			MetalHeart  int `json:"metal_heart"`
			Glacite     int `json:"glacite"`
			Umber       int `json:"umber"`
			Tungsten    int `json:"tungsten"`
		}{
			Mithril:     m.Collections.Mithril - otherData.Collections.Mithril,
			Gemstone:    m.Collections.Gemstone - otherData.Collections.Gemstone,
			Gold_Ingot:  m.Collections.Gold_Ingot - otherData.Collections.Gold_Ingot,
			Netherrack:  m.Collections.Netherrack - otherData.Collections.Netherrack,
			Diamond:     m.Collections.Diamond - otherData.Collections.Diamond,
			Ice:         m.Collections.Ice - otherData.Collections.Ice,
			Redstone:    m.Collections.Redstone - otherData.Collections.Redstone,
			Lapis:       m.Collections.Lapis - otherData.Collections.Lapis,
			Sulphur:     m.Collections.Sulphur - otherData.Collections.Sulphur,
			Coal:        m.Collections.Coal - otherData.Collections.Coal,
			Emerald:     m.Collections.Emerald - otherData.Collections.Emerald,
			EndStone:    m.Collections.EndStone - otherData.Collections.EndStone,
			Glowstone:   m.Collections.Glowstone - otherData.Collections.Glowstone,
			Gravel:      m.Collections.Gravel - otherData.Collections.Gravel,
			IronIngot:   m.Collections.IronIngot - otherData.Collections.IronIngot,
			Mycelium:    m.Collections.Mycelium - otherData.Collections.Mycelium,
			Quartz:      m.Collections.Quartz - otherData.Collections.Quartz,
			Obsidian:    m.Collections.Obsidian - otherData.Collections.Obsidian,
			RedSand:     m.Collections.RedSand - otherData.Collections.RedSand,
			Sand:        m.Collections.Sand - otherData.Collections.Sand,
			Cobblestone: m.Collections.Cobblestone - otherData.Collections.Cobblestone,
			HardStone:   m.Collections.HardStone - otherData.Collections.HardStone,
			MetalHeart:  m.Collections.MetalHeart - otherData.Collections.MetalHeart,
			Glacite:     m.Collections.Glacite - otherData.Collections.Glacite,
			Umber:       m.Collections.Umber - otherData.Collections.Umber,
			Tungsten:    m.Collections.Tungsten - otherData.Collections.Tungsten,
		},
	}

	return result, nil
}
