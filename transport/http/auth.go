package http

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"gitlab.com/renodesper/gokit-microservices/endpoint"
	ctxUtil "gitlab.com/renodesper/gokit-microservices/util/ctx"
	"gitlab.com/renodesper/gokit-microservices/util/errors"
	resp "gitlab.com/renodesper/gokit-microservices/util/response"
	"golang.org/x/oauth2"
)

func encodeGoogleLoginAuthResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	requestID := ctxUtil.GetRequestID(ctx)

	config := response.(*oauth2.Config)
	oauthState := googleGenerateOauthStateCookie(w)
	url := config.AuthCodeURL(oauthState)

	return json.NewEncoder(w).Encode(&resp.SuccessResponse{
		Data: url,
		Meta: resp.PopulateMeta(requestID),
	})
}

func decodeGoogleCallbackAuthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	oauthState, _ := r.Cookie("oauthState")

	state := r.FormValue("state")
	code := r.FormValue("code")

	if oauthState == nil || state != oauthState.Value {
		return nil, errors.InvalidGoogleOauthState
	}

	var req endpoint.CallbackAuthRequest
	req.Code = code

	return req, nil
}

func googleGenerateOauthStateCookie(w http.ResponseWriter) string {
	expiration := time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}

	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthState", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func decodeLoginAuthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.LoginAuthRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.UnparsableJSON
	}
	defer r.Body.Close()

	validate = validator.New()
	if err := validate.Struct(req); err != nil {
		return nil, errors.InvalidRequest
	}

	return req, nil
}

func decodeLogoutAuthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.LogoutAuthRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.UnparsableJSON
	}
	defer r.Body.Close()

	validate = validator.New()
	if err := validate.Struct(req); err != nil {
		return nil, errors.InvalidRequest
	}

	return req, nil
}

func decodeRegisterAuthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.RegisterAuthRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.UnparsableJSON
	}
	defer r.Body.Close()

	validate = validator.New()
	if err := validate.Struct(req); err != nil {
		return nil, errors.InvalidRequest
	}

	return req, nil
}

func decodeRequestResetPasswordAuthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.RequestResetPasswordAuthRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.UnparsableJSON
	}
	defer r.Body.Close()

	validate = validator.New()
	if err := validate.Struct(req); err != nil {
		return nil, errors.InvalidRequest
	}

	return req, nil
}
