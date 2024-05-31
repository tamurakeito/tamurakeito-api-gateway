package usecase

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/tamurakeito/tamurakeito-api-gateway/src/domain"
	"github.com/tamurakeito/tamurakeito-api-gateway/src/utils"
)

type ProxyUsecase interface {
	SetupProxies()
}

type proxyUsecase struct {
	repo domain.ProxyConfigRepository
}

func NewProxyUsecase(repo domain.ProxyConfigRepository) ProxyUsecase {
	proxyUsecase := proxyUsecase{repo: repo}
	return &proxyUsecase
}

func (p *proxyUsecase) SetupProxies() {
	configs, err := p.repo.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, config := range configs {
		setupProxy(config)
	}
}

func setupProxy(config domain.ProxyConfig) {
	targetURL, err := url.Parse(config.Target)
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
		req.URL.Path = utils.SingleJoiningSlash(targetURL.Path, strings.TrimPrefix(req.URL.Path, config.Path))
	}

	http.HandleFunc(config.Path, func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})
}
