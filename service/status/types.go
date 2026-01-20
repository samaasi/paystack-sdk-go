package status

// Summary represents the summary of the status page
type Summary struct {
	Page                  Page                   `json:"page"`
	Components            []Component            `json:"components"`
	Incidents             []Incident             `json:"incidents"`
	ScheduledMaintenances []ScheduledMaintenance `json:"scheduled_maintenances"`
	Status                Status                 `json:"status"`
}

// Page represents the page information
type Page struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	TimeZone  string `json:"time_zone"`
	UpdatedAt string `json:"updated_at"`
}

// Component represents a system component
type Component struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Status             string `json:"status"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	Position           int    `json:"position"`
	Description        string `json:"description"`
	Showcase           bool   `json:"showcase"`
	StartDate          string `json:"start_date"`
	GroupID            string `json:"group_id"`
	OnlyShowIfDegraded bool   `json:"only_show_if_degraded"`
}

// Incident represents an incident
type Incident struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	Status          string           `json:"status"`
	CreatedAt       string           `json:"created_at"`
	UpdatedAt       string           `json:"updated_at"`
	MonitoringAt    string           `json:"monitoring_at"`
	ResolvedAt      string           `json:"resolved_at"`
	Impact          string           `json:"impact"`
	Shortlink       string           `json:"shortlink"`
	StartedAt       string           `json:"started_at"`
	PageID          string           `json:"page_id"`
	IncidentUpdates []IncidentUpdate `json:"incident_updates"`
	Components      []Component      `json:"components"`
}

// IncidentUpdate represents an update to an incident
type IncidentUpdate struct {
	ID                   string      `json:"id"`
	Status               string      `json:"status"`
	Body                 string      `json:"body"`
	IncidentID           string      `json:"incident_id"`
	CreatedAt            string      `json:"created_at"`
	UpdatedAt            string      `json:"updated_at"`
	DisplayAt            string      `json:"display_at"`
	AffectedComponents   []Component `json:"affected_components"`
	DeliverNotifications bool        `json:"deliver_notifications"`
	CustomTweet          string      `json:"custom_tweet"`
	TweetID              string      `json:"tweet_id"`
}

// ScheduledMaintenance represents a scheduled maintenance
type ScheduledMaintenance struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	Status          string           `json:"status"`
	CreatedAt       string           `json:"created_at"`
	UpdatedAt       string           `json:"updated_at"`
	MonitoringAt    string           `json:"monitoring_at"`
	ResolvedAt      string           `json:"resolved_at"`
	Impact          string           `json:"impact"`
	Shortlink       string           `json:"shortlink"`
	StartedAt       string           `json:"started_at"`
	PageID          string           `json:"page_id"`
	IncidentUpdates []IncidentUpdate `json:"incident_updates"`
	Components      []Component      `json:"components"`
	ScheduledFor    string           `json:"scheduled_for"`
	ScheduledUntil  string           `json:"scheduled_until"`
}

// Status represents the overall status
type Status struct {
	Indicator   string `json:"indicator"`
	Description string `json:"description"`
}
