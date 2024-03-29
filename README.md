
# ADSTXT
---------

### Project structure
The cmd folder contains the adstxt server.
- `./adsd` default http port is 8080 but you can change it with `-http {port}`

The pkg folder folder contains reusable code, there's also 2 adstxt parsing functions, one that is looking at each character and the other one using pattern.

### Usage
Once the server is started you have 2 http endpoints that you can use, one to save ads.txt file to the db and an other one to get the publisher ads.txt content.

- To save `http://localhost:8080/ads/populate` ex: bloomberg.com
you need to send a json payload looking like this.
```json
{
  "domain": "bloomberg.com"
}
```
- To get content `http://localhost:8080/ads/{domain name}` ex: nytimes
you will receive a json response containing all the sellers for that website.
```json
{
	"sellers": [{
	    "publisherName": "nytimes.com",
	    "domain": "amazon-adsystem.com",
	    "accountID": " 3030",
	    "typeOfAccount": "DIRECT",
	    "certAuthID": ""
	  },
          {
            "publisherName": "nytimes.com",
            "domain": "pubmatic.com",
            "accountID": " 158573",
            "typeOfAccount": "DIRECT",
            "certAuthID": " 5d62403b186f2ace"
          }]
}
```
