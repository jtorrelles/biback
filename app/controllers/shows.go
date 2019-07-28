package controllers

import (
	"biback/app/models"
	"biback/app/services"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"
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
	e.GET("/shows", handler.GetShows)
	e.GET("/shows/:id", handler.GetShowByID)
	e.POST("/shows", handler.Store)
	e.PUT("/shows/:id", handler.Update)
	//e.GET("/articles/:id", handler.GetByID)
	//e.DELETE("/articles/:id", handler.Delete)
}

// FetchArticle will fetch the article based on given params
func (a *ShowHandler) GetShows(c echo.Context) error {
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

func (a *ShowHandler) GetShowByID(c echo.Context) error {

	idS := c.Param("id")
	id, _ := strconv.Atoi(idS)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	listAr, err := a.ShowService.GetByID(ctx, int64(id))

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listAr)
}

func (a *ShowHandler) Store(c echo.Context) error {
	var show models.Show
	err := c.Bind(&show)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&show); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = a.ShowService.Store(ctx, &show)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, show)
}

func (a *ShowHandler) Update(c echo.Context) error {

	idS := c.Param("id")
	id, _ := strconv.Atoi(idS)

	var show models.Show
	err := c.Bind(&show)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&show); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = a.ShowService.Update(ctx, int64(id), &show)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, show)
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

func isRequestValid(m *models.Show) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
