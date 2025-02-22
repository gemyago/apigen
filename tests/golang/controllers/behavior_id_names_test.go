package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestBehaviorIdNames(t *testing.T) {
	fake := faker.New()

	type testCase = routeTestCase[*behaviorIDNamesControllerTestActions]

	setupRouter := func(tc testCase) (*behaviorIDNamesControllerTestActions, http.Handler) {
		testActions := tc.setupActions(&behaviorIDNamesControllerTestActions{})
		controller := &behaviorIDNamesController{testActions}
		rootHandler := handlers.
			NewRootHandler(
				&routerAdapter{mux: http.NewServeMux()},
				tc.appendRootHandlerOpts(handlers.WithLogger(newLogger()))...,
			).
			RegisterBehaviorIdNamesRoutes(controller)
		return testActions, rootHandler
	}

	t.Run("behaviorNamesWithId", func(t *testing.T) {
		runRouteTestCase(t, "should handle parsing errors with default handler", setupRouter, func() testCase {
			wantData := &models.BehaviorNamesWithIdData{
				ID:               fmt.Sprintf("id-%s", fake.Lorem().Word()),
				EndsWithID:       fmt.Sprintf("ends-with-id-%s", fake.Lorem().Word()),
				TheIDInTheMiddle: fmt.Sprintf("the-id-in-the-middle-%s", fake.Lorem().Word()),
			}
			wantParams := &handlers.BehaviorIdNamesBehaviorNamesWithIDRequest{
				ID:                    fmt.Sprintf("id-%s", fake.Lorem().Word()),
				EndsWithID:            fmt.Sprintf("ends-with-id-%s", fake.Lorem().Word()),
				TheIDInTheMiddle:      fmt.Sprintf("the-id-in-the-middle-%s", fake.Lorem().Word()),
				QueryEndsWithID:       fmt.Sprintf("query-ends-with-id-%s", fake.Lorem().Word()),
				QueryTheIDInTheMiddle: fmt.Sprintf("query-the-id-in-the-middle-%s", fake.Lorem().Word()),
				Payload:               wantData,
			}
			return testCase{
				method: http.MethodPost,
				setupActions: func(a *behaviorIDNamesControllerTestActions) *behaviorIDNamesControllerTestActions {
					a.behaviorNamesWithID.nextResult = wantData
					return a
				},
				path: fmt.Sprintf(
					"/behavior/id-names/%s/%s/%s",
					wantParams.ID, wantParams.EndsWithID, wantParams.TheIDInTheMiddle,
				),
				query: url.Values{
					"queryEndsWithId":       []string{wantParams.QueryEndsWithID},
					"queryTheIdInTheMiddle": []string{wantParams.QueryTheIDInTheMiddle},
				},
				body: marshalJSONDataAsReader(t, wantData),
				expect: func(t *testing.T, testActions *behaviorIDNamesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 200, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.behaviorNamesWithID.calls, 1)
					assert.Equal(t, wantParams, testActions.behaviorNamesWithID.calls[0].params)

					gotResponse := unmarshalData[models.BehaviorNamesWithIdData](t, recorder.Body)
					assert.Equal(t, wantData, &gotResponse)
				},
			}
		})
	})
}
