package rest

import (
	"bytes"
	"fmt"
	"net/url"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
)

type InfoResponse struct {
	Status
	Info models.Info `json:"info"`
}

type SlashCommandsResponse struct {
	Status
	Commands []models.SlashCommand `json:"commands"`
}

type ExecuteSlashCommandResponse struct {
	Status
}

// GetServerInfo a simple method, requires no authentication,
// that returns information about the server including version information.
//
// https://rocket.chat/docs/developer-guides/rest-api/miscellaneous/info
func (c *Client) GetServerInfo() (*models.Info, error) {
	response := new(InfoResponse)
	if err := c.Get("info", nil, response); err != nil {
		return nil, err
	}

	return &response.Info, nil
}

type DirectoryResponse struct {
	Status
	models.Directory
}

// GetDirectory a method, that searches by users or channels on all users and channels available on server.
// It supports the Offset, Count, and Sort Query Parameters along with Query and Fields Query Parameters.
//
// https://rocket.chat/docs/developer-guides/rest-api/miscellaneous/directory
func (c *Client) GetDirectory(params url.Values) (*models.Directory, error) {
	response := new(DirectoryResponse)
	if err := c.Get("directory", params, response); err != nil {
		return nil, err
	}

	return &response.Directory, nil
}

type SpotlightResponse struct {
	Status
	models.Spotlight
}

// GetSpotlight searches for users or rooms that are visible to the user.
// WARNING: It will only return rooms that user didn’t join yet.
//
// https://rocket.chat/docs/developer-guides/rest-api/miscellaneous/spotlight
func (c *Client) GetSpotlight(params url.Values) (*models.Spotlight, error) {
	response := new(SpotlightResponse)
	if err := c.Get("spotlight", params, response); err != nil {
		return nil, err
	}

	return &response.Spotlight, nil
}

type StatisticsResponse struct {
	Status
	models.StatisticsInfo
}

// GetStatistics
// Statistics about the Rocket.Chat server.
//
// https://rocket.chat/docs/developer-guides/rest-api/miscellaneous/statistics
func (c *Client) GetStatistics() (*models.StatisticsInfo, error) {
	response := new(StatisticsResponse)
	if err := c.Get("statistics", nil, response); err != nil {
		return nil, err
	}

	return &response.StatisticsInfo, nil
}

type StatisticsListResponse struct {
	Status
	models.StatisticsList
}

// GetStatisticsList
// Selectable statistics about the Rocket.Chat server.
// It supports the Offset, Count and Sort Query Parameters along with just the Fields and Query Parameters.
//
// https://rocket.chat/docs/developer-guides/rest-api/miscellaneous/statistics.list
func (c *Client) GetStatisticsList(params url.Values) (*models.StatisticsList, error) {
	response := new(StatisticsListResponse)
	if err := c.Get("statistics.list", params, response); err != nil {
		return nil, err
	}

	return &response.StatisticsList, nil
}

// GetSlashCommandsList
// Slash Commands available in the Rocket.Chat server.
// It supports the offset, count and Sort Query Parameters along with just the Fields and Query Parameters.
//
// https://developer.rocket.chat/reference/api/rest-api/endpoints/core-endpoints/commands-endpoints/list
func (c *Client) GetSlashCommandsList(params url.Values) ([]models.SlashCommand, error) {
	response := new(SlashCommandsResponse)
	if err := c.Get("commands.list", params, response); err != nil {
		return nil, err
	}
	return response.Commands, nil
}

// ExecuteSlashCommand
// Execute a slash command in a room in the Rocket.Chat server
// command and roomId are required params is optional it depends upon command
//
// https://developer.rocket.chat/reference/api/rest-api/endpoints/core-endpoints/commands-endpoints/execute-a-slash-command
func (c *Client) ExecuteSlashCommand(channel *models.ChannelSubscription, command string, params string) (ExecuteSlashCommandResponse, error) {
	response := new(ExecuteSlashCommandResponse)
	var body = fmt.Sprintf(`{ "command": "%s", "roomId": "%s", "params": "%s"}`, command, string(channel.RoomId), params)
	if err := c.Post("commands.run", bytes.NewBufferString(body), response); err != nil {
		return *response, err
	}
	return *response, nil
}
