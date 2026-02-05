package service

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"yangfan-ui/internal/auth"
	"yangfan-ui/internal/platformapi"
	"yangfan-ui/internal/platformclient"
)

type PlatformService struct {
	client *platformclient.Client
	auth   *auth.Manager
}

type loginData struct {
	User      map[string]any `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}

type userInfoData struct {
	UserInfo map[string]any `json:"userInfo"`
}

func NewPlatformService(client *platformclient.Client, authManager *auth.Manager) *PlatformService {
	return &PlatformService{
		client: client,
		auth:   authManager,
	}
}

func (s *PlatformService) Captcha(ctx context.Context) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointCaptcha, nil, nil)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return map[string]any{}, nil
	}
	return data, nil
}

func (s *PlatformService) Login(ctx context.Context, username string, password string, captcha string, captchaId string, node string) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[loginData](ctx, s.client, platformapi.EndpointLogin, map[string]any{
		"username":  username,
		"password":  password,
		"captcha":   captcha,
		"captchaId": captchaId,
		"node":      node,
		"source":    "ui",
	}, nil)
	if err != nil {
		return nil, err
	}
	if err := s.auth.Set(data.Token, data.ExpiresAt); err != nil {
		return nil, err
	}
	return data.User, nil
}

func (s *PlatformService) GetUserInfo(ctx context.Context) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[userInfoData](ctx, s.client, platformapi.EndpointUserInfo, nil, nil)
	if err != nil {
		return nil, err
	}
	return data.UserInfo, nil
}

func (s *PlatformService) SetUserAuthority(ctx context.Context, authorityId uint, projectId uint) (map[string]any, error) {
	headers, _, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointSetAuth, map[string]any{
		"authorityId": authorityId,
		"projectId":   projectId,
	}, nil)
	if err == nil {
		_ = s.auth.ApplyNewTokenFromHeaders(headers)
	}
	if err != nil {
		return nil, err
	}
	return s.GetUserInfo(ctx)
}

func (s *PlatformService) SetSelfInfo(ctx context.Context, info map[string]any) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointSetSelf, info, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *PlatformService) ChangePassword(ctx context.Context, password string, newPassword string) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointChangePwd, map[string]any{
		"password":    password,
		"newPassword": newPassword,
	}, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *PlatformService) GetUINodeMenuTree(ctx context.Context) ([]map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[[]map[string]any](ctx, s.client, platformapi.EndpointUINodeMenuTree, map[string]any{}, nil)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return []map[string]any{}, nil
	}
	return data, nil
}

func (s *PlatformService) GetAndroidDeviceOptionsList(ctx context.Context, query map[string]any) (map[string]any, error) {
	spec, err := platformapi.GetSpec(platformapi.EndpointAndroidDeviceOptionsList)
	if err != nil {
		return nil, err
	}
	pathWithQuery, err := appendQuery(spec.Path, query)
	if err != nil {
		return nil, err
	}
	_, body, err := s.client.Do(ctx, spec.Method, pathWithQuery, nil, nil)
	if err != nil {
		return nil, err
	}
	r, err := platformclient.DecodeAPIResponse[map[string]any](body)
	if err != nil {
		return nil, err
	}
	if r.Code != 0 {
		return nil, errors.New(r.Msg)
	}
	if r.Data == nil {
		return map[string]any{}, nil
	}
	return r.Data, nil
}

func (s *PlatformService) CreateAndroidDeviceOptions(ctx context.Context, payload map[string]any) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointAndroidDeviceOptionsCreate, payload, nil)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return map[string]any{}, nil
	}
	return data, nil
}

func (s *PlatformService) UpdateAndroidDeviceOptions(ctx context.Context, payload map[string]any) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointAndroidDeviceOptionsUpdate, payload, nil)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return map[string]any{}, nil
	}
	return data, nil
}

func (s *PlatformService) DeleteAndroidDeviceOptions(ctx context.Context, id uint) error {
	spec, err := platformapi.GetSpec(platformapi.EndpointAndroidDeviceOptionsDelete)
	if err != nil {
		return err
	}
	pathWithQuery, err := appendQuery(spec.Path, map[string]any{"ID": id})
	if err != nil {
		return err
	}
	_, body, err := s.client.Do(ctx, spec.Method, pathWithQuery, nil, nil)
	if err != nil {
		return err
	}
	r, err := platformclient.DecodeAPIResponse[map[string]any](body)
	if err != nil {
		return err
	}
	if r.Code != 0 {
		return errors.New(r.Msg)
	}
	return nil
}

func (s *PlatformService) GetRunConfigList(ctx context.Context, query map[string]any) (map[string]any, error) {
	return s.callQueryAndDecodeMap(ctx, platformapi.EndpointRunConfigList, query)
}

func (s *PlatformService) CreateRunConfig(ctx context.Context, payload map[string]any) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointRunConfigCreate, payload, nil)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return map[string]any{}, nil
	}
	return data, nil
}

func (s *PlatformService) UpdateRunConfig(ctx context.Context, payload map[string]any) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointRunConfigUpdate, payload, nil)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return map[string]any{}, nil
	}
	return data, nil
}

func (s *PlatformService) DeleteRunConfig(ctx context.Context, id uint) error {
	return s.callQueryAndDecodeEmpty(ctx, platformapi.EndpointRunConfigDelete, map[string]any{"ID": id})
}

func (s *PlatformService) GetAutoStepList(ctx context.Context, query map[string]any) (map[string]any, error) {
	return s.callQueryAndDecodeMap(ctx, platformapi.EndpointAutoStepList, query)
}

func (s *PlatformService) CreateAutoStep(ctx context.Context, payload map[string]any) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointAutoStepCreate, payload, nil)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return map[string]any{}, nil
	}
	return data, nil
}

func (s *PlatformService) UpdateAutoStep(ctx context.Context, payload map[string]any) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointAutoStepUpdate, payload, nil)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return map[string]any{}, nil
	}
	return data, nil
}

func (s *PlatformService) DeleteAutoStep(ctx context.Context, id uint) error {
	return s.callQueryAndDecodeEmpty(ctx, platformapi.EndpointAutoStepDelete, map[string]any{"ID": id})
}

func (s *PlatformService) GetAutoCaseStepList(ctx context.Context, query map[string]any) (map[string]any, error) {
	return s.callQueryAndDecodeMap(ctx, platformapi.EndpointAutoCaseStepList, query)
}

func (s *PlatformService) CreateAutoCaseStep(ctx context.Context, payload map[string]any) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointAutoCaseStepCreate, payload, nil)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return map[string]any{}, nil
	}
	return data, nil
}

func (s *PlatformService) UpdateAutoCaseStep(ctx context.Context, payload map[string]any) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointAutoCaseStepUpdate, payload, nil)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return map[string]any{}, nil
	}
	return data, nil
}

func (s *PlatformService) DeleteAutoCaseStep(ctx context.Context, id uint) error {
	return s.callQueryAndDecodeEmpty(ctx, platformapi.EndpointAutoCaseStepDelete, map[string]any{"ID": id})
}

func (s *PlatformService) GetAutoCaseList(ctx context.Context, query map[string]any) (map[string]any, error) {
	return s.callQueryAndDecodeMap(ctx, platformapi.EndpointAutoCaseList, query)
}

func (s *PlatformService) CreateAutoCase(ctx context.Context, payload map[string]any) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointAutoCaseCreate, payload, nil)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return map[string]any{}, nil
	}
	return data, nil
}

func (s *PlatformService) UpdateAutoCase(ctx context.Context, payload map[string]any) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointAutoCaseUpdate, payload, nil)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return map[string]any{}, nil
	}
	return data, nil
}

func (s *PlatformService) DeleteAutoCase(ctx context.Context, id uint) error {
	return s.callQueryAndDecodeEmpty(ctx, platformapi.EndpointAutoCaseDelete, map[string]any{"ID": id})
}

func (s *PlatformService) GetTimerTaskList(ctx context.Context, query map[string]any) (map[string]any, error) {
	return s.callQueryAndDecodeMap(ctx, platformapi.EndpointTimerTaskList, query)
}

func (s *PlatformService) CreateTimerTask(ctx context.Context, payload map[string]any) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointTimerTaskCreate, payload, nil)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return map[string]any{}, nil
	}
	return data, nil
}

func (s *PlatformService) UpdateTimerTask(ctx context.Context, payload map[string]any) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointTimerTaskUpdate, payload, nil)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return map[string]any{}, nil
	}
	return data, nil
}

func (s *PlatformService) DeleteTimerTask(ctx context.Context, id uint) error {
	return s.callQueryAndDecodeEmpty(ctx, platformapi.EndpointTimerTaskDelete, map[string]any{"ID": id})
}

func (s *PlatformService) GetAutoReportList(ctx context.Context, query map[string]any) (map[string]any, error) {
	return s.callQueryAndDecodeMap(ctx, platformapi.EndpointAutoReportList, query)
}

func (s *PlatformService) FindAutoReport(ctx context.Context, id uint) (map[string]any, error) {
	return s.callQueryAndDecodeMap(ctx, platformapi.EndpointAutoReportFind, map[string]any{"ID": id})
}

func (s *PlatformService) FindAutoCaseStepApis(ctx context.Context, autoCaseStepId uint) ([]map[string]any, error) {
	return s.callQueryAndDecodeSlice(ctx, platformapi.EndpointAutoCaseStepFindApis, map[string]any{"ID": autoCaseStepId})
}

func (s *PlatformService) AddAutoCaseStepApi(ctx context.Context, autoCaseStepId uint, apiId uint, sort uint) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointAutoCaseStepAddApi, map[string]any{
		"id":     autoCaseStepId,
		"api_id": apiId,
		"sort":   sort,
	}, nil)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return map[string]any{}, nil
	}
	return data, nil
}

func (s *PlatformService) DeleteAutoCaseStepApi(ctx context.Context, autoStepId uint) error {
	return s.callQueryAndDecodeEmpty(ctx, platformapi.EndpointAutoCaseStepDeleteApi, map[string]any{"ID": autoStepId})
}

func (s *PlatformService) SortAutoCaseStepApis(ctx context.Context, data []map[string]any) error {
	_, _, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointAutoCaseStepSortApis, map[string]any{
		"data": data,
	}, nil)
	return err
}

func (s *PlatformService) GetAutoCaseSteps(ctx context.Context, autoCaseId uint) ([]map[string]any, error) {
	return s.callQueryAndDecodeSlice(ctx, platformapi.EndpointAutoCaseGetSteps, map[string]any{"ID": autoCaseId})
}

func (s *PlatformService) AddAutoCaseStep(ctx context.Context, caseId uint, stepId uint) error {
	_, _, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointAutoCaseAddStep, map[string]any{
		"caseId": caseId,
		"stepId": stepId,
	}, nil)
	return err
}

func (s *PlatformService) DeleteAutoCaseStepRef(ctx context.Context, refId uint) error {
	return s.callQueryAndDecodeEmpty(ctx, platformapi.EndpointAutoCaseDelStep, map[string]any{"ID": refId})
}

func (s *PlatformService) SortAutoCaseSteps(ctx context.Context, caseId uint, data []map[string]any) error {
	_, _, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointAutoCaseSortSteps, map[string]any{
		"caseId": caseId,
		"data":   data,
	}, nil)
	return err
}

func (s *PlatformService) SetAutoCaseStepConfig(ctx context.Context, refId uint, isConfig bool, isStepConfig bool) error {
	_, _, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointAutoCaseSetStepConfig, map[string]any{
		"ID":           refId,
		"isConfig":     isConfig,
		"isStepConfig": isStepConfig,
	}, nil)
	return err
}

func (s *PlatformService) GetTimerTaskCases(ctx context.Context, taskId uint) ([]map[string]any, error) {
	return s.callQueryAndDecodeSlice(ctx, platformapi.EndpointTimerTaskGetCases, map[string]any{"ID": taskId})
}

func (s *PlatformService) AddTimerTaskCase(ctx context.Context, taskId uint, caseId uint) error {
	_, _, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointTimerTaskAddCase, map[string]any{
		"taskId": taskId,
		"caseId": caseId,
	}, nil)
	return err
}

func (s *PlatformService) DeleteTimerTaskCaseRef(ctx context.Context, refId uint) error {
	return s.callQueryAndDecodeEmpty(ctx, platformapi.EndpointTimerTaskDelCase, map[string]any{"ID": refId})
}

func (s *PlatformService) SortTimerTaskCases(ctx context.Context, taskId uint, data []map[string]any) error {
	_, _, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointTimerTaskSortCases, map[string]any{
		"taskId": taskId,
		"data":   data,
	}, nil)
	return err
}

func (s *PlatformService) callQueryAndDecodeMap(ctx context.Context, endpoint platformapi.Endpoint, query map[string]any) (map[string]any, error) {
	spec, err := platformapi.GetSpec(endpoint)
	if err != nil {
		return nil, err
	}
	pathWithQuery, err := appendQuery(spec.Path, query)
	if err != nil {
		return nil, err
	}
	_, body, err := s.client.Do(ctx, spec.Method, pathWithQuery, nil, nil)
	if err != nil {
		return nil, err
	}
	r, err := platformclient.DecodeAPIResponse[map[string]any](body)
	if err != nil {
		return nil, err
	}
	if r.Code != 0 {
		return nil, errors.New(r.Msg)
	}
	if r.Data == nil {
		return map[string]any{}, nil
	}
	return r.Data, nil
}

func (s *PlatformService) callQueryAndDecodeEmpty(ctx context.Context, endpoint platformapi.Endpoint, query map[string]any) error {
	spec, err := platformapi.GetSpec(endpoint)
	if err != nil {
		return err
	}
	pathWithQuery, err := appendQuery(spec.Path, query)
	if err != nil {
		return err
	}
	_, body, err := s.client.Do(ctx, spec.Method, pathWithQuery, nil, nil)
	if err != nil {
		return err
	}
	r, err := platformclient.DecodeAPIResponse[map[string]any](body)
	if err != nil {
		return err
	}
	if r.Code != 0 {
		return errors.New(r.Msg)
	}
	return nil
}

func (s *PlatformService) callQueryAndDecodeSlice(ctx context.Context, endpoint platformapi.Endpoint, query map[string]any) ([]map[string]any, error) {
	spec, err := platformapi.GetSpec(endpoint)
	if err != nil {
		return nil, err
	}
	pathWithQuery, err := appendQuery(spec.Path, query)
	if err != nil {
		return nil, err
	}
	_, body, err := s.client.Do(ctx, spec.Method, pathWithQuery, nil, nil)
	if err != nil {
		return nil, err
	}
	r, err := platformclient.DecodeAPIResponse[[]map[string]any](body)
	if err != nil {
		return nil, err
	}
	if r.Code != 0 {
		return nil, errors.New(r.Msg)
	}
	if r.Data == nil {
		return []map[string]any{}, nil
	}
	return r.Data, nil
}

func appendQuery(path string, query map[string]any) (string, error) {
	if len(query) == 0 {
		return path, nil
	}
	values := url.Values{}
	for k, v := range query {
		if k == "" || v == nil {
			continue
		}
		switch vv := v.(type) {
		case string:
			if vv != "" {
				values.Add(k, vv)
			}
		case []string:
			for _, item := range vv {
				if item != "" {
					values.Add(k, item)
				}
			}
		case int:
			values.Add(k, strconv.Itoa(vv))
		case int64:
			values.Add(k, strconv.FormatInt(vv, 10))
		case uint:
			values.Add(k, strconv.FormatUint(uint64(vv), 10))
		case uint64:
			values.Add(k, strconv.FormatUint(vv, 10))
		case float64:
			values.Add(k, fmt.Sprintf("%v", vv))
		case bool:
			if vv {
				values.Add(k, "true")
			} else {
				values.Add(k, "false")
			}
		default:
			return "", fmt.Errorf("不支持的 query 类型: %s", k)
		}
	}
	encoded := values.Encode()
	if encoded == "" {
		return path, nil
	}
	return path + "?" + encoded, nil
}
