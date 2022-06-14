package template

import (
	"os"
	helpers "podcasts/helpers"
)

func saveTemplate(template []byte) {
	os.Mkdir("build", 0777)

	pwd, dirErr := os.Getwd()
	helpers.HandleError(dirErr, "getting current dir")

	filePath := pwd + "/build/index.html"

	file, cErr := os.Create(filePath)
	helpers.HandleError(cErr, "creating file")
	defer file.Close()

	err := os.WriteFile(filePath, template, 0777)
	helpers.HandleError(err, "writing file")
}

func BuildTemplate(query string) {
	videos := FetchVideos(query)
	template := BuildHTML(videos)
	saveTemplate(template)
}
