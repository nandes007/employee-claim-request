CREATE TABLE `claim_requests`
(
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT,
    claim_category VARCHAR(225),
    claim_date DATE,
    currency VARCHAR(10),
    claim_amount DECIMAL,
    description TEXT,
    support_document VARCHAR(225),
    status VARCHAR(20),
    created_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);