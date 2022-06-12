package template

import "fmt"

type videoData struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Img         string `json:"img"`
}

// todo: CSS should be reused from other project

func BuildHTML(videos Response) []byte {

	channel := videos.Items[0].Snippet.ChannelTitle

	html := fmt.Sprintf("%s%s", GetHeader(channel), GetBody(videos))

	fmt.Println(string(html))

	return []byte(html)
}
