package attack

// only the host needs to be replaced for the attack to work
const host = "63648acf-cafe-4be9-9680-97d259e80a60.idocker.vuln.land"

const base_url = "https://" + host + "/12001/blindsql_case1/auth_blindsql1/register"
const originalURL = "https%253A%252F%252F576aeb1e-a75d-4d45-8dda-48f5e4b834a3.idocker.vuln.land%252F12001%252Fblindsql_case1%252Fblindsql1%252Fcontroller%253Faction%253Dprofile"

var headers = map[string]string{
	"Host":                      "cef2b763-6dc7-4d60-ba5e-643da1e9c53f.idocker.vuln.land",
	"Cookie":                    "ACookie=12352",
	"User-Agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/110.0",
	"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
	"Accept-Language":           "en-GB,en;q=0.5",
	"Content-Type":              "application/x-www-form-urlencoded",
	"Content-Length":            "200",
	"Origin":                    "https://cef2b763-6dc7-4d60-ba5e-643da1e9c53f.idocker.vuln.land",
	"Referer":                   "https://cef2b763-6dc7-4d60-ba5e-643da1e9c53f.idocker.vuln.land/12001/blindsql_case1/auth_blindsql1/login?originalURL=https%253A%252F%252Fglocken.vm.vuln.land%252F12001%252Fblindsql_case1%252Fblindsql1%252Fcontroller%253Faction%253Dprofile",
	"Upgrade-Insecure-Requests": "1",
	"Sec-Fetch-Dest":            "frame",
	"Sec-Fetch-Mode":            "navigate",
	"Sec-Fetch-Site":            "same-origin",
	"Sec-Fetch-User":            "?1",
	"Te":                        "trailers",
}
