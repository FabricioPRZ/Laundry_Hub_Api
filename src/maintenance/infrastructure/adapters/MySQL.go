package adapters

import (
	"database/sql"
	"fmt"
	"laundry-hub-api/src/maintenance/domain/entities"
	"math"
	"time"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(record *entities.MaintenanceRecord) (*entities.MaintenanceRecord, error) {
	query := `INSERT INTO maintenance_records (machine_id, description) VALUES (?, ?)`

	result, err := m.conn.Exec(query, record.MachineID, record.Description)
	if err != nil {
		return nil, fmt.Errorf("error al guardar registro: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error al obtener ID: %v", err)
	}

	record.ID = int(id)
	return m.GetByID(record.ID)
}

func (m *MySQL) GetAll() ([]*entities.MaintenanceRecord, error) {
	query := `
		SELECT mr.id, mr.machine_id, ma.name, mr.description,
		       mr.is_resolved, mr.resolved_at, mr.created_at
		FROM maintenance_records mr
		JOIN machines ma ON ma.id = mr.machine_id
		ORDER BY mr.created_at DESC
	`

	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener registros: %v", err)
	}
	defer rows.Close()

	var records []*entities.MaintenanceRecord
	for rows.Next() {
		var r entities.MaintenanceRecord
		if err := rows.Scan(
			&r.ID, &r.MachineID, &r.MachineName,
			&r.Description, &r.IsResolved, &r.ResolvedAt, &r.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("error al escanear registro: %v", err)
		}
		records = append(records, &r)
	}

	return records, nil
}

func (m *MySQL) GetByID(id int) (*entities.MaintenanceRecord, error) {
	query := `
		SELECT mr.id, mr.machine_id, ma.name, mr.description,
		       mr.is_resolved, mr.resolved_at, mr.created_at
		FROM maintenance_records mr
		JOIN machines ma ON ma.id = mr.machine_id
		WHERE mr.id = ?
	`

	var r entities.MaintenanceRecord
	err := m.conn.QueryRow(query, id).Scan(
		&r.ID, &r.MachineID, &r.MachineName,
		&r.Description, &r.IsResolved, &r.ResolvedAt, &r.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar registro: %v", err)
	}

	return &r, nil
}

func (m *MySQL) Resolve(id int) error {
	query := `UPDATE maintenance_records SET is_resolved = TRUE, resolved_at = ? WHERE id = ?`
	_, err := m.conn.Exec(query, time.Now(), id)
	if err != nil {
		return fmt.Errorf("error al resolver registro: %v", err)
	}
	return nil
}

func (m *MySQL) Delete(id int) error {
	query := `DELETE FROM maintenance_records WHERE id = ?`
	_, err := m.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar registro: %v", err)
	}
	return nil
}

func DaysElapsed(t time.Time) int {
	return int(math.Floor(time.Since(t).Hours() / 24))
}