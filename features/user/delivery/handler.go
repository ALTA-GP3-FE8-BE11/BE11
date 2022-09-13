package delivery

import (
	"net/http"
	"project/kutsuya/features/user"
	"project/kutsuya/utils/helper"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userUsecase user.UsecaseInterface
}

func New(e *echo.Echo, usecase user.UsecaseInterface) {
	handler := &UserDelivery{
		userUsecase: usecase,
	}

	// e.GET("/users", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/login", handler.GetLoginUser)
	e.POST("/users", handler.PostData)

}

func (delivery *UserDelivery) PostData(c echo.Context) error {
	var userRequestData UserRequest
	errBind := c.Bind(&userRequestData)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind user data"))
	}

	row, err := delivery.userUsecase.PostData(ToCore(userRequestData))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail input user data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("insert row affected is not 1"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("post new user data"))

}

func (delivery *UserDelivery) GetLoginUser(c echo.Context) error {
	var userRequest_Login UserRequest
	errBind := c.Bind(&userRequest_Login)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind user data"))
	}

	Token_JWT, err := delivery.userUsecase.GetLogin(ToCore(userRequest_Login))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail input user data"))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("get all user data", Token_JWT))

}
