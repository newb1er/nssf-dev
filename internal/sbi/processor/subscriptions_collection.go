/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package processor

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/free5gc/nssf/internal/logger"
	"github.com/free5gc/nssf/internal/sbi/nssaiavailability"
	"github.com/free5gc/openapi"
	. "github.com/free5gc/openapi/models"
)

func (p *Processor) HTTPNSSAIAvailabilityPost(c *gin.Context) {
	var createData NssfEventSubscriptionCreateData

	requestBody, err := c.GetRawData()
	if err != nil {
		problemDetail := ProblemDetails{
			Title:  "System failure",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Cause:  "SYSTEM_FAILURE",
		}
		logger.NssaiavailLog.Errorf("Get Request Body error: %+v", err)
		c.JSON(http.StatusInternalServerError, problemDetail)
		return
	}

	err = openapi.Deserialize(&createData, requestBody, "application/json")
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.NssaiavailLog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	createdData, problemDetails := nssaiavailability.SubscriptionCreate(createData)

	if problemDetails != nil {
		c.JSON(int(problemDetails.Status), problemDetails)
		return
	}

	// TODO: Based on TS 29.531 5.3.2.3.1, add location header

	c.JSON(http.StatusCreated, createdData)
}