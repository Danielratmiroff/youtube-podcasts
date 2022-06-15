## Fetch your favourite podcasts

=========

![Build Status]((https://github.com/github/docs/actions/workflows/main.yml/badge.svg)

While searching for podcasts in youtube, I get distracted by all of its rich features.

To addressed this problem, I created this simple website that fetchs uniquely my favourite podcasts.

#### Currently supported podcasts:

- [Lex Fridman Podcasts](https://www.youtube.com/c/lexfridman)

## Usage

1. Clone the project into a local folder
2. Create a Youtube API Key - [How to](https://developers.google.com/youtube/registering_an_application)
3. Set you Youtube API Key as a OS environmental variable

```bash
sudo nano ~/.bashrc
```

_Add:_

```bash
export YOUTUBE_API_KEY=YOUR_API_KEY_HERE
```

4. Run run-locally.sh

```bash
sh run-locally.sh
```

5. Access [http://localhost:9000/](http://localhost:9000/)
