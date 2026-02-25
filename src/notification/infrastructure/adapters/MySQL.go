package adapters

import (
	"database/sql"
	"fmt"
	"laundry-hub-api/src/notification/domain/entities"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{conn: conn}
}

func (m *MySQL) GetByID(id int) (*entities.Notification, error) {
	query := `SELECT id, user_id, reservation_id, message, type, is_read, created_at FROM notifications WHERE id = ?`

	var n entities.Notification
	err := m.conn.QueryRow(query, id).Scan(
		&n.ID,
		&n.UserID,
		&n.ReservationID,
		&n.Message,
		&n.Type,
		&n.IsRead,
		&n.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar notificación: %v", err)
	}

	return &n, nil
}

func (m *MySQL) Save(notification *entities.Notification) (*entities.Notification, error) {
	query := `INSERT INTO notifications (user_id, reservation_id, message, type, is_read) VALUES (?, ?, ?, ?, ?)`

	result, err := m.conn.Exec(query,
		notification.UserID,
		notification.ReservationID,
		notification.Message,
		notification.Type,
		false,
	)
	if err != nil {
		return nil, fmt.Errorf("error al guardar notificación: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error al obtener ID: %v", err)
	}

	return m.GetByID(int(id))
}

func (m *MySQL) GetByUserID(userID int) ([]*entities.Notification, error) {
	query := `
		SELECT id, user_id, reservation_id, message, type, is_read, created_at
		FROM notifications WHERE user_id = ? ORDER BY created_at DESC
	`

	rows, err := m.conn.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener notificaciones: %v", err)
	}
	defer rows.Close()

	var notifications []*entities.Notification
	for rows.Next() {
		var n entities.Notification
		err := rows.Scan(
			&n.ID,
			&n.UserID,
			&n.ReservationID,
			&n.Message,
			&n.Type,
			&n.IsRead,
			&n.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear notificación: %v", err)
		}
		notifications = append(notifications, &n)
	}

	return notifications, nil
}

func (m *MySQL) MarkAsRead(id int) error {
	query := `UPDATE notifications SET is_read = TRUE WHERE id = ?`

	_, err := m.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al marcar notificación como leída: %v", err)
	}

	return nil
}

func (m *MySQL) MarkAllAsRead(userID int) error {
	query := `UPDATE notifications SET is_read = TRUE WHERE user_id = ? AND is_read = FALSE`

	_, err := m.conn.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("error al marcar notificaciones como leídas: %v", err)
	}

	return nil
}
