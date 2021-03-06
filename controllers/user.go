package controllers

import (
	"net/http"

	"github.com/deputadosemfoco/go-libs/messages"
	"github.com/deputadosemfoco/users/interactors"
	"github.com/labstack/echo"
)

// UserCtrl ...
type UserCtrl struct {
	Interactor interactors.RegistrationInteractorContract
}

// Post 
func (ctrl *UserCtrl) Post(c echo.Context) error {
	req := new(interactors.RegistrationRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	res, err := ctrl.Interactor.Register(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, messages.Error{Message: err.Error()})
	}

	code := http.StatusOK
	if res.Created {
		code = http.StatusCreated
	}

	return c.JSON(code, res.User)
}
