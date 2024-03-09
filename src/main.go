package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/peterhellberg/link"
)

type LicenseDetail struct {
	Key                  string   `json:"key"`
	ShortName            string   `json:"short_name"`
	Name                 string   `json:"name"`
	Category             string   `json:"category"`
	Owner                string   `json:"owner"`
	HomepageURL          string   `json:"homepage_url"`
	Notes                string   `json:"notes"`
	SpdxLicenseKey       string   `json:"spdx_license_key"`
	OtherSpdxLicenseKeys []string `json:"other_spdx_license_keys"`
	OsiLicenseKey        string   `json:"osi_license_key"`
	TextUrls             []string `json:"text_urls"`
	OsiURL               string   `json:"osi_url"`
	FaqURL               string   `json:"faq_url"`
	OtherUrls            []string `json:"other_urls"`
	IgnorableCopyrights  []string `json:"ignorable_copyrights"`
	IgnorableHolders     []string `json:"ignorable_holders"`
	Text                 string   `json:"text"`
}

type LicenseList []struct {
	LicenseKey     string `json:"license_key"`
	Category       string `json:"category"`
	SpdxLicenseKey string `json:"spdx_license_key"`
	IsException    bool   `json:"is_exception"`
	IsDeprecated   bool   `json:"is_deprecated"`
	JSON           string `json:"json"`
	Yaml           string `json:"yaml"`
	HTML           string `json:"html"`
	License        string `json:"license"`
}

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
		//License             []string `json:"license"`
		License          string   `json:"license"`
		LicenseName      string   `json:"license_name"`
		LicenseLink      string   `json:"license_link"`
		Multilinguality  []string `json:"multilinguality"`
		SizeCategories   []string `json:"size_categories"`
		SourceDatasets   []string `json:"source_datasets"`
		TaskCategories   []string `json:"task_categories"`
		TaskIds          []string `json:"task_ids"`
		PaperswithcodeID string   `json:"paperswithcode_id"`
		PrettyName       string   `json:"pretty_name"`
		TrainEvalIndex   []struct {
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

func license_dump() {
	var licenses LicenseList
	var licenseDetail LicenseDetail
	file, err := os.Create("license.csv")
	filepath := "license.csv"
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	w1 := csv.NewWriter(file)
	w1.Write([]string{
		"id",
		"license_name",
		"license_fullname",
		"license_link",
		"license_category",
		"license_text"})
	w1.Flush()
	url := "https://scancode-licensedb.aboutcode.org/index.json"
	method := "GET"
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)
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

	// for _, l := range link.ParseResponse(res) {
	// 	//fmt.Printf("URI: %q, Rel: %q, Extra: %+v\n", l.URI, l.Rel, l.Extra)
	// 	// URI: "https://api.github.com/search/code?q=Println+user%3Agolang&page=2", Rel: "next", Extra: map[]
	// 	// URI: "https://api.github.com/search/code?q=Println+user%3Agolang&page=34", Rel: "last", Extra: map[]
	// 	if l.Rel == "next" {
	// 		url = l.URI
	// 	}
	// }
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(body, &licenses)
	for i, license := range licenses {
		url := "https://scancode-licensedb.aboutcode.org/" + license.LicenseKey + ".json"
		method := "GET"
		client := &http.Client{}
		req, _ := http.NewRequest(method, url, nil)
		res1, err := client.Do(req)
		fmt.Println("new dump:" + url)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res1.Body.Close()
		body1, err := io.ReadAll(res1.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		json.Unmarshal(body1, &licenseDetail)
		w1.Write([]string{
			strconv.Itoa(i),
			license.LicenseKey,
			license.SpdxLicenseKey,
			licenseDetail.HomepageURL,
			licenseDetail.Category,
			licenseDetail.Text,
		})
		w1.Flush()
	}
	if link.ParseResponse(res) == nil {
		w1.Flush()
		return
	}
}

func huggingface_dump() {
	var Datasets []HF_Dataset
	strType := ""
	sizeCatagory := "0"
	file, err := os.Create("hf_dataset.csv")
	filepath := "hf_dataset.csv"
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	w1 := csv.NewWriter(file)

	w1.Write([]string{
		"id",
		"dataset_fullname",
		"dataset_name",
		"dataset_version",
		"license_id",
		//"license_set",
		"license",
		"license_name",
		"license_link",
		"licensor",
		"license_from",
		"license_location",
		"license_content",
		"origin",
		"downloaded_outlet",
		"outlet_licensed",
		"hash_code",
		"data_size",
		"task_type",
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
		"available",
		"size_catagory",
		"downloads",
		"like",
		"update_at"})
	w1.Flush()
	url := "https://huggingface.co/api/datasets?full=full"
	method := "GET"
	client := &http.Client{}
	for {
		if url == "" {
			return
		}
		fmt.Println("new dump:" + url)
		req, _ := http.NewRequest(method, url, nil)

		req.Header.Add("Authorization", "XXXXXXXXXXXXXXXXXXXXXXXX")
		file, err = os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
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

		for _, l := range link.ParseResponse(res) {
			//fmt.Printf("URI: %q, Rel: %q, Extra: %+v\n", l.URI, l.Rel, l.Extra)
			// URI: "https://api.github.com/search/code?q=Println+user%3Agolang&page=2", Rel: "next", Extra: map[]
			// URI: "https://api.github.com/search/code?q=Println+user%3Agolang&page=34", Rel: "last", Extra: map[]
			if l.Rel == "next" {
				url = l.URI
			}
		}

		//aa := res.Header["Link"]

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		json.Unmarshal(body, &Datasets)
		for _, dataset := range Datasets {
			for _, tag := range dataset.CardData.TaskCategories {
				if strType != "" {
					strType = strType + ","
				}
				strType = strType + tag
			}

			if len(dataset.CardData.SizeCategories) != 0 {
				//fmt.Println(dataset.CardData.SizeCategories)
				sizeCatagory = dataset.CardData.SizeCategories[0]
			}

			w1.Write([]string{
				dataset.ID,
				dataset.CardData.PrettyName,
				dataset.Name,
				"",
				"",
				//strings.Join(dataset.CardData.License, ","),
				//strings.Join(dataset.CardData.LicenseName, ","),
				//strings.Join(dataset.CardData.LicenseLink, ","),
				//strings.Join(dataset.CardData.LicenseSet, ","),
				dataset.CardData.License,
				dataset.CardData.LicenseName,
				dataset.CardData.LicenseLink,
				dataset.Citation,
				"HuggingFace",
				"",
				"",
				"",
				"HuggingFace",
				"True",
				dataset.Sha,
				strconv.Itoa(dataset.CardData.DatasetInfo.DownloadSize),
				strType,
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
				strconv.Itoa(Bool2int(dataset.Disabled)),
				sizeCatagory,
				strconv.Itoa(dataset.Downloads),
				strconv.Itoa(dataset.Likes),
				dataset.LastModified.String(),
			})
			w1.Flush()
			//fmt.Println("Now dump ...")
			strType = ""
		}
		if link.ParseResponse(res) == nil {
			w1.Flush()
			return
		}
		fmt.Println(url)
	}
}

func kaggle_dump() {
	var Datasets []KG_Dataset
	strType := ""
	file, err := os.Create("kg_dataset.csv")
	filepath := "kg_dataset.csv"
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
		"task_type",
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

	for i := 1; i <= 1; i++ {
		url := "https://www.kaggle.com/api/v1/datasets/list?page=" + strconv.Itoa(i)
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)
		req.SetBasicAuth("clementlee1987", "54c805ae4953af785d83978f72bd4fe9")

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

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		json.Unmarshal(body, &Datasets)
		for _, dataset := range Datasets {
			for _, tag := range dataset.Tags {
				if strType != "" {
					strType = strType + ","
				}
				strType = strType + tag.Name
			}
			w1.Write([]string{
				strconv.Itoa(dataset.ID),
				dataset.Title,
				strconv.Itoa(dataset.CurrentVersionNumber),
				"",
				dataset.LicenseName,
				dataset.CreatorName,
				"",
				"",
				dataset.LicenseName,
				dataset.URL,
				"Kaggle",
				"Kaggle Term of use",
				"",
				strconv.Itoa(dataset.TotalBytes),
				strType,
				"",
				dataset.Description,
				strconv.FormatBool(dataset.IsPrivate),
				strconv.Itoa(dataset.CurrentVersionNumber),
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				fmt.Sprintf("%f", dataset.UsabilityRating)})
			w1.Flush()
			fmt.Println("Now dump " + strconv.Itoa(dataset.ID))
			strType = ""
		}

	}
}

func Bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {
	//initdump()
	//huggingface_dump()
	//kaggle_dump()
	license_dump()
}
