package repo

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/team-18-project/InnoSync/backend/models"
)

type ChatRepository struct {
	db *sqlx.DB
}

func NewChatRepository(db *sqlx.DB) *ChatRepository {
	return &ChatRepository{db: db}
}

func (r *ChatRepository) GetProjectChat(ctx context.Context, projectID int64) (*models.GroupChat, error) {
	chat := &models.GroupChat{}
	query := `SELECT * FROM group_chat WHERE project_id = $1`
	err := r.db.GetContext(ctx, chat, query, projectID)
	return chat, err
}

func (r *ChatRepository) CreateMessage(ctx context.Context, message *models.Message) error {
	query := `INSERT INTO message (sender_id, group_chat_id, content) 
              VALUES ($1, $2, $3) RETURNING id, sent_at`
	return r.db.QueryRowContext(ctx, query,
		message.SenderID,
		message.GroupChatID,
		message.Content).
		Scan(&message.ID, &message.SentAt)
}

func (r *ChatRepository) GetChatMessages(ctx context.Context, chatID int64) ([]models.Message, error) {
	var messages []models.Message
	query := `SELECT * FROM message WHERE group_chat_id = $1 ORDER BY sent_at ASC`
	err := r.db.SelectContext(ctx, &messages, query, chatID)
	return messages, err
}
