package models

import "time"

type GroupChat struct {
	ID        int64     `json:"id" db:"id"`
	ProjectID int64     `json:"project_id" db:"project_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Message struct {
	ID          int64     `json:"id" db:"id"`
	SenderID    int64     `json:"sender_id" db:"sender_id"`
	GroupChatID *int64    `json:"group_chat_id,omitempty" db:"group_chat_id"`
	Content     string    `json:"content" db:"content"`
	SentAt      time.Time `json:"sent_at" db:"sent_at"`
}
