package response

import "api.fitness.com/bean/models"

type LoginResponse struct {
	User      models.Users `json:"user"`
	Token     string       `json:"token"`
	ExpiresAt int64        `json:"expiresAt"`
}
