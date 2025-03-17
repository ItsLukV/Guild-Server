package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ItsLukV/Guild-Server/src/model"
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
	GlacitePlayerData      GlacitePlayerData        `json:"glacite_player_data"`
	Events                 interface{}              `json:"events"`
	GardenPlayerData       interface{}              `json:"garden_player_data"`
	AccessoryBagStorage    interface{}              `json:"accessory_bag_storage"`
	Leveling               Leveling                 `json:"leveling"`
	ItemData               interface{}              `json:"item_data"`
	JacobsContest          interface{}              `json:"jacobs_contest"`
	Currencies             interface{}              `json:"currencies"`
	Dungeons               SkyblockDungeons         `json:"dungeons"`
	Profile                interface{}              `json:"profile"`
	PetsData               interface{}              `json:"pets_data"`
	PlayerId               string                   `json:"player_id"`
	NetherIslandPlayerData interface{}              `json:"nether_island_player_data"`
	Experimentation        interface{}              `json:"experimentation"`
	MiningCore             MiningCore               `json:"mining_core"`
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
	Collection             Collection               `json:"collection"`
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
	WormKills       int `json:"worm_5"`
	ScathaKills     int `json:"scatha_10"`
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

type GlacitePlayerData struct {
	FossilsDonated []string `json:"fossils_donated"`
	FossilDust     float64  `json:"fossil_dust"`
	CorpsesLooted  struct {
		Tungsten int `json:"tungsten"`
		Umber    int `json:"umber"`
		Lapis    int `json:"lapis"`
		Vanguard int `json:"vanguard"`
	} `json:"corpses_looted"`
	MineshaftsEntered int `json:"mineshafts_entered"`
}

type MiningCore struct {
	Nodes struct {
		Special0             int `json:"special_0"`
		MiningMadness        int `json:"mining_madness"`
		EfficientMiner       int `json:"efficient_miner"`
		MiningFortune        int `json:"mining_fortune"`
		MiningSpeed          int `json:"mining_speed"`
		Mole                 int `json:"mole"`
		Professional         int `json:"professional"`
		PickaxeToss          int `json:"pickaxe_toss"`
		ForgeTime            int `json:"forge_time"`
		Fortunate            int `json:"fortunate"`
		FrontLoaded          int `json:"front_loaded"`
		MiningExperience     int `json:"mining_experience"`
		Crystalline          int `json:"crystalline"`
		EagerAdventurer      int `json:"eager_adventurer"`
		GiftsFromTheDeparted int `json:"gifts_from_the_departed"`
		HungryForMore        int `json:"hungry_for_more"`
		KeepItCool           int `json:"keep_it_cool"`
		MiningFortune2       int `json:"mining_fortune_2"`
		MiningMaster         int `json:"mining_master"`
		MiningSpeed2         int `json:"mining_speed_2"`
		NoStoneUnturned      int `json:"no_stone_unturned"`
		PowderBuff           int `json:"powder_buff"`
		RagsToRiches         int `json:"rags_to_riches"`
		SteadyHand           int `json:"steady_hand"`
		StrongArm            int `json:"strong_arm"`
		Surveyor             int `json:"surveyor"`
		WarmHearted          int `json:"warm_hearted"`
	} `json:"nodes"`
	ReceivedFreeTier            bool    `json:"received_free_tier"`
	Tokens                      int     `json:"tokens"`
	TokensSpent                 int     `json:"tokens_spent"`
	PowderMithril               int     `json:"powder_mithril"`
	PowderMithrilTotal          int     `json:"powder_mithril_total"`
	PowderSpentMithril          int     `json:"powder_spent_mithril"`
	Experience                  float64 `json:"experience"`
	RetroactiveTier2Token       bool    `json:"retroactive_tier2_token"`
	DailyOresMinedDayMithrilOre int     `json:"daily_ores_mined_day_mithril_ore"`
	DailyOresMinedMithrilOre    int     `json:"daily_ores_mined_mithril_ore"`
	LastReset                   int64   `json:"last_reset"`
	GreaterMinesLastAccess      int64   `json:"greater_mines_last_access"`
	Crystals                    struct {
		JadeCrystal struct {
			State       string `json:"state"`
			TotalPlaced int    `json:"total_placed"`
			TotalFound  int    `json:"total_found"`
		} `json:"jade_crystal"`
		AmberCrystal struct {
			State       string `json:"state"`
			TotalPlaced int    `json:"total_placed"`
			TotalFound  int    `json:"total_found"`
		} `json:"amber_crystal"`
		TopazCrystal struct {
			State       string `json:"state"`
			TotalPlaced int    `json:"total_placed"`
			TotalFound  int    `json:"total_found"`
		} `json:"topaz_crystal"`
		SapphireCrystal struct {
			State       string `json:"state"`
			TotalPlaced int    `json:"total_placed"`
			TotalFound  int    `json:"total_found"`
		} `json:"sapphire_crystal"`
		AmethystCrystal struct {
			State       string `json:"state"`
			TotalPlaced int    `json:"total_placed"`
			TotalFound  int    `json:"total_found"`
		} `json:"amethyst_crystal"`
		JasperCrystal struct {
			State      string `json:"state"`
			TotalFound int    `json:"total_found"`
		} `json:"jasper_crystal"`
		RubyCrystal struct {
			State      string `json:"state"`
			TotalFound int    `json:"total_found"`
		} `json:"ruby_crystal"`
		CitrineCrystal struct {
			State      string `json:"state"`
			TotalFound int    `json:"total_found"`
		} `json:"citrine_crystal"`
		OpalCrystal struct {
			State      string `json:"state"`
			TotalFound int    `json:"total_found"`
		} `json:"opal_crystal"`
		PeridotCrystal struct {
			State      string `json:"state"`
			TotalFound int    `json:"total_found"`
		} `json:"peridot_crystal"`
		AquamarineCrystal struct {
			State      string `json:"state"`
			TotalFound int    `json:"total_found"`
		} `json:"aquamarine_crystal"`
		OnyxCrystal struct {
			State      string `json:"state"`
			TotalFound int    `json:"total_found"`
		} `json:"onyx_crystal"`
	} `json:"crystals"`
	Biomes struct {
		Dwarven struct {
		} `json:"dwarven"`
		Precursor struct {
			ClaimingWithPrecursorApparatus bool `json:"claiming_with_precursor_apparatus"`
		} `json:"precursor"`
		Goblin struct {
			KingQuestActive     bool `json:"king_quest_active"`
			KingQuestsCompleted int  `json:"king_quests_completed"`
		} `json:"goblin"`
		Jungle struct {
			JungleTempleOpen      bool `json:"jungle_temple_open"`
			JungleTempleChestUses int  `json:"jungle_temple_chest_uses"`
		} `json:"jungle"`
	} `json:"biomes"`
	PowderGemstone            int    `json:"powder_gemstone"`
	PowderGemstoneTotal       int    `json:"powder_gemstone_total"`
	PowderSpentGemstone       int    `json:"powder_spent_gemstone"`
	DailyOresMinedDayGemstone int    `json:"daily_ores_mined_day_gemstone"`
	DailyOresMinedGemstone    int    `json:"daily_ores_mined_gemstone"`
	DailyOresMinedDayGlacite  int    `json:"daily_ores_mined_day_glacite"`
	DailyOresMinedGlacite     int    `json:"daily_ores_mined_glacite"`
	PowderGlacite             int    `json:"powder_glacite"`
	PowderGlaciteTotal        int    `json:"powder_glacite_total"`
	PowderSpentGlacite        int    `json:"powder_spent_glacite"`
	DailyOresMined            int    `json:"daily_ores_mined"`
	DailyOresMinedDay         int    `json:"daily_ores_mined_day"`
	SelectedPickaxeAbility    string `json:"selected_pickaxe_ability"`
}
type Collection struct {
	Cobblestone          int `json:"COBBLESTONE"`
	RottenFlesh          int `json:"ROTTEN_FLESH"`
	Wheat                int `json:"WHEAT"`
	Seeds                int `json:"SEEDS"`
	CarrotItem           int `json:"CARROT_ITEM"`
	PotatoItem           int `json:"POTATO_ITEM"`
	Log                  int `json:"LOG"`
	LOG2                 int `json:"LOG:2"`
	LOG1                 int `json:"LOG:1"`
	LOG21                int `json:"LOG_2:1"`
	Pumpkin              int `json:"PUMPKIN"`
	Melon                int `json:"MELON"`
	Pork                 int `json:"PORK"`
	Feather              int `json:"FEATHER"`
	RawChicken           int `json:"RAW_CHICKEN"`
	Leather              int `json:"LEATHER"`
	Cactus               int `json:"CACTUS"`
	MushroomCollection   int `json:"MUSHROOM_COLLECTION"`
	Coal                 int `json:"COAL"`
	INKSACK3             int `json:"INK_SACK:3"`
	SugarCane            int `json:"SUGAR_CANE"`
	Rabbit               int `json:"RABBIT"`
	Mutton               int `json:"MUTTON"`
	IronIngot            int `json:"IRON_INGOT"`
	GoldIngot            int `json:"GOLD_INGOT"`
	INKSACK4             int `json:"INK_SACK:4"`
	Redstone             int `json:"REDSTONE"`
	Emerald              int `json:"EMERALD"`
	SlimeBall            int `json:"SLIME_BALL"`
	Diamond              int `json:"DIAMOND"`
	Obsidian             int `json:"OBSIDIAN"`
	NetherStalk          int `json:"NETHER_STALK"`
	String               int `json:"STRING"`
	Bone                 int `json:"BONE"`
	ClayBall             int `json:"CLAY_BALL"`
	RawFish              int `json:"RAW_FISH"`
	InkSack              int `json:"INK_SACK"`
	WaterLily            int `json:"WATER_LILY"`
	PrismarineShard      int `json:"PRISMARINE_SHARD"`
	PrismarineCrystals   int `json:"PRISMARINE_CRYSTALS"`
	SpiderEye            int `json:"SPIDER_EYE"`
	EnchantedRedstone    int `json:"ENCHANTED_REDSTONE"`
	RAWFISH2             int `json:"RAW_FISH:2"`
	RAWFISH1             int `json:"RAW_FISH:1"`
	RAWFISH3             int `json:"RAW_FISH:3"`
	Sponge               int `json:"SPONGE"`
	EnderPearl           int `json:"ENDER_PEARL"`
	EnderStone           int `json:"ENDER_STONE"`
	Sand                 int `json:"SAND"`
	MithrilOre           int `json:"MITHRIL_ORE"`
	LOG3                 int `json:"LOG:3"`
	Log2                 int `json:"LOG_2"`
	Sulphur              int `json:"SULPHUR"`
	GemstoneCollection   int `json:"GEMSTONE_COLLECTION"`
	HardStone            int `json:"HARD_STONE"`
	Quartz               int `json:"QUARTZ"`
	Ice                  int `json:"ICE"`
	BlazeRod             int `json:"BLAZE_ROD"`
	MagmaCream           int `json:"MAGMA_CREAM"`
	SulphurOre           int `json:"SULPHUR_ORE"`
	GlowstoneDust        int `json:"GLOWSTONE_DUST"`
	Gravel               int `json:"GRAVEL"`
	EnchantedLapisLazuli int `json:"ENCHANTED_LAPIS_LAZULI"`
	EnchantedDarkOakLog  int `json:"ENCHANTED_DARK_OAK_LOG"`
	RawBeef              int `json:"RAW_BEEF"`
	Wool                 int `json:"WOOL"`
	EnchantedBone        int `json:"ENCHANTED_BONE"`
	Netherrack           int `json:"NETHERRACK"`
	Mycel                int `json:"MYCEL"`
	SAND1                int `json:"SAND:1"`
	GhastTear            int `json:"GHAST_TEAR"`
	MagmaFish            int `json:"MAGMA_FISH"`
	WiltedBerberis       int `json:"WILTED_BERBERIS"`
	CaducousStem         int `json:"CADUCOUS_STEM"`
	AgaricusCap          int `json:"AGARICUS_CAP"`
	MetalHeart           int `json:"METAL_HEART"`
	HalfEatenCarrot      int `json:"HALF_EATEN_CARROT"`
	Hemovibe             int `json:"HEMOVIBE"`
	Umber                int `json:"UMBER"`
	Glacite              int `json:"GLACITE"`
	Tungsten             int `json:"TUNGSTEN"`
	Timite               int `json:"TIMITE"`
	ChiliPepper          int `json:"CHILI_PEPPER"`
}

type Leveling struct {
	Experience  int `json:"leveling"`
	Completions struct {
		NucleusRuns int `json:"NUCLEUS_RUNS"`
	} `json:"completions"`
	CompletedTasks              []string `json:"completed_tasks"`
	Migrated                    bool     `json:"migrated"`
	HighestPetScore             int      `json:"hightest_pet_score"`
	CategoryExpanded            bool     `json:"category_expanded"`
	MiningFiestaOresMined       int      `json:"mining_fiesta_ores_mined"`
	ClaimedTalisman             bool     `json:"claimed_talisman"`
	BopBonus                    string   `json:"bop_bonus"`
	FishingFestivalSharksKilled int      `json:"fishing_festival_sharks_killed"`
	EmblemUnlocks               []string `json:"emblem_unlocks"`
	TaskSort                    string   `json:"task_sort"`
	LastViewedTasks             []string `json:"last_viewed_tasks"`
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
	url := fmt.Sprintf("https://api.minecraftservices.com/minecraft/profile/lookup/name/%s", ign)
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

func IntoDianaData(data SkyblockPlayerData, userId string) model.DianaData {
	dianaData := model.DianaData{
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

	kills := data.Profile.Members[userId].Bestiary.Kills

	dianaData.GaiaConstruct = kills.GaiaConstruct
	dianaData.MinosChampion = kills.MinosChampion
	dianaData.MinosHunter = kills.MinosHunter
	dianaData.MinosInquisitor = kills.MinosInquisitor
	dianaData.Minotaur = kills.Minotaur
	dianaData.SiameseLynx = kills.SiameseLynx

	burrowData := data.Profile.Members[userId].PlayerStats.Mythos

	dianaData.BurrowsCombat = burrowData.BurrowsDugCombat.Legendary
	dianaData.BurrowsTreasure = burrowData.BurrowsDugTreasure.Legendary

	return dianaData
}

func IntoDungeonsData(data SkyblockPlayerData, userId string) model.DungeonsData {
	dungeonApiData := data.Profile.Members[userId].Dungeons
	playerClass := dungeonApiData.PlayerClasses
	return model.DungeonsData{
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

func IntoMiningData(data SkyblockPlayerData, userId string) model.MiningData {
	miningApiData := data.Profile.Members[userId]
	glaciteData := miningApiData.GlacitePlayerData
	miningCore := miningApiData.MiningCore
	collection := miningApiData.Collection
	return model.MiningData{
		UserId:         userId,
		Mineshaft:      glaciteData.MineshaftsEntered,
		FossilDust:     float32(glaciteData.FossilDust),
		TungstenCorpse: glaciteData.CorpsesLooted.Tungsten,
		UmberCorpse:    glaciteData.CorpsesLooted.Umber,
		LapisCorpse:    glaciteData.CorpsesLooted.Lapis,
		VanguardCorpse: glaciteData.CorpsesLooted.Vanguard,
		NucleusRuns:    data.Profile.Members[userId].Leveling.Completions.NucleusRuns,
		MithrilPowder:  miningCore.PowderMithrilTotal,
		GemstonePowder: miningCore.PowderGemstoneTotal,
		GlacitePowder:  miningCore.PowderGlaciteTotal,
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
			Mithril:     collection.MithrilOre,
			Gemstone:    collection.GemstoneCollection,
			Gold_Ingot:  collection.GoldIngot,
			Netherrack:  collection.Netherrack,
			Diamond:     collection.Diamond,
			Ice:         collection.Ice,
			Redstone:    collection.Redstone,
			Lapis:       collection.INKSACK4,
			Sulphur:     collection.SulphurOre,
			Coal:        collection.Coal,
			Emerald:     collection.Emerald,
			EndStone:    collection.EnderStone,
			Glowstone:   collection.GlowstoneDust,
			Gravel:      collection.Gravel,
			IronIngot:   collection.IronIngot,
			Mycelium:    collection.Mycel,
			Quartz:      collection.Quartz,
			Obsidian:    collection.Obsidian,
			RedSand:     collection.SAND1,
			Sand:        collection.SAND1,
			Cobblestone: collection.Cobblestone,
			HardStone:   collection.HardStone,
			MetalHeart:  collection.MetalHeart,
			Glacite:     collection.Glacite,
			Umber:       collection.Umber,
			Tungsten:    collection.Tungsten,
		},
		ScathaKills: data.Profile.Members[userId].Bestiary.Kills.ScathaKills,
		WormKills:   data.Profile.Members[userId].Bestiary.Kills.WormKills,
	}
}
