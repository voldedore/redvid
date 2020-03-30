package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/buger/jsonparser"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get video URL",
	Long:  `Usage: redvid get https://reddit-url`,
	Run: func(cmd *cobra.Command, args []string) {
		wantedUrl := parseJson(getJson(args[0] + ".json"))
		fmt.Println(wantedUrl)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func parseJson(json []byte) string {
	if json == nil {
		return ""
	}

	// Check if error was received
	output, err := jsonparser.GetString(json, "message")
	if err != nil {
		log.Print("Getting error message failed, everything seems right")
	} else {
		return output
	}

	output, err = jsonparser.GetString(json, "[0]", "data", "children", "[0]", "data", "secure_media", "reddit_video", "fallback_url")
	if err != nil {
		// log.Print(string(json[:]))
		log.Print(err)
	} else {
		return output
	}

	return ""
}

func getJson(url string) []byte {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
		return nil
	}

	req.Header.Set("User-Agent", "telegram-bot:reddit-video:v1.0.0")

	resp, err := client.Do(req)

	if err != nil {
		log.Print(err)
	} else {

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Print(err)
		} else {
			return body
		}
	}

	return nil
}
