package actions

import (
	"testing"

	"github.com/gobuffalo/authrecipe/models"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/suite"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	s, err := suite.NewActionWithFixtures(App(), packr.NewBox("../fixtures"))
	if err != nil {
		t.Fatal(err)
	}
	as := &ActionSuite{s}
	suite.Run(t, as)
}

func (a *ActionSuite) Login() *models.User {
	a.LoadFixture("user")
	u := &models.User{}
	a.NoError(a.DB.First(u))
	a.Session.Set("current_user_id", u.ID)
	a.NoError(a.Session.Save())
	return u
}
