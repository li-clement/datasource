package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type KG_Dataset struct {
	SubtitleNullable             string    `json:"subtitleNullable"`
	CreatorNameNullable          string    `json:"creatorNameNullable"`
	CreatorURLNullable           string    `json:"creatorUrlNullable"`
	TotalBytesNullable           int       `json:"totalBytesNullable"`
	URLNullable                  string    `json:"urlNullable"`
	LicenseNameNullable          string    `json:"licenseNameNullable"`
	DescriptionNullable          string    `json:"descriptionNullable"`
	OwnerNameNullable            string    `json:"ownerNameNullable"`
	OwnerRefNullable             string    `json:"ownerRefNullable"`
	TitleNullable                string    `json:"titleNullable"`
	CurrentVersionNumberNullable int       `json:"currentVersionNumberNullable"`
	UsabilityRatingNullable      float64   `json:"usabilityRatingNullable"`
	ID                           int       `json:"id"`
	Ref                          string    `json:"ref"`
	Subtitle                     string    `json:"subtitle"`
	HasSubtitle                  bool      `json:"hasSubtitle"`
	CreatorName                  string    `json:"creatorName"`
	HasCreatorName               bool      `json:"hasCreatorName"`
	CreatorURL                   string    `json:"creatorUrl"`
	HasCreatorURL                bool      `json:"hasCreatorUrl"`
	TotalBytes                   int       `json:"totalBytes"`
	HasTotalBytes                bool      `json:"hasTotalBytes"`
	URL                          string    `json:"url"`
	HasURL                       bool      `json:"hasUrl"`
	LastUpdated                  time.Time `json:"lastUpdated"`
	DownloadCount                int       `json:"downloadCount"`
	IsPrivate                    bool      `json:"isPrivate"`
	IsFeatured                   bool      `json:"isFeatured"`
	LicenseName                  string    `json:"licenseName"`
	HasLicenseName               bool      `json:"hasLicenseName"`
	Description                  string    `json:"description"`
	HasDescription               bool      `json:"hasDescription"`
	OwnerName                    string    `json:"ownerName"`
	HasOwnerName                 bool      `json:"hasOwnerName"`
	OwnerRef                     string    `json:"ownerRef"`
	HasOwnerRef                  bool      `json:"hasOwnerRef"`
	KernelCount                  int       `json:"kernelCount"`
	Title                        string    `json:"title"`
	HasTitle                     bool      `json:"hasTitle"`
	TopicCount                   int       `json:"topicCount"`
	ViewCount                    int       `json:"viewCount"`
	VoteCount                    int       `json:"voteCount"`
	CurrentVersionNumber         int       `json:"currentVersionNumber"`
	HasCurrentVersionNumber      bool      `json:"hasCurrentVersionNumber"`
	UsabilityRating              float64   `json:"usabilityRating"`
	HasUsabilityRating           bool      `json:"hasUsabilityRating"`
	Tags                         []struct {
		NameNullable        string `json:"nameNullable"`
		DescriptionNullable string `json:"descriptionNullable"`
		FullPathNullable    string `json:"fullPathNullable"`
		Ref                 string `json:"ref"`
		Name                string `json:"name"`
		HasName             bool   `json:"hasName"`
		Description         string `json:"description"`
		HasDescription      bool   `json:"hasDescription"`
		FullPath            string `json:"fullPath"`
		HasFullPath         bool   `json:"hasFullPath"`
		CompetitionCount    int    `json:"competitionCount"`
		DatasetCount        int    `json:"datasetCount"`
		ScriptCount         int    `json:"scriptCount"`
		TotalCount          int    `json:"totalCount"`
	} `json:"tags"`
}

type HF_Dataset struct {
	ID           string    `json:"_id"`
	Name         string    `json:"id"`
	Sha          string    `json:"sha"`
	LastModified time.Time `json:"lastModified"`
	Private      bool      `json:"private"`
	Gated        bool      `json:"gated"`
	Disabled     bool      `json:"disabled"`
	Citation     string    `json:"citation"`
	Description  string    `json:"description"`
	Downloads    int       `json:"downloads"`
	CardData     struct {
		AnnotationsCreators []string `json:"annotations_creators"`
		LanguageCreators    []string `json:"language_creators"`
		Language            []string `json:"language"`
		License             []string `json:"license"`
		Multilinguality     []string `json:"multilinguality"`
		SizeCategories      []string `json:"size_categories"`
		SourceDatasets      []string `json:"source_datasets"`
		TaskCategories      []string `json:"task_categories"`
		TaskIds             []string `json:"task_ids"`
		PaperswithcodeID    string   `json:"paperswithcode_id"`
		PrettyName          string   `json:"pretty_name"`
		TrainEvalIndex      []struct {
			Config string `json:"config"`
			Task   string `json:"task"`
			TaskID string `json:"task_id"`
			Splits struct {
				EvalSplit string `json:"eval_split"`
			} `json:"splits"`
			ColMapping struct {
				Tokens string `json:"tokens"`
				Labels string `json:"labels"`
			} `json:"col_mapping"`
		} `json:"train-eval-index"`
		Tags        []string `json:"tags"`
		DatasetInfo struct {
			Features []struct {
				Name     string `json:"name"`
				Dtype    string `json:"dtype,omitempty"`
				Sequence string `json:"sequence,omitempty"`
			} `json:"features"`
			Splits []struct {
				Name        string `json:"name"`
				NumBytes    int    `json:"num_bytes"`
				NumExamples int    `json:"num_examples"`
			} `json:"splits"`
			DownloadSize int `json:"download_size"`
			DatasetSize  int `json:"dataset_size"`
		} `json:"dataset_info"`
	} `json:"cardData"`
	Likes            int      `json:"likes"`
	PaperswithcodeID string   `json:"paperswithcode_id"`
	Tags             []string `json:"tags"`
	Siblings         string   `json:"siblings"`
	Key              string   `json:"key"`
}

func huggingface_dump() {
	var Datasets []HF_Dataset
	file, err := os.Create("src/hf_dataset.csv")
	filepath := "src/hf_dataset.csv"
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	w1 := csv.NewWriter(file)

	w1.Write([]string{
		"id",
		"dataset_name",
		"dataset_version",
		"license_id",
		"license_name",
		"licensor",
		"license_from",
		"license_location",
		"license_content",
		"origin",
		"downloaded_outlet",
		"outlet_licensed",
		"hash_code",
		"data_size",
		"format",
		"description",
		"is_personal_data",
		"currentVersionNumber",
		"is_additional_verify",
		"is_offensive_content",
		"collection_process",
		"restrictions",
		"is_comply",
		"restriction_notes",
		"dataset_collect_method",
		"additional_notes",
		"challenges",
		"available"})
	w1.Flush()

	url := "https://huggingface.co/api/datasets?full=full"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("Authorization", "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")
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
		w1.Write([]string{
			dataset.ID,
			dataset.CardData.PrettyName,
			"",
			"",
			strings.Join(dataset.CardData.License, ","),
			dataset.Citation,
			"HuggingFace",
			"",
			"",
			"",
			"HuggingFace",
			"True",
			dataset.Sha,
			strconv.Itoa(dataset.CardData.DatasetInfo.DownloadSize),
			"",
			dataset.Description,
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			fmt.Sprintf("%f", dataset.Disabled)})
		w1.Flush()
		fmt.Println("Now dump ...")
	}
}

func kaggle_dump() {
	var Datasets []KG_Dataset
	filepath := "src/kg_dataset.csv"
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
			w1.Write([]string{
				strconv.Itoa(dataset.Id),
				dataset.Ref,
				dataset.Subtitle,
				dataset.CreatorName,
				dataset.CreatorUrl,
				strconv.Itoa(dataset.TotalBytes),
				dataset.Url,
				dataset.LastUpdated,
				strconv.Itoa(dataset.DownloadCount),
				strconv.FormatBool(dataset.IsPrivate),
				strconv.FormatBool(dataset.IsFeatured),
				dataset.LicenseName,
				dataset.Description,
				dataset.OwnerName,
				dataset.OwnerRef,
				strconv.Itoa(dataset.KernelCount),
				dataset.Title,
				strconv.Itoa(dataset.CurrentVersionNumber),
				fmt.Sprintf("%f", dataset.UsabilityRating)})
			w1.Flush()
			fmt.Println("Now dump " + strconv.Itoa(dataset.Id))
		}

	}
}

func main() {
	//initdump()
	//huggingface_dump()
	kaggle_dump()
}
