## Fetch your favourite podcasts

![Build Status](https://github.com/Danielratmiroff/yt-podcasts/actions/workflows/main.yml/badge.svg)

While searching for podcasts in youtube, I get distracted by all of its rich features.

To addressed this problem, I created this simple website that fetchs uniquely my favourite podcasts.

#### Currently supported podcasts:

- [Lex Fridman Podcast](https://www.youtube.com/c/lexfridman)

## How to use

#### Clone the project into your local folder

#### Create a Youtube API Key - [How to create API key](https://developers.google.com/youtube/registering_an_application)

#### Set you Youtube API Key as an OS environmental variable

```bash
sudo vim ~/.bashrc
```

Add in the file:

```bash
export YOUTUBE_API_KEY=YOUR_API_KEY_HERE
```

#### Run it with:

```bash
sh run-locally.sh
```

#### Access:

[http://localhost:9000](http://localhost:9000/)
