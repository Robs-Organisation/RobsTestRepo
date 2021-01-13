package actions

import (
	"net/http"

	"buffalo_calc/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gofrs/uuid"
)

// HistoryDataHistory default implementation.
func HistoryDataHistory(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("history.html"))
}

func readData(c buffalo.Context) error {

	histories := models.Histories{}
	uid := c.Session().Get("current_user_id").(uuid.UUID)

	err := models.DB.Where("by_user=?", uid).All(&histories)
	if err != nil {
		c.Flash().Add("danger", err.Error())
	}

	//select * from histroys where useruuid='uid';

	// err := models.DB.All(&histories)
	// if err != nil {
	// 	return errors.WithMessage(err, "database err")
	// }

	c.Set("histories", histories)

	return c.Render(http.StatusOK, r.HTML("history.html"))
}
