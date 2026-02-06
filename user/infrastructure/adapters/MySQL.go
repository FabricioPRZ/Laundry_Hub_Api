package adapters

import (
	"database/sql"
	"errors"

	"laundry-hub-api/user/domain/entities"
)

type UserMySQLRepository struct {
	db *sql.DB
}

func NewUserMySQLRepository(db *sql.DB) *UserMySQLRepository {
	return &UserMySQLRepository{db: db}
}

// Create crea un nuevo usuario en la base de datos
func (r *UserMySQLRepository) Create(user *entities.User) error {
	query := `
		INSERT INTO users (id, name, email, password, role, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(query,
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}

// FindByID busca un usuario por ID
func (r *UserMySQLRepository) FindByID(id string) (*entities.User, error) {
	query := `
		SELECT id, name, email, password, role, created_at, updated_at
		FROM users
		WHERE id = ?
	`

	user := &entities.User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindByEmail busca un usuario por email
func (r *UserMySQLRepository) FindByEmail(email string) (*entities.User, error) {
	query := `
		SELECT id, name, email, password, role, created_at, updated_at
		FROM users
		WHERE email = ?
	`

	user := &entities.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Update actualiza un usuario existente
func (r *UserMySQLRepository) Update(user *entities.User) error {
	query := `
		UPDATE users
		SET name = ?, email = ?, role = ?, updated_at = ?
		WHERE id = ?
	`

	_, err := r.db.Exec(query,
		user.Name,
		user.Email,
		user.Role,
		user.UpdatedAt,
		user.ID,
	)

	return err
}

// Delete elimina un usuario
func (r *UserMySQLRepository) Delete(id string) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}