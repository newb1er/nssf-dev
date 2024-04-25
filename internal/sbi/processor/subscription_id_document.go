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
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/util/httpwrapper"
)

func (p *Processor) HTTPNSSAIAvailabilityUnsubscribe(c *gin.Context) {
	// Due to conflict of route matching, 'subscriptions' in the route is replaced with the existing wildcard ':nfId'
	nfID := c.Param("nfId")
	if nfID != "subscriptions" {
		c.JSON(http.StatusNotFound, gin.H{})
		logger.NssaiavailLog.Infof("404 Not Found")
		return
	}

	req := httpwrapper.NewRequest(c.Request, nil)
	req.Params["subscriptionId"] = c.Params.ByName("subscriptionId")

	subscriptionId := c.Params.ByName("subscriptionId")
	if subscriptionId == "" {
		problemDetails := &models.ProblemDetails{
			Status: http.StatusBadRequest,
			Cause:  "UNSPECIFIED", // TODO: Check if this is the correct cause
		}

		c.JSON(http.StatusBadRequest, problemDetails)
		return
	}

	problemDetails := nssaiavailability.SubscriptionUnsubscribe(subscriptionId)

	if problemDetails != nil {
		c.JSON(int(problemDetails.Status), problemDetails)
		return
	}

	c.Status(http.StatusNoContent)
}