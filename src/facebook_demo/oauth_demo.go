package main

import (
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"net/http"
)

func main() {
	conf := &oauth2.Config{
		ClientID:     "738278743355174",
		ClientSecret: "ab3977bfd65f00a114a7aa85bc37aa5b",
		Scopes:       []string{"public_profile"},
		//Endpoint: oauth2.Endpoint{
		//	AuthURL:  "https://provider.com/o/oauth2/auth",
		//	TokenURL: "https://provider.com/o/oauth2/token",
		//},
		Endpoint:facebook.Endpoint,
		RedirectURL:  "http://localhost:8094",
	}
	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v \n", url)

	resp, err :=http.Get(url)
	if(err!=nil){
		fmt.Println("error ",err)
	}
	fmt.Println(resp)

	fmt.Println(resp.Header)
	fmt.Println(resp.Body)


	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	//var code string="EAAKfdeoZBLyYBACSDuQ9ZBx5tnCtzQeDVJjn4mKw0EIY9AOdm4ddAFDBwnJr51abSlDxMXp6Oe8OOC2xRMeAUI600zubQZCBOcPVCpO7dtFgCfUGBmAgu9WszynLzLxaqq3YFitOFRs1uLZCBJpoe59DYDytBNeca24doPZCX4CRZC66GeRt4t22NxFm5Q67sZD"
	//if _, err := fmt.Scan(&code); err != nil {
	//	log.Fatal(err)
	//}
	//ctx := context.Background()
	//tok, err := conf.Exchange(ctx, code)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//client := conf.Client(ctx, tok)
	//client.Get("...")
}