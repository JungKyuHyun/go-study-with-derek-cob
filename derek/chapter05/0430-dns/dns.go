package dns

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
)

// dns 정보를 담을 구조체
type Lookup struct {
	cname string
	hosts []string
}

// dns 정보 프린트
func (d *Lookup) String() string {
	result := ""
	for _, host := range d.hosts {
		result += fmt.Sprintf("%s IN A %s\n", d.cname, host)
	}
	return result
}

// cname and host 반환
func LookupAddress(address string) (*Lookup, error) {
	cname, err := net.LookupCNAME(address)
	if err != nil {
		return nil, errors.Wrap(err, "error looking up CNAME")
	}
	hosts, err := net.LookupHost(address)
	if err != nil {
		return nil, errors.Wrap(err, "error looking up HOST")
	}

	return &Lookup{cname: cname, hosts: hosts}, nil
}