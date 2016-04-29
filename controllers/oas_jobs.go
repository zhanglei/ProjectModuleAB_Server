package controllers

import (
	"fmt"
	"moduleab_server/models"
	"net/http"

	"github.com/astaxie/beego"
)

type OasJobsController struct {
	beego.Controller
}

// @Title getOAS
// @router /:job_id [get]
func (a *OasJobsController) Get() {
	jobId := a.GetString(":job_id")
	beego.Debug("[C] Got job id:", jobId)
	if jobId != "" {
		oasJob := &models.OasJobs{
			JobId: jobId,
		}
		oasJobs, err := models.GetOasJobs(oasJob, 0, 0)
		if err != nil {
			a.Data["json"] = map[string]string{
				"message": fmt.Sprint("Failed to get with job id:", jobId),
				"error":   err.Error(),
			}
			beego.Warn("[C] Got error:", err)
			a.Ctx.Output.SetStatus(http.StatusInternalServerError)
			a.ServeJSON()
			return
		}
		a.Data["json"] = oasJobs
		if len(oasJobs) == 0 {
			beego.Debug("[C] Got nothing with job id:", jobId)
			a.Ctx.Output.SetStatus(http.StatusNotFound)
			a.ServeJSON()
			return
		} else {
			a.Ctx.Output.SetStatus(http.StatusOK)
			a.ServeJSON()
			return
		}
	}
}

// @Title listOAS
// @router / [get]
func (a *OasJobsController) GetAll() {
	limit, _ := a.GetInt("limit", 0)
	index, _ := a.GetInt("index", 0)

	oasJob := &models.OasJobs{}
	oasJobs, err := models.GetOasJobs(oasJob, limit, index)
	if err != nil {
		a.Data["json"] = map[string]string{
			"message": fmt.Sprint("Failed to get"),
			"error":   err.Error(),
		}
		beego.Warn("[C] Got error:", err)
		a.Ctx.Output.SetStatus(http.StatusInternalServerError)
		a.ServeJSON()
		return
	}
	a.Data["json"] = oasJobs
	if len(oasJobs) == 0 {
		beego.Debug("[C] Got nothing")
		a.Ctx.Output.SetStatus(http.StatusNotFound)
		a.ServeJSON()
		return
	} else {
		a.Ctx.Output.SetStatus(http.StatusOK)
		a.ServeJSON()
		return
	}
}
