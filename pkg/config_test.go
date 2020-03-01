package pkg

import (
	"log"
	"os"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReadConfig(t *testing.T) {

	Convey("Given existing enviroment", t, func() {
		os.Setenv("APP_ENV", "test")
		config, err := ReadConfig()
		Convey("It should retrieve their config", func() {
			if err != nil {
				t.Log(err.Error())
				t.Fatal("Config could not be readed")
			}
			if !config.testing {
				t.Fatal("Config testing could not be readed")
			}
		})
		Convey("It should have a server config", func() {
			gotType := reflect.TypeOf(config.server).String()
			wantType := "pkg.ServerConfig"
			if gotType != wantType {
				t.Fatalf("Invalid type: got %v, want %v", gotType, wantType)
			}
		})
		Convey("It should have a database config", func() {
			gotType := reflect.TypeOf(config.database).String()
			wantType := "pkg.DatabaseConfig"
			if gotType != wantType {
				t.Fatalf("Invalid type: got %v, want %v", gotType, wantType)
			}
		})
	})

	Convey("Given unexistent enviroment", t, func() {
		os.Unsetenv("APP_ENV")
		config, err := ReadConfig()
		Convey("It should retrieve a default config", func() {
			if err != nil {
				t.Log(err.Error())
				t.Fatal("Config could not be readed")
			}
			if config.testing {
				t.Fatal("Config testing could not be readed")
			}
		})
		os.Setenv("APP_ENV", "test")
	})

	Convey("Given enviroment with unexistent config file", t, func() {
		local_path := "config/config_local.yaml"
		tmp_path := "config/config_tmp.yaml"
		os.Setenv("APP_ENV", "local")
		if fileExists(local_path) {
			err := os.Rename(local_path, tmp_path)
			if err != nil {
				log.Fatal(err)
			}
		}

		Convey("It should retrieve an error", func() {
			_, err := ReadConfig()
			if err == nil {
				t.Fatal("ReadConfig must return an error")
			}
		})

		if fileExists(tmp_path) {
			err := os.Rename(tmp_path, local_path)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.Setenv("APP_ENV", "test")
	})
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
