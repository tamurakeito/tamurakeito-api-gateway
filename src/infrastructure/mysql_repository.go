package infrastructure

import (
	"github.com/tamurakeito/tamurakeito-api-gateway/src/domain"
)

type MySQLProxyConfigRepository struct {
	SqlHandler
}

func NewMySQLProxyConfigRepository(sqlHandler SqlHandler) domain.ProxyConfigRepository {
	proxyRepository := MySQLProxyConfigRepository{sqlHandler}
	return &proxyRepository
}

func (m *MySQLProxyConfigRepository) FindAll() ([]domain.ProxyConfig, error) {
	rows, err := m.SqlHandler.Conn.Query("SELECT path, target FROM proxy_settings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var configs []domain.ProxyConfig
	for rows.Next() {
		var config domain.ProxyConfig
		if err := rows.Scan(&config.Path, &config.Target); err != nil {
			return nil, err
		}
		configs = append(configs, config)
	}
	return configs, nil
}
