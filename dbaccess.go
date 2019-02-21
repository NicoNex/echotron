package core

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	Host string `json:"host"`
	User string `json:"user"`
	Pass string `json:"password"`
	Name string `json:"dbname"`
}


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
	database, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", db.User, db.Pass, db.Host, db.Name))

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
	database, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/", db.User, db.Pass, db.Host))

	if err != nil {
		fmt.Println(err)
	} else {
		database.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", db.Name))
	}
	defer database.Close()
}


func (db *DB) CreateTable(name string, columns string) (string, error) {
	return db.executeQuery(fmt.Sprintf("CREATE TABLE %s (%s) DEFAULT CHARSET=utf8", name, columns))
}


func (db *DB) AddColumnToTable(tableName string, columnName string) (string, error) {
	return db.executeQuery(fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s", tableName, columnName))
}


func (db *DB) InsertRecord(tableName string, columns string, values string) (string, error) {
	return db.executeQuery(fmt.Sprintf("INSERT INTO %s (%s) VALUES (\"%s\")", tableName, columns, values))
}


func (db *DB) SelectRecord(tableName string, columnName string, value string) (string, error) {
	return db.executeQuery(fmt.Sprintf("SELECT %s FROM %s WHERE %s = \"%s\"", columnName, tableName, columnName, value))
}


func (db *DB) TableExists(tableName string) bool {
	queryResult, err := db.executeQuery(fmt.Sprintf("SELECT table_name FROM information_schema.tables WHERE table_name = \"%s\"", tableName))

	if err != nil {
		return false
	}

	if len(queryResult) > 0 {
		return true
	} else {
		return false
	}
}


func (db *DB) UpdateRecord(tableName string, columnName1 string, value1 string, columnName2 string, value2 string) (string, error) {
	return db.executeQuery(fmt.Sprintf("UPDATE %s SET %s = \"%s\" WHERE %s = \"%s\"", tableName, columnName1, value1, columnName2, value2))
}


func (db *DB) DeleteRecord(tableName string, columnName string, value string) (string, error) {
	return db.executeQuery(fmt.Sprintf("DELETE FROM %s WHERE %s = \"%s\"", tableName, columnName, value))
}
