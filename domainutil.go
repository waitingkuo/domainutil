package domainutil

import (
	"errors"
	"net/url"
	"regexp"
)

type Domain struct {
	fullDomain string
	RootDomain string
	SubDomain  string
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

	// Parse Root domain & Sub Domain
	reg := regexp.MustCompile(`((?:[A-Za-z0-9][A-Za-z0-9-]{1,62}\.)?)([A-Za-z0-9][A-Za-z0-9-]{1,62}\.[A-Za-z]{2,6})$`)
	result := reg.FindStringSubmatch(host)
	if len(result) == 0 {
		return nil, errors.New("domainutil: incorect domain")
	}

	fullDomain := host
	rootDomain := result[2]
	subDomain := result[1]

	if len(subDomain) > 0 {
		subDomain = subDomain[:len(subDomain)-1]
	}

	return &Domain{fullDomain, rootDomain, subDomain}, nil
}
func (d *Domain) String() string {
	return d.fullDomain
}
