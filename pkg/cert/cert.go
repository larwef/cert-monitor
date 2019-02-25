package cert

import (
	"time"
)

const (
	CommfidesProdURL = "ldap.commfides.com"
	CommfidesTestURL = "ldap.test.commfides.com"
	CommfidesBaseDN  = "ou=Signature,ou=Enterprise,dc=commfides,dc=com"
	CommfidesFilter  = "(serialNumber=%s)"
)

// Request is used when requesting certificates registered for an organization.
type Request struct {
	OrgNo  string `json:"orgNo"`
	URL    string `json:"url"`
	BaseDN string `json:"baseDn"`
	Filter string `json:"filter"`
}

// Response holds the certificates returned by the search.
type Response struct {
	Certs []*Cert `json:"certs"`
}

// Cert holds information about a certificate.
type Cert struct {
	Organization   string    `json:"organization"`
	SerialNumber   string    `json:"serialNumber"`
	Issuer         string    `json:"issuer"`
	ValidFrom      time.Time `json:"validFrom"`
	ValidTo        time.Time `json:"validTo"`
	NonRepudiation bool      `json:"nonRepudiation"`
}

// BuypassProdRequest returns a request for Buypass production with the specified orgNo
func BuypassProdRequest(orgNo string) *Request {
	return &Request{
		OrgNo:  orgNo,
		URL:    "ldap.buypass.no",
		BaseDN: "dc=buypass,dc=NO,cn=Buypass Class 3 CA",
		Filter: "(&(sn=%s))",
	}
}

// BuypassTestRequest returns a request for Buypass test with the specified orgNo
func BuypassTestRequest(orgNo string) *Request {
	return &Request{
		OrgNo:  orgNo,
		URL:    "ldap.test4.buypass.no",
		BaseDN: "dc=buypass,dc=NO",
		Filter: "(&(sn=%s))",
	}
}

// CommfidesProdRequest returns a request for Commfides production with the specified orgNo
func CommfidesProdRequest(orgNo string) *Request {
	return &Request{
		OrgNo:  orgNo,
		URL:    "ldap.commfides.com",
		BaseDN: "ou=Signature,ou=Enterprise,dc=commfides,dc=com",
		Filter: "(serialNumber=%s)",
	}
}

// CommfidesTestRequest returns a request for Commfides test with the specified orgNo
func CommfidesTestRequest(orgNo string) *Request {
	return &Request{
		OrgNo:  orgNo,
		URL:    "ldap.test.commfides.com",
		BaseDN: "ou=Signature,ou=Enterprise,dc=commfides,dc=com",
		Filter: "(serialNumber=%s)",
	}
}
