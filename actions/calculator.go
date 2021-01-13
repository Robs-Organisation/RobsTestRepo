package actions

import (
	"buffalo_calc/models"
	"errors"
	"log"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
)

// HomeHandler is a default handler to serve up
// a home page.
func CalculatorHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("calculator.html"))
}

func input(c buffalo.Context) error {

	input := c.Param("task")
	res := calc(input)
	log.Println(res)
	c.Set("result", []string{res})

	a, err := pop.Connect("development")
	if err != nil {
		err := errors.New("no transection found")
		c.Flash().Add("danger", err.Error())
	}

	uid := c.Session().Get("current_user_id").(uuid.UUID)

	dataset := &models.History{
		Task:   input,
		Result: res,
		ByUser: uid,
	}

	err = a.Create(dataset)
	if err != nil {
		c.Flash().Add("danger", err.Error())
	}

	return c.Render(http.StatusOK, r.HTML("calculator.html"))
}
