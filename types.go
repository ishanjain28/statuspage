package statuspage

import (
	"fmt"
	"time"
)

type Incident struct {
	Name                                      string            `json:"name"`
	Status                                    string            `json:"status"`
	ImpactOverride                            string            `json:"impact_override"`
	ScheduledFor                              string            `json:"scheduled_for"`
	ScheduledUntil                            string            `json:"scheduled_until"`
	ScheduledRemindPrior                      bool              `json:"scheduled_remind_prior"`
	ScheduledAutoInProgress                   bool              `json:"scheduled_auto_in_progress"`
	ScheduledAutoCompleted                    bool              `json:"scheduled_auto_completed"`
	DeliverNotifications                      bool              `json:"deliver_notifications"`
	AutoTransitionDeliverNotificationsAtEnd   bool              `json:"auto_transition_deliver_notifications_at_end"`
	AutoTransitionDeliverNotificationsAtStart bool              `json:"auto_transition_deliver_notifications_at_start"`
	AutoTransitionToMaintenanceState          bool              `json:"auto_transition_to_maintenance_state"`
	AutoTransitionToOperationalState          bool              `json:"auto_transition_to_operational_state"`
	AutoTweetAtBeginning                      bool              `json:"auto_tweet_at_beginning"`
	AutoTweetOnCompletion                     bool              `json:"auto_tweet_on_completion"`
	AutoTweetOnCreation                       bool              `json:"auto_tweet_on_creation"`
	AutoTweetOneHourBefore                    bool              `json:"auto_tweet_one_hour_before"`
	BackfillDate                              string            `json:"backfill_date"`
	Backfilled                                bool              `json:"backfilled"`
	Body                                      string            `json:"body"`
	Components                                map[string]string `json:"components"`
	ComponentIds                              []string          `json:"component_ids"`
	ScheduledAutoTransition                   bool              `json:"scheduled_auto_transition"`
}

type DetailedIncident struct {
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

type Components struct {
	ID                 string    `json:"id"`
	PageID             string    `json:"page_id"`
	GroupID            string    `json:"group_id"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	Group              bool      `json:"group"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	Position           int       `json:"position"`
	Status             string    `json:"status"`
	Showcase           bool      `json:"showcase"`
	OnlyShowIfDegraded bool      `json:"only_show_if_degraded"`
	AutomationEmail    string    `json:"automation_email"`
}
type IncidentUpdates struct {
	ID                   string    `json:"id"`
	IncidentID           string    `json:"incident_id"`
	AffectedComponents   []string  `json:"affected_components"`
	Body                 string    `json:"body"`
	CreatedAt            time.Time `json:"created_at"`
	CustomTweet          string    `json:"custom_tweet"`
	DeliverNotifications bool      `json:"deliver_notifications"`
	DisplayAt            time.Time `json:"display_at"`
	Status               string    `json:"status"`
	TweetID              string    `json:"tweet_id"`
	TwitterUpdatedAt     time.Time `json:"twitter_updated_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	WantsTwitterUpdate   bool      `json:"wants_twitter_update"`
}

type FailedResponse struct {
	StatusCode int
	Message    string
}

func (f FailedResponse) Error() string {
	return fmt.Sprintf("%s(%d)", f.Message, f.StatusCode)
}

type DetailedIncidentUpdates struct {
	ID                   string               `json:"id"`
	IncidentID           string               `json:"incident_id"`
	AffectedComponents   []AffectedComponents `json:"affected_components"`
	Body                 string               `json:"body"`
	CreatedAt            time.Time            `json:"created_at"`
	CustomTweet          string               `json:"custom_tweet"`
	DeliverNotifications bool                 `json:"deliver_notifications"`
	DisplayAt            time.Time            `json:"display_at"`
	Status               string               `json:"status"`
	TweetID              string               `json:"tweet_id"`
	TwitterUpdatedAt     time.Time            `json:"twitter_updated_at"`
	UpdatedAt            time.Time            `json:"updated_at"`
	WantsTwitterUpdate   bool                 `json:"wants_twitter_update"`
}

type AffectedComponents struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	OldStatus string `json:"old_status"`
	NewStatus string `json:"new_status"`
}
