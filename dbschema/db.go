package dbschema

import (
	"database/sql"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	gwisql "gwi/platform2.0-go-challenge/internal/repositories/sql"

	"gwi/platform2.0-go-challenge/environment"
)

// DBSeed : Seeds `/seed/`.
func DBSeed() {
	db := newClient()

	rootDir := rootDir()
	path := rootDir + "/dbschema/seed/"

	files, err := allSeedFilenames()
	if err != nil {
		log.Fatalf("Error in db seed %s", err.Error())
	}

	for _, f := range files {
		execSQLFile(db, path+"/"+f)
	}
}

// DBClear : Clears all data from db.
func DBClear() {
	db := newClient()

	files, err := allSeedFilenames()
	if err != nil {
		log.Fatalf("Error in db seed clear %s", err.Error())
	}

	for _, f := range files {
		tableName := strings.ReplaceAll(f, ".seed.sql", "")
		_, err := db.Exec("DELETE FROM " + tableName)
		if err != nil {
			log.Printf("Error in db seed clear %s", err.Error())
		}
	}
}

// DBObjects : Creates db schema - objects.
func DBObjects() {
	db := newClient()

	rootDir := rootDir()
	path := rootDir + "/dbschema/objects/"

	files, err := allObjectFilenames()
	if err != nil {
		log.Fatalf("Error in db objects creation %s", err.Error())
	}

	for _, f := range files {
		execSQLFile(db, path+"/"+f)
	}
}

func newClient() *sql.DB {
	conf := environment.LoadConfig()
	client, err := gwisql.NewDBClient(conf)
	if err != nil {
		log.Fatal(err)
	}

	return client.DB
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func allSeedFilenames() ([]string, error) {
	path := rootDir() + "/dbschema/seed/"
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	filenames := make([]string, 0, 10)
	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".sql" {
			filenames = append(filenames, f.Name())
		}
	}

	return filenames, nil
}

func allObjectFilenames() ([]string, error) {
	path := rootDir() + "/dbschema/objects/"
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	filenames := make([]string, 0, 10)
	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".sql" {
			filenames = append(filenames, f.Name())
		}
	}

	return filenames, nil
}

func execSQLFile(db *sql.DB, f string) {
	query, err := os.ReadFile(f)
	if err != nil {
		log.Fatalf("read query from file failed: %s", err.Error())
	}

	if len(query) == 0 {
		return
	}

	_, err = db.Exec(string(query))
	if err != nil {
		log.Fatalf("exec sql file failed: %s", err.Error())
	}
}
