package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ishanjain28/statuspage"
)

func main() {
	spPageID := os.Getenv("STATUS_PAGE_PAGE_ID")
	spAPIKey := os.Getenv("STATUS_PAGE_API_KEY")

	sp := statuspage.WithAPIKey(spAPIKey)

	resp, err := sp.DeleteIncident(spPageID, "")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%#+v", resp)
}
