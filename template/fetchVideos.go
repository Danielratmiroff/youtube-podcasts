package template

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	helpers "podcasts/helpers"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	helpers.HandleError(err, "Error loading .env file")
	return os.Getenv(key)
}

func FetchVideos(query string) Response {
	req, reqErr := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/search", nil)
	helpers.HandleError(reqErr, "handling new request")

	apiKey := goDotEnvVariable("YOUTUBE_API_KEY")

	q := req.URL.Query()

	if len(query) > 0 {
		parsed, err := url.ParseQuery(query)
		helpers.HandleError(err, "parsing query")
		q.Add("q", parsed.Get("q"))
	}

	q.Add("key", apiKey)
	q.Add("channelId", "UCSHZKyawb77ixDdsGog4iWA")
	q.Add("part", "snippet")
	q.Add("order", "viewCount")
	q.Add("type", "video")
	q.Add("maxResults", "50")

	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, doErr := client.Do(req)
	helpers.HandleError(doErr, "fetching videos")

	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)

	var response Response
	jsonErr := json.Unmarshal(body, &response)
	helpers.HandleError(jsonErr, "fetching videos")

	return response
}
