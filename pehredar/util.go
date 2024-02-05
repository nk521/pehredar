package pehredar

import (
	"log"
	"os"
	"path/filepath"
)

func getDefaultPath() string {
	confighome, present := os.LookupEnv("APPDATA")

	if !present {
		confighome, present = os.LookupEnv("XDG_CONFIG_HOME")
	}

	if !present {
		confighome = os.Getenv("HOME")
		if len(confighome) < 1 {
			log.Panic("couldn't determine a place to put config ... ")
		}
		confighome = filepath.Join(confighome, ".config")
	}

	return filepath.Join(confighome, "pehredar", "database.toml")
}

func GetOrCreateDefaultDatabasePath() string {
	database_path := getDefaultPath()

	if _, err := os.Stat(database_path); os.IsNotExist(err) {
		// create the folder for database
		os.MkdirAll(filepath.Dir(database_path), os.ModePerm)

		// init empty database file
		err = os.WriteFile(database_path, []byte{}, os.ModePerm)
		PanicIfNotNil(err)
	}

	return database_path
}

func PanicIfNotNil(e error) {
	if e != nil {
		log.Panicln(e)
	}
}
