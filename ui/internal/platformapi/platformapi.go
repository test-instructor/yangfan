package platformapi

import (
	"context"
	"errors"
	"net/http"

	"yangfan-ui/internal/platformclient"
)

type Endpoint string

const (
	EndpointHealthAPI      Endpoint = "health_api"
	EndpointHealth         Endpoint = "health"
	EndpointCaptcha        Endpoint = "captcha"
	EndpointLogin          Endpoint = "login"
	EndpointUserInfo       Endpoint = "user_info"
	EndpointSetAuth        Endpoint = "set_user_authority"
	EndpointSetSelf        Endpoint = "set_self_info"
	EndpointChangePwd      Endpoint = "change_password"
	EndpointUINodeMenuTree Endpoint = "ui_node_menu_tree"

	EndpointAndroidDeviceOptionsList   Endpoint = "android_device_options_list"
	EndpointAndroidDeviceOptionsCreate Endpoint = "android_device_options_create"
	EndpointAndroidDeviceOptionsUpdate Endpoint = "android_device_options_update"
	EndpointAndroidDeviceOptionsDelete Endpoint = "android_device_options_delete"

	EndpointRunConfigList   Endpoint = "runconfig_list"
	EndpointRunConfigCreate Endpoint = "runconfig_create"
	EndpointRunConfigUpdate Endpoint = "runconfig_update"
	EndpointRunConfigDelete Endpoint = "runconfig_delete"

	EndpointAutoStepList   Endpoint = "autostep_list"
	EndpointAutoStepCreate Endpoint = "autostep_create"
	EndpointAutoStepUpdate Endpoint = "autostep_update"
	EndpointAutoStepDelete Endpoint = "autostep_delete"

	EndpointAutoCaseStepList      Endpoint = "autocasestep_list"
	EndpointAutoCaseStepCreate    Endpoint = "autocasestep_create"
	EndpointAutoCaseStepUpdate    Endpoint = "autocasestep_update"
	EndpointAutoCaseStepDelete    Endpoint = "autocasestep_delete"
	EndpointAutoCaseStepFindApis  Endpoint = "autocasestep_find_apis"
	EndpointAutoCaseStepAddApi    Endpoint = "autocasestep_add_api"
	EndpointAutoCaseStepDeleteApi Endpoint = "autocasestep_delete_api"
	EndpointAutoCaseStepSortApis  Endpoint = "autocasestep_sort_apis"

	EndpointAutoCaseList          Endpoint = "autocase_list"
	EndpointAutoCaseCreate        Endpoint = "autocase_create"
	EndpointAutoCaseUpdate        Endpoint = "autocase_update"
	EndpointAutoCaseDelete        Endpoint = "autocase_delete"
	EndpointAutoCaseGetSteps      Endpoint = "autocase_get_steps"
	EndpointAutoCaseAddStep       Endpoint = "autocase_add_step"
	EndpointAutoCaseDelStep       Endpoint = "autocase_del_step"
	EndpointAutoCaseSortSteps     Endpoint = "autocase_sort_steps"
	EndpointAutoCaseSetStepConfig Endpoint = "autocase_set_step_config"

	EndpointTimerTaskList      Endpoint = "timertask_list"
	EndpointTimerTaskCreate    Endpoint = "timertask_create"
	EndpointTimerTaskUpdate    Endpoint = "timertask_update"
	EndpointTimerTaskDelete    Endpoint = "timertask_delete"
	EndpointTimerTaskGetCases  Endpoint = "timertask_get_cases"
	EndpointTimerTaskAddCase   Endpoint = "timertask_add_case"
	EndpointTimerTaskDelCase   Endpoint = "timertask_del_case"
	EndpointTimerTaskSortCases Endpoint = "timertask_sort_cases"

	EndpointAutoReportList Endpoint = "autoreport_list"
	EndpointAutoReportFind Endpoint = "autoreport_find"
)

type Spec struct {
	Method string
	Path   string
}

var specs = map[Endpoint]Spec{
	EndpointHealthAPI: {Method: http.MethodGet, Path: "/api/health/yangfan/server"},
	EndpointHealth:    {Method: http.MethodGet, Path: "/health"},
	EndpointCaptcha:   {Method: http.MethodPost, Path: "/api/base/captcha"},
	EndpointLogin:     {Method: http.MethodPost, Path: "/api/base/login"},
	EndpointUserInfo:  {Method: http.MethodGet, Path: "/api/user/getUserInfo"},
	EndpointSetAuth:   {Method: http.MethodPost, Path: "/api/user/setUserAuthority"},
	EndpointSetSelf:   {Method: http.MethodPut, Path: "/api/user/setSelfInfo"},
	EndpointChangePwd: {Method: http.MethodPost, Path: "/api/user/changePassword"},

	EndpointUINodeMenuTree: {Method: http.MethodPost, Path: "/api/ui/node/menu/getMenuTree"},

	EndpointAndroidDeviceOptionsList:   {Method: http.MethodGet, Path: "/api/ado/getAndroidDeviceOptionsList"},
	EndpointAndroidDeviceOptionsCreate: {Method: http.MethodPost, Path: "/api/ado/createAndroidDeviceOptions"},
	EndpointAndroidDeviceOptionsUpdate: {Method: http.MethodPut, Path: "/api/ado/updateAndroidDeviceOptions"},
	EndpointAndroidDeviceOptionsDelete: {Method: http.MethodDelete, Path: "/api/ado/deleteAndroidDeviceOptions"},

	EndpointRunConfigList:   {Method: http.MethodGet, Path: "/api/rc/getRunConfigList"},
	EndpointRunConfigCreate: {Method: http.MethodPost, Path: "/api/rc/createRunConfig"},
	EndpointRunConfigUpdate: {Method: http.MethodPut, Path: "/api/rc/updateRunConfig"},
	EndpointRunConfigDelete: {Method: http.MethodDelete, Path: "/api/rc/deleteRunConfig"},

	EndpointAutoStepList:   {Method: http.MethodGet, Path: "/api/as/getAutoStepList"},
	EndpointAutoStepCreate: {Method: http.MethodPost, Path: "/api/as/createAutoStep"},
	EndpointAutoStepUpdate: {Method: http.MethodPut, Path: "/api/as/updateAutoStep"},
	EndpointAutoStepDelete: {Method: http.MethodDelete, Path: "/api/as/deleteAutoStep"},

	EndpointAutoCaseStepList:      {Method: http.MethodGet, Path: "/api/acs/getAutoCaseStepList"},
	EndpointAutoCaseStepCreate:    {Method: http.MethodPost, Path: "/api/acs/createAutoCaseStep"},
	EndpointAutoCaseStepUpdate:    {Method: http.MethodPut, Path: "/api/acs/updateAutoCaseStep"},
	EndpointAutoCaseStepDelete:    {Method: http.MethodDelete, Path: "/api/acs/deleteAutoCaseStep"},
	EndpointAutoCaseStepFindApis:  {Method: http.MethodGet, Path: "/api/acs/findAutoCaseStepApi"},
	EndpointAutoCaseStepAddApi:    {Method: http.MethodPost, Path: "/api/acs/addAutoCaseStepApi"},
	EndpointAutoCaseStepDeleteApi: {Method: http.MethodDelete, Path: "/api/acs/deleteAutoCaseStepApi"},
	EndpointAutoCaseStepSortApis:  {Method: http.MethodPost, Path: "/api/acs/sortAutoCaseStepApi"},

	EndpointAutoCaseList:          {Method: http.MethodGet, Path: "/api/ac/getAutoCaseList"},
	EndpointAutoCaseCreate:        {Method: http.MethodPost, Path: "/api/ac/createAutoCase"},
	EndpointAutoCaseUpdate:        {Method: http.MethodPut, Path: "/api/ac/updateAutoCase"},
	EndpointAutoCaseDelete:        {Method: http.MethodDelete, Path: "/api/ac/deleteAutoCase"},
	EndpointAutoCaseGetSteps:      {Method: http.MethodGet, Path: "/api/ac/getAutoCaseSteps"},
	EndpointAutoCaseAddStep:       {Method: http.MethodPost, Path: "/api/ac/addAutoCaseStep"},
	EndpointAutoCaseDelStep:       {Method: http.MethodDelete, Path: "/api/ac/delAutoCaseStep"},
	EndpointAutoCaseSortSteps:     {Method: http.MethodPost, Path: "/api/ac/sortAutoCaseStep"},
	EndpointAutoCaseSetStepConfig: {Method: http.MethodPut, Path: "/api/ac/setStepConfig"},

	EndpointTimerTaskList:      {Method: http.MethodGet, Path: "/api/tk/getTimerTaskList"},
	EndpointTimerTaskCreate:    {Method: http.MethodPost, Path: "/api/tk/createTimerTask"},
	EndpointTimerTaskUpdate:    {Method: http.MethodPut, Path: "/api/tk/updateTimerTask"},
	EndpointTimerTaskDelete:    {Method: http.MethodDelete, Path: "/api/tk/deleteTimerTask"},
	EndpointTimerTaskGetCases:  {Method: http.MethodGet, Path: "/api/tk/getTimerTaskCases"},
	EndpointTimerTaskAddCase:   {Method: http.MethodPost, Path: "/api/tk/addTimerTaskCase"},
	EndpointTimerTaskDelCase:   {Method: http.MethodDelete, Path: "/api/tk/delTimerTaskCase"},
	EndpointTimerTaskSortCases: {Method: http.MethodPost, Path: "/api/tk/sortTimerTaskCase"},

	EndpointAutoReportList: {Method: http.MethodGet, Path: "/api/ar/getAutoReportList"},
	EndpointAutoReportFind: {Method: http.MethodGet, Path: "/api/ar/findAutoReport"},
}

func GetSpec(e Endpoint) (Spec, error) {
	s, ok := specs[e]
	if !ok {
		return Spec{}, errors.New("未知的接口标识")
	}
	if s.Method == "" || s.Path == "" {
		return Spec{}, errors.New("接口配置不完整")
	}
	return s, nil
}

func Call(ctx context.Context, client *platformclient.Client, endpoint Endpoint, reqBody any, extraHeaders map[string]string) (http.Header, []byte, error) {
	if client == nil {
		return nil, nil, errors.New("client 为空")
	}
	s, err := GetSpec(endpoint)
	if err != nil {
		return nil, nil, err
	}
	return client.Do(ctx, s.Method, s.Path, reqBody, extraHeaders)
}

func CallAndDecodeData[T any](ctx context.Context, client *platformclient.Client, endpoint Endpoint, reqBody any, extraHeaders map[string]string) (http.Header, T, error) {
	headers, body, err := Call(ctx, client, endpoint, reqBody, extraHeaders)
	if err != nil {
		var zero T
		return headers, zero, err
	}
	r, err := platformclient.DecodeAPIResponse[T](body)
	if err != nil {
		var zero T
		return headers, zero, err
	}
	if r.Code != 0 {
		var zero T
		return headers, zero, errors.New(r.Msg)
	}
	return headers, r.Data, nil
}
