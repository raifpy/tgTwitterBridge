package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var myTwitterHeaders = map[string]string{
	"User-Agent":                "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:76.0) Gecko/20100101 Firefox/76.0",
	"Content-Type":              "application/x-www-form-urlencoded",
	"x-twitter-auth-type":       "OAuth2Session",
	"x-twitter-client-language": "tr",
	"x-twitter-active-user":     "yes",
	"x-csrf-token":              "9b50e5a20d6710a0b00333c098f885baee5d14e41422b5798865bba82ac4ab0c53f706c95567f572a8d98793d75424fb44700ad4e0595deedaed9600926e4c28708e8a4af38bd7c9c7b33074625fd774",
	"Origin":                    "https://twitter.com",
	"DNT":                       "1",
	"authorization":             "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
}

var myTwitterData = url.Values{
	"include_profile_interstitial_type": {"1"},
	"include_blocking":                  {"1"},
	"include_blocked_by":                {"1"},
	"include_followed_by":               {"1"},
	"include_want_retweets":             {"1"},
	"include_mute_edge":                 {"1"},
	"include_can_dm":                    {"1"},
	"include_can_media_tag":             {"1"},
	"skip_status":                       {"1"},
	"cards_platform":                    {"Web-12"},
	"include_cards":                     {"1"},
	"include_ext_alt_text":              {"true"},
	"include_quote_count":               {"true"},
	"include_reply_count":               {"1"},
	"tweet_mode":                        {"extended"},
	"simple_quoted_tweet":               {"true"},
	"trim_user":                         {"false"},
	"include_ext_media_color":           {"true"},
	"include_ext_media_availability":    {"true"},
	"auto_populate_reply_metadata":      {"false"},
	"batch_mode":                        {"off"},
	"status":                            {"t.me/raifBlog"}, // @raifpy
}

const cookieString = `personalization_id="v1_1U+zrhMyMnU5gbGJnVuKQQ=="; guest_id=v1%3A160640758891419766; dnt=1; ads_prefs="HBISAAA="; kdt=Cu3xmlGdXzTpqGsoIiV3KDjCMbguFydjStBoIM2t; remember_checked_on=1; auth_token=b17aff41f35fdb6b1bf56bd7c6ac447191ff7aea; twid=u%3D1074605921865793536; ct0=9b50e5a20d6710a0b00333c098f885baee5d14e41422b5798865bba82ac4ab0c53f706c95567f572a8d98793d75424fb44700ad4e0595deedaed9600926e4c28708e8a4af38bd7c9c7b33074625fd774; at_check=true; des_opt_in=N; mbox=PC#3cc25cc130f84d858d8271900c0edeee.37_0#1675973649|session#203431e43e2b498ab16a57f7cfc4c1df#1612729387; lang=tr`

func sendPost(text string) {

	myTwitterData["status"][0] = text
	request, err := http.NewRequest("POST", twitterPostURL, strings.NewReader(myTwitterData.Encode()))
	if err != nil {
		panic(err)
	}
	for key, value := range myTwitterHeaders {
		request.Header.Set(key, value)
	}
	/*for key, value := range myTwitterCookies {
		request.AddCookie(&http.Cookie{Name: key, Value: value})
	}*/
	request.Header.Add("Cookie", cookieString)

	fmt.Println(request.Header)

	/*client := http.Client{}
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	if response.StatusCode != 200 {
		log.Println(response.StatusCode)
	}

	byteResponse, _ := ioutil.ReadAll(response.Body)
	//stringResponse := string(byteResponse)

	fmt.Println(string(byteResponse))*/

}

func amain() {
	/*var text string

	fmt.Print("Text : ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text = scanner.Text()

	}*/
	sendPost("oye")
}
