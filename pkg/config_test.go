package pkg

import (
	"log"
	"os"
	"testing"
)

func useTestDir(t *testing.T) {
	err := os.Chdir("../test")
	if err != nil {
		t.Log(err.Error())
		t.Fatal("Can not change pwd to /test dir")
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func backupLocalConfig(src string, dst string) {
	if fileExists(src) {
		err := os.Rename(src, dst)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func restoreLocalConfig(src string, dst string) {
	if fileExists(src) {
		err := os.Rename(src, dst)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestReadConfigWithValidEnv(t *testing.T) {
	useTestDir(t)
	os.Setenv("APP_ENV", "test")

	config, err := ReadConfig()

	if err != nil {
		t.Log(err.Error())
		t.Fatal("Config could not be readed")
	}
	if !config.Testing {
		t.Fatal("Config testing could not be readed")
	}
}

func TestReadConfigWithInvalidEnv(t *testing.T) {
	useTestDir(t)
	os.Setenv("APP_ENV", "no-valid-env")

	config, err := ReadConfig()

	if err != nil {
		t.Log(err.Error())
		t.Fatal("Config could not be readed")
	}
	if config.Testing {
		t.Fatal("Default config could not be readed")
	}
}

func TestReadConfigWithoutEnv(t *testing.T) {
	useTestDir(t)
	os.Unsetenv("APP_ENV")

	config, err := ReadConfig()

	if err != nil {
		t.Log(err.Error())
		t.Fatal("Config could not be readed")
	}
	if config.Testing {
		t.Fatal("Default config could not be readed")
	}
	os.Setenv("APP_ENV", "test")
}

func TestReadConfigWithNonExistantFile(t *testing.T) {
	localPath := "config_local.yaml"
	tmpPath := "config_tmp.yaml"
	useTestDir(t)
	backupLocalConfig(localPath, tmpPath)
	os.Setenv("APP_ENV", "local")

	_, err := ReadConfig()

	if err == nil {
		t.Fatal("ReadConfig must return an error")
	}
	restoreLocalConfig(tmpPath, localPath)
	os.Setenv("APP_ENV", "test")
}

func TestWriteConfig(t *testing.T) {
	useTestDir(t)
	os.Setenv("APP_ENV", "local")
	os.RemoveAll("config_local.yaml")

	err := WriteConfig()

	if err != nil {
		t.Fatal("Config could not be writed")
	}
	os.Setenv("APP_ENV", "test")
	os.RemoveAll("config_local.yaml")
}
