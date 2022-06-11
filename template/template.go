package template

import (
	"os"
	helpers "podcasts/helpers"
)

func saveTemplate(template []byte) {
	os.Mkdir("tmp", 0644)

	pwd, dirErr := os.Getwd()
	helpers.HandleError(dirErr, "getting current dir")

	filePath := pwd + "/tmp/index.html"

	file, cErr := os.Create(filePath)
	helpers.HandleError(cErr, "creating file")
	defer file.Close()

	err := os.WriteFile(filePath, template, 0644)
	helpers.HandleError(err, "writing file")
}

func BuildTemplate() {
	videos := FetchVideos()
	template := BuildHTML(videos)
	saveTemplate(template)

}
