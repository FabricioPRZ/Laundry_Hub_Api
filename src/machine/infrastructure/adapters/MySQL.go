package adapters

import (
	"database/sql"
	"errors"
	"fmt"
	"laundry-hub-api/src/machine/domain/entities"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(machine *entities.Machine) (*entities.Machine, error) {
	query := `INSERT INTO machines (name, status, capacity, location) VALUES (?, ?, ?, ?)`

	result, err := m.conn.Exec(query, machine.Name, machine.Status, machine.Capacity, machine.Location)
	if err != nil {
		return nil, fmt.Errorf("error al guardar máquina: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error al obtener ID: %v", err)
	}

	machine.ID = int(id)
	return machine, nil
}

func (m *MySQL) GetByID(id int) (*entities.Machine, error) {
	query := `SELECT id, name, status, capacity, location, created_at, updated_at FROM machines WHERE id = ?`

	var machine entities.Machine
	err := m.conn.QueryRow(query, id).Scan(
		&machine.ID,
		&machine.Name,
		&machine.Status,
		&machine.Capacity,
		&machine.Location,
		&machine.CreatedAt,
		&machine.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar máquina: %v", err)
	}

	return &machine, nil
}

func (m *MySQL) GetAll() ([]*entities.Machine, error) {
	query := `SELECT id, name, status, capacity, location, created_at, updated_at FROM machines ORDER BY created_at DESC`

	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener máquinas: %v", err)
	}
	defer rows.Close()

	var machines []*entities.Machine
	for rows.Next() {
		var machine entities.Machine
		err := rows.Scan(
			&machine.ID,
			&machine.Name,
			&machine.Status,
			&machine.Capacity,
			&machine.Location,
			&machine.CreatedAt,
			&machine.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear máquina: %v", err)
		}
		machines = append(machines, &machine)
	}

	return machines, nil
}

func (m *MySQL) Update(machine *entities.Machine) error {
	query := `UPDATE machines SET name = ?, status = ?, capacity = ?, location = ? WHERE id = ?`

	result, err := m.conn.Exec(query, machine.Name, machine.Status, machine.Capacity, machine.Location, machine.ID)
	if err != nil {
		return fmt.Errorf("error al actualizar máquina: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("máquina no encontrada")
	}

	return nil
}

func (m *MySQL) Delete(id int) error {
	query := `DELETE FROM machines WHERE id = ?`

	result, err := m.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar máquina: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("máquina no encontrada")
	}

	return nil
}
