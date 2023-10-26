CREATE TABLE `users`
(
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(225),
    email VARCHAR(255),
    password VARCHAR(255),
    company_id BIGINT,
    is_admin TINYINT,
    created_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);