package template

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	helpers "podcasts/helpers"
)

func FetchVideos() Response {
	req, reqErr := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/search", nil)
	helpers.HandleError(reqErr, "handling new request")

	// todo: disabled for now
	// apiKey := os.Getenv("YOUTUBE_API_KEY")

	query := req.URL.Query()

	query.Add("key", "AIzaSyDwGJ6LNP1vy2O1NyBip0SYEgvbxgTJkqo")
	query.Add("channelId", "UCSHZKyawb77ixDdsGog4iWA")
	query.Add("part", "snippet")
	query.Add("order", "viewCount")
	query.Add("type", "video")
	query.Add("maxResults", "3")

	req.URL.RawQuery = query.Encode()

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
