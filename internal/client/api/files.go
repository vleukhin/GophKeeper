package api

import "github.com/vleukhin/GophKeeper/internal/models"

const filesEndpoint = "api/files"

func (c *HTTPClient) GetFiles(accessToken string) (files []models.File, err error) {
	if err := c.getEntities(&files, accessToken, filesEndpoint); err != nil {
		return nil, err
	}

	return files, nil
}

func (c *HTTPClient) StoreFile(accessToken string, file models.File) error {
	return c.addEntity(file, accessToken, filesEndpoint)
}

func (c *HTTPClient) DelFile(accessToken, fileID string) error {
	return c.delEntity(accessToken, filesEndpoint, fileID)
}