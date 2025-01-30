package places

import (
	"database/sql"
	"errors"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(p *Place) error {
	query := `
        INSERT INTO places (
            name, type, address, district, subdistrict,
            city, latitude, longitude, size_m2
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING id, created_at, updated_at
    `
	err := r.db.QueryRow(query,
		p.Name, p.Type, p.Address, p.District, p.Subdistrict,
		p.City, p.Latitude, p.Longitude, p.SizeM2,
	).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	return err
}

func (r *Repository) GetAll() ([]Place, error) {
	rows, err := r.db.Query(`
        SELECT id, name, type, address, district, subdistrict,
               city, latitude, longitude, size_m2,
               created_at, updated_at
          FROM places
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var places []Place
	for rows.Next() {
		var p Place
		err = rows.Scan(
			&p.ID, &p.Name, &p.Type, &p.Address, &p.District, &p.Subdistrict,
			&p.City, &p.Latitude, &p.Longitude, &p.SizeM2,
			&p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		places = append(places, p)
	}
	return places, rows.Err()
}

func (r *Repository) GetByID(id string) (*Place, error) {
	var p Place
	query := `
        SELECT id, name, type, address, district, subdistrict,
               city, latitude, longitude, size_m2,
               created_at, updated_at
          FROM places
         WHERE id = $1
    `
	err := r.db.QueryRow(query, id).Scan(
		&p.ID, &p.Name, &p.Type, &p.Address, &p.District, &p.Subdistrict,
		&p.City, &p.Latitude, &p.Longitude, &p.SizeM2,
		&p.CreatedAt, &p.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *Repository) Update(id string, p *Place) error {
	query := `
        UPDATE places
           SET name = $2,
               type = $3,
               address = $4,
               district = $5,
               subdistrict = $6,
               city = $7,
               latitude = $8,
               longitude = $9,
               size_m2 = $10,
               updated_at = NOW()
         WHERE id = $1
     RETURNING created_at, updated_at
    `
	err := r.db.QueryRow(query,
		id,
		p.Name, p.Type, p.Address, p.District, p.Subdistrict,
		p.City, p.Latitude, p.Longitude, p.SizeM2,
	).Scan(&p.CreatedAt, &p.UpdatedAt)
	return err
}

func (r *Repository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM places WHERE id = $1", id)
	return err
}
