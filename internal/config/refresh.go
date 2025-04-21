package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// tenantTokenResponse 用于解析飞书租户 token 接口的响应
type tenantTokenResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TenantAccessToken string `json:"tenant_access_token"`
		Expire            int    `json:"expire"`
	} `json:"data"`
}

// FetchTenantAccessToken 调用飞书开放平台获取新的 tenant_access_token
func FetchTenantAccessToken() (string, error) {
	reqBody := map[string]string{
		"app_id":     C.AppID,
		"app_secret": C.AppSecret,
	}
	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("marshal request body failed: %w", err)
	}
	resp, err := http.Post(
		"https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal/",
		"application/json",
		bytes.NewReader(bodyBytes),
	)
	if err != nil {
		return "", fmt.Errorf("request tenant_access_token failed: %w", err)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read response body failed: %w", err)
	}
	var result tenantTokenResponse
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return "", fmt.Errorf("unmarshal response failed: %w", err)
	}
	if result.Code != 0 {
		return "", fmt.Errorf("fetch tenant_access_token error: code=%d msg=%s", result.Code, result.Msg)
	}
	return result.Data.TenantAccessToken, nil
}

// FetchUserAccessToken 暂时从配置中读取 user_access_token，后续可实现真正的刷新逻辑
func FetchUserAccessToken() (string, error) {
	if C.UserAccessToken == "" {
		return "", fmt.Errorf("user_access_token 未配置")
	}
	return C.UserAccessToken, nil
}

// RefreshKey 根据 key 分发到对应的刷新函数
func RefreshKey(key string) (string, error) {
	switch key {
	case "tenant_access_token":
		return FetchTenantAccessToken()
	case "user_access_token":
		return FetchUserAccessToken()
	default:
		return "", fmt.Errorf("不可刷新配置项：%s", key)
	}
}
