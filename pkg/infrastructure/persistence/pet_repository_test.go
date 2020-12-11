package persistence

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPetRepositoy(t *testing.T) {

	Convey("Given a pet creation", t, func() {

		Convey("When pet entity has correct values", func() {

			Convey("It is created on dabatase", nil)

		})

		Convey("When pet entity has no values", func() {

			Convey("It is not created", nil)

			Convey("It retrieve a list of erros", nil)

		})

	})

	Convey("Given pet search by ID", t, func() {

		Convey("When it exists", func() {

			Convey("It is retrived", nil)

		})

		Convey("When it does not exists", func() {

			Convey("It is retrived a empty pet", nil)

		})

	})

	Convey("Given a pets fectch", t, func() {

		Convey("When it exists", func() {

			Convey("It is retrived", nil)

		})

		Convey("When it does not exists", func() {

			Convey("It is retrived a empty list of pets", nil)

		})

	})

}
