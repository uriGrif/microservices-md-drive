CREATE TABLE permission (
    file_id VARCHAR(255),
    user_id VARCHAR(255),
    permission_level INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (file_id, user_id)
);