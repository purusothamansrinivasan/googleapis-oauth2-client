package main

import (
	"github.com/google-oauth2-api/mcp-server/config"
	"github.com/google-oauth2-api/mcp-server/models"
	tools_oauth2 "github.com/google-oauth2-api/mcp-server/tools/oauth2"
	tools_userinfo "github.com/google-oauth2-api/mcp-server/tools/userinfo"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_oauth2.CreateOauth2_tokeninfoTool(cfg),
		tools_userinfo.CreateOauth2_userinfo_getTool(cfg),
		tools_userinfo.CreateOauth2_userinfo_v2_me_getTool(cfg),
	}
}
