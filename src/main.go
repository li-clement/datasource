package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type Dataset struct {
	Id                   int     `json:"id"`
	Ref                  string  `json:"ref,omitempty"`
	Subtitle             string  `json:"subtitle,omitempty"`
	CreatorName          string  `json:"creatorName,omitempty"`
	CreatorUrl           string  `json:"creatorUrl,omitempty"`
	TotalBytes           int     `json:"totalBytes,omitempty"`
	Url                  string  `json:"url,omitempty"`
	LastUpdated          string  `json:"lastUpdated,omitempty"`
	DownloadCount        int     `json:"downloadCount,omitempty"`
	IsPrivate            bool    `json:"isPrivate,omitempty"`
	IsFeatured           bool    `json:"isFeatured,omitempty"`
	LicenseName          string  `json:"licenseName,omitempty"`
	Description          string  `json:"description,omitempty"`
	OwnerName            string  `json:"ownerName,omitempty"`
	OwnerRef             string  `json:"ownerRef,omitempty"`
	KernelCount          int     `json:"kernelCount,omitempty"`
	Title                string  `json:"title,omitempty"`
	CurrentVersionNumber int     `json:"currentVersionNumber,omitempty"`
	UsabilityRating      float32 `json:"usabilityRating,omitempty"`
}

func initdump() {
	var Datasets []Dataset
	file, err := os.Create("src/dataset.csv")
	filepath := "src/dataset.csv"
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	w1 := csv.NewWriter(file)

	w1.Write([]string{"id", "ref", "subtitle", "creatorName", "creatorUrl", "totalBytes", "url", "lastUpdated", "downloadCount", "isPrivate", "isFeatured", "licenseName", "description", "ownerName", "ownerRef", "kernelCount", "title", "currentVersionNumber", "usabilityRating"})
	w1.Flush()

	for i := 1; i <= 1; i++ {
		url := "https://www.kaggle.com/api/v1/datasets/list?page=" + strconv.Itoa(i)
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)
		req.SetBasicAuth("clementlee1987", "986d3695525944f46e186c051bfc30b8")

		file, err = os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0666)
		w1 = csv.NewWriter(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		json.Unmarshal(body, &Datasets)
		for _, dataset := range Datasets {
			w1.Write([]string{strconv.Itoa(dataset.Id), dataset.Ref, dataset.Subtitle,
				dataset.CreatorName, dataset.CreatorUrl, strconv.Itoa(dataset.TotalBytes),
				dataset.Url, dataset.LastUpdated, strconv.Itoa(dataset.DownloadCount), strconv.FormatBool(dataset.IsPrivate),
				strconv.FormatBool(dataset.IsFeatured), dataset.LicenseName, dataset.Description,
				dataset.OwnerName, dataset.OwnerRef, strconv.Itoa(dataset.KernelCount), dataset.Title, strconv.Itoa(dataset.CurrentVersionNumber), fmt.Sprintf("%f", dataset.UsabilityRating)})
			w1.Flush()
			fmt.Println("Now dump " + strconv.Itoa(dataset.Id))
		}
	}
}

func redump() {
	var Datasets []Dataset
	filepath := "src/dataset.csv"
	for i := 1; i <= 1; i++ {
		url := "https://www.kaggle.com/api/v1/datasets/list?page=" + strconv.Itoa(i)
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)
		req.SetBasicAuth("clementlee1987", "986d3695525944f46e186c051bfc30b8")

		file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0666)
		w1 := csv.NewWriter(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		json.Unmarshal(body, &Datasets)
		for _, dataset := range Datasets {
			w1.Write([]string{strconv.Itoa(dataset.Id), dataset.Ref, dataset.Subtitle,
				dataset.CreatorName, dataset.CreatorUrl, strconv.Itoa(dataset.TotalBytes),
				dataset.Url, dataset.LastUpdated, strconv.Itoa(dataset.DownloadCount), strconv.FormatBool(dataset.IsPrivate),
				strconv.FormatBool(dataset.IsFeatured), dataset.LicenseName, dataset.Description,
				dataset.OwnerName, dataset.OwnerRef, strconv.Itoa(dataset.KernelCount), dataset.Title, strconv.Itoa(dataset.CurrentVersionNumber), fmt.Sprintf("%f", dataset.UsabilityRating)})
			w1.Flush()
			fmt.Println("Now dump " + strconv.Itoa(dataset.Id))
		}

	}
}

func main() {
	initdump()
	//redump()
}
