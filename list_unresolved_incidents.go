package statuspage

import (
	"fmt"
	"net/http"
	"strconv"
)

type ListUnresolvedIncidentsInput struct {
	page    int
	perPage int
}

func (c Client) ListUnresolvedIncidents(pageID string, input *ListUnresolvedIncidentsInput) (*[]DetailedIncident, error) {
	if pageID == "" {
		return nil, fmt.Errorf("page_id can not be empty")
	}

	url := fmt.Sprintf("%s/v1/pages/%s/incidents/unresolved", c.url, pageID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error in constructing request: %v", err)
	}

	query := req.URL.Query()
	if input.page != 0 {
		query.Set("page", strconv.Itoa(input.page))
	}
	if input.perPage != 0 {
		query.Set("per_page", strconv.Itoa(input.perPage))
	}
	req.URL.RawQuery = query.Encode()

	response := new([]DetailedIncident)

	err = c.makeRequest(req, apiKeyQueryParam, &response, http.StatusOK)

	return response, err
}
