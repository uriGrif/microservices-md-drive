CREATE TABLE permission (
    file_id VARCHAR(255),
    user_id VARCHAR(255),
    permission_level int,
    PRIMARY KEY (file_id, user_id)
);