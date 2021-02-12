package imgprocess

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)



func TestImageProcessing_Single_ShouldGetSingleFilenameFromExternalServer(t *testing.T) {
	imageApi := New("http://localhost:11111")
	ctx := context.Background()

	picture := "https://play-lh.googleusercontent.com/6hiR7BeQTcyhQoOoNeasulh0WM8tKFLF09naV7t4LAGFjMs20PxKU2RONu1cZIVxY7yy"

	res, err := imageApi.single(ctx, picture)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	t.Log(res)
}

func TestImageProcessing_Batch_ShouldReturnMultiplyResultsFromExternalServer(t *testing.T) {
	imageApi := New("http://localhost:11111")
	ctx := context.Background()

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

	res, err := imageApi.batch(ctx, pictures)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, len(pictures), len(res))

	t.Log(res)
}
