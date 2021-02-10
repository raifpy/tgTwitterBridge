package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

const twitterPostURL = "https://twitter.com/i/api/1.1/statuses/update.json"

type twitterHeader struct {
	//Lang string `json:"x-twitter-client-language" url:"x-twitter-client-language"`
	//Origin                 string
	//DNT           string
	Authorization string `json:"authorization" url:"authorization"`
	//Xtwitterauthtype       string
	//Xtwitterclientlanguage string
	//Xcsrftoken string
	UserAgent string `json:"User-Agent" url:"UserAgent"`
	//ContentType            string
	//Xtwitteractiveuser     string
}

type twitterCookie string

func (t twitterCookie) getValue(key string) string {
	veri := strings.Split(string(t), ";")
	for _, listValue := range veri {
		listValue = strings.TrimLeft(listValue, " ")
		listValue2 := strings.Split(listValue, "=")
		if listValue2[0] == key {
			if len(listValue2) == 1 {
				return ""
			}
			return listValue2[1]

		}

	}
	return ""

}

//type twitterDefaultData map[string]string
type twitterDefaultData struct {
	IncludeOprofileOinterstitialOtype string `json:"include_profile_interstitial_type" url:"include_profile_interstitial_type"`
	IncludeOblocking                  string `json:"include_blocking" url:"include_blocking"`
	IncludeOblockedOby                string `json:"include_blocked_by" url:"include_blocked_by"`
	IncludeOfollowedOby               string `json:"include_followed_by" url:"include_followed_by"`
	IncludeOwantOretweets             string `json:"include_want_retweets" url:"include_want_retweets"`
	IncludeOmuteOedge                 string `json:"include_mute_edge" url:"include_mute_edge"`
	IncludeOcanOdm                    string `json:"include_can_dm" url:"include_can_dm"`
	IncludeOcanOmediaOtag             string `json:"include_can_media_tag" url:"include_can_media_tag"`
	SkipOstatus                       string `json:"skip_status" url:"skip_status"`
	CardsOplatform                    string `json:"cards_platform" url:"cards_platform"`
	IncludeOcards                     string `json:"include_cards" url:"include_cards"`
	IncludeOextOaltOtext              string `json:"include_ext_alt_text" url:"include_ext_alt_text"`
	IncludeOquoteOcount               string `json:"include_quote_count" url:"include_quote_count"`
	IncludeOreplyOcount               string `json:"include_reply_count" url:"include_reply_count"`
	TweetOmode                        string `json:"tweet_mode" url:"tweet_mode"`
	SimpleOquotedOtweet               string `json:"simple_quoted_tweet" url:"simple_quoted_tweet"`
	TrimOuser                         string `json:"trim_user" url:"trim_user"`
	IncludeOextOmediaOcolor           string `json:"include_ext_media_color" url:"include_ext_media_color"`
	IncludeOextOmediaOavailability    string `json:"include_ext_media_availability" url:"include_ext_media_availability"`
	AutoOpopulateOreplyOmetadata      string `json:"auto_populate_reply_metadata" url:"auto_populate_reply_metadata"`
	BatchOmode                        string `json:"batch_mode" url:"batch_mode"`
	//Status                            string `json:"status" url:"status"`
}

type twitterClient struct {
	twitterHeader      twitterHeader
	twitterCookie      twitterCookie
	twitterDefaultData twitterDefaultData
	apiURL             string
}

func (t twitterClient) getDefaultRequest(url string, data url.Values) *http.Request {
	request, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		tgLog(err.Error(), logID)
		return nil
	}

	request.Header.Add("Origin", "https://twitter.com")
	request.Header.Add("Referer", "https://twitter.com/home")
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	request.Header.Add("User-Agent", t.twitterHeader.UserAgent)
	//request.Header.Add("DNT", t.twitterHeader.DNT)
	request.Header.Add("DNT", "1")

	request.Header.Add("x-csrf-token", t.twitterCookie.getValue("ct0"))
	request.Header.Add("x-twitter-active-user", "yes")
	request.Header.Add("x-twitter-auth-type", "OAuth2Session")
	request.Header.Add("x-twitter-client-language", t.twitterCookie.getValue("lang"))

	request.Header.Add("authorization", t.twitterHeader.Authorization)
	request.Header.Add("Cookie", string(t.twitterCookie))

	return request

}

func (t twitterClient) getDefaultRequestIO(wwwurl string, data io.Reader) *http.Request {
	request, err := http.NewRequest("POST", twitterPostURL, data)
	if err != nil {
		tgLog(err.Error(), logID)
		return nil
	}

	request.Header.Add("Origin", "https://twitter.com")
	request.Header.Add("Referer", "https://twitter.com/home")
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	request.Header.Add("User-Agent", t.twitterHeader.UserAgent)
	//request.Header.Add("DNT", t.twitterHeader.DNT)
	request.Header.Add("DNT", "1")

	request.Header.Add("x-csrf-token", t.twitterCookie.getValue("ct0"))
	request.Header.Add("x-twitter-active-user", "yes")
	request.Header.Add("x-twitter-auth-type", "OAuth2Session")
	request.Header.Add("x-twitter-client-language", t.twitterCookie.getValue("lang"))

	request.Header.Add("authorization", t.twitterHeader.Authorization)
	request.Header.Add("Cookie", string(t.twitterCookie))

	return request
}

func (t twitterClient) parseResponse(petrol *http.Response) {
	// Future
}

func (t twitterClient) sendTweet(text string) {

	values, err := query.Values(t.twitterDefaultData)
	if err != nil {
		tgLog(err.Error(), logID)
		return
	}
	values.Add("status", text)
	client := http.Client{}
	response, err := client.Do(t.getDefaultRequest(twitterPostURL, values))
	if err != nil {
		tgLog(err.Error(), logID)
		return
	}
	fmt.Println(string(res2String(response)))
	//t.parseResponse(response) // future
}

func (t twitterClient) uploadMedia(ext string, data []byte) (string, error) {
	var tweetMedia string
	switch ext {
	case "png", "jpg", "jpeg":
		tweetMedia = "image"
	case "gif":
		tweetMedia = "gif"
	case "mp4":
		tweetMedia = "video"
	}
	//fmt.Println(t.req2String( ), nil)))
	client := http.Client{}
	response, err := client.Do(t.getDefaultRequest("https://upload.twitter.com/i/media/upload.json?command=INIT&total_bytes="+strconv.Itoa(len(data))+"&media_type=image."+ext+"&media_category=tweet_"+tweetMedia, nil))
	if err != nil {
		tgLog(err.Error(), logID)
		return "", err
	}

	mediaData := res2String(response)
	response.Body.Close()

	var media mediaStruct
	err = json.Unmarshal(mediaData, &media)
	if err != nil {
		tgLog(err.Error(), logID)
		return "", err
	}

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	d, err := writer.CreateFormFile("media", "blob")
	if err != nil {
		tgLog(err.Error(), logID)
		return "", err
	}

	io.Copy(d, bytes.NewReader(data))
	/*for key, value := range data {
		err = writer.WriteField(key, value)
		if err != nil {
			return nil, err
		}
	}*/
	writer.Close()

	request, err := http.NewRequest("POST", "https://upload.twitter.com/i/media/upload.json?command=APPEND&media_id="+media.MediaIDString+"&segment_index=0", buf)
	fmt.Println(4)
	buf = nil // free memory
	d = nil

	if err != nil {
		tgLog(err.Error(), logID)
		return "", err
	}

	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("Origin", "https://twitter.com")
	request.Header.Add("Referer", "https://twitter.com/home")
	request.Header.Add("User-Agent", t.twitterHeader.UserAgent)
	request.Header.Add("DNT", "1")
	request.Header.Add("x-csrf-token", t.twitterCookie.getValue("ct0"))
	request.Header.Add("x-twitter-active-user", "yes")
	request.Header.Add("x-twitter-auth-type", "OAuth2Session")
	request.Header.Add("x-twitter-client-language", t.twitterCookie.getValue("lang"))
	request.Header.Add("authorization", t.twitterHeader.Authorization)
	request.Header.Add("Cookie", string(t.twitterCookie))

	/*values, err := query.Values(t.twitterDefaultData)
	if err != nil {
		return "", err
	}
	values.Add("media_id", media.MediaIDString)
	values.Add("status", "Merhaba DUnya resim 228")*/
	fmt.Println(req2String(request))

	//tekrarla:
	response, err = client.Do(t.getDefaultRequest("https://upload.twitter.com/i/media/upload.json?command=FINALIZE&media_id="+media.MediaIDString, nil))
	if err != nil {
		tgLog(err.Error(), logID)
		return "", err
	}
	finazlizeData := res2String(response)
	var finalizeStruct twitterUploadFinalizeResponse
	err = json.Unmarshal(finazlizeData, &finalizeStruct)
	if err != nil {
		tgLog(err.Error(), logID)
		return "", err
	}
	if finalizeStruct.ProcessingInfo.CheckAfterSecs != 0 {
		fmt.Println("Waiting", finalizeStruct.ProcessingInfo.CheckAfterSecs, " secs.")
		time.Sleep(time.Second * time.Duration(finalizeStruct.ProcessingInfo.CheckAfterSecs))
		//goto tekrarla
		/*response, err := client.Do(t.getDefaultRequest("https://upload.twitter.com/i/media/upload.json?command=FINALIZE&media_id="+media.MediaIDString, nil))
		if err != nil {
			return "", err
		}
		fmt.Println(string(res2String(response)))*/
	}

	return media.MediaIDString, nil
}

func (t twitterClient) uploadMedias(medias map[string][]byte) ([]string, error) {
	list := make([]string, len(medias))
	var err error = nil
	var mediaid string
	for key, value := range medias {
		mediaid, err = t.uploadMedia(key, value)
		if err == nil {
			list = append(list, mediaid)
		}
	}
	return list, err
}

func (t twitterClient) sendMedia(text string, mediaIds []string) {

	values, err := query.Values(t.twitterDefaultData)
	if err != nil {
		tgLog(err.Error(), logID)
		return
	}
	values.Add("status", text)
	for _, mediaid := range mediaIds {
		values.Add("media_ids", mediaid)
	}
	fmt.Println(req2String(t.getDefaultRequest(twitterPostURL, values)))

	/*url.Values{
	"status":    {text},
	"media_ids": mediaIds}*/

}

func newTwitter(defaultJSONPath, headerJSONpath, cookiePATH string) (*twitterClient, error) {
	var defaultJSON twitterDefaultData
	defaultFile, err := os.Open(defaultJSONPath)
	if err != nil {
		tgLog(err.Error(), logID)
		return nil, err
	}
	err = json.NewDecoder(defaultFile).Decode(&defaultJSON)
	defaultFile.Close()
	if err != nil {
		tgLog(err.Error(), logID)
		return nil, err
	}

	var headerJSON twitterHeader
	headerFile, err := os.Open(headerJSONpath)
	if err != nil {
		tgLog(err.Error(), logID)
		return nil, err
	}
	err = json.NewDecoder(headerFile).Decode(&headerJSON)
	headerFile.Close()
	if err != nil {
		tgLog(err.Error(), logID)
		return nil, err
	}

	cookie, err := ioutil.ReadFile(cookiePATH)
	if err != nil {
		tgLog(err.Error(), logID)
		return nil, err
	}
	var tCookie = twitterCookie(string(cookie))

	cookie = nil

	if tCookie.getValue("kdt") == "" {
		tgLog(cookiePATH+" > kdt=XXXXX not found.", logID)
	}
	return &twitterClient{
		twitterHeader:      headerJSON,
		twitterCookie:      tCookie,
		twitterDefaultData: defaultJSON,
	}, nil

}

// ÇORBA
