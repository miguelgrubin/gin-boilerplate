package services_test

import (
	"log"
	"os"
	"testing"

	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"
	"github.com/stretchr/testify/suite"
)

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

type ConfigServiceTestSuite struct {
	suite.Suite
}

func (suite *ConfigServiceTestSuite) SetupTestSuite() {
	os.Chdir("../../test")
}

func (s *ConfigServiceTestSuite) TestReadConfigWithValidEnv() {
	os.Setenv("APP_ENV", "test")

	cs := services.NewConfigService()
	config, err := cs.ReadConfig()

	s.NoError(err)
	s.NotNil(config)
}

func (s *ConfigServiceTestSuite) TestReadConfigWithInvalidEnv() {
	os.Setenv("APP_ENV", "no-valid-env")

	cs := services.NewConfigService()
	config, err := cs.ReadConfig()

	s.NoError(err)
	s.NotNil(config)
}

func (s *ConfigServiceTestSuite) TestReadConfigWithoutEnv() {
	os.Unsetenv("APP_ENV")

	cs := services.NewConfigService()
	config, err := cs.ReadConfig()

	s.NoError(err)
	s.NotNil(config)
	os.Setenv("APP_ENV", "test")
}

func (s *ConfigServiceTestSuite) TestReadConfigWithNonExistantFile() {
	localPath := "config_local.yaml"
	tmpPath := "config_tmp.yaml"
	backupLocalConfig(localPath, tmpPath)
	os.Setenv("APP_ENV", "local")

	cs := services.NewConfigService()
	_, err := cs.ReadConfig()

	s.NoError(err)
	restoreLocalConfig(tmpPath, localPath)
	os.Setenv("APP_ENV", "test")
}

func (s *ConfigServiceTestSuite) TestWriteConfig() {
	os.Setenv("APP_ENV", "local")
	os.RemoveAll("config_local.yaml")

	cs := services.NewConfigService()
	err := cs.WriteConfig()

	s.NoError(err)
	os.Setenv("APP_ENV", "test")
	os.RemoveAll("config_local.yaml")
}

func TestConfigService(t *testing.T) {
	csts := new(ConfigServiceTestSuite)
	csts.SetupTestSuite()
	suite.Run(t, csts)
}
