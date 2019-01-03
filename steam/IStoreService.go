package steam

import (
	"encoding/json"
	"net/url"
	"strconv"
)

func (s Steam) GetAppList(limit int, offset int) (apps AppList, bytes []byte, err error) {

	q := url.Values{}
	q.Set("include_games", "1")
	q.Set("include_dlc", "1")
	q.Set("include_software", "1")
	q.Set("include_videos", "1")
	q.Set("include_hardware", "1")
	//q.Set("if_modified_since", "")

	if offset > 0 {
		q.Set("last_appid", strconv.Itoa(offset))
	}
	if limit > 0 {
		q.Set("max_results", strconv.Itoa(limit))
	}

	bytes, err = s.getFromAPI("IStoreService/GetAppList/v1", q)
	if err != nil {
		return apps, bytes, err
	}

	var resp AppListResponse
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return apps, bytes, err
	}

	return resp.AppListResponseInner, bytes, nil
}

type AppListResponse struct {
	AppListResponseInner AppList `json:"response"`
}

type AppList struct {
	Apps            []App `json:"apps"`
	HaveMoreResults bool  `json:"have_more_results"`
	LastAppID       int   `json:"last_appid"`
}

type App struct {
	AppID             int    `json:"appid"`
	Name              string `json:"name"`
	LastModified      int64  `json:"last_modified"`
	PriceChangeNumber int    `json:"price_change_number"`
}
