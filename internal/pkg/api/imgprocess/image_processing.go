package imgprocess

import (
	"Samurai/internal/pkg/api/inhuman"
	"context"
	"errors"
	"log"
	"strings"
)

type ImageProcessing struct {
	apiurl string
}

func (i *ImageProcessing) Process(ctx context.Context, image []string) ([]string, error) {
	var result []string
	if len(image) == 1 {
		filename, err := i.single(ctx, image[0])
		if err != nil {
			return nil, err
		}
		result = append(result, filename)
	} else {
		filenames, err := i.batch(ctx, image)
		if err != nil {
			return nil, err
		}
		result = append(result, filenames...)
	}

	return result, nil
}

func (i *ImageProcessing) single(ctx context.Context, image string) (string, error) {
	var imgstruct = &struct{
		Filename string `json:"filename,omitempty"`
	}{}
	err := inhuman.Request("load", "get", inhuman.WithQueryParams(map[string]interface{}{
		"link": image,
	}), inhuman.WithResponseType(&imgstruct), inhuman.WithUrl(i.apiurl))
	if err != nil {
		return "", err
	}

	img := imgstruct.Filename
	if img != "" {
		img = RemoveExtension(img)
	} else {
		return "", errors.New("response empty")
	}

	return img, nil
}

func (i *ImageProcessing) batch(ctx context.Context, images []string) ([]string, error) {
	var imgs map[string]string
	err := inhuman.Request("load/batch", "post",
		inhuman.WithData(images), inhuman.WithResponseType(&imgs), inhuman.WithUrl(i.apiurl))
	if err != nil {
		return nil, err
	}

	resultImgs := make([]string, len(imgs))
	index := 0
	for k, v := range imgs {
		if !strings.Contains(v, "none") {
			resultImgs[index] = RemoveExtension(v)
		} else {
			log.Printf("can not porcess image with url %s, use fallback solution", k)
			resultImgs[index] = k
		}
		index++
	}

	return resultImgs, nil
}

func New(url string) *ImageProcessing {
	return &ImageProcessing{
		apiurl: url,
	}
}