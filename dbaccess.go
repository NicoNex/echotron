/*
 * Echotron
 * Copyright (C) 2019  Nicolò Santamaria, Michele Dimaggio
 *
 * Echotron is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Echotron is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package echotron


import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	_ "github.com/go-sql-driver/mysql"
)


// DB is the struct which contains the hostname of the database,
// the login data and the name of the database that will be used
// for the various operations.
type DB struct {
	host string `json:"host"`
	user string `json:"user"`
	pass string `json:"password"`
	name string `json:"dbname"`
}


// NewDB creates a new DB instance.
// It reads the data needed by the DB struct from a file (credFile),
// stores them in the DB struct, tries to create a new database if
// it doesn't exist and returns a DB instance.
func NewDB(credFile string) (*DB, error) {
	jsonFile, err := ioutil.ReadFile(credFile)

	if err != nil {
		return nil, err
	}

	var db DB
	json.Unmarshal(jsonFile, &db)

	db.createDB()

	return &db, nil
}


func (db *DB) openDB() (*sql.DB, error) {
	database, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", db.user, db.pass, db.host, db.name))

	if err != nil {
		return nil, err
	}

	return database, nil
}


func (db *DB) executeQuery(query string) (string, error) {
	database, err := db.openDB()

	if err != nil {
		return "", err
	}
	defer database.Close()

	rows, err := database.Query(query)

	if err != nil {
		return "", err
	}
	defer rows.Close()

	var queryResult string

	for rows.Next() {
		err := rows.Scan(&queryResult)

		if err != nil {
			return "", err
		}
	}

	return queryResult, nil
}


func (db *DB) createDB() {
	database, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/", db.user, db.pass, db.host))

	if err != nil {
		fmt.Println(err)
	} else {
		database.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", db.name))
	}
	defer database.Close()
}


// CreateTable creates a new table in the database.
func (db *DB) CreateTable(name string, columns string) (string, error) {
	return db.executeQuery(fmt.Sprintf("CREATE TABLE %s (%s) DEFAULT CHARSET=utf8", name, columns))
}


// AddColumnToTable adds a new column to a specific table.
func (db *DB) AddColumnToTable(tableName string, columnName string) (string, error) {
	return db.executeQuery(fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s", tableName, columnName))
}


// InsertRecord inserts data into the database.
func (db *DB) InsertRecord(tableName string, columns string, values string) (string, error) {
	return db.executeQuery(fmt.Sprintf("INSERT INTO %s (%s) VALUES (\"%s\")", tableName, columns, values))
}


// SelectRecord reads data from the database and returns it.
func (db *DB) SelectRecord(tableName string, columnName string, value string) (string, error) {
	return db.executeQuery(fmt.Sprintf("SELECT %s FROM %s WHERE %s = \"%s\"", columnName, tableName, columnName, value))
}


// Similar to SelectRecord, but in addition allows
// to compare some known data with a specific column.
func (db *DB) SelectRecordCompared(tableName string, columnName string, columnToCompare string, value string) (string, error) {
	return db.executeQuery(fmt.Sprintf("SELECT %s FROM %s WHERE %s = \"%s\"", columnName, tableName, columnToCompare, value))
}


// TableExists is used to check whether a specific table exists.
func (db *DB) TableExists(tableName string) bool {
	queryResult, err := db.executeQuery(fmt.Sprintf("SELECT table_name FROM information_schema.tables WHERE table_name = \"%s\"", tableName))

	if err != nil {
		return false
	}

	return len(queryResult) > 0
}


// UpdateRecord edits some specific data in the database.
func (db *DB) UpdateRecord(tableName string, columnName1 string, value1 string, columnName2 string, value2 string) (string, error) {
	return db.executeQuery(fmt.Sprintf("UPDATE %s SET %s = \"%s\" WHERE %s = \"%s\"", tableName, columnName1, value1, columnName2, value2))
}


// DeleteRecord deletes some specific data from the database.
func (db *DB) DeleteRecord(tableName string, columnName string, value string) (string, error) {
	return db.executeQuery(fmt.Sprintf("DELETE FROM %s WHERE %s = \"%s\"", tableName, columnName, value))
}
