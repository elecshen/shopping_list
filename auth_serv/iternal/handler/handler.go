package handler

import (
	"github.com/elecshen/auth_service/iternal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/errors"
	oaSrv "github.com/go-oauth2/oauth2/v4/server"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Handler struct {
	services     *service.Service
	oauthHandler *oaSrv.Server
}

func NewHandler(services *service.Service, oauthHandler *oaSrv.Server) *Handler {
	return &Handler{services: services, oauthHandler: oauthHandler}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	oauth := router.Group("/oauth")
	{
		oauth.GET("/authorize", h.userIdentity, h.oauthAuthorize)
		oauth.POST("/token", h.oauthToken)
		oauth.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, struct{}{})
		})
	}

	h.oauthHandler.SetUserAuthorizationHandler(func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
		idInt := r.Context().Value(userCtx)
		if idInt == nil {
			return "", errors.New("user id not found")
		}

		userID = strconv.Itoa(idInt.(int))
		return
	})

	h.oauthHandler.SetClientInfoHandler(func(r *http.Request) (clientID, clientSecret string, err error) {
		clientID = r.FormValue("client_id")
		clientSecret = r.FormValue("client_secret")
		return
	})

	h.oauthHandler.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		logrus.Println("Internal Error:", err.Error())
		return
	})
	h.oauthHandler.SetResponseErrorHandler(func(re *errors.Response) {
		logrus.Println("Response Error:", re.Error.Error())
	})

	return router
}
