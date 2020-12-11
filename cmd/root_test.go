package cmd

import (
	"log"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCmd(t *testing.T) {
	Convey("Test cmd root", t, func() {
		gin.SetMode(gin.TestMode)
		args := []string{}
		rootCmd.SetArgs(args)
		err := rootCmd.Execute()
		if err != nil {
			log.Fatal(err)
		}
	})
	Convey("Test cmd serve", t, func() {
		gin.SetMode(gin.TestMode)
		args := []string{"serve"}
		rootCmd.SetArgs(args)
		err := rootCmd.Execute()
		if err != nil {
			log.Fatal(err)
		}
	})
	Convey("Test cmd migrate exists", t, func() {
		args := []string{"migrate"}
		rootCmd.SetArgs(args)
		err := rootCmd.Execute()
		if err != nil {
			log.Fatal(err)
		}
	})
	Convey("Test cmd seed exists", t, func() {
		args := []string{"seed"}
		rootCmd.SetArgs(args)
		err := rootCmd.Execute()
		if err != nil {
			log.Fatal(err)
		}
	})
	Convey("Test cmd create exists", t, func() {
		args := []string{"create"}
		rootCmd.SetArgs(args)
		err := rootCmd.Execute()
		if err != nil {
			log.Fatal(err)
		}
	})
	Convey("Test cmd create config exists", t, func() {
		args := []string{"create", "config"}
		rootCmd.SetArgs(args)
		err := rootCmd.Execute()
		if err != nil {
			log.Fatal(err)
		}
	})
	Convey("Test Execute", t, func() {
		gin.SetMode(gin.TestMode)
		Execute()
	})
}
