package Users

import (
	"database/sql"
	"fmt"
)

const (
	// UserDBTableName is the name of the table for the UserDB model
	UserDBTableName = "UserDB"
	// UserDBFirstNameCol is the column name of the model's first name
	UserDBFirstNameCol = "first_name"
	// UserDBLastNameCol is the column name of the model's last name
	UserDBLastNameCol = "last_name"
	// UserDBAgeCol is the column name of the model's age
	UserDBAgeCol = "age"
)

// UserDB is the database model for a UserDB
type UserDB struct {
	FirstName string
	LastName  string
	Age       uint
}

// CreateUserDBTable uses db to create a new table for UserDB models, and returns the result
func CreateUserDBTable(db *sql.DB) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("CREATE TABLE %s (%s varchar(255), %s varchar(255), %s int)",
			UserDBTableName,
			UserDBFirstNameCol,
			UserDBLastNameCol,
			UserDBAgeCol,
		),
	)
}

// InsertUserDB inserts UserDB into db
func InsertUserDB(db *sql.DB, UserDB UserDB) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("INSERT INTO %s VALUES(?, ?, ?)", UserDBTableName),
		UserDB.FirstName,
		UserDB.LastName,
		UserDB.Age,
	)
}

// SelectUserDB selects a UserDB with the given first & last names and age. On success, writes the result into result and on failure, returns a non-nil error and makes no modifications to result
func SelectUserDB(db *sql.DB, firstName, lastName string, age uint, result *UserDB) error {
	row := db.QueryRow(
		fmt.Sprintf(
			"SELECT * FROM %s WHERE %s=? AND %s=? AND %s=?",
			UserDBTableName,
			UserDBFirstNameCol,
			UserDBLastNameCol,
			UserDBAgeCol,
		),
		firstName,
		lastName,
		age,
	)
	var retFirstName, retLastName string
	var retAge uint
	if err := row.Scan(&retFirstName, &retLastName, &retAge); err != nil {
		return err
	}
	result.FirstName = retFirstName
	result.LastName = retLastName
	result.Age = retAge
	return nil
}

// UpdateUserDB updates the UserDB with the given first & last names and age with newUserDB. Returns a non-nil error if the update failed, and nil if the update succeeded
func UpdateUserDB(db *sql.DB, firstName, lastName string, age uint, newUserDB UserDB) error {
	_, err := db.Exec(
		fmt.Sprintf(
			"UPDATE %s SET %s=?,%s=?,%s=? WHERE %s=? AND %s=? AND %s=?",
			UserDBTableName,
			UserDBFirstNameCol,
			UserDBLastNameCol,
			UserDBAgeCol,
			UserDBFirstNameCol,
			UserDBLastNameCol,
			UserDBAgeCol,
		),
		newUserDB.FirstName,
		newUserDB.LastName,
		newUserDB.Age,
		firstName,
		lastName,
		age,
	)
	return err
}

// DeleteUserDB deletes the UserDB with the given first & last names and age. Returns a non-nil error if the delete failed, and nil if the delete succeeded
func DeleteUserDB(db *sql.DB, firstName, lastName string, age uint) error {
	_, err := db.Exec(
		fmt.Sprintf(
			"DELETE FROM %s WHERE %s=? AND %s=? AND %s=?",
			UserDBTableName,
			UserDBFirstNameCol,
			UserDBLastNameCol,
			UserDBAgeCol,
		),
		firstName,
		lastName,
		age,
	)
	return err
}
