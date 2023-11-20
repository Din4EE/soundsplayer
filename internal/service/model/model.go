package model

type Track struct {
	Id            uint64 `json:"id"`
	Audio         string `json:"audio"`
	Title         string `json:"title"`
	Artist        string `json:"artist"`
	Time          uint64 `json:"time"`
	FormattedTime string `json:"formatted_time"`
	Cover         string `json:"cover"`
}
