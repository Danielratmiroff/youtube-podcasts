package template

import "time"

type Response struct {
	Kind          string   `json:"kind"`
	Etag          string   `json:"etag"`
	NextPageToken string   `json:"nextPageToken"`
	RegionCode    string   `json:"regionCode"`
	PageInfo      PageInfo `json:"pageInfo"`
	Items         []Items  `json:"items"`
}
type PageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}
type ID struct {
	Kind    string `json:"kind"`
	VideoID string `json:"videoId"`
}
type Default struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
type Medium struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
type High struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
type Thumbnails struct {
	Default Default `json:"default"`
	Medium  Medium  `json:"medium"`
	High    High    `json:"high"`
}
type Snippet struct {
	PublishedAt          time.Time  `json:"publishedAt"`
	ChannelID            string     `json:"channelId"`
	Title                string     `json:"title"`
	Description          string     `json:"description"`
	Thumbnails           Thumbnails `json:"thumbnails"`
	ChannelTitle         string     `json:"channelTitle"`
	LiveBroadcastContent string     `json:"liveBroadcastContent"`
	PublishTime          time.Time  `json:"publishTime"`
}
type Items struct {
	Kind    string  `json:"kind"`
	Etag    string  `json:"etag"`
	ID      ID      `json:"id"`
	Snippet Snippet `json:"snippet"`
}
