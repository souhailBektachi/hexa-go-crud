package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/souhailBektachi/hexa-go-crud/internal/core/domain"

	_ "github.com/lib/pq"
)

type PostgressRepository struct {
	db *sql.DB
}

func NewPostgressRepository() *PostgressRepository {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	return &PostgressRepository{db: db}
}

func (r *PostgressRepository) Save(somthing domain.Somthing) error {
	query := "INSERT INTO somthing (name) VALUES ($1) RETURNING id"
	err := r.db.QueryRow(query, somthing.Name).Scan(&somthing.ID)
	if err != nil {
		return fmt.Errorf("could not save somthing: %v", err)
	}
	return nil
}

func (r *PostgressRepository) FindById(id int) (domain.Somthing, error) {
	var somthing domain.Somthing
	query := "SELECT id, name FROM somthing WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&somthing.ID, &somthing.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return somthing, errors.New("somthing not found")
		}
		return somthing, fmt.Errorf("could not find somthing: %v", err)
	}
	return somthing, nil
}

func (r *PostgressRepository) FindAll() ([]domain.Somthing, error) {
	var somthings []domain.Somthing
	query := "SELECT id, name FROM somthing"
	rows, err := r.db.Query(query)
	if err != nil {
		return somthings, fmt.Errorf("could not find somthings: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var somthing domain.Somthing
		if err := rows.Scan(&somthing.ID, &somthing.Name); err != nil {
			return somthings, fmt.Errorf("could not scan somthing: %v", err)
		}
		somthings = append(somthings, somthing)
	}
	return somthings, nil
}

func (r *PostgressRepository) DeleteById(id int) error {
	query := "DELETE FROM somthing WHERE id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not delete somthing: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return errors.New("somthing not found")
	}
	return nil
}

func (r *PostgressRepository) Update(somthing domain.Somthing) error {
	query := "UPDATE somthing SET name = $1 WHERE id = $2"
	result, err := r.db.Exec(query, somthing.Name, somthing.ID)
	if err != nil {
		return fmt.Errorf("could not update somthing: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return errors.New("somthing not found")
	}
	return nil
}
