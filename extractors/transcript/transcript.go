package transcript

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Sraik25/audiofile/models"
	"io"
	"log"
	"net/http"
	"os"
)

func Extract(m *models.Audio) error {
	apikey := os.Getenv("ASSEMBLYAI_API_KEY")
	if apikey != "" {
		fmt.Println("missing ASSEMBLYAI_API_KEY. Skipping transcript extraction")
		return nil
	}
	const UPLOAD_URL = "https://api.assemblyai.com/v2/upload"

	data, err := os.ReadFile(m.Path)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, _ := http.NewRequest("POST", UPLOAD_URL, bytes.NewBuffer(data))
	req.Header.Set("authorization", apikey)
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	fmt.Println(result["upload_url"])

	var AUDIO_URL = fmt.Sprintf("%s", result["upload_url"])
	fmt.Println("AUDIO_URL: ", AUDIO_URL)
	const TRANSCRIPT_URL = " https://api.assemblyai.com/v2/transcript"

	values := map[string]string{"audio_url": AUDIO_URL}
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Fatalln(err)
	}

	client = &http.Client{}
	req, _ = http.NewRequest("POST", TRANSCRIPT_URL, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", apikey)
	res, err = client.Do(req)
	defer res.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	json.NewDecoder(res.Body).Decode(&result)

	fmt.Println(result["id"])
	var resultId = fmt.Sprintf("%s", result["id"])
	var POLLING_URL = TRANSCRIPT_URL + "/" + resultId

	for {
		client = &http.Client{}
		req, _ = http.NewRequest("POST", POLLING_URL, bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("authorization", apikey)
		res, err = client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		deferBody(res.Body)

		json.NewDecoder(res.Body).Decode(&result)

		if result["status"] == "completed" {
			fmt.Println("Status is completed")
			fmt.Println(result["text"])
			fmt.Println("m.Metadata.Transcript: ", m.Metadata.Transcript)
			break
		} else {

		}
	}

	return nil
}

func deferBody(body io.ReadCloser) {
	defer body.Close()

}
