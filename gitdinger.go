package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

var (
	ConfigFile = flag.String("config", "config.json", "Location of the config file")
	Cwd, _     = os.Getwd()
)

type Config struct {
	Repos  []Repo `json:"repos"`
	Period int64  `json:"period"`
}

type Repo struct {
	Path      string `json:"path"`
	Branch    string `json:"branch"`
	Autofetch bool   `json:"autofetch"`
	Soundfile string `json:"soundfile"`
}

type Notification struct {
	Type    int64
	Message string
}

func ParseConfig(configpath string) *Config {
	file, _ := os.Open(configpath)
	decoder := json.NewDecoder(file)
	configuration := &Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}
	return configuration
}

func checkCommand(name string, executable string, params string) {
	err := exec.Command(executable, params).Run()
	if err != nil {
		log.Fatal(fmt.Sprintf("You must have %s installed in your PATH", name))
	}
}

func Notifier(c chan Notification) {
	checkCommand("say", "say", "-v?")
	checkCommand("afplay", "which", "afplay")
	for msg := range c {
		if msg.Type == 1 {
			log.Printf(" Saying: %s\n", msg.Message)
			exec.Command("say", fmt.Sprintf("%s", msg.Message)).Run()
		} else if msg.Type == 2 {
			play_cmd := exec.Command("afplay", msg.Message)
			play_cmd.Dir = Cwd
			err := play_cmd.Run()
			if err != nil {
				log.Printf(" Could not play file %s", msg.Message)
			}
		}
	}
}

func WatchRepo(notify chan Notification, ticker *time.Ticker, r Repo) {
	var repo_state string
	for _ = range ticker.C {
		name := path.Base(r.Path)
		log.Printf("Checking %s:%s\n", name, r.Branch)

		if r.Autofetch == true {
			cmd := exec.Command("git", "fetch", "--all")
			cmd.Dir = r.Path
			fetch_err := cmd.Run()
			if fetch_err != nil {
				log.Println(" - Could not fetch, skipping", fetch_err)
				continue
			}
		}

		state_cmd := exec.Command("git", "log", "--pretty=format:%h", "-n1", r.Branch)
		state_cmd.Dir = r.Path
		current_state, err_current_state := state_cmd.Output()
		if err_current_state != nil {
			log.Fatal(err_current_state)
		}
		if repo_state == "" {
			log.Printf(" Setting initial state on %s to %s\n", name, current_state)
		} else {
			diff_cmd := exec.Command("git", "log", "--pretty=format:\"%an's commit %s\"", fmt.Sprintf("%s..%s", repo_state, current_state))
			diff_cmd.Dir = r.Path
			diff, _ := diff_cmd.Output()
			differences := strings.Split(string(diff), "\n")
			if len(diff) > 0 {
				difference_size := len(differences)
				log.Printf(" Dinging %d times.\n", difference_size)
				for i := 0; i < difference_size; i++ {
					notify <- Notification{2, r.Soundfile}
				}
				for i := 0; i < difference_size; i++ {
					notify <- Notification{1, fmt.Sprintf(differences[i])}
				}
			}
		}
		repo_state = string(current_state)
	}
}

func main() {
	flag.Parse()
	config := ParseConfig(*ConfigFile)
	notification_channel := make(chan Notification)
	checkCommand("git", "git", "--version")

	for repo := range config.Repos {
		r := config.Repos[repo]
		log.Printf("Setting up listener for repo '%s':%s to check every %d seconds.\n", r.Path, r.Branch, config.Period)
		go WatchRepo(notification_channel, time.NewTicker(time.Duration(config.Period)*time.Second), r)
	}

	Notifier(notification_channel)
}
