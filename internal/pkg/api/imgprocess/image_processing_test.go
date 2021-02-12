package imgprocess_test

import (
	"Samurai/config"
	"Samurai/internal/pkg/api"
	"Samurai/internal/pkg/api/imgprocess"
	"Samurai/internal/pkg/api/inhuman"
	"context"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var pictures = []string{
	"https://play-lh.googleusercontent.com/6hiR7BeQTcyhQoOoNeasulh0WM8tKFLF09naV7t4LAGFjMs20PxKU2RONu1cZIVxY7yy",
	"https://play-lh.googleusercontent.com/2nDcr8fp7m-kLiSN31yUGruOHAZGMHO9u5h-V5EdXqsFwe6FEVU_saDX9ji2d6TNvw",
	"https://play-lh.googleusercontent.com/kAxtdEKXBf5bASmakmC_8qTnUue0vqUMkvP_fofgyWGnfaiQRnEQmLS8sTkvtoVDyww",
	"https://play-lh.googleusercontent.com/0aWZJEYQ5EsvbdaUQO4uSAMu8NODT0XwufU90i1N85BEsZy25CH9rXe6QbnDZkCRQLc",
	"https://play-lh.googleusercontent.com/xz_RxPMYZi9_AugnrHuNnmdrRrZDvALPK5CmauMlgZt9K5bBRfHtFjh2tFSiR3deTg",
	"https://play-lh.googleusercontent.com/5BricghADLXngbOmZArsPqCdzgmtDmq5ES1pPNjQ-ns1mvzG5XTdqPfUEzMHXFKgiA",
	"https://is1-ssl.mzstatic.com/image/thumb/Purple124/v4/47/67/e0/4767e03c-a8cd-c5f3-a4d0-899c83740bad/pr_source.png/2048x2732w.png",
	"https://is3-ssl.mzstatic.com/image/thumb/Purple124/v4/1f/fa/10/1ffa10ba-01db-015e-eb30-678770d1a04f/pr_source.png/2048x2732w.png",
	"https://is4-ssl.mzstatic.com/image/thumb/Purple114/v4/04/e9/d4/04e9d4f5-7cf1-b669-5973-28a4eb770a0f/pr_source.png/2048x2732w.png",
	"https://is3-ssl.mzstatic.com/image/thumb/Purple124/v4/91/3e/7b/913e7be3-4616-6de5-957a-b5f11b2b68ba/pr_source.png/2048x2732w.png",
	"https://is3-ssl.mzstatic.com/image/thumb/Purple114/v4/33/f9/39/33f939dc-bba3-7894-5213-236ef967678f/pr_source.png/2048x2732w.png",
}

func TestImageProcessing_Process_ShouldReturnErrorBecauseIncorrectUrl(t *testing.T) {
	imgapi := imgprocess.New("http://localhost:11111")
	ctx := context.Background()

	wrongUrl := "http://source/hi/123.png"

	var tt = []struct {
		name            string
		data            []string
		expectError     bool
		expectBodyLen   int
		expectNoneValue string
	}{
		{
			name:            "send data with nil array",
			data:            nil,
			expectError:     true,
			expectBodyLen:   0,
			expectNoneValue: "",
		},
		{
			name:            "Should Return Error Because Incorrect Url",
			data:            []string{wrongUrl},
			expectError:     true,
			expectBodyLen:   0,
			expectNoneValue: "",
		},
		{
			name:            "should return value in slice with default url",
			data:            append(pictures, wrongUrl),
			expectError:     false,
			expectBodyLen:   len(pictures) + 1,
			expectNoneValue: wrongUrl,
		},
		{
			name:            "should return default success result",
			data:            pictures,
			expectError:     false,
			expectBodyLen:   len(pictures),
			expectNoneValue: "",
		},
		{
			name:            "should return single success result",
			data:            pictures[2:3],
			expectError:     false,
			expectBodyLen:   1,
			expectNoneValue: "",
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			res, err := imgapi.Process(ctx, test.data)
			assert.Equal(t, test.expectError, err != nil)

			if test.expectError {
				return
			}

			if test.expectNoneValue != "" {
				flag := 0
				for _, v := range res {
					if v == test.expectNoneValue {
						flag += 1
					}
				}
				assert.Greater(t, flag, 0)
			}

			assert.Equal(t, test.expectBodyLen, len(res))
		})
	}
}

func TestImageProcessing_ProcessWithGooglePlay_ShouldReturnModelWithNewPictureAndScreenshots(t *testing.T) {
	imgapi := imgprocess.New("http://localhost:11111")
	ctx := context.Background()
	c := config.New("../../../../config/dev.yml")

	var tt = []struct {
		appApi api.ExternalApi
		bundle string
	}{
		{
			appApi: inhuman.NewApiPlay( inhuman.Config{
				Url:        c.Api.Url,
				Key:        c.Api.Key,
				Hl:         "ru",
				Gl:         "ru",
				ItemsCount: 250,
			}),
			bundle: "com.orcacorp.wargame",
		},
		{
			appApi: inhuman.NewApiStore(inhuman.Config{
				Url:        c.Api.Url,
				Key:        c.Api.Key,
				Hl:         "de",
				Gl:         "de",
				ItemsCount: 200,
			}),
			bundle: "1504417378",
		},
	}

	for _, test := range tt {
		res, err := test.appApi.App(test.bundle)
		assert.NoError(t, err)
		assert.NotNil(t, res)

		picture := res.Picture
		var screenshots []string
		for _, v := range res.Screenshots {
			screenshots = append(screenshots, v.Screens...)
		}

		pic, err := imgapi.Process(ctx, []string{picture})
		assert.NoError(t, err)
		assert.NotNil(t, pic)
		res.Picture = pic[0]

		assert.NotEqual(t, picture, res.Picture)

		mpic, err := imgapi.Process(ctx, screenshots)
		assert.NoError(t, err)
		assert.NotNil(t, mpic)

		index := 0
		length := 0
		for _, v := range res.Screenshots {
			for i := range v.Screens {
				v.Screens[i] = mpic[index]
				index++

				assert.False(t, strings.Contains(v.Screens[i], "http"))
			}
			length += len(v.Screens)
		}

		assert.Equal(t, len(screenshots), length)
		assert.NotEqual(t, screenshots, mpic)

		t.Log(res.Screenshots)
	}
}
