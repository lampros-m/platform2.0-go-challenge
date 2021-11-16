package test

import (
	"net/http"
	"testing"

	seed "gwi/platform2.0-go-challenge/dbschema"
	"gwi/platform2.0-go-challenge/test"
)

func TestGetSeoAnalytics(t *testing.T) {
	seed.DBClear()
	seed.DBSeed()

	client := test.NewHTTPClient(t)

	// Health
	client.PostAndCompare(
		`{}`,
		"/health",
		map[string]string{},
		test.HTTPExpectedResponse{
			ExpectedCode: http.StatusOK,
			LogResponse:  false,
		},
	)

	// User login
	respLogin, _ := client.PostAndCompare(
		`{
			"username": "oni",
			"password": "1234"
		}`,
		"/auth/login",
		map[string]string{},
		test.HTTPExpectedResponse{
			ExpectedCode:     http.StatusOK,
			ExpectedResponse: test.JSONFile(`login_response.json`),
			LogResponse:      false,
		},
	)
	userToken := respLogin["info"].([]interface{})[0].(map[string]interface{})["token"].(string)

	// User not enriched assets
	client.PostAndCompare(
		`{
			"enriched_view" : false
		}`,
		"/dashboard/userassets",
		map[string]string{"Authorization": "Bearer" + " " + userToken},
		test.HTTPExpectedResponse{
			ExpectedCode:     http.StatusOK,
			ExpectedResponse: test.JSONFile(`user_assets_non_enriched.json`),
			LogResponse:      false,
		},
	)
}
