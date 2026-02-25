package adapters

import (
	"database/sql"
	"fmt"
	"laundry-hub-api/src/reservation/domain/entities"
	"time"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(reservation *entities.Reservation) (*entities.Reservation, error) {
	query := `INSERT INTO reservations (user_id, machine_id, status) VALUES (?, ?, ?)`

	result, err := m.conn.Exec(query, reservation.UserID, reservation.MachineID, reservation.Status)
	if err != nil {
		return nil, fmt.Errorf("error al guardar reservación: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error al obtener ID: %v", err)
	}

	return m.GetByID(int(id))
}

func (m *MySQL) GetByID(id int) (*entities.Reservation, error) {
	query := `SELECT id, user_id, machine_id, status, started_at, ended_at FROM reservations WHERE id = ?`

	var reservation entities.Reservation
	err := m.conn.QueryRow(query, id).Scan(
		&reservation.ID,
		&reservation.UserID,
		&reservation.MachineID,
		&reservation.Status,
		&reservation.StartedAt,
		&reservation.EndedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar reservación: %v", err)
	}

	return &reservation, nil
}

func (m *MySQL) GetByUserID(userID int) ([]*entities.Reservation, error) {
	query := `SELECT id, user_id, machine_id, status, started_at, ended_at FROM reservations WHERE user_id = ? ORDER BY started_at DESC`

	rows, err := m.conn.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener reservaciones: %v", err)
	}
	defer rows.Close()

	var reservations []*entities.Reservation
	for rows.Next() {
		var reservation entities.Reservation
		err := rows.Scan(
			&reservation.ID,
			&reservation.UserID,
			&reservation.MachineID,
			&reservation.Status,
			&reservation.StartedAt,
			&reservation.EndedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear reservación: %v", err)
		}
		reservations = append(reservations, &reservation)
	}

	return reservations, nil
}

func (m *MySQL) UpdateStatus(id int, status string, endedAt *time.Time) error {
	query := `UPDATE reservations SET status = ?, ended_at = ? WHERE id = ?`

	_, err := m.conn.Exec(query, status, endedAt, id)
	if err != nil {
		return fmt.Errorf("error al actualizar reservación: %v", err)
	}

	return nil
}
