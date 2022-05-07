package API

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Dict360 struct {
	Data struct {
		Explain struct {
			EnglishExplain []interface{} `json:"english_explain"`
			Word           string        `json:"word"`
			Caiyun         struct {
				Info struct {
					Lbsynonym    []string      `json:"lbsynonym"`
					Antonym      []interface{} `json:"antonym"`
					WordExchange []interface{} `json:"word_exchange"`
				} `json:"info"`
			} `json:"caiyun"`
			RelatedWords    []string      `json:"related_words"`
			WordLevel       []string      `json:"word_level"`
			Exsentence      []interface{} `json:"exsentence"`
			Phonetic        interface{}   `json:"phonetic"`
			WebTranslations []struct {
				Translation string `json:"translation"`
				Example     string `json:"example"`
			} `json:"web_translations"`
			Translation []string `json:"translation"`
			Speech      struct {
				AM string `json:"美"`
				EN string `json:"英"`
			} `json:"speech"`
		} `json:"explain"`
		Fanyi    string `json:"fanyi"`
		SpeakURL struct {
			SpeakURL     string `json:"speak_url"`
			TSpeakURL    string `json:"tSpeak_url"`
			WordSpeakURL string `json:"word_speak_url"`
			WordType     string `json:"word_type"`
		} `json:"speak_url"`
		Vendor string `json:"vendor"`
	} `json:"data"`
	Error int    `json:"error"`
	Msg   string `json:"msg"`
}

func Query360(word string) {
	client := &http.Client{}
	url := "https://fanyi.so.com/index/search?eng=1&validate=&ignore_trans=0&query=" + word + "%0A"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "fanyi.so.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "en-CA,en;q=0.9,zh-CN;q=0.8,zh;q=0.7,ja-JP;q=0.6,ja;q=0.5,en-GB;q=0.4,en-US;q=0.3")
	req.Header.Set("content-length", "0")
	req.Header.Set("cookie", "QiHooGUID=EB3EA779F8D8534229C4270684C308A7.1651960229814; Q_UDID=9056f68f-8e0b-05f7-6535-e810ff2f04c0; __guid=144965027.2444922924669585400.1651960233696.3865; count=1; gtHuid=1; __huid=11mGTN84vpbzeoTjlTaVYLflkZgnPlnf8afplkGKgQV7M%3D")
	req.Header.Set("origin", "https://fanyi.so.com")
	req.Header.Set("pro", "fanyi")
	req.Header.Set("referer", "https://fanyi.so.com/")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", bodyText)
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	var dictResponse Dict360
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(word + ":")
	for _, meaning := range dictResponse.Data.Explain.Translation {
		fmt.Println("\t" + meaning)
	}
}
