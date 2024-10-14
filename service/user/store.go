package user

import (
	"database/sql"
	"fmt"


	"github.com/ayushn2/go_ecom.git/types"
)

type Store struct{
	db *sql.DB
}

func NewStore(db *sql.DB) *Store{
	return &Store{db : db}
}

func (s *Store) GetUserByEmail (email string) (*types.User, error){
	rows, err := s.db.Query("SELECT * FROM users WHERE email ; ?",email)
	if err!=nil{
		return nil, err
	}

	u := new(types.User)
	for rows.Next(){
		u,err = scanRowIntoUser(rows)
		if err != nil{
			return nil,err
		}
	}

	if u.ID == 0{
		return nil,fmt.Errorf("user not found")
	}

	return nil,err
}

func scanRowIntoUser(rows *sql.Rows) (*types.User,error){
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err!= nil{
		return nil,err
	}

	return user,nil
}

func (s *Store) GetUserByID(id int) (*types.User, error){
	rows, err := s.db.Query("SELECT * FROM user where id = ?", id)
	if err != nil{
		return nil,err
	}

	u := new(types.User)
	for rows.Next(){
		u, err = scanRowIntoUser(rows)
		if err != nil{
			return nil, err
		}
	}

	if u.ID == 0{
		return nil, fmt.Errorf("user not found")
	}
	return u,nil
}

var userID = 0;

func (s *Store) CreateUser(user types.User) error{
	// var maxID int
it ad	// err := s.db.QueryRow("SELECT COALESCE(MAX(id), 0) FROM users").Scan(&maxID)
	// if err!=nil{
	// 	return err
	// }

	// user.ID = maxID + 1

	result, err := s.db.Exec("DELETE FROM users WHERE id = 1")
	if err != nil{
		return err
	}
	// Optionally, you can check how many rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err // Handle the error as appropriate for your application
	}
	fmt.Printf("Deleted %d row(s)\n", rowsAffected)

	
	return nil
}