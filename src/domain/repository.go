package domain

type ProxyConfigRepository interface {
	FindAll() ([]ProxyConfig, error)
}
