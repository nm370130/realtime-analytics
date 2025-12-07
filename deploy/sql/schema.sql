-- ===========================================================
--   DATABASE: realtime_analytics
--   SCHEMA FOR ASSIGNMENT
-- ===========================================================

CREATE DATABASE IF NOT EXISTS realtime_analytics;
USE realtime_analytics;

SET FOREIGN_KEY_CHECKS = 0;


-- ===========================================================
--   TABLE: modules
-- ===========================================================

CREATE TABLE IF NOT EXISTS modules (
    id                       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    module_name              VARCHAR(255) NOT NULL,
    current_version          VARCHAR(50) NOT NULL,
    last_deployed_at         DATETIME NULL,
    upcoming_version         VARCHAR(50) NULL,
    upcoming_release_date    DATE NULL,
    
    created_at               DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at               DATETIME NULL ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- ===========================================================
--   TABLE: projects
-- ===========================================================

CREATE TABLE IF NOT EXISTS projects (
    id          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    is_live     BOOLEAN NOT NULL DEFAULT FALSE,

    created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME NULL ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- ===========================================================
--   TABLE: sensors
-- ===========================================================

CREATE TABLE IF NOT EXISTS sensors (
    id          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    project_id  BIGINT UNSIGNED NOT NULL,
    type        VARCHAR(100) NOT NULL,
    status      ENUM('online','offline') NOT NULL,

    created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME NULL ON UPDATE CURRENT_TIMESTAMP,

    INDEX idx_sensor_type (type),
    INDEX idx_sensor_status (status),

    CONSTRAINT fk_sensor_project
        FOREIGN KEY (project_id)
        REFERENCES projects (id)
        ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- ===========================================================
--   TABLE: users (optional, for analytics)
-- ===========================================================

CREATE TABLE IF NOT EXISTS users (
    id          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username    VARCHAR(255) NOT NULL,
    platform    VARCHAR(100) NOT NULL,

    created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME NULL ON UPDATE CURRENT_TIMESTAMP,

    INDEX idx_users_platform (platform)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- ===========================================================
--   TABLE: metrics_history
-- ===========================================================

CREATE TABLE IF NOT EXISTS metrics_history (
    id           BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    metric_type  VARCHAR(100) NOT NULL,
    ts           DATETIME NOT NULL,
    value        BIGINT NOT NULL,

    created_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    INDEX idx_metrics_metric_type (metric_type),
    INDEX idx_metrics_ts (ts),
    INDEX idx_metrics_metric_ts (metric_type, ts)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


SET FOREIGN_KEY_CHECKS = 1;

-- ===========================================================
-- Done
-- ===========================================================

SELECT 'Database schema created successfully!' AS message;