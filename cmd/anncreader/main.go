package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"

	"github.com/BenJetson/canvas-announcement-reader/canvas"
)

type config struct {
	token string
	host  string
}

func getConfigFromEnvironment() *config {
	var cfg config
	cfg.token = os.Getenv("CANVAS_TOKEN")
	cfg.host = os.Getenv("CANVAS_HOST")

	if len(cfg.token) < 1 || len(cfg.host) < 1 {
		return nil
	}
	return &cfg
}

func getConfigInteractive() *config {
	var cfg config

	fmt.Println("Enter a the Canvas hostname for your school.")
	fmt.Println(`For example, clemson.instructure.com for Clemson University.`)
	fmt.Print(">>> ")
	fmt.Scanln(&cfg.host)

	if len(cfg.host) < 1 {
		err := errors.New("Hostname cannot be blank.")
		panic(err)
	}

	fmt.Println("Enter a Canvas access token for your account.")
	fmt.Print(">>> ")
	fmt.Scanln(&cfg.token)

	if len(cfg.token) < 1 {
		err := ("Access token cannot be blank.")
		panic(err)
	}

	fmt.Println()
	return &cfg
}

func main() {
	cfg := getConfigFromEnvironment()
	if cfg == nil {
		cfg = getConfigInteractive()
	}

	api := canvas.NewAPI(cfg.host, cfg.token)

	courses, err := api.GetActiveCourses()
	if err != nil {
		panic(err)
	}

	for _, course := range courses {
		fmt.Printf("Course ID %d: %s\n", course.ID, course.Name)

		announcements, err := api.GetAnnouncementsForCourse(course.ID)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Found %d announcements!\n", len(announcements))
		for _, a := range announcements {
			fmt.Printf("Annc ID: %d: %s", a.ID, a.Title)

			if a.ReadState != "read" {
				if err = api.MarkAnnouncementAsRead(course.ID, a.ID); err != nil {
					panic(err)
				}
				fmt.Println(" ... marked as read!")
			} else {
				fmt.Println(" ... already read.")
			}
		}

		fmt.Println()
	}
}
