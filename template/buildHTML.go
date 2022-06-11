package template

import "fmt"

type videoData struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Img         string `json:"img"`
}

func BuildHTML(videos Response) []byte {

	channel := videos.Items[0].Snippet.ChannelTitle
	Header(channel)

	var body []byte
	for _, item := range videos.Items {
		videoData := videoData{
			Id:          item.ID.VideoID,
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
			Img:         item.Snippet.Thumbnails.Default.URL,
		}
		video := Body(videoData)
		body = append(body, video...)
	}

	fmt.Println(string(body))

	// html := append(Header((channel), body))
	return Header(channel)
}
