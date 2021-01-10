package types

type Channel struct {
	Id string `json:"_id"`
	Name string `json:"name"`
	DisplayName string `json:"displayName"`
	Joined bool `json:"joined"`
	Plan string `json:"plan"`
}

type ChannelResponse struct {
	Status int16 `json:"status"`
	Channel Channel `json:"channel"`
}
