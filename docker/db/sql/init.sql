CREATE DATABASE IF NOT EXISTS proxy_config;
USE proxy_config;

CREATE TABLE IF NOT EXISTS proxy_settings (
    id INT AUTO_INCREMENT PRIMARY KEY,
    path VARCHAR(255) NOT NULL,
    target VARCHAR(255) NOT NULL
);

INSERT INTO proxy_settings (path, target) VALUES
('/memo-app/', 'http://35.233.218.140');