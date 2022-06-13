package template

import (
	"fmt"
	"strings"
)

// replace misc css file path with dynamic one
func GetHeader(channel string) []byte {
	str := fmt.Sprintf(`<!DOCTYPE html>
	<html lang="en">
	  <head>
		<meta charset="UTF-8" />
		<meta name="description" content="Youtube Podcasts Content">
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<meta http-equiv="X-UA-Compatible" content="ie=edge" />
		<link rel="stylesheet" type="text/css" href="https://css-storage.s3.eu-central-1.amazonaws.com/common-vendor.b8ecfc406ac0b5f77a26.css">
		<link rel="stylesheet" type="text/css" href="https://css-storage.s3.eu-central-1.amazonaws.com/fretboard.f32f2a8d5293869f0195.css">
		<link rel="stylesheet" type="text/css" href="https://css-storage.s3.eu-central-1.amazonaws.com/pretty-vendor.83ac49e057c3eac4fce3.css">
		<link rel="stylesheet" type="text/css" href="https://css-storage.s3.eu-central-1.amazonaws.com/pretty.0ae3265014f89d9850bf.css">
		<link rel="stylesheet" type="text/css" href="/styles.css">
		<title>%s</title>
	  </head>
	  <body>
	.`, channel)

	return []byte(str)
}

// func GetNavBar() []byte {
// 	html := "hey"
// 	return []byte(html)
// }

func GetBody(videos Response) []byte {
	startBody := `<section class="videos" id="featured-videos">
					<div class="video-grid front-page" id="front-page-videos">
      				<ul class="video-list featured">`

	endBody := `</ul>
				</div>
				</section>
				<script></script>
				</body>
			  </html>`

	var body []byte
	for _, item := range videos.Items {
		videoData := videoData{
			Id:          item.ID.VideoID,
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
			Img:         item.Snippet.Thumbnails.Default.URL,
		}
		videoUrl := strings.Split(videoData.Img, "/")[4]
		videoThumbnail := bodyList(videoData, videoUrl)

		body = append(body, videoThumbnail...)
	}

	html := fmt.Sprintf("%s%s%s", startBody, body, endBody)

	return []byte(html)
}

func bodyList(metadata videoData, videoUrl string) string {
	str := fmt.Sprintf(`
        <li class="video featured">
          <a target="_blank" href="https://www.youtube.com/watch?v=%s" class="featured-video">
            <figure style="background-image: url(%s);">
              <img src="https://s3-us-west-2.amazonaws.com/s.cdpn.io/50598/video-thumb-placeholder-16-9.png" />
              <figcaption>%s</figcaption>
            </figure>
          </a>
		  <p>%s</p>`, videoUrl, metadata.Img, metadata.Title, metadata.Description)
	return str
}
