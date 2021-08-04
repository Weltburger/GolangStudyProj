package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"forRoma/pkg/server"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func formBody(data map[string]interface{}, buf *bytes.Buffer) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	buf.Reset()
	if _, err := buf.Write(b); err != nil {
		return err
	}
	return nil
}

func ServerStart() *gin.Engine {
	server, err := server.NewTest("./config/config_test.json")
	if err != nil {
		log.Fatal(err)
	}
	return server
}

/*func TestLogin(t *testing.T) {

	server := ServerStart()

	testCase := []struct {
		Email string
		Password string
		Result int
	}{
		{
			Email: "abirvalg",
			Password: "",
			Result: 400,
		},
		{
			Email: "pak@gmail.com",
			Password: "12345",
			Result: 200,
		},
		{
			Email: "suppa@ro.ru",
			Password: "715",
			Result: 200,
		},
	}

	gin.SetMode(gin.ReleaseMode)
	ts := httptest.NewServer(server)

	client := http.Client{}

	buf := &bytes.Buffer{}

	for _, tc := range testCase {
		t.Run(tc.Email, func(t *testing.T) {

			err := formBody(map[string]interface{}{
				"email": tc.Email,
				"password": tc.Password,
			}, buf)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := client.Post(fmt.Sprintf("%s/login", ts.URL), "application/json", buf)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			assert.Equal(t, tc.Result, resp.StatusCode)
		})
	}
}*/

/*func TestRegister(t *testing.T) {

	server := ServerStart()

	testCase := []struct {
		Name string
		Email string
		Password string
		Result int
	}{
		{
			Name: "Kapitan",
			Email: "lotto@ro.ru",
			Password: "12345",
			Result: 200,
		},
		{
			Name: "GaGGGabella",
			Email: "tarkov55@shit.lol",
			Password: "5648",
			Result: 200,
		},
		{
			Name: "",
			Email: "pak@gmail.com",
			Password: "12345",
			Result: 400,
		},
		{
			Name: "Kapitan",
			Email: "",
			Password: "715",
			Result: 400,
		},
		{
			Name: "Kapitan",
			Email: "aeae1@yandex.ru",
			Password: "",
			Result: 400,
		},
		{
			Name: "Kapitan12",
			Email: "aeae2@yandex.ru",
			Password: " ",
			Result: 200,
		},
		{
			Name: " ",
			Email: "oi@ro.ru",
			Password: "5648",
			Result: 200,
		},
	}

	ts := httptest.NewServer(server)

	client := http.Client{}

	buf := &bytes.Buffer{}

	for _, tc := range testCase {
		t.Run(tc.Email, func(t *testing.T) {

			err := formBody(map[string]interface{}{
				"name": tc.Name,
				"email": tc.Email,
				"password": tc.Password,
			}, buf)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := client.Post(fmt.Sprintf("%s/registration", ts.URL), "application/json", buf)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tc.Result {
				buf.Reset()
				buf.ReadFrom(resp.Body)
				t.Error(buf.String())
			}
		})
	}
}*/

/*func TestArticles(t *testing.T) {

	server := ServerStart()

	testCase := []struct {
		Take int64
		Skip int64
		Result int
	}{
		{
			Take: 10,
			Skip: 0,
			Result: 200,
		},
	}

	ts := httptest.NewServer(server)

	client := http.Client{}

	buf := &bytes.Buffer{}

	for i, tc := range testCase {
		t.Run(strconv.Itoa(i), func(t *testing.T) {

			err := formBody(map[string]interface{}{
				"take": tc.Take,
				"skip": tc.Skip,
			}, buf)
			if err != nil {
				t.Fatal(err)
			}

			request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/articles?take=%d&skip=%d",
				ts.URL,
				tc.Take,
				tc.Skip), nil)
			if err != nil {
				t.Fatal(err)
			}
			request.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjIwNDA3MDMsIm5iZiI6MTYyMTc4MTUwMywidG9rZW4iOiJkMjg1NWQyZWVlMjYwNDU1YTM2Y2U4M2Q1YTEyZGIwNTEzMTdlZTAxY2IyZjE3ZTI3NjQ1YWI1NzIxOGEzMDVmIn0.W-UCGGRTMHp_L1CLlPK75ZbQXAt-vqND6eoPPNuOaPLd6dZEiCmxALrqnzyS3_24jh7-MdUYTLewPqnp5emgZg")
			resp, err := client.Do(request)
			if err != nil {
				t.Fatal(err)
			}

			defer resp.Body.Close()

			if resp.StatusCode != tc.Result {
				buf.Reset()
				buf.ReadFrom(resp.Body)
				t.Error(buf.String())
			}
		})
	}
}*/

/*func TestCreateArticle(t *testing.T) {

	server := ServerStart()

	newUser := models.User{UUID: uuid.NewV1()}

	testCase := []struct {
		User *models.User
		Title string
		Text string
		Result int
	}{
		{
			User: &newUser,
			Title: "abirvalg",
			Text: "",
			Result: 400,
		},
		{
			User: &newUser,
			Title: "sdfsdfsdfsd",
			Text: "12345",
			Result: 400,
		},
		{
			User: &newUser,
			Title: "",
			Text: "rtyrtrty",
			Result: 400,
		},
	}

	ts := httptest.NewServer(server)

	client := http.Client{}

	buf := &bytes.Buffer{}

	for i, tc := range testCase {
		t.Run(strconv.Itoa(i), func(t *testing.T) {

			err := formBody(map[string]interface{}{
				"user_uuid": newUser.UUID,
				"title": tc.Title,
				"text": tc.Text,
			}, buf)
			if err != nil {
				t.Fatal(err)
			}

			request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/create/article", ts.URL), nil)
			if err != nil {
				t.Fatal(err)
			}
			request.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjIwNDA3MDMsIm5iZiI6MTYyMTc4MTUwMywidG9rZW4iOiJkMjg1NWQyZWVlMjYwNDU1YTM2Y2U4M2Q1YTEyZGIwNTEzMTdlZTAxY2IyZjE3ZTI3NjQ1YWI1NzIxOGEzMDVmIn0.W-UCGGRTMHp_L1CLlPK75ZbQXAt-vqND6eoPPNuOaPLd6dZEiCmxALrqnzyS3_24jh7-MdUYTLewPqnp5emgZg")
			resp, err := client.Do(request)
			if err != nil {
				t.Fatal(err)
			}

			defer resp.Body.Close()

			if resp.StatusCode != tc.Result {
				buf.Reset()
				buf.ReadFrom(resp.Body)
				t.Error(buf.String())
			}
		})
	}
}*/

func TestLikeUnlikeArticle(t *testing.T) {

	server := ServerStart()

	testCase := []struct {
		UUID string
		Result int
	}{
		{
			UUID: "25046252-9d47-11eb-a1f2-977cc2cd5fa1",
			Result: 200,
		},
		{
			UUID: "25046252-9d47-11eb-a1f2",
			Result: 400,
		},
	}

	ts := httptest.NewServer(server)

	client := http.Client{}

	buf := &bytes.Buffer{}

	for i, tc := range testCase {
		t.Run(strconv.Itoa(i), func(t *testing.T) {

			err := formBody(map[string]interface{}{
				"uuid": tc.UUID,
			}, buf)
			if err != nil {
				t.Fatal(err)
			}

			request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/article/25046252-9d47-11eb-a1f2-977cc2cd5fa1/like_unlike", ts.URL), nil)
			if err != nil {
				t.Fatal(err)
			}
			request.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjA3MzczNzIsIm5iZiI6MTYyMDQ3ODE3MiwidG9rZW4iOiJkZjdiMzM5NzcxOWY1ZjI1YzUyZDc2NzFjZjVmN2Q4MjMxMTBiZmU1ZTg2Y2Q2Y2FhOWY1ZTk5NTA4M2UyYWI3In0.YTBn-uO5MeGRRZ5J4vqY7P8tyvaXkFUSMs5yOHgtNEbzf62PFg-fIVyEnr7nHBy77wsIZttP9uNnfJ-HgYMG_A")
			resp, err := client.Do(request)
			if err != nil {
				t.Fatal(err)
			}

			defer resp.Body.Close()

			if resp.StatusCode != tc.Result {
				buf.Reset()
				buf.ReadFrom(resp.Body)
				t.Error(buf.String())
			}

			defer resp.Body.Close()

			if resp.StatusCode != tc.Result {
				buf.Reset()
				buf.ReadFrom(resp.Body)
				t.Error(buf.String())
			}
		})
	}
}
