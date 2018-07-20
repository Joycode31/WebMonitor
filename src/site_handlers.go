package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type SiteMonitor struct {
	SiteName     string `json:"siteName"`
	SiteStatus	 string `json:"SiteStatus"`
}

func CheckHeartbeat() {
		for k, _ := range sitesMap {
			SiteMonitor := SiteMonitor{}
			SiteMonitor.SiteName = k
			timeout := time.Duration(800 * time.Millisecond)
			client := http.Client{
				Timeout: timeout,
			}
			resp, err := client.Get(k)
				if err != nil {
					print(err.Error())
					SiteMonitor.SiteStatus = "DOWN"
				} else {
					print(string(resp.StatusCode) + resp.Status)
					SiteMonitor.SiteStatus = "UP"
				}
				sitesMap[k] = SiteMonitor
		}
}
/* [Comment] Joy Code : Thu 19 Jul 2018 05:37:03 PM
	This routing periodically calls CheckHeartbeat() after 300 seconds to know 
	wether Url is active or not
*/

func CheckUrlStatus() {
	for {
		<-time.After(300 * time.Second)
		go CheckHeartbeat()
	}

/* [Comment] Joy Code : Thu 19 Jul 2018 08:38:09 PM
	The function gets all the entries from the map and converts it to json 
	format and send it to UI.

*/
}

func getSiteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In getSiteHandler")

	values := []SiteMonitor{}
	for _, value := range sitesMap {
		values = append(values, value)
	}
	siteListBytes, err := json.Marshal(values)
	fmt.Printf("siteListBytes : %s",siteListBytes);

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(siteListBytes)
}

/* [Comment] Joy Code : Thu 19 Jul 2018 05:43:00 PM
TODO : It is not redirecting to getSiteHandler to publish entries after deletion
It should perform following tasks
1) fetch the name of the site to delete and remove it from sitesMap
2) Once the entry is deleted.send call to get to publish the new list after
deletion of site.

*/

func DeleteUrlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In DeleteUrlHandler")
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the SiteMonitor from the form info
	_, err = url.ParseRequestURI(r.Form.Get("siteName"))
	if err != nil {
		fmt.Fprintf(w, "Invalid Uri Entered !!!")
		return
	}
	fmt.Println(r.Form.Get("siteName"))
	delete(sitesMap,r.Form.Get("siteName"));
	http.Redirect(w, r, "/views/", http.StatusFound)   //TODO : It's not redirecting here
}

/* [Comment] Joy Code : Thu 19 Jul 2018 08:40:15 PM
	The callback reads the values from http request and stores the entires in map
*/

func createSiteHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of SiteMonitor
	fmt.Println("In createSiteHandler")
	SiteMonitor := SiteMonitor{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Validate the Url 
	_, err = url.ParseRequestURI(r.Form.Get("siteName"))
	if err != nil {
		fmt.Fprintf(w, "Invalid Uri Entered !!!")
		return
	}
	SiteMonitor.SiteName = r.Form.Get("siteName")
	timeout := time.Duration(800 * time.Millisecond)
	client := http.Client{
		Timeout: timeout,
	}
	_, err = client.Get(SiteMonitor.SiteName)
	if err != nil {
		print(err.Error())
		SiteMonitor.SiteStatus = "DOWN"
	} else {
		SiteMonitor.SiteStatus = "UP"
	}
	sitesMap[SiteMonitor.SiteName] =  SiteMonitor

	http.Redirect(w, r, "/views/", http.StatusFound)
}
