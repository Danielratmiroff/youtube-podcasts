package template

import "fmt"

type videoData struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Img         string `json:"img"`
}

func BuildHTML(videos Response) []byte {
	channel := "Youtube Podcasts"
	if len(videos.Items) > 0 {
		channel = videos.Items[0].Snippet.ChannelTitle
	}

	html := fmt.Sprintf("%s%s", GetHeader(channel), GetBody(videos))
	return []byte(html)
}
