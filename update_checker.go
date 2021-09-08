package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

type versions struct {
	Builds []struct {
		Version   string `json:"version"`
		Build     int    `json:"build"`
		Downloads struct {
			Application struct {
				Name string `json:"name"`
			} `json:"application"`
		} `json:"downloads"`
	} `json:"builds"`
}

type lastUpdate struct {
	Build   int
	Version string
}

func main() {
	url := "https://papermc.io/api/v2/projects/paper/version_group/" + os.Getenv("VERSION_GROUP") + "/builds"
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode == http.StatusNotFound {
		log.Println("cannot check for new paper version, network error")
		os.Exit(1)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("cannot check for new paper version, cannot read body fully")
		os.Exit(1)
	}
	err = resp.Body.Close()
	if err != nil {
		log.Println("cannot check for new paper version, cannot close body reader (update check)")
		os.Exit(1)
	}
	versions := versions{}
	err = json.Unmarshal(body, &versions)
	if err != nil {
		log.Println("cannot check for new paper version, json syntax error")
		os.Exit(1)
	}
	latestBuild := versions.Builds[len(versions.Builds)-1]
	build := latestBuild.Build
	filename := latestBuild.Downloads.Application.Name
	version := latestBuild.Version
	url = "https://papermc.io/api/v2/projects/paper/versions/" + version + "/builds/" + strconv.Itoa(build) + "/downloads/" + filename
	resp, err = http.Get(url)
	if err != nil || resp.StatusCode == http.StatusNotFound {
		log.Println("cannot check for new paper version, download failed")
		os.Exit(1)
	}
	out, err := os.Create(filename)
	if err != nil {
		log.Println("cannot check for new paper version, create new file failed")
		os.Exit(1)
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Println("cannot check for new paper version, write new version to file failed")
		os.Exit(1)
	}
	err = out.Close()
	if err != nil {
		log.Println("cannot check for new paper version, cannot close file after write to new file")
		os.Exit(1)
	}
	err = resp.Body.Close()
	if err != nil {
		log.Println("cannot check for new paper version, cannot close body reader (downloaded file)")
		os.Exit(1)
	}
	lastUpdate := lastUpdate{}
	lastUpdate.Version = version
	lastUpdate.Build = build
	if _, err := os.Stat("last_update.json"); err == nil {
		err = os.Remove("last_update.json")
		if err != nil {
			log.Println("cannot check for new paper version, delete file last_update.json failed")
			os.Exit(1)
		}
	}
	lastUpdateJson, err := json.MarshalIndent(lastUpdate, "", "    ")
	if err != nil {
		log.Println("cannot check for new paper version, last update json marshal failed")
		os.Exit(1)
	}
	err = ioutil.WriteFile("last_update.json", lastUpdateJson, 0644)
	if err != nil {
		log.Println("cannot check for new paper version, write last_update.json failed")
		os.Exit(1)
	}

	if os.Getenv("MOVE") == "true" {
		cmd := exec.Command("mv", filename, "server_files/minecraft_server.jar")
		err = cmd.Run()
		if err != nil {
			log.Println("cannot check for new paper version, mv " + filename + " server_files/minecraft_server.jar failed")
			os.Exit(1)
		}
	}
}
