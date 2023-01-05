/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_cost
	@Date    2022/5/25 11:05
	@Desc
*/

package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/go-chi/render"
	"github.com/heshaofeng1991/common/util/httpresponse"
	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	chlLogic "github.com/heshaofeng1991/ddd-sample/interfaces/channel"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	UTC = "2006-01-02T15:04:05Z"
)

//nolint:revive
func (h HTTPServer) PostCostsChannelCostBatchIdUpload(w http.ResponseWriter, r *http.Request,
	channelCostBatchId int64,
) {
	rsp := chlLogic.UploadChannelCostRsp{
		Data: &chlLogic.UploadChannelCostData{},
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "read upload file error"+err.Error())

		return
	}

	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "open excel file error"+err.Error())

		return
	}

	results, err := ReadExcel(xlsx)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "parse excel file error"+err.Error())

		return
	}

	cost, err := h.application.Queries.GetChannelCostBatchByID.Handle(context.Background(), channelCostBatchId)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "query channel cost batch error"+err.Error())

		return
	}

	channelCosts := make([]*domainEntity.ChannelCost, 0)

	for _, val := range results {
		channelCost := domainEntity.NewChannelCost(
			0, channelCostBatchId, cost.ChannelID(), val.Mode(),
			val.CountryCode(), val.Zone(), val.StartWeight(), val.EndWeight(),
			val.FirstWeight(), val.FirstWeightFee(), val.UnitWeight(), val.UnitWeightFee(),
			val.FuelFee(), val.ProcessingFee(), val.RegistrationFee(), val.MiscFee(),
			time.Time{}, time.Time{}, nil, val.MinNormalDays(), val.MaxNormalDays(),
			1, val.AverageDays())

		channelCosts = append(channelCosts, channelCost)
	}

	if err = h.application.Commands.CreateChannelCost.Handle(context.Background(), channelCosts); err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "create channel cost batch error"+err.Error())

		return
	}

	rsp.Data.Status = 1

	render.Respond(w, r, rsp)

	return
}

const (
	BitSize = 10
)

func ReadExcel(xlsx *excelize.File) ([]*domainEntity.UploadChannel, error) {
	results := make([]*domainEntity.UploadChannel, 0)

	rows, err := xlsx.GetRows(xlsx.GetSheetName(xlsx.GetActiveSheetIndex()))
	if err != nil {
		return results, errors.Wrap(err, "")
	}

	if len(rows) == 1 {
		return results, errors.New("empty excel content")
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}

		startWeight, _ := strconv.Atoi(row[2])
		endWeight, _ := strconv.Atoi(row[3])
		firstWeight, _ := strconv.Atoi(row[4])
		firstWeightFee, _ := strconv.ParseFloat(row[5], BitSize)
		unitWeight, _ := strconv.Atoi(row[6])
		unitWeightFee, _ := strconv.ParseFloat(row[7], BitSize)
		fuelFee, _ := strconv.ParseFloat(row[8], BitSize)
		processingFee, _ := strconv.ParseFloat(row[9], BitSize)
		registrationFee, _ := strconv.ParseFloat(row[BitSize], BitSize)
		miscFee, _ := strconv.ParseFloat(row[11], BitSize)
		minNormalDays, _ := strconv.Atoi(row[15])
		maxNormalDays, _ := strconv.Atoi(row[16])
		averageDays, _ := strconv.Atoi(row[17])

		data := domainEntity.NewUploadChannel(
			row[0],
			row[1],
			startWeight,
			endWeight,
			firstWeight,
			firstWeightFee,
			unitWeight,
			unitWeightFee,
			fuelFee,
			processingFee,
			registrationFee,
			miscFee,
			Mode[row[12]],
			row[13],
			row[14],
			minNormalDays,
			maxNormalDays,
			averageDays)

		results = append(results, data)
	}

	return results, nil
}

var Mode = map[string]int8{
	"总价模式":  1,
	"单价模式":  2,
	"续单价模式": 3,
}

func (h HTTPServer) GetCostsDownloadTemplatesTemplateName(w http.ResponseWriter, r *http.Request, templateName string) {
	if templateName != "quotation" {
		httpresponse.ErrRsp(w, r, http.StatusBadRequest, "invalid template name")

		return
	}

	file, err := OpenTemplateExcel("quotation.xlsx")
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "open excel file error"+err.Error())

		return
	}

	data, length, err := WriteExcel(file, "Sheet1")
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "write excel error"+err.Error())

		return
	}

	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Disposition", "attachment; filename=quotation.xlsx")
	w.Header().Add("Content-Length", fmt.Sprintf("%d", length))

	if _, err = w.Write(data); err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "write response error"+err.Error())

		return
	}

	render.Respond(w, r, nil)

	return
}

func OpenTemplateExcel(filename string) (*excelize.File, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		logrus.Error("Get current directory failed", err)

		return nil, errors.Wrap(err, "")
	}

	templatePath := filepath.Join(rootDir, "resources/templates/"+filename)

	return OpenExcel(templatePath)
}

func OpenExcel(filePath string) (f *excelize.File, err error) {
	f, err = excelize.OpenFile(filePath)
	if err != nil {
		logrus.Error("open file error", err)

		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			logrus.Error("close file error: ", err)
		}
	}()

	return
}

func WriteExcel(f *excelize.File, sheetName string) (data []byte, length int64, err error) {
	index := f.GetSheetIndex(sheetName)
	f.SetActiveSheet(index)

	fileRes, err := f.WriteToBuffer()
	if err != nil {
		logrus.Errorf("write to buffer error: %v", err)

		return
	}

	reader := bytes.NewReader(fileRes.Bytes())

	return fileRes.Bytes(), reader.Size(), nil
}

//nolint:revive
func (h HTTPServer) GetChannelIdCostsChannelCostBatchId(w http.ResponseWriter, r *http.Request,
	channelId int64, channelCostBatchId int64,
	params chlLogic.GetChannelIdCostsChannelCostBatchIdParams,
) {
	rsp := chlLogic.ChannelCostSearchRsp{
		Data: &chlLogic.ChannelCostSearchInfo{
			Meta: chlLogic.MetaData{},
		},
	}

	pageSize, current := getPageSizeAndCurrent(params.Current, params.PageSize)

	results, total, err := h.application.Queries.GetChannelCosts.Handle(
		context.Background(),
		params.CountryCode,
		channelId,
		channelCostBatchId,
		current,
		pageSize)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "internal error get channel cost"+err.Error())

		return
	}

	for _, val := range results {
		unitWeight := val.UnitWeight()
		unitWeightFee := val.UnitWeightFee()

		cost := chlLogic.ChannelCostInfo{
			ChannelCode:        val.Channel().Code(),
			ChannelCostBatchId: val.ChannelCostBatchID(),
			ChannelId:          val.ChannelID(),
			CountryCode:        val.CountryCode(),
			EndWeight:          val.EndWeight(),
			FirstWeight:        val.FirstWeight(),
			FirstWeightFee:     val.FirstWeightFee(),
			FuelFee:            val.FuelFee(),
			Id:                 val.ID(),
			MaxNormalDays:      val.MaxNormalDays(),
			MinNormalDays:      val.MinNormalDays(),
			MiscFee:            val.MiscFee(),
			Mode:               val.Mode(),
			ProcessingFee:      val.ProcessingFee(),
			RegistrationFee:    val.RegistrationFee(),
			StartWeight:        val.StartWeight(),
			UnitWeight:         &unitWeight,
			UnitWeightFee:      &unitWeightFee,
			Zone:               val.Zone(),
			Status:             val.Status(),
		}

		rsp.Data.List = append(rsp.Data.List, cost)
	}

	rsp.Data.Meta.Total = total
	rsp.Data.Meta.Current = current
	rsp.Data.Meta.PageSize = pageSize

	render.Respond(w, r, rsp)

	return
}

func (h HTTPServer) PutCosts(w http.ResponseWriter, r *http.Request) {
	body := chlLogic.UpdateChannelCostStatusRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpresponse.ErrRsp(w, r, http.StatusBadRequest, "invalid request params"+err.Error())

		return
	}

	result, err := h.application.Commands.
		UpdateChannelCostStatus.
		Handle(context.Background(), body.Ids, body.CountryCodes, bool(body.Status))
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "internal error"+err.Error())

		return
	}

	render.Respond(w, r, chlLogic.UpdateChannelCostStatusRsp{
		Data: &chlLogic.UpdateCustomerConfigData{
			Status: result,
		},
	})
}
