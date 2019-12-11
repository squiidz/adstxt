package pkg

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// ProcessDomain fetch the domain ads.txt file and process the data to build
// a Publisher structure
func ProcessDomain(name string) (*Publisher, error) {
	var domain = name
	if !strings.Contains(name, "/ads.txt") {
		domain = fmt.Sprintf("%s/ads.txt", name)
	}
	var resp = &http.Response{}
	resp, err := http.Get(fmt.Sprintf("http://%s", domain))
	if err != nil || resp.StatusCode == http.StatusNotFound {
		resp, err = http.Get(fmt.Sprintf("https://%s", domain))
		if err != nil || resp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("ads.txt not found")
		}
	}
	defer resp.Body.Close()
	return parseAdsTxt(name, resp.Body)
}

// parseAdsTxt is parsing the adsTxt file without looking at each character
func parseAdsTxt(name string, r io.Reader) (*Publisher, error) {
	p := &Publisher{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		ss := strings.Split(scanner.Text(), ",")
		if len(ss) < 2 || strings.Contains(ss[0], "#") {
			continue
		}
		s := &Seller{
			PublisherName: name,
			Domain:        strings.TrimSpace(ss[0]),
			AccountID:     strings.TrimSpace(ss[1]),
		}

		toas := strings.Split(ss[2], "#")
		s.TypeOfAccount = strings.TrimSpace(toas[0])

		if len(ss) > 3 && len(toas) == 1 {
			s.CertAuthID = strings.TrimSpace(strings.Split(ss[3], "#")[0])
		}
		p.Sellers = append(p.Sellers, s)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return p, nil
}

// parseAdsTxt2 is looking at each character, seems a bit faster
func parseAdsTxt2(name string, r io.Reader) (*Publisher, error) {
	p := &Publisher{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if seller, err := parseLine(scanner.Text(), name); err == nil {
			p.Sellers = append(p.Sellers, seller)
		}
	}
	return p, nil
}

func parseLine(line string, name string) (*Seller, error) {
	var value string
	for i := 0; i < len(line); i++ {
		if line[i] == '#' {
			break
		}
		if line[i] == ' ' {
			continue
		}
		value += string(line[i])
	}
	values := strings.Split(value, ",")
	if len(values) > 2 {
		s := &Seller{
			PublisherName: name,
			Domain:        values[0],
			AccountID:     values[1],
			TypeOfAccount: values[2],
		}
		if len(values) > 3 {
			s.CertAuthID = values[3]
		}
		return s, nil
	}
	return nil, fmt.Errorf("invalid line")
}
