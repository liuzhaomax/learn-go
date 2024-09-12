-- prismer.`user` definition

CREATE TABLE `user` (
  `id` varchar(22) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `phone` varchar(11) COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `salt` varchar(16) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `invitation_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_user_phone` (`phone`),
  KEY `idx_user_deleted_at` (`deleted_at`),
  KEY `idx_phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- prismer.user_info definition

CREATE TABLE `user_info` (
  `id` varchar(22) COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(50) COLLATE utf8mb4_general_ci NOT NULL,
  `avatar` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `info` json DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_user_user_info` FOREIGN KEY (`id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- prismer.user_wechat_info definition

CREATE TABLE `user_wechat_info` (
  `id` varchar(22) COLLATE utf8mb4_general_ci NOT NULL,
  `openid` varchar(50) COLLATE utf8mb4_general_ci NOT NULL,
  `nickname` varchar(50) COLLATE utf8mb4_general_ci NOT NULL,
  `sex` bigint NOT NULL DEFAULT '0',
  `province` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `city` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `country` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `headimgurl` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `privilege` json DEFAULT NULL,
  `union_id` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_user_wechat_info_openid` (`openid`),
  KEY `idx_openid` (`openid`),
  CONSTRAINT `fk_user_user_wechat_info` FOREIGN KEY (`id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- prismer.tag definition

CREATE TABLE `tag` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
  `desc` varchar(1024) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `parent_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_tag_name` (`name`),
  KEY `idx_name` (`name`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_tag_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=163 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- prismer.user_info_tag definition

CREATE TABLE `user_info_tag` (
  `tag_id` bigint unsigned NOT NULL,
  `user_info_id` varchar(22) COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`tag_id`,`user_info_id`),
  KEY `fk_user_info_tag_user_info` (`user_info_id`),
  CONSTRAINT `fk_user_info_tag_tag` FOREIGN KEY (`tag_id`) REFERENCES `tag` (`id`),
  CONSTRAINT `fk_user_info_tag_user_info` FOREIGN KEY (`user_info_id`) REFERENCES `user_info` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

