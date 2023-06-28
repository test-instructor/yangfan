package interfacecase

import (
	"encoding/json"

	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"gorm.io/datatypes"
)

func resetApiCaseStep(apicase ...*interfacecase.ApiStep) {
	for i := 0; i < len(apicase); i++ {
		resetApiCaseStepModel(apicase[i])
		if apicase[i].Request != nil {
			resetRequest(apicase[i].Request)
		}
		if apicase[i].Grpc != nil {
			resetGrpc(apicase[i].Grpc)
		}
	}
}

func resetReport(report *interfacecase.ApiReport) {
	for i := 0; i < len(report.Details); i++ {
		report.Details[i].RootDir = ""
		report.Details[i].InOut = nil
		for j := 0; j < len(report.Details[i].Records); j++ {
			report.Details[i].Records[j].ExportVars = nil
			for k := 0; k < len(report.Details[i].Records[j].Data); k++ {
				report.Details[i].Records[j].Data[k].ExportVars = nil
				report.Details[i].Records[j].Data[k].HttpStat = nil
				data, err := resetReportDetailData(report.Details[i].Records[j].Data[k].Data)
				if err == nil {
					report.Details[i].Records[j].Data[k].Data = data
				}
			}
		}
	}
}

func resetApiCaseStepModel(apicase *interfacecase.ApiStep) {
	apicase.CreatedBy = nil
	apicase.UpdateBy = nil
	apicase.DeleteBy = nil
	apicase.ProjectID = 0
	apicase.ApiType = 0
	apicase.Transaction = nil
	apicase.Rendezvous = nil
	apicase.ThinkTime = nil
	apicase.ThinkTimeID = 0
	apicase.TransactionID = 0
	apicase.RendezvousID = 0
	apicase.RequestID = 0
	apicase.GrpcID = 0
	apicase.Variables = nil
	apicase.Extract = nil
	apicase.Validate = nil
	apicase.ValidateJson = nil
	apicase.ExtractJson = nil
	apicase.VariablesJson = nil
	apicase.SetupHooks = nil
	apicase.TeardownHooks = nil
	apicase.ExportHeader = nil
	apicase.ExportParameter = nil
}

func resetGrpc(gRPC *interfacecase.ApiGrpc) {
	gRPC.Headers = nil
	gRPC.HeadersJson = nil
	gRPC.Body = nil
	gRPC.Detail = nil
}

func resetRequest(request *interfacecase.ApiRequest) {
	request.Agreement = ""
	request.Params = nil
	request.Headers = nil
	request.Data = nil
	request.ParamsJson = nil
	request.HeadersJson = nil
	request.DataJson = nil
	request.Json = nil
	request.Timeout = 0
	request.AllowRedirects = false
	request.Verify = false
}

type DataDetail struct {
	Success  bool     `json:"success"`
	ReqResps ReqResps `json:"req_resps"`
}

type ReqResps struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}

type Request struct {
	URL    string `json:"url"`
	Method string `json:"method"`
}

type Response struct {
	Proto string `json:"proto"`
}

func resetReportDetailData(data datatypes.JSON) (dataNew datatypes.JSON, err error) {
	var dataStruct DataDetail
	err = json.Unmarshal(data, &dataStruct)
	if err != nil {
		global.GVA_LOG.Warn("格式转换错误")
		return
	}

	dataNew, err = json.Marshal(dataStruct)
	if err != nil {
		global.GVA_LOG.Warn("格式转换错误")
		return
	}

	return
}
