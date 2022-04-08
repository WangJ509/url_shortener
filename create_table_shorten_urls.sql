DROP TABLE IF EXISTS `shorten_urls`;

CREATE TABLE `shorten_urls` (
    `url` varchar(255) NOT NULL UNIQUE,
    `original_url` varchar(255) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `expired_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);