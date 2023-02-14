package main

import (
	"os/exec"
	"fmt"
	"io"
	"net/http"
	"os"
)


func main() {

	// URL of the file to download
	url := "https://cheemsbread505.github.io/Cheems-Bread-Does-Code/video/hacker.mp4"

	// Create an HTTP client
	client := &http.Client{}

	// Send a GET request to the server
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer resp.Body.Close()

	// Create a new file to save the downloaded data
	out, err := os.Create("hacker.mp4")
	if err != nil {
		fmt.Println("Error while creating", "hacker.mp4", "-", err)
		return
	}
	defer out.Close()

	// Copy the downloaded data to the new file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

	fmt.Println("Downloaded", url, "to", "hacker.mp4")

	
	filePath := "hacker.mp4"

	// Loop 5 times
	for i := 0; i < 5; i++ {
	// Start a new process to open the video file in the default video player.
	cmd := exec.Command("xdg-open", filePath)
	err := cmd.Run()
	if err != nil {
		panic(err)
		
		}
	}
}
