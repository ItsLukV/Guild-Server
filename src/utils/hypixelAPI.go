package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ItsLukV/Guild-Server/src/app"
)

// -----------------------------------------------
// -------------- Hypixel API types --------------
// -----------------------------------------------

type SkyblockPlayerData struct {
	Success bool                `json:"success"`
	Profile SkyblockProfileData `json:"profile"`
}

type SkyblockProfileData struct {
	ProfileId         string                               `json:"profile_id"`
	CommunityUpgrades interface{}                          `json:"community_upgrades"`
	Members           map[string]SkyblockProfileMemberData `json:"members"`
	Banking           interface{}                          `json:"banking"`
}

type SkyblockProfileMemberData struct {
	Rift                   interface{}              `json:"rift"`
	PlayerData             interface{}              `json:"player_data"`
	GlacitePlayerData      interface{}              `json:"glacite_player_data"`
	Events                 interface{}              `json:"events"`
	GardenPlayerData       interface{}              `json:"garden_player_data"`
	AccessoryBagStorage    interface{}              `json:"accessory_bag_storage"`
	Leveling               interface{}              `json:"leveling"`
	ItemData               interface{}              `json:"item_data"`
	JacobsContest          interface{}              `json:"jacobs_contest"`
	Currencies             interface{}              `json:"currencies"`
	Dungeons               SkyblockDungeons         `json:"dungeons"`
	Profile                interface{}              `json:"profile"`
	PetsData               interface{}              `json:"pets_data"`
	PlayerId               string                   `json:"player_id"`
	NetherIslandPlayerData interface{}              `json:"nether_island_player_data"`
	Experimentation        interface{}              `json:"experimentation"`
	MiningCore             interface{}              `json:"mining_core"`
	Bestiary               Bestiary                 `json:"bestiary"`
	Quests                 interface{}              `json:"quests"`
	PlayerStats            SkyblockPlayerStats      `json:"player_stats"`
	WinterPlayerData       interface{}              `json:"winter_player_data"`
	Forge                  interface{}              `json:"forge"`
	FairySoul              interface{}              `json:"fairy_soul"`
	Slayer                 SkyblockPlayerSlayerData `json:"slayer"`
	TrophyFish             interface{}              `json:"trophy_fish"`
	Objectives             interface{}              `json:"objectives"`
	Inventory              interface{}              `json:"inventory"`
	SharedInventory        interface{}              `json:"shared_inventory"`
	Collection             interface{}              `json:"collection"`
}

type SkyblockPlayerStats struct {
	Mythos SkyblockMythos `json:"mythos"`
}

type SkyblockPlayerSlayerData struct {
	SlayerQuest  interface{}                   `json:"slayer_quest"`
	SlayerBosses map[string]SkyblockSlayerBoss `json:"slayer_bosses"`
}

type SkyblockSlayerBoss struct {
	ClaimedLevels     map[string]bool `json:"claimed_levels"`
	BossKillsTier0    int             `json:"boss_kills_tier_0"`
	Xp                int             `json:"xp"`
	BossKillsTier1    int             `json:"boss_kills_tier_1"`
	BossKillsTier2    int             `json:"boss_kills_tier_2"`
	BossKillsTier3    int             `json:"boss_kills_tier_3"`
	BossKillsTier4    int             `json:"boss_kills_tier_4"`
	BossAttemptsTier0 int             `json:"boss_attempts_tier_0"`
	BossAttemptsTier1 int             `json:"boss_attempts_tier_1"`
	BossAttemptsTier2 int             `json:"boss_attempts_tier_2"`
	BossAttemptsTier3 int             `json:"boss_attempts_tier_3"`
	BossAttemptsTier4 int             `json:"boss_attempts_tier_4"`
}

type HypixelPlayerSocialMedia struct {
	Player SocialMediaPlayer `json:"player"`
}

type SocialMediaPlayer struct {
	SocialMedia SocialMediaLinks `json:"socialMedia"`
}

type SocialMediaLinks struct {
	Links Links `json:"links"`
}

type Links struct {
	Discord string `json:"DISCORD"`
}

type SkyblockProfilesByPlayerData struct {
	Success  bool      `json:"success"`
	Profiles []Profile `json:"profiles"`
}

type Profile struct {
	ProfileID         string                 `json:"profile_id"`
	CommunityUpgrades interface{}            `json:"community_upgrades"`
	Members           map[string]interface{} `json:"members"`
	CuteName          string                 `json:"cute_name"`
	Selected          bool                   `json:"selected"`
}

type Bestiary struct {
	Kills BestiaryKills `json:"kills"`
}

type BestiaryKills struct {
	GaiaConstruct   int `json:"gaia_construct_260"`
	MinosChampion   int `json:"minos_champion_310"`
	MinosHunter     int `json:"minos_hunter_125"`
	MinosInquisitor int `json:"minos_inquisitor_750"`
	Minotaur        int `json:"minotaur_210"`
	SiameseLynx     int `json:"siamese_lynx_155"`
}

type SkyblockMythos struct {
	Kills              float32                  `json:"kills"`
	BurrowsDugTreasure MythosBurrowsDugTreasure `json:"burrows_dug_treasure"`
	BurrowsDugCombat   MythosBurrowsDugCombat   `json:"burrows_dug_combat"`
}
type MythosBurrowsDugTreasure struct {
	TotalBurrows float32 `json:"total"`
	Legendary    float32 `json:"LEGENDARY"`
}
type MythosBurrowsDugCombat struct {
	TotalBurrows float32 `json:"total"`
	Legendary    float32 `json:"LEGENDARY"`
}

type SkyblockDungeons struct {
	DungeonTypes           map[string]DungeonTypes `json:"dungeon_types"`
	PlayerClasses          PlayerClasses           `json:"player_classes"`
	DungeonJournal         interface{}             `json:"dungeon_journal"`
	DungeonsBlahBlah       interface{}             `json:"dungeons_blah_blah"`
	SelectedDungeonClass   string                  `json:"selected_dungeon_class"`
	DailyRuns              interface{}             `json:"daily_runs"`
	Treasures              interface{}             `json:"treasures"`
	DungeonHubRaceSettings interface{}             `json:"dungeon_hub_race_settings"`
	LastDungeonRun         string                  `json:"last_dungeon_run"`
	Secrets                int                     `json:"secrets"`
}

type DungeonTypes struct {
	TimesPlayed          interface{}        `json:"times_played"`
	Experience           float64            `json:"experience"`
	TierCompletions      map[string]float32 `json:"tier_completions"`
	FastestTime          interface{}        `json:"fastest_time"`
	BestRuns             interface{}        `json:"best_runs"`
	BestScore            interface{}        `json:"best_score"`
	MobsKilled           interface{}        `json:"mobs_killed"`
	MostMobsKilled       interface{}        `json:"most_mobs_killed"`
	MostDamageBerserk    interface{}        `json:"most_damage_berserk"`
	MostHealing          interface{}        `json:"most_healing"`
	WatcherKills         interface{}        `json:"watcher_kills"`
	HighestTierCompleted int                `json:"highest_tier_completed"`
	FastestTimeS         interface{}        `json:"fastest_time_s"`
	MostDamageArcher     interface{}        `json:"most_damage_archer"`
	FastestTimeSPlus     interface{}        `json:"fastest_time_s_plus"`
	MostDamageMage       interface{}        `json:"most_damage_mage"`
	MilestoneCompletions interface{}        `json:"milestone_completions"`
	MostDamageHealer     interface{}        `json:"most_damage_healer"`
	MostDamageTank       interface{}        `json:"most_damage_tank"`
}

type PlayerClasses struct {
	Healer struct {
		Experience float64 `json:"experience"`
	} `json:"healer"`
	Mage struct {
		Experience float64 `json:"experience"`
	} `json:"mage"`
	Berserk struct {
		Experience float64 `json:"experience"`
	} `json:"berserk"`
	Archer struct {
		Experience float64 `json:"experience"`
	} `json:"archer"`
	Tank struct {
		Experience float64 `json:"experience"`
	} `json:"tank"`
}

// -----------------------------------------------
// ------------ Hypixel API Functions ------------
// -----------------------------------------------

func FetchPlayerData(uuid string, profile string) (*SkyblockPlayerData, error) {
	// Make the GET request
	key := os.Getenv("HYPIXEL_API")
	url := fmt.Sprintf("https://api.hypixel.net/v2/skyblock/profile?key=%s&uuid=%s&profile=%s", key, uuid, profile)
	return fetchApi[SkyblockPlayerData](url)
}

func FetchPlayerSlayerData(uuid string, profile string) (*SkyblockPlayerSlayerData, error) {
	data, err := FetchPlayerData(uuid, profile)
	if err != nil {
		return nil, err
	}

	member, exists := data.Profile.Members[uuid]
	if !exists {
		return nil, fmt.Errorf("player %s not found in profile", uuid)
	}

	return &member.Slayer, nil
}

func CheckUserName(minecraftUUID string) (*string, error) {
	key := os.Getenv("HYPIXEL_API")
	url := fmt.Sprintf("https://api.hypixel.net/v2/player?key=%v&uuid=%v", key, minecraftUUID)
	data, err := fetchApi[HypixelPlayerSocialMedia](url)
	if err != nil {
		return nil, err
	}
	return &data.Player.SocialMedia.Links.Discord, nil
}

func FetchActivePlayerProfile(minecraftUUID string) (string, error) {
	key := os.Getenv("HYPIXEL_API")
	url := fmt.Sprintf("https://api.hypixel.net/v2/skyblock/profiles?key=%v&uuid=%v", key, minecraftUUID)

	data, err := fetchApi[SkyblockProfilesByPlayerData](url)
	if err != nil {
		return "", err
	}
	// Call `getSelectedProfileID` to get the profile ID.
	profileID := getSelectedProfileID(data)
	if profileID == "" {
		return "", fmt.Errorf("no active profile found for player with UUID: %s", minecraftUUID)
	}

	return profileID, nil
}

func getSelectedProfileID(data *SkyblockProfilesByPlayerData) string {
	for _, profile := range data.Profiles {
		if profile.Selected {
			return profile.ProfileID
		}
	}
	// Return an empty string if no profile is selected
	return ""
}

// -----------------------------------------------
// ----------- Minecraft API functions -----------
// -----------------------------------------------

type McUUID struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// GetMCUUID Returns the uuid of a minecraft player
func GetMCUUID(ign string) (*McUUID, error) {
	url := fmt.Sprintf("https://api.mojang.com/users/profiles/minecraft/%s", ign)
	return fetchApi[McUUID](url)
}

// -----------------------------------------------
// ----------- Generic API functions -------------
// -----------------------------------------------

func fetchApi[T interface{}](url string) (*T, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check for a successful response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: Unable to fetch data. Status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the JSON data
	var data T
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func IntoDianaData(data SkyblockPlayerData, userId string, uuid string) app.DianaData {
	dianaData := app.DianaData{
		UserId:          userId,
		BurrowsTreasure: 0,
		BurrowsCombat:   0,
		GaiaConstruct:   0,
		MinosChampion:   0,
		MinosHunter:     0,
		MinosInquisitor: 0,
		Minotaur:        0,
		SiameseLynx:     0,
	}

	kills := data.Profile.Members[uuid].Bestiary.Kills

	dianaData.GaiaConstruct = kills.GaiaConstruct
	dianaData.MinosChampion = kills.MinosChampion
	dianaData.MinosHunter = kills.MinosHunter
	dianaData.MinosInquisitor = kills.MinosInquisitor
	dianaData.Minotaur = kills.Minotaur
	dianaData.SiameseLynx = kills.SiameseLynx

	burrowData := data.Profile.Members[uuid].PlayerStats.Mythos

	dianaData.BurrowsCombat = burrowData.BurrowsDugCombat.Legendary
	dianaData.BurrowsTreasure = burrowData.BurrowsDugTreasure.Legendary

	return dianaData
}

func IntoDungeonsData(data SkyblockPlayerData, userId string, uuid string) app.DungeonsData {
	dungeonApiData := data.Profile.Members[uuid].Dungeons
	playerClass := dungeonApiData.PlayerClasses
	return app.DungeonsData{
		UserId:     userId,
		Experience: dungeonApiData.DungeonTypes["catacombs"].Experience,
		Completions: map[string]float32{
			"0": dungeonApiData.DungeonTypes["catacombs"].TierCompletions["0"],
			"1": dungeonApiData.DungeonTypes["catacombs"].TierCompletions["1"],
			"2": dungeonApiData.DungeonTypes["catacombs"].TierCompletions["2"],
			"3": dungeonApiData.DungeonTypes["catacombs"].TierCompletions["3"],
			"4": dungeonApiData.DungeonTypes["catacombs"].TierCompletions["4"],
			"5": dungeonApiData.DungeonTypes["catacombs"].TierCompletions["5"],
			"6": dungeonApiData.DungeonTypes["catacombs"].TierCompletions["6"],
			"7": dungeonApiData.DungeonTypes["catacombs"].TierCompletions["7"],
		},
		MasterCompletions: map[string]float32{
			"1": dungeonApiData.DungeonTypes["master_catacombs"].TierCompletions["1"],
			"2": dungeonApiData.DungeonTypes["master_catacombs"].TierCompletions["2"],
			"3": dungeonApiData.DungeonTypes["master_catacombs"].TierCompletions["3"],
			"4": dungeonApiData.DungeonTypes["master_catacombs"].TierCompletions["4"],
			"5": dungeonApiData.DungeonTypes["master_catacombs"].TierCompletions["5"],
			"6": dungeonApiData.DungeonTypes["master_catacombs"].TierCompletions["6"],
			"7": dungeonApiData.DungeonTypes["master_catacombs"].TierCompletions["7"],
		},
		ClassXp: map[string]float64{
			"healer":  playerClass.Healer.Experience,
			"mage":    playerClass.Mage.Experience,
			"berserk": playerClass.Berserk.Experience,
			"archer":  playerClass.Archer.Experience,
			"tank":    playerClass.Tank.Experience,
		},
		Secrets: dungeonApiData.Secrets,
	}
}
