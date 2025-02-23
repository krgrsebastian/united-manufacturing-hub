package controllers

/*


import (
	"github.com/gin-gonic/gin"
	"github.com/united-manufacturing-hub/united-manufacturing-hub/cmd/factoryinsight/helpers"
	"github.com/united-manufacturing-hub/united-manufacturing-hub/cmd/factoryinsight/v3/models"
	"github.com/united-manufacturing-hub/united-manufacturing-hub/cmd/factoryinsight/v3/services"
	"net/http"
)

func GetTagGroupsHandler(c *gin.Context) {
	var request models.GetTagGroupsRequest
	var tagGroups models.GetTagGroupsResponse

	err := c.BindUri(&request)
	if err != nil {
		helpers.HandleInvalidInputError(c, err)
		return
	}

	// Check if the user has access to that resource
	err = helpers.CheckIfUserIsAllowed(c, request.EnterpriseName)
	if err != nil {
		return
	}

	// Fetch data from database
	tagGroups, err = services.GetTagGroups(
		request.EnterpriseName,
		request.SiteName,
		request.AreaName,
		request.ProductionLineName,
		request.WorkCellName)
	if err != nil {
		helpers.HandleInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, tagGroups)
}

func GetTagsHandler(c *gin.Context) {
	var request models.GetTagsRequest
	var tags models.GetTagsResponse

	err := c.BindUri(&request)
	if err != nil {
		helpers.HandleInvalidInputError(c, err)
		return
	}

	// Check if the user has access to that resource
	err = helpers.CheckIfUserIsAllowed(c, request.EnterpriseName)
	if err != nil {
		return
	}

	switch request.TagGroupName {
	case models.CustomTagGroup:
		tags, err = services.GetStandardTags()
	case models.StandardTagGroup:
		tags, err = services.GetCustomTags(
			request.EnterpriseName,
			request.SiteName,
			request.AreaName,
			request.ProductionLineName,
			request.WorkCellName,
			request.TagGroupName)
	default:
		helpers.HandleInvalidInputError(c, err)
		return
	}

	if err != nil {
		helpers.HandleInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, tags)
}

func GetTagsDataHandler(c *gin.Context) {
	var request models.GetTagsDataRequest

	err := c.BindUri(&request)
	if err != nil {
		helpers.HandleInvalidInputError(c, err)
		return
	}

	// Check if the user has access to that resource
	err = helpers.CheckIfUserIsAllowed(c, request.EnterpriseName)
	if err != nil {
		return
	}

	switch request.TagGroupName {
	case models.StandardTagGroup:
		switch request.TagName {
		case models.JobsStandardTag:
			services.ProcessJobTagRequest(c, request)
		case models.OutputStandardTag:
			services.ProcessOutputTagRequest(c, request)
		case models.ShiftsStandardTag:
			services.ProcessShiftsTagRequest(c, request)
		case models.StateStandardTag:
			services.ProcessStateTagRequest(c, request)
		case models.ThroughputStandardTag:
			services.ProcessThroughputTagRequest(c, request)

		default:
			helpers.HandleInvalidInputError(c, err)
			return
		}
	case models.CustomTagGroup:
		// TODO: Implement custom tags
	default:
		helpers.HandleInvalidInputError(c, err)
		return
	}
}
*/
