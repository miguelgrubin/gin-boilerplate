package services_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"
	"github.com/stretchr/testify/suite"
)

type ConfigServiceTestSuite struct {
	configPath string
	suite.Suite
}

func (suite *ConfigServiceTestSuite) TestReadConfigWithValidEnv() {
	os.Setenv("APP_ENV", "test")

	cs := services.NewConfigServiceWithPath(&suite.configPath)
	config, err := cs.ReadConfig()

	suite.NoError(err)
	suite.NotNil(config)
}

func (suite *ConfigServiceTestSuite) TestReadConfigWithInvalidEnv() {
	os.Setenv("APP_ENV", "no-valid-env")

	cs := services.NewConfigServiceWithPath(&suite.configPath)
	config, err := cs.ReadConfig()

	suite.NoError(err)
	suite.NotNil(config)
}

func (suite *ConfigServiceTestSuite) TestReadConfigWithoutEnv() {
	os.Unsetenv("APP_ENV")

	cs := services.NewConfigServiceWithPath(&suite.configPath)
	config, err := cs.ReadConfig()

	suite.NoError(err)
	suite.NotNil(config)
	os.Setenv("APP_ENV", "test")
}

func (suite *ConfigServiceTestSuite) TestReadConfigWithNonExistantFile() {
	os.Setenv("APP_ENV", "local")
	os.Remove(suite.configPath + "/config_local.yaml")

	cs := services.NewConfigServiceWithPath(&suite.configPath)
	_, err := cs.ReadConfig()

	suite.Error(err)
	os.Setenv("APP_ENV", "test")
}

func (suite *ConfigServiceTestSuite) TestWriteConfig() {
	os.Setenv("APP_ENV", "local")
	os.Remove(suite.configPath + "/config_local.yaml")

	cs := services.NewConfigServiceWithPath(&suite.configPath)
	err := cs.WriteConfig()

	suite.NoError(err)
	os.Setenv("APP_ENV", "test")
}

func TestConfigService(t *testing.T) {
	csts := new(ConfigServiceTestSuite)
	tp, _ := filepath.Abs("../../../test/")
	csts.configPath = tp
	suite.Run(t, csts)
}
