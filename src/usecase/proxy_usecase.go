package usecase

import (
	"github.com/tamurakeito/tamurakeito-api-gateway/src/domain"
)

// プロキシ設定データを取得するためのインターフェースを提供する
// プロキシ設定の具体的な実装はpresentation層HTTPハンドラーに記述

type ProxyUsecase interface {
	GetProxies() ([]domain.ProxyConfig, error)
}

type proxyUsecase struct {
	repo domain.ProxyConfigRepository
}

func NewProxyUsecase(repo domain.ProxyConfigRepository) ProxyUsecase {
	return &proxyUsecase{repo: repo}
}

func (p *proxyUsecase) GetProxies() ([]domain.ProxyConfig, error) {
	return p.repo.FindAll()
}
