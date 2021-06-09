package http

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"gitlab.com/renodesper/gokit-microservices/endpoint"
	ctxutil "gitlab.com/renodesper/gokit-microservices/util/ctx"
	"gitlab.com/renodesper/gokit-microservices/util/errors"
	errs "gitlab.com/renodesper/gokit-microservices/util/errors"
	resp "gitlab.com/renodesper/gokit-microservices/util/response"
	"golang.org/x/oauth2"
)

func decodeCallbackAuthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	oauthState, _ := r.Cookie("oauthState")

	state := r.FormValue("state")
	code := r.FormValue("code")

	if state != oauthState.Value {
		return nil, errors.InvalidGoogleOauthState
	}

	var req endpoint.CallbackAuthRequest
	req.Code = code

	return req, nil
}

func encodeLoginAuthResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	requestID := ctxutil.GetRequestID(ctx)

	config := response.(*oauth2.Config)
	oauthState := googleGenerateOauthStateCookie(w)
	url := config.AuthCodeURL(oauthState)

	return json.NewEncoder(w).Encode(&resp.SuccessResponse{
		Data: url,
		Meta: resp.PopulateMeta(requestID),
	})
}

func googleGenerateOauthStateCookie(w http.ResponseWriter) string {
	expiration := time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)

	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthState", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func decodeLogoutAuthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.LogoutAuthRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errs.UnparsableJSON
	}
	defer r.Body.Close()

	validate = validator.New()
	if err := validate.Struct(req); err != nil {
		return nil, errs.InvalidRequest
	}

	return req, nil
}
