package maps

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)

	websites["LinkedIn"] = "https://linkedin.com"

	delete(websites, "Google")

}
