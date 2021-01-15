package imgprocess

import (
	"Samurai/internal/pkg/api/request"
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
	err := request.Request("load", "get", request.WithQueryParams(map[string]interface{}{
		"link": image,
	}), request.WithResponseType(&imgstruct), request.WithUrl(i.apiurl))
	if err != nil {
		return "", err
	}

	img := imgstruct.Filename
	if img != "" {
		img = RemoveExtension(img)
	} else {
		return "", errors.New("response empty")
	}

	return i.apiurl + img, nil
}

func (i *ImageProcessing) batch(ctx context.Context, images []string) ([]string, error) {
	var imgs map[string]string
	err := request.Request("load/batch", "post",
		request.WithData(images), request.WithResponseType(&imgs), request.WithUrl(i.apiurl))
	if err != nil {
		return nil, err
	}

	resultImgs := make([]string, len(imgs))
	for index, v := range images {
		value := imgs[v]
		if !strings.Contains(value, "none") {
			resultImgs[index] = i.apiurl + RemoveExtension(value)
		} else {
			log.Printf("can not porcess image with url %s, use fallback solution", v)
			resultImgs[index] = v
		}
	}

	return resultImgs, nil
}

func New(url string) *ImageProcessing {
	return &ImageProcessing{
		apiurl: url,
	}
}