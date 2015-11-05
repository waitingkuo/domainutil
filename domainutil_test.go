package domainutil

import (
	"net/url"
	"testing"
)

func TestParseFromRawURL(t *testing.T) {

	rawUrl := "http://www.google.com"

	domain, err := ParseFromRawURL(rawUrl)
	if err != nil {
		t.Error("Failed to parse rawurl ", rawUrl)
	}

	domainString := domain.String()
	if domainString != "www.google.com" {
		t.Error("Expected www.google.com, got ", domainString)
	}

	if domain.RootDomain != "google.com" {
		t.Error("Expected www.google.com, got ", domain.RootDomain)
	}

	if domain.SubDomain != "www" {
		t.Error("Expected www.google.com, got ", domain.SubDomain)
	}

	if domain.Suffix != "com" {
		t.Error("Expected www.google.com, got ", domain.SubDomain)
	}

}

func TestParseFromURL(t *testing.T) {

	u, _ := url.Parse("http://www.google.com")

	domain, err := ParseFromURL(u)
	if err != nil {
		t.Error("Failed to parse *url.URL ", u)
	}

	domainString := domain.String()
	if domainString != "www.google.com" {
		t.Error("Expected www.google.com, got ", domainString)
	}

	if domain.RootDomain != "google.com" {
		t.Error("Expected www.google.com, got ", domain.RootDomain)
	}

	if domain.SubDomain != "www" {
		t.Error("Expected www.google.com, got ", domain.SubDomain)
	}

}

func TestParseFromHost(t *testing.T) {

	host := "www.google.com"

	domain, err := ParseFromHost(host)
	if err != nil {
		t.Error("Failed to parse host ", host)
	}

	domainString := domain.String()
	if domainString != "www.google.com" {
		t.Error("Expected www.google.com, got ", domainString)
	}

	if domain.RootDomain != "google.com" {
		t.Error("Expected www.google.com, got ", domain.RootDomain)
	}

	if domain.SubDomain != "www" {
		t.Error("Expected www, got ", domain.SubDomain)
	}

}

// parse domain without subdomain
func TestParseShortDomain(t *testing.T) {

	rawUrl := "http://google.com"

	domain, err := ParseFromRawURL(rawUrl)
	if err != nil {
		t.Error("Failed to parse rawurl ", rawUrl)
	}

	domainString := domain.String()
	if domainString != "google.com" {
		t.Error("Expected google.com, got ", domainString)
	}

	if domain.RootDomain != "google.com" {
		t.Error("Expected google.com, got ", domain.RootDomain)
	}

	if domain.SubDomain != "" {
		t.Error("Expected empty string, got ", domain.SubDomain)
	}

}

func TestParseLongDomain(t *testing.T) {

	rawUrl := "http://subsub.www.google.com"

	domain, err := ParseFromRawURL(rawUrl)
	if err != nil {
		t.Error("Failed to parse rawurl ", rawUrl)
	}

	domainString := domain.String()
	if domainString != "subsub.www.google.com" {
		t.Error("Expected subsub.www.google.com, got ", domainString)
	}

	if domain.RootDomain != "google.com" {
		t.Error("Expected google.com, got ", domain.RootDomain)
	}

	if domain.SubDomain != "www" {
		t.Error("Expected www, got ", domain.SubDomain)
	}

}

/*
func TestParseIncorrectDomain(t *testing.T) {

	rawUrl := "http://www.google.c"
	_, err := ParseFromRawURL(rawUrl)
	if err == nil {
		t.Error("Should Fail to parse rawurl ", rawUrl)
	}

	rawUrl = "http://com"
	_, err = ParseFromRawURL(rawUrl)
	if err == nil {
		t.Error("Should Fail to parse rawurl ", rawUrl)
	}

}
*/

func TestParsePublisSuffix(t *testing.T) {

	rawUrl := "http://www.google.com.tw"
	domain, _ := ParseFromRawURL(rawUrl)
	if domain.SubDomain != "www" {
		t.Error("Expected www, got ", domain.SubDomain)
	}
	if domain.RootDomain != "google.com.tw" {
		t.Error("Expected google.com.tw, got ", domain.RootDomain)
	}
	if domain.Suffix != "com.tw" {
		t.Error("Expected com.tw, got ", domain.Suffix)
	}

	rawUrl = "http://www.google.com.au"
	domain, _ = ParseFromRawURL(rawUrl)
	if domain.Suffix != "com.au" {
		t.Error("Expected com.au, got ", domain.Suffix)
	}
}
