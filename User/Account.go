package User

import (
	"database/sql"
	"fmt"
)

// Account Table

const (
	// AccountTableName is the name of the table for the Account model
	AccountTableName = "account"
	// AccountIdNumberCol is the column name of the id
	AccountIdNumberCol = "IdNumber"
	// AccountEmailCol is the column name of the email
	AccountEmailCol = "Email"
	// AccountDeviceIdCol is the column name of the DeviceId
	AccountDeviceIdCol = "DeviceId"
)

///Users/martinsinnuss/go/pkg/mod/github.com/arschles/go-in-5-minutes@v0.0.0-20200709150023-eb8196a64257/episode22/models/models.go

// Account is the database model for a Account
type Account struct {
	IdNumber int64
	Email    string
	DeviceId string
}

// CreateAccountTable uses db to create a new table for Account, and returns the result
func CreateAccountTable(db *sql.DB) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("CREATE TABLE %s (%s int, %s varchar(255), %s varchar(255))",
			AccountTableName,
			AccountIdNumberCol,
			AccountEmailCol,
			AccountDeviceIdCol,
		),
	)
}

// InsertAccount inserts Account into db
func InsertAccount(db *sql.DB, account Account) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("INSERT INTO %s VALUES($1, $2, $3)", AccountTableName),
		account.IdNumber,
		account.Email,
		account.DeviceId,
	)
}

// SelectAccount selects a Account with the given id and email and DeviceId. On success, writes the result into result and on failure, returns a non-nil error and makes no modifications to result
func SelectAccount(db *sql.DB, IdNumber int64, Email string, DeviceId string, result *Account) error {
	row := db.QueryRow(
		fmt.Sprintf(
			"SELECT * FROM %s WHERE %s=$1 AND %s=$2 AND %s=$3 ",
			AccountTableName,
			AccountIdNumberCol,
			AccountEmailCol,
			AccountDeviceIdCol,
		),
		IdNumber,
		Email,
		DeviceId,
	)

	var retEmail, retDeviceId string
	var retIdNumber int64
	if err := row.Scan(&retIdNumber, &retEmail, &retDeviceId); err != nil {
		return err
	}
	result.IdNumber = retIdNumber
	result.Email = retEmail
	result.DeviceId = retDeviceId
	return nil
}

// UpdateAccount updates the Account with the id, email and DeviceId with newAccount. Returns a non-nil error if the update failed, and nil if the update succeeded
func UpdateAccount(db *sql.DB, IdNumber int64, Email string, DeviceId string, newAccount Account) error {
	_, err := db.Exec(
		fmt.Sprintf(
			//"UPDATE %s SET %s,%s,%s WHERE %s AND %s AND %s",
			"UPDATE %s SET %s=$1,%s=$2,%s=$3 WHERE %s=$4 AND %s=$5 AND %s=$6",
			AccountTableName,
			AccountIdNumberCol,
			AccountEmailCol,
			AccountDeviceIdCol,
			AccountIdNumberCol,
			AccountEmailCol,
			AccountDeviceIdCol,
		),
		newAccount.IdNumber,
		newAccount.Email,
		newAccount.DeviceId,
		IdNumber,
		Email,
		DeviceId,
	)
	return err
}

// DeleteAccount deletes the Account with the given id, email and DeviceId. Returns a non-nil error if the delete failed, and nil if the delete succeeded
func DeleteAccount(db *sql.DB, IdNumber int64, Email string, DeviceId string) error {
	_, err := db.Exec(
		fmt.Sprintf(
			//"DELETE FROM $1 WHERE $2 AND $3 AND $4",
			"DELETE FROM %s WHERE %s=$1 AND %s=$2 AND %s=$3",
			AccountTableName,
			AccountIdNumberCol,
			AccountEmailCol,
			AccountDeviceIdCol,
		),
		IdNumber,
		Email,
		DeviceId,
	)
	return err
}

// DeviceId Table

const (
	// AccountTableName is the name of the table for the Account model
	DeviceIdTableName = "DeviceId"
	// AccountIdNumberCol is the column name of the id
	DeviceIdCol = "DeviceId"
)

type DeviceId struct {
	DeviceId string
}

func CreateDeviceIdTable(db *sql.DB) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("CREATE TABLE %s (%s varchar(255))",
			DeviceIdTableName,
			DeviceIdCol,
		),
	)
}

// Email Table

const (
	// AccountTableName is the name of the table for the Account model
	EmailTableName = "Email"
	// AccountIdNumberCol is the column name of the id
	EmailCol = "Email"
)

type Email struct {
	Email string
}

func CreateEmailTable(db *sql.DB) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("CREATE TABLE %s (%s varchar(255))",
			EmailTableName,
			EmailCol,
		),
	)
}
