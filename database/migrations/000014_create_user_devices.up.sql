CREATE TABLE user_devices (
    id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    fmc_token VARCHAR(255) NOT NULL,
    device_type ENUM('ios', 'android') NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE, -- Tambah koma di sini
    INDEX idx_user_id (user_id) -- Membuat index langsung di dalam tabel
);