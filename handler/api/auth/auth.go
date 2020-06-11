/**
 * Created by zc on 2020/6/9.
 */
package auth

import (
	"context"
	"net/http"
	"stone/global"
	"stone/pkg/api"
	"stone/pkg/api/store"
	"stone/pkg/ctr"
	"stone/service"
)

func Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params api.RegisterParams
		if err := ctr.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		user := &store.User{
			Username: params.Username,
			Email:    params.Email,
			Pwd:      params.Password,
		}
		if err := service.New().User().Create(context.Background(), user); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params api.LoginParams
		if err := ctr.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		user, err := service.New().User().FindByNameAndPwd(context.Background(), params.Username, params.Password)
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		userInfo := ctr.JwtUserInfo{
			UID:      user.UID,
			Username: user.Username,
			Email:    user.Email,
		}
		token, err := ctr.JwtCreate(ctr.JwtClaims{User: userInfo}, global.Cfg().Serve.Secret)
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.OK(w, api.LoginResult{
			Username: user.Username,
			Email:    user.Email,
			Token:    token,
		})
	}
}

func Info() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := ctr.ContextUserFrom(r.Context())
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.OK(w, user)
	}
}
