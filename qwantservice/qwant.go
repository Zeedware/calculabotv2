package qwantservice

import (
	"errors"
	"github.com/imroc/req"
)

func GetImage(query string) (string, []byte, error) {
	title, url, err := searchImage(query)
	if err != nil {
		return "", nil, err
	}
	image, err := downloadImage(url)
	if err != nil {
		return "", nil, err
	}
	return title, image, err
}

func searchImage(query string) (string, string, error) {

	qwantUrl := "https://api.qwant.com/api/search/images"
	param := req.Param{
		"count":      "1",
		"t":          "images",
		"uiv":        "4",
		"safeSearch": "0",
		"q":          query,
	}

	res, err := req.New().Get(qwantUrl, param)
	if err != nil {
		return "", "", err
	}
	var x qwantSearchResponse

	err = res.ToJSON(&x)
	if err != nil {
		return "", "", err
	}

	if len(x.Data.Result.Items) < 1 {
		return "", "", errors.New("no image found")
	}

	return x.Data.Result.Items[0].Title, x.Data.Result.Items[0].Media, nil
}

func downloadImage(url string) ([]byte, error) {
	res, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	return res.Bytes(), nil
}

type qwantSearchResponse struct {
	Status string `json:"status"`
	Data   struct {
		Result struct {
			Total int `json:"total"`
			Items []struct {
				Title         string `json:"title"`
				Type          string `json:"type"`
				Media         string `json:"media"`
				Desc          string `json:"desc"`
				Thumbnail     string `json:"thumbnail"`
				ThumbWidth    int    `json:"thumb_width"`
				ThumbHeight   int    `json:"thumb_height"`
				Width         string `json:"width"`
				Height        string `json:"height"`
				Size          string `json:"size"`
				URL           string `json:"url"`
				ID            string `json:"_id"`
				BID           string `json:"b_id"`
				MediaFullsize string `json:"media_fullsize"`
				ThumbType     string `json:"thumb_type"`
				Count         int    `json:"count"`
				MediaPreview  string `json:"media_preview"`
			} `json:"items"`
			Domain string `json:"domain"`
		} `json:"result"`
	} `json:"data"`
}
