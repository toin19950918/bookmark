-- +migrate Up
CREATE TABLE `user` (
     `id` int NOT NULL AUTO_INCREMENT,
     `user_account` VARCHAR(36) NOT NULL COMMENT '使用者帳號',
     `gender` ENUM('male', 'female') NOT NULL COMMENT '性別',
     `created_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
     `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
     PRIMARY KEY(`id`)
) CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='使用者列表';
-- +migrate Down
SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `user`;