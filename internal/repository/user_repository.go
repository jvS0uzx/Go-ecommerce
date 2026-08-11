package repository

import(
	"database/sql"
	"ecommerce-api/internal/model"
)


type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository{
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user *model.User) error{

	query := `INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id`
	
	err := r.db.QueryRow(query, user.Email, user.PasswordHash).Scan(&user.ID)

	if err != nil{
		return err
	}

	return nil
}