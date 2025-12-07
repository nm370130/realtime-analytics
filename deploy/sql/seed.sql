SET FOREIGN_KEY_CHECKS = 0;
TRUNCATE TABLE modules;
TRUNCATE TABLE projects;
TRUNCATE TABLE sensors;
TRUNCATE TABLE metrics_history;
TRUNCATE TABLE users;
SET FOREIGN_KEY_CHECKS = 1;

-- =====================================================================
-- MODULES (20 rows)
-- Different versions, some missing upcoming release, mixed dates
-- =====================================================================

INSERT INTO modules (module_name, current_version, last_deployed_at, upcoming_version, upcoming_release_date, created_at)
VALUES
('auth-service', '1.4.2', NOW() - INTERVAL 20 DAY, '1.5.0', '2025-01-10', NOW()),
('payment-service', '2.1.0', NOW() - INTERVAL 18 DAY, '2.2.0', '2025-02-01', NOW()),
('sensor-ingestion', '3.0.1', NOW() - INTERVAL 10 DAY, NULL, NULL, NOW()),
('reporting-service', '1.0.0', NOW() - INTERVAL 30 DAY, '1.1.0', '2025-03-05', NOW()),
('alert-engine', '1.3.4', NOW() - INTERVAL 7 DAY, '1.4.0', '2025-01-20', NOW()),
('device-sync', '4.0.0', NOW() - INTERVAL 15 DAY, NULL, NULL, NOW()),
('rules-engine', '2.3.1', NOW() - INTERVAL 60 DAY, '2.4.0', '2025-04-10', NOW()),
('config-service', '1.9.9', NOW() - INTERVAL 2 DAY, '2.0.0', '2025-01-08', NOW()),
('ui-portal', '5.1.0', NOW() - INTERVAL 25 DAY, NULL, NULL, NOW()),
('scheduler', '3.2.1', NOW() - INTERVAL 29 DAY, '3.3.0', '2025-03-22', NOW()),
('batch-processor', '1.2.2', NOW() - INTERVAL 10 DAY, NULL, NULL, NOW()),
('pipeline-api', '2.2.1', NOW() - INTERVAL 8 DAY, '2.3.0', '2025-02-12', NOW()),
('asset-service', '1.0.1', NOW() - INTERVAL 6 DAY, NULL, NULL, NOW()),
('sync-service', '7.3.2', NOW() - INTERVAL 4 DAY, '7.4.0', '2025-01-30', NOW()),
('geo-service', '0.9.8', NOW() - INTERVAL 40 DAY, NULL, NULL, NOW()),
('push-notifications', '1.3.0', NOW() - INTERVAL 90 DAY, '1.4.0', '2025-05-05', NOW()),
('tenant-service', '2.0.3', NOW() - INTERVAL 11 DAY, NULL, NULL, NOW()),
('gateway-proxy', '9.0.1', NOW() - INTERVAL 12 DAY, '9.1.0', '2025-04-15', NOW()),
('voice-service', '1.0.0', NOW() - INTERVAL 15 DAY, NULL, NULL, NOW()),
('email-service', '2.3.3', NOW() - INTERVAL 5 DAY, '2.4.0', '2025-06-10', NOW());

-- =====================================================================
-- PROJECTS (20 rows)
-- Mixed live & offline projects, different creation times
-- =====================================================================

INSERT INTO projects (name, is_live, created_at)
VALUES
('Smart Factory Dashboard', TRUE, NOW() - INTERVAL 2 DAY),
('City Traffic IoT', TRUE, NOW() - INTERVAL 6 DAY),
('Smart Irrigation', FALSE, NOW() - INTERVAL 10 DAY),
('Parking Automation', TRUE, NOW() - INTERVAL 12 DAY),
('Warehouse Control', FALSE, NOW() - INTERVAL 1 DAY),
('Building Automation', TRUE, NOW() - INTERVAL 20 DAY),
('Drone Tracking', TRUE, NOW() - INTERVAL 3 DAY),
('Campus Security', FALSE, NOW() - INTERVAL 15 DAY),
('Cold Storage Monitoring', TRUE, NOW() - INTERVAL 7 DAY),
('Fleet Management', TRUE, NOW() - INTERVAL 14 DAY),
('Hospital IoT System', FALSE, NOW() - INTERVAL 30 DAY),
('Energy Grid Monitor', TRUE, NOW() - INTERVAL 25 DAY),
('Water Treatment', FALSE, NOW() - INTERVAL 4 DAY),
('Server Room Sensors', TRUE, NOW() - INTERVAL 9 DAY),
('Smart Elevators', TRUE, NOW() - INTERVAL 18 DAY),
('HVAC Analytics', FALSE, NOW() - INTERVAL 11 DAY),
('Environmental Monitoring', TRUE, NOW() - INTERVAL 13 DAY),
('Public Transport IoT', FALSE, NOW() - INTERVAL 17 DAY),
('School Automation', TRUE, NOW() - INTERVAL 19 DAY),
('Apartment Complex IoT', TRUE, NOW() - INTERVAL 22 DAY);

-- =====================================================================
-- SENSORS (20 rows)
-- Covers all sensor types, online/offline, various projects
-- =====================================================================

INSERT INTO sensors (project_id, type, status, created_at)
VALUES
(1, 'temperature', 'online', NOW()),
(1, 'temperature', 'offline', NOW()),
(2, 'humidity', 'online', NOW()),
(2, 'motion', 'offline', NOW()),
(3, 'airQuality', 'online', NOW()),
(3, 'temperature', 'online', NOW()),
(4, 'vibration', 'offline', NOW()),
(4, 'humidity', 'online', NOW()),
(5, 'motion', 'online', NOW()),
(5, 'temperature', 'offline', NOW()),
(6, 'sound', 'online', NOW()),
(6, 'airQuality', 'offline', NOW()),
(7, 'waterLevel', 'online', NOW()),
(7, 'temperature', 'online', NOW()),
(8, 'humidity', 'offline', NOW()),
(9, 'motion', 'online', NOW()),
(10, 'temperature', 'online', NOW()),
(10, 'motion', 'offline', NOW()),
(11, 'airQuality', 'online', NOW()),
(12, 'humidity', 'online', NOW());

-- =====================================================================
-- USERS (20 rows)
-- Mixed platforms for activeUsersByPlatform testing
-- =====================================================================

INSERT INTO users (username, platform, created_at)
VALUES
('john_web', 'web', NOW()),
('nancy_web', 'web', NOW()),
('alex_web', 'web', NOW()),
('marie_app1', 'mobile-app1', NOW()),
('user1_app1', 'mobile-app1', NOW()),
('user2_app1', 'mobile-app1', NOW()),
('mike_app2', 'mobile-app2', NOW()),
('priya_app2', 'mobile-app2', NOW()),
('dev_api', 'api', NOW()),
('test_api', 'api', NOW()),
('tester1', 'web', NOW()),
('tester2', 'mobile-app1', NOW()),
('tester3', 'mobile-app2', NOW()),
('qa_user', 'web', NOW()),
('ops_user', 'mobile-app2', NOW()),
('admin_user', 'web', NOW()),
('guest1', 'api', NOW()),
('guest2', 'web', NOW()),
('guest3', 'mobile-app1', NOW()),
('guest4', 'api', NOW());

-- =====================================================================
-- METRICS HISTORY (20+ rows)
-- For both activeUsers and apiRejectedCount
-- =====================================================================

INSERT INTO metrics_history (metric_type, ts, value, created_at)
VALUES
('activeUsers', NOW() - INTERVAL 19 MINUTE, 12, NOW()),
('activeUsers', NOW() - INTERVAL 18 MINUTE, 15, NOW()),
('activeUsers', NOW() - INTERVAL 17 MINUTE, 18, NOW()),
('activeUsers', NOW() - INTERVAL 16 MINUTE, 20, NOW()),
('activeUsers', NOW() - INTERVAL 15 MINUTE, 22, NOW()),
('activeUsers', NOW() - INTERVAL 14 MINUTE, 30, NOW()),
('activeUsers', NOW() - INTERVAL 13 MINUTE, 31, NOW()),
('activeUsers', NOW() - INTERVAL 12 MINUTE, 33, NOW()),
('activeUsers', NOW() - INTERVAL 11 MINUTE, 35, NOW()),
('activeUsers', NOW() - INTERVAL 10 MINUTE, 37, NOW()),
('activeUsers', NOW() - INTERVAL 9 MINUTE, 39, NOW()),
('activeUsers', NOW() - INTERVAL 8 MINUTE, 40, NOW()),
('activeUsers', NOW() - INTERVAL 7 MINUTE, 38, NOW()),
('activeUsers', NOW() - INTERVAL 6 MINUTE, 36, NOW()),
('activeUsers', NOW() - INTERVAL 5 MINUTE, 42, NOW()),
('activeUsers', NOW() - INTERVAL 4 MINUTE, 45, NOW()),
('activeUsers', NOW() - INTERVAL 3 MINUTE, 47, NOW()),
('activeUsers', NOW() - INTERVAL 2 MINUTE, 49, NOW()),
('activeUsers', NOW() - INTERVAL 1 MINUTE, 50, NOW()),
('activeUsers', NOW(), 48, NOW()),

('apiRejectedCount', NOW() - INTERVAL 5 MINUTE, 2, NOW()),
('apiRejectedCount', NOW() - INTERVAL 4 MINUTE, 3, NOW()),
('apiRejectedCount', NOW() - INTERVAL 3 MINUTE, 4, NOW()),
('apiRejectedCount', NOW() - INTERVAL 2 MINUTE, 5, NOW()),
('apiRejectedCount', NOW() - INTERVAL 1 MINUTE, 3, NOW());

SELECT 'Seed Data Loaded Successfully' AS message;