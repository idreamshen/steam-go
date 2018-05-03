package steam

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// Retrieves the global achievement percentages for the specified app.
func GetGlobalAchievementPercentagesForApp(appID int) (percentages GlobalAchievementPercentages, bytes []byte, err error) {

	options := url.Values{}
	options.Set("gameid", strconv.Itoa(appID))

	bytes, err = get("ISteamUserStats/GetGlobalAchievementPercentagesForApp/v2", options)
	if err != nil {
		return percentages, bytes, err
	}

	var resp GlobalAchievementPercentagesResponse
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return percentages, bytes, err
	}

	return resp.GlobalAchievementPercentagesOuter, bytes, nil
}

type GlobalAchievementPercentagesResponse struct {
	GlobalAchievementPercentagesOuter GlobalAchievementPercentages `json:"achievementpercentages"`
}

type GlobalAchievementPercentages struct {
	GlobalAchievementPercentage []AchievementPercentage `json:"achievements"`
}

type AchievementPercentage struct {
	Name    string  `json:"name"`
	Percent float64 `json:"percent"`
}

// Gets the total number of players currently active in the specified app on Steam.
func GetNumberOfCurrentPlayers(appID int) (players int, bytes []byte, err error) {

	options := url.Values{}
	options.Set("appid", strconv.Itoa(appID))

	bytes, err = get("ISteamUserStats/GetNumberOfCurrentPlayers/v1", options)
	if err != nil {
		return players, bytes, err
	}

	var resp NumberOfCurrentPlayersResponse
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return players, bytes, err
	}

	return resp.Response.PlayerCount, bytes, nil
}

type NumberOfCurrentPlayersResponse struct {
	Response NumberOfCurrentPlayers `json:"response"`
}

type NumberOfCurrentPlayers struct {
	PlayerCount int `json:"player_count"`
	Result      int `json:"result"`
}

// Gets the complete list of stats and achievements for the specified game.
func GetSchemaForGame(appID int) (schema SchemaForGame, bytes []byte, err error) {

	options := url.Values{}
	options.Set("appid", strconv.Itoa(appID))
	options.Set("l", "english")

	bytes, err = get("ISteamUserStats/GetSchemaForGame/v2", options)
	if err != nil {
		return schema, bytes, err
	}

	var resp SchemaForGameResponse
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return schema, bytes, err
	}

	return resp.Game, bytes, nil
}

type SchemaForGameResponse struct {
	Game SchemaForGame `json:"game"`
}

type SchemaForGame struct {
	Name               string             `json:"gameName"`
	Version            string             `json:"gameVersion"`
	AvailableGameStats SchemaForGameGroup `json:"availableGameStats"`
}

type SchemaForGameGroup struct {
	Stats        []SchemaForGameStats        `json:"stats"`
	Achievements []SchemaForGameAchievements `json:"achievements"`
}

type SchemaForGameStats struct {
	Name         string `json:"name"`
	DefaultValue int    `json:"defaultvalue"`
	DisplayName  string `json:"displayName"`
}

type SchemaForGameAchievements struct {
	Name         string `json:"name"`
	DefaultValue int    `json:"defaultvalue"`
	DisplayName  string `json:"displayName"`
	Hidden       int8    `json:"hidden"`
	Description  string `json:"description"`
	Icon         string `json:"icon"`
	IconGray     string `json:"icongray"`
}
