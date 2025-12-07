-- PROJECTS
CREATE TABLE IF NOT EXISTS projects (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    is_live BOOLEAN DEFAULT TRUE,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

-- USERS
CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    platform ENUM('web', 'mobile-app1', 'mobile-app2') NOT NULL,
    last_active_at DATETIME NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

-- SENSORS
CREATE TABLE IF NOT EXISTS sensors (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    project_id BIGINT NOT NULL,
    type VARCHAR(50) NOT NULL,
    status ENUM('online', 'offline') NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id)
);

-- MODULES
CREATE TABLE IF NOT EXISTS modules (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    module_name VARCHAR(100) NOT NULL,
    current_version VARCHAR(50) NOT NULL,
    last_deployed_at DATETIME,
    upcoming_version VARCHAR(50),
    upcoming_release_date DATETIME,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

-- METRICS HISTORY
CREATE TABLE IF NOT EXISTS metrics_history (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    metric_type VARCHAR(100) NOT NULL,
    ts DATETIME NOT NULL,
    value BIGINT NOT NULL,
    created_at DATETIME NOT NULL
);