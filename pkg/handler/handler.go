package handler

import (
	"context"
	"crypto/x509"
	"fmt"
	"log"
	"strings"

	"github.com/larwef/cert-monitor/pkg/cert"
	"gopkg.in/ldap.v2"
)

// Handler is the lambda handler function
func Handler(ctx context.Context, req cert.Request) (cert.Response, error) {
	log.Println("Invoked")

	log.Printf("Request: %+v", req)

	certs, err := search(&req)
	for _, elem := range certs {
		log.Printf("%+v\n", elem)
	}

	if len(certs) == 0 {
		log.Println("Found no certificates")
	}

	defer log.Println("Finished")
	return cert.Response{
			Certs: certs,
		},
		err
}

func search(req *cert.Request) ([]*cert.Cert, error) {
	var certs []*cert.Cert

	sr, err := searchCerts(req)
	if err != nil {
		return nil, err
	}

	for _, entry := range sr.Entries {
		crt, err := x509.ParseCertificate(entry.GetRawAttributeValue("userCertificate;binary"))
		if err != nil {
			log.Printf("Error parsing certificate: %v", err)
			continue
		}

		certs = append(certs, mapCert(crt))
	}

	return certs, nil
}

func searchCerts(req *cert.Request) (*ldap.SearchResult, error) {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", req.URL, 389))
	if err != nil {
		return nil, err
	}
	defer l.Close()

	searchRequest := ldap.NewSearchRequest(
		req.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(req.Filter, req.OrgNo),
		[]string{"dn", "cn", "userCertificate;binary"},
		nil,
	)

	return l.Search(searchRequest)
}

func mapCert(crt *x509.Certificate) *cert.Cert {
	return &cert.Cert{
		Organization:   strings.Join(crt.Subject.Organization, ", "),
		SerialNumber:   fmt.Sprintf("%x", crt.SerialNumber),
		Issuer:         strings.Join(crt.Issuer.Organization, ", "),
		ValidFrom:      crt.NotBefore,
		ValidTo:        crt.NotAfter,
		NonRepudiation: x509.KeyUsage(crt.KeyUsage) == x509.KeyUsageContentCommitment,
	}
}
