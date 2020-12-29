package types

type Track struct {
	ProviderId string `json:"providerId"`
	Provider   string `json:"provider"`
	Duration   int    `json:"duration"`
	Title      string `json:"title"`
	Artist     string `json:"artist"`
	URL        string `json:"url"`
}

type User struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	ProviderId  string `json:"providerId"`
	Provider    string `json:"provider"`
}

type Song struct {
	Id        string `json:"_id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Track     Track  `json:"track"`
	User      User   `json:"user"`
	Position  int    `json:"_position"`
}

type QueueResponse struct {
	Total           int      `json:"_total"`
	CurrentSong     Song     `json:"_currentSong"`
	RequestsEnabled bool     `json:"_requestsEnabled"`
	Providers       []string `json:"_providers"`
	SearchProvider  string   `json:"_searchProvider"`
	Status          int      `json:"status"`
	Queue           []Song   `json:"queue"`
}
