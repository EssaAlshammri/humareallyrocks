package main

import "context"

func (app *application) healthcheckHandler(ctx context.Context, I *struct{}) (*HealthCheckOutput, error) {
	resp := &HealthCheckOutput{}
	resp.Body.Status = "available"
	resp.Body.Environment = app.config.env
	resp.Body.Version = version
	return resp, nil
}
