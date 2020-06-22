package statuspage

import (
	"fmt"
	"net/http"
	"strconv"
)

type ListAllIncidentsInput struct {
	query      string
	limit      int
	pageOffset int
}

func (c Client) ListAllIncidents(pageID string, input *ListAllIncidentsInput) (*[]DetailedIncident, error) {
	if pageID == "" {
		return nil, fmt.Errorf("page_id can not be empty")
	}

	url := fmt.Sprintf("%s/v1/pages/%s/incidents", c.url, pageID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error in constructing request: %v", err)
	}

	query := req.URL.Query()
	if input.query != "" {
		query.Set("q", input.query)
	}
	if input.limit != 0 {
		query.Set("limit", strconv.Itoa(input.limit))
	}
	if input.pageOffset != 0 {
		query.Set("page", strconv.Itoa(input.pageOffset))
	}
	req.URL.RawQuery = query.Encode()

	response := new([]DetailedIncident)

	err = c.makeRequest(req, apiKeyQueryParam, &response, http.StatusOK)
	return response, err
}
