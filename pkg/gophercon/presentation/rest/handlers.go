package rest

import (
	"net/http"

	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/domain"
	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/usecase"
	"github.com/gin-gonic/gin"
)

// AcceptedContentTypes is a list of all the accepted content types
var AcceptedContentTypes = []string{"application/json", "application/x-www-form-urlencoded"}

// PresentationHandlersImpl represents the presentation implementation object
type PresentationHandlersImpl struct {
	usecases usecase.IUsecase
}

// NewPresentationHandlers initializes a new rest handlers
func NewPresentationHandlers(u usecase.IUsecase) *PresentationHandlersImpl {
	return &PresentationHandlersImpl{
		usecases: u,
	}
}

func (h *PresentationHandlersImpl) HandleCreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := domain.User{}
		if err := c.ShouldBindJSON(&payload); err != nil {
			jsonErrorResponse(c, http.StatusBadRequest, err)
			return
		}

		output, err := h.usecases.CreateUser(c.Request.Context(), &payload)
		if err != nil {
			jsonErrorResponse(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, output)
	}
}

func jsonErrorResponse(c *gin.Context, statusCode int, err error) {
	c.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
}
