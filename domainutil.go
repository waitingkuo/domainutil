package domainutil

import (
	"errors"
	"golang.org/x/net/publicsuffix"
	"net/url"
	"regexp"
	"strings"
)

type Domain struct {
	fullDomain string
	RootDomain string
	SubDomain  string
	Suffix     string
}

func ParseFromRawURL(rawurl string) (*Domain, error) {

	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	return parse(u.Host)
}

func ParseFromHost(host string) (*Domain, error) {
	return parse(host)
}

func ParseFromURL(u *url.URL) (*Domain, error) {
	return parse(u.Host)
}

func parse(host string) (*Domain, error) {

	fullDomain := host

	suffix, _ := publicsuffix.PublicSuffix(host)
	host = strings.TrimSuffix(host, suffix)
	host = strings.TrimSuffix(host, ".")

	// Parse Root domain & Sub Domain
	reg := regexp.MustCompile(`((?:[A-Za-z0-9][A-Za-z0-9-]{1,62}\.)?)([A-Za-z0-9][A-Za-z0-9-]{1,62})$`)
	//reg := regexp.MustCompile(`((?:[A-Za-z0-9][A-Za-z0-9-]{1,62}\.)?)([A-Za-z0-9][A-Za-z0-9-]{1,62}\.[A-Za-z]{2,6})$`)
	result := reg.FindStringSubmatch(host)
	if len(result) == 0 {
		return nil, errors.New("domainutil: incorect domain")
	}
	rootDomain := strings.Join([]string{result[2], suffix}, ".")
	subDomain := result[1]

	if len(subDomain) > 0 {
		subDomain = subDomain[:len(subDomain)-1]
	}

	return &Domain{fullDomain, rootDomain, subDomain, suffix}, nil
}
func (d *Domain) String() string {
	return d.fullDomain
}
