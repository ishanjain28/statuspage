package statuspage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CreateIncidentInput struct {
	Incident Incident `json:"incident"`
}

type CreateIncidentOutput struct {
	ID                            string                    `json:"id"`
	Components                    []Components              `json:"components"`
	CreatedAt                     time.Time                 `json:"created_at"`
	Impact                        string                    `json:"impact"`
	ImpactOverride                string                    `json:"impact_override"`
	IncidentUpdates               []DetailedIncidentUpdates `json:"incident_updates"`
	MonitoringAt                  time.Time                 `json:"monitoring_at"`
	Name                          string                    `json:"name"`
	PageID                        string                    `json:"page_id"`
	PostmortemBody                string                    `json:"postmortem_body"`
	PostmortemBodyLastUpdatedAt   time.Time                 `json:"postmortem_body_last_updated_at"`
	PostmortemIgnored             bool                      `json:"postmortem_ignored"`
	PostmortemNotifiedSubscribers bool                      `json:"postmortem_notified_subscribers"`
	PostmortemNotifiedTwitter     bool                      `json:"postmortem_notified_twitter"`
	PostmortemPublishedAt         bool                      `json:"postmortem_published_at"`
	ResolvedAt                    time.Time                 `json:"resolved_at"`
	ScheduledAutoCompleted        bool                      `json:"scheduled_auto_completed"`
	ScheduledAutoInProgress       bool                      `json:"scheduled_auto_in_progress"`
	ScheduledFor                  time.Time                 `json:"scheduled_for"`
	ScheduledRemindPrior          bool                      `json:"scheduled_remind_prior"`
	ScheduledRemindedAt           time.Time                 `json:"scheduled_reminded_at"`
	ScheduledUntil                time.Time                 `json:"scheduled_until"`
	Shortlink                     string                    `json:"shortlink"`
	Status                        string                    `json:"status"`
	UpdatedAt                     time.Time                 `json:"updated_at"`
}

func (c Client) CreateIncident(pageID string, input *CreateIncidentInput) (*CreateIncidentOutput, error) {
	if pageID == "" {
		return nil, fmt.Errorf("page_id can not be empty")
	}

	url := fmt.Sprintf("%s/v1/pages/%s/incidents", c.url, pageID)

	body, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("error in marshalling create incident input: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("error in constructing request: %v", err)
	}

	response := new(CreateIncidentOutput)

	err = c.makeRequest(req, authOauthHeader, &response, http.StatusCreated)
	return response, err
}
