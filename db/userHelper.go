package db

import (
	"ApiBeeGo/models"
	"errors"
)

func CreateUser(user *models.User) (int64, error) {
	db := DBConnection()
	if AvailableUser(user.Email) {
		db.Close()
		// return fmt.Errorf("AllReady user exist in db")
		return 0, errors.New("AllReady user exist in db")
	}
	rst, err := db.Exec("INSERT INTO User VALUES (?, ?, ?, ?, ?)", nil, user.FirstName, user.LastName, user.Email, user.Password)
	db.Close()
	id, err := rst.LastInsertId()
	return id, err
}

func GetUser(email string) (models.User, error) {
	db := DBConnection()
	var user models.User
	rows, err := db.Query("SELECT * FROM user where email = ?", email)
	if err != nil {
		db.Close()
		return models.User{}, err
	}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err != nil {
			db.Close()
			return models.User{}, err
		}
	}
	db.Close()
	return user, nil
}
func GetAllUsers() ([]models.User, error) {
	db := DBConnection()
	// Execute the query
	results, err := db.Query("SELECT * FROM user")
	db.Close()
	if err != nil {
		return nil, err // proper error handling instead of log in your app
	}
	var users []models.User
	for results.Next() {
		var user models.User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err != nil {
			return nil, err // proper error handling instead of log in your app
		}
		// and then print out the tag's Name attribute
		user.Password = "can not allowed to out side..."
		users = append(users, user)
	}
	db.Close()
	return users, nil
}

func AvailableUser(email string) bool {
	db := DBConnection()
	rows, err := db.Query("SELECT * FROM user where email = ? LIMIT 1", email)
	if err != nil {
		db.Close()
		return false
	}
	db.Close()
	return rows.Next()
}

func Delete(id string) error {
	db := DBConnection()
	_, err := db.Exec("DELETE FROM user WHERE id= ?", id)
	db.Close()

	return err
}

func UpdateUser(user *models.User) error {
	db := DBConnection()
	_, err := db.Exec("Update User SET first_name=?, last_name=?, email=?, password=?  WHERE id= ?", user.FirstName, user.LastName, user.Email, user.Password, user.Id)
	db.Close()

	return err
}
