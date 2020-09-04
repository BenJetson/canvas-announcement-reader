package main

import (
	"fmt"
	"os"

	"github.com/BenJetson/canvas-announcement-reader/canvas"
)

func main() {
	token := os.Getenv("CANVAS_TOKEN")
	host := os.Getenv("CANVAS_HOST")

	api := canvas.NewAPI(host, token)

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
