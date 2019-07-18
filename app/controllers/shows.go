package controllers

import (
	"biback/app/services"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ArticleHandler  represent the httphandler for article
type ShowHandler struct {
	ShowService services.Showservice
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewShowHandler(e *echo.Echo, us services.Showservice) {
	handler := &ShowHandler{
		ShowService: us,
	}
	e.GET("/shows", handler.FetchShow)
	//e.POST("/articles", handler.Store)
	//e.GET("/articles/:id", handler.GetByID)
	//e.DELETE("/articles/:id", handler.Delete)
}

// FetchArticle will fetch the article based on given params
func (a *ShowHandler) FetchShow(c echo.Context) error {
	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)
	cursor := c.QueryParam("cursor")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listAr, nextCursor, err := a.ShowService.Fetch(ctx, cursor, int64(num))

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, listAr)
}

/*func GetShows(c echo.Context) error {
	result := models.GetShows()
	return c.JSON(http.StatusOK, result)
}

func GetShowById(c echo.Context) error {
	result := models.GetShow(c.Param("id"))
	return c.JSON(http.StatusOK, result)
}

func NewShow(c echo.Context) (err error) {
	show := &models.Show{}
	if err = c.Bind(show); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	result := show.AddShow(show)

	return c.JSON(http.StatusCreated, result)
}*/

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	} else {
		return http.StatusInternalServerError
	}
	/*logrus.Error(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}*/
}
