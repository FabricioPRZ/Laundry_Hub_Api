package adapters

import (
	"database/sql"
	"errors"
	"fmt"
	"laundry-hub-api/src/user/domain/entities"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(user *entities.User) (*entities.User, error) {
	query := `
		INSERT INTO users (
			name, second_name, paternal_surname, maternal_surname,
			email, password, image_profile, oauth_provider, oauth_id, role
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := m.conn.Exec(
		query,
		user.Name,
		user.SecondName,
		user.PaternalSurname,
		user.MaternalSurname,
		user.Email,
		user.Password,
		user.ImageProfile,
		user.OAuthProvider,
		user.OAuthID,
		user.Role,
	)
	if err != nil {
		return nil, fmt.Errorf("error al guardar usuario: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error al obtener ID: %v", err)
	}

	return m.GetByID(int(id))
}

func (m *MySQL) GetByEmail(email string) (*entities.User, error) {
	query := `
		SELECT id, name, second_name, paternal_surname, maternal_surname,
		       email, password, image_profile, oauth_provider, oauth_id, role,
		       created_at, updated_at
		FROM users WHERE email = ?
	`

	var user entities.User
	err := m.conn.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.SecondName,
		&user.PaternalSurname,
		&user.MaternalSurname,
		&user.Email,
		&user.Password,
		&user.ImageProfile,
		&user.OAuthProvider,
		&user.OAuthID,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar usuario por email: %v", err)
	}

	return &user, nil
}

func (m *MySQL) GetByID(id int) (*entities.User, error) {
	query := `
		SELECT id, name, second_name, paternal_surname, maternal_surname,
		       email, password, image_profile, oauth_provider, oauth_id, role,
		       created_at, updated_at
		FROM users WHERE id = ?
	`

	var user entities.User
	err := m.conn.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.SecondName,
		&user.PaternalSurname,
		&user.MaternalSurname,
		&user.Email,
		&user.Password,
		&user.ImageProfile,
		&user.OAuthProvider,
		&user.OAuthID,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar usuario por ID: %v", err)
	}

	return &user, nil
}

func (m *MySQL) GetAll() ([]*entities.User, error) {
	query := `
		SELECT id, name, second_name, paternal_surname, maternal_surname,
		       email, password, image_profile, oauth_provider, oauth_id, role,
		       created_at, updated_at
		FROM users ORDER BY created_at DESC
	`

	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener usuarios: %v", err)
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		var user entities.User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.SecondName,
			&user.PaternalSurname,
			&user.MaternalSurname,
			&user.Email,
			&user.Password,
			&user.ImageProfile,
			&user.OAuthProvider,
			&user.OAuthID,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear usuario: %v", err)
		}
		users = append(users, &user)
	}

	return users, nil
}

func (m *MySQL) Update(user *entities.User) error {
	query := `
		UPDATE users SET
			name = ?,
			second_name = ?,
			paternal_surname = ?,
			maternal_surname = ?,
			email = ?,
			image_profile = ?,
			role = ?
		WHERE id = ?
	`

	result, err := m.conn.Exec(
		query,
		user.Name,
		user.SecondName,
		user.PaternalSurname,
		user.MaternalSurname,
		user.Email,
		user.ImageProfile,
		user.Role,
		user.ID,
	)
	if err != nil {
		return fmt.Errorf("error al actualizar usuario: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("usuario no encontrado")
	}

	return nil
}

func (m *MySQL) Delete(id int) error {
	query := `DELETE FROM users WHERE id = ?`

	result, err := m.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar usuario: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return errors.New("usuario no encontrado")
	}

	return nil
}
