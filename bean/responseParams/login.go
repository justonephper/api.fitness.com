package responseParams

import "fitness/bean/models"

type LoginResponse struct {
	User      models.Users `json:"user"`
	Token     string       `json:"jwt"`
	ExpiresAt int64        `json:"expiresAt"`
}
