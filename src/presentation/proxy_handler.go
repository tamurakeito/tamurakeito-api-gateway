package presentation

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/tamurakeito/tamurakeito-api-gateway/src/domain"
	"github.com/tamurakeito/tamurakeito-api-gateway/src/usecase"
	"github.com/tamurakeito/tamurakeito-api-gateway/src/utils"
)

// 以下ではHTTPハンドラーを配置
// HTTPリクエストを受け取り適切なユースケースを呼び出しレスポンスを返す

func RegisterProxies(proxyUsecase usecase.ProxyUsecase) {
	configs, err := proxyUsecase.GetProxies() // このメソッドを usecase に追加する必要があります
	if err != nil {
		log.Fatal(err) // 本番環境では適切なエラーハンドリングを行う
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
