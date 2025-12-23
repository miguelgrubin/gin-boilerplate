package pkg_test

import (
	"testing"

	"github.com/miguelgrubin/gin-boilerplate/pkg"
)

func TestMigrateAll(t *testing.T) {
	err := pkg.MigrateAll()
	if err != nil {
		t.Error("Migration failed:", err.Error())
	}
}

func TestSeedAll(t *testing.T) {
	err := pkg.SeedAll()
	if err != nil {
		t.Error("Seeding failed:", err.Error())
	}
}
