package adapters

import (
	"database/sql"
	"errors"

	"laundry-hub-api/machine/domain/entities"
)

type MachineMySQLRepository struct {
	db *sql.DB
}

func NewMachineMySQLRepository(db *sql.DB) *MachineMySQLRepository {
	return &MachineMySQLRepository{db: db}
}

// Create crea una nueva máquina en la base de datos
func (r *MachineMySQLRepository) Create(machine *entities.Machine) error {
	query := `
		INSERT INTO machines (id, name, status, capacity, location, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(query,
		machine.ID,
		machine.Name,
		machine.Status,
		machine.Capacity,
		machine.Location,
		machine.CreatedAt,
		machine.UpdatedAt,
	)

	return err
}

// FindAll obtiene todas las máquinas
func (r *MachineMySQLRepository) FindAll() ([]*entities.Machine, error) {
	query := `
		SELECT id, name, status, capacity, location, created_at, updated_at
		FROM machines
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var machines []*entities.Machine

	for rows.Next() {
		machine := &entities.Machine{}
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
			return nil, err
		}
		machines = append(machines, machine)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return machines, nil
}

// FindByID busca una máquina por ID
func (r *MachineMySQLRepository) FindByID(id string) (*entities.Machine, error) {
	query := `
		SELECT id, name, status, capacity, location, created_at, updated_at
		FROM machines
		WHERE id = ?
	`

	machine := &entities.Machine{}
	err := r.db.QueryRow(query, id).Scan(
		&machine.ID,
		&machine.Name,
		&machine.Status,
		&machine.Capacity,
		&machine.Location,
		&machine.CreatedAt,
		&machine.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("machine not found")
	}
	if err != nil {
		return nil, err
	}

	return machine, nil
}

// Update actualiza una máquina existente
func (r *MachineMySQLRepository) Update(machine *entities.Machine) error {
	query := `
		UPDATE machines
		SET name = ?, status = ?, capacity = ?, location = ?, updated_at = ?
		WHERE id = ?
	`

	result, err := r.db.Exec(query,
		machine.Name,
		machine.Status,
		machine.Capacity,
		machine.Location,
		machine.UpdatedAt,
		machine.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("machine not found")
	}

	return nil
}

// Delete elimina una máquina
func (r *MachineMySQLRepository) Delete(id string) error {
	query := `DELETE FROM machines WHERE id = ?`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("machine not found")
	}

	return nil
}