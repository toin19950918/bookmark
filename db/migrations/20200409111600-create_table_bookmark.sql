-- +migrate Up
CREATE TABLE `bookmark` (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL,
    `name` VARCHAR(36) NOT NULL COMMENT '名稱',
    `url` TEXT NOT NULL COMMENT 'url',
    `created_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE,
    PRIMARY KEY(`id`)
) CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='書簽列表';
-- +migrate Down
SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `bookmark`;