package model

type MessageUser struct {
	UserID string      `json:"user_id"`
	Data   MessageData `json:"data"`
}

type MessageUsers struct {
	UserIDs []string    `json:"user_ids"`
	Data    MessageData `json:"data"`
}

type MessageData map[string]string
