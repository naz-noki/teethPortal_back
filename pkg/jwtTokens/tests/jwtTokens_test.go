package jwtTokens_test

import (
	"MySotre/pkg/jwtTokens"
	"log"
	"testing"
	"time"
)

type user struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func TestMain(m *testing.M) {
	log.Println("Start TestMain", "package: jwtTokens_test")
	m.Run()
	log.Println("End TestMain", "package: jwtTokens_test")
}

func TestAccessTokens(t *testing.T) {
	t.Parallel()

	t.Run("TestAccessTokens - success - 1", func(t *testing.T) {
		secret := "lkjSDFLKSJDHFL*(SYFS*&S*&F)"
		u := user{
			Id:       1,
			Login:    "sadklfj",
			Password: "asl;kdfj",
		}
		nu := new(user)

		token, errCreateAccess := jwtTokens.CreateAccess(secret, &u, time.Second*3)

		if errCreateAccess != nil {
			t.Error(errCreateAccess)
			return
		}

		if errCheckAccess := jwtTokens.CheckAccess(token, secret, nu); errCheckAccess != nil {
			t.Error(errCheckAccess)
			return
		}

		if nu.Id != u.Id || nu.Login != u.Login || nu.Password != u.Password {
			t.Errorf("structures do not match, initial structure: %v resulting structure: %v", u, *nu)
		}
	})
	t.Run("TestAccessTokens - error - 2", func(t *testing.T) {
		secret1 := "asdf*(SYFS*sadfdf&sadfasS*&F)"
		secret2 := "lkjSDFLKSJDH234234234FLSD*(SsadYFS*&S*&F)"
		u := user{
			Id:       1,
			Login:    "sadklfj",
			Password: "asl;kdfj",
		}
		nu := new(user)

		token, errCreateAccess := jwtTokens.CreateAccess(secret1, &u, time.Second*3)

		if errCreateAccess != nil {
			t.Error(errCreateAccess)
			return
		}

		if errCheckAccess := jwtTokens.CheckAccess(token, secret2, nu); errCheckAccess == nil {
			t.Error(errCheckAccess)
			return
		}
	})
	t.Run("TestAccessTokens - error - 3", func(t *testing.T) {
		secret := "asdf*(a*sdfasdrf23423&sadfasS*&F)"
		u := user{
			Id:       1,
			Login:    "sadklfj",
			Password: "asl;kdfj",
		}
		nu := new(user)

		token, errCreateAccess := jwtTokens.CreateAccess(secret, &u, time.Second*2)

		time.Sleep(time.Second * 3)

		if errCreateAccess != nil {
			t.Error(errCreateAccess)
			return
		}

		if errCheckAccess := jwtTokens.CheckAccess(token, secret, nu); errCheckAccess == nil {
			t.Error(errCheckAccess)
			return
		}
	})
}
