package template

import "fmt"

func Header(channel string) []byte {
	str := fmt.Sprintf(`<!DOCTYPE html>
	<html lang="en">
	  <head>
		<meta charset="UTF-8" />
		<meta name="description" content="Youtube Podcasts Content">
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<meta http-equiv="X-UA-Compatible" content="ie=edge" />
		<title>%s</title>
	  </head>
	  <body>
		<script></script>
	  </body>
	</html>.`, channel)

	return []byte(str)
}

func Body(metadata videoData) []byte {
	// todo: replace LINK
	str := fmt.Sprintf(`<section class="videos" id="featured-videos">
    <div class="video-grid front-page" id="front-page-videos">
      <ul class="video-list featured">
        <li class="video featured">
          <a href="#" class="featured-video">
            <figure style="background-image: url(%s);">
              <img src="https://s3-us-west-2.amazonaws.com/s.cdpn.io/50598/video-thumb-placeholder-16-9.png" />
              <figcaption>%s</figcaption>
            </figure>
          </a>
		  <p>%s</p>
        </li>
      </ul>
    </div
  </section>`, metadata.Img, metadata.Title, metadata.Description)
	return []byte(str)
}
