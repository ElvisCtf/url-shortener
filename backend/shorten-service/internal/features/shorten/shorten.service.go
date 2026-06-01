package shorten

import (
	"errors"
	"net"
	"net/url"

	"example.com/shared"
)

var ErrSSRF = errors.New("URL targets a private or reserved address")

// cgnat is the one range not covered by Go's stdlib IP methods (RFC 6598).
var _, cgnat, _ = net.ParseCIDR("100.64.0.0/10")

type ShortenService struct {
	repo   *ShortenRepository
	config *shared.Config
}

func NewShortenService(repo *ShortenRepository, config *shared.Config) *ShortenService {
	return &ShortenService{repo: repo, config: config}
}

func (s *ShortenService) Create(originalURL string) (*ShortenResponse, error) {
	if !isSafeURL(originalURL) {
		return nil, ErrSSRF
	}
	code, err := s.repo.Save(originalURL)
	if err != nil {
		return nil, err
	}
	return &ShortenResponse{
		OriginalURL: originalURL,
		ShortenURL:  s.config.Redirect + "/" + code,
	}, nil
}

// isPublicIP allowlists only globally routable unicast IPs.
// It rejects loopback, private, link-local, multicast, unspecified, and CGNAT.
func isPublicIP(ip net.IP) bool {
	return !ip.IsLoopback() &&
		!ip.IsPrivate() &&
		!ip.IsLinkLocalUnicast() &&
		!ip.IsMulticast() &&
		!ip.IsUnspecified() &&
		!cgnat.Contains(ip)
}

func isSafeURL(rawURL string) bool {
	u, err := url.Parse(rawURL)
	if err != nil {
		return false
	}
	// Allowlist: only http and https schemes
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}
	host := u.Hostname()
	addrs, err := net.LookupHost(host)
	if err != nil {
		return false
	}
	// Every resolved IP must be a public address
	for _, addr := range addrs {
		ip := net.ParseIP(addr)
		if ip == nil || !isPublicIP(ip) {
			return false
		}
	}
	return true
}
