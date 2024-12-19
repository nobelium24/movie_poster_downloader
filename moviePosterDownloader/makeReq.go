package movieposterdownloader

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// https://www.omdbapi.com/?apikey=b138ed31&t=Inception

func FetchMovies() {
	movieTitle := os.Args[1]
	url := "https://www.omdbapi.com/?apikey=b138ed31&t=" + movieTitle

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	// defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to make request: %d", resp.StatusCode)
	}

	var movieResponse OMDBResponse
	err = json.NewDecoder(resp.Body).Decode(&movieResponse)
	if err != nil {
		log.Fatalf("Error decoding request: %v", err)
	}

	posterUrl := movieResponse.Poster

	req, err = http.NewRequest("GET", posterUrl, nil)
	if err != nil {
		log.Fatalf("Error fetching image: %v", err)
	}
	resp, err = client.Do(req)
	if err != nil {
		log.Fatalf("Error making img request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to make request: %d", resp.StatusCode)
	}

	file, err := os.Create("poster.jpg")
	if err != nil {
		log.Printf("An error occurred during file creation: %v", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatalf("Error writing image to file: %v", err)
	}

	log.Println("Poster downloaded successfully as poster.jpg")

}
