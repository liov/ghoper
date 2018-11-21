DROP TABLE IF EXISTS `moment_tag`;
CREATE TABLE `moment_tag` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `moment_id` int(11) unsigned NOT NULL,
  `tag_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



DROP TABLE IF EXISTS `moment`;
CREATE TABLE `moment` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `content` longtext,
  `mood_id`  int(11) unsigned,
  `user_id` int(11) unsigned NOT NULL,
  `parent_id` int(11) unsigned DEFAULT NULL,
  `desc_flag` int(11) unsigned  NOT NULL DEFAULT '0',
  `browse_count` int(11) unsigned NOT NULL DEFAULT '0',
  `comment_count` int(11) unsigned NOT NULL DEFAULT '0',
  `collect_count` int(11) unsigned NOT NULL DEFAULT '0',
  `like_count`  int(11) unsigned NOT NULL DEFAULT '0',
  `status` int(11) unsigned NOT NULL DEFAULT '0',
  `permission` int(11) unsigned NOT NULL DEFAULT '0',
  `modify_times` int(11) unsigned NOT NULL DEFAULT '0',
  `image_url` varchar(200)  DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `moment_history`;
CREATE TABLE `moment_history` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `content` longtext,
  `mood_id`  int(11) unsigned,
  `user_id` int(11) unsigned NOT NULL,
  `root_id` int(11) unsigned NOT NULL,
  `parent_id` int(11) unsigned DEFAULT NULL,
  `like_count`  int(11) unsigned NOT NULL DEFAULT '0',
  `status` int(11) unsigned NOT NULL DEFAULT '0',
  `permission` int(11) unsigned NOT NULL DEFAULT '0',
  `modify_times` int(11) unsigned NOT NULL DEFAULT '0',
  `image_url` varchar(200)  DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(200)  NOT NULL,
  `created_by` varchar(200)  NOT NULL,
  `status`  int(11) unsigned NOT NULL DEFAULT '0',
  `moment_count` int(11) unsigned NOT NULL DEFAULT '0',
  `article_count`int(11) unsigned NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
)

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(200)  NOT NULL,
  `last_name` varchar(200)  NOT NULL,
  `password` varchar(200)  NOT NULL,
  `email` varchar(50) DEFAULT '',
  `score` int(11) unsigned NOT NULL,
  `phone` varchar(50) DEFAULT NULL,
  `sex` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `location` varchar(200) DEFAULT NULL,
  `introduce` varchar(500) DEFAULT NULL,
  `role` int(11) NOT NULL,
  `avatar_url` varchar(500) NOT NULL DEFAULT '',
  `cover_url` varchar(500) DEFAULT NULL,
  `status`  int(11) unsigned NOT NULL DEFAULT '0',
  `moment_count` int(11) unsigned NOT NULL DEFAULT '0',
  `article_count`int(11) unsigned NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
)

DROP TABLE IF EXISTS `collection`;
CREATE TABLE `collection` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL,
  `created_by` varchar(200)  NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
)

DROP TABLE IF EXISTS `collection_article`;
CREATE TABLE `collection` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `collection_id` int(11) unsigned NOT NULL,
  `article_id` int(11) unsigned NOT NULL,
  `created_by` varchar(200)  NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
)

DROP TABLE IF EXISTS `collection_article`;
CREATE TABLE `collection` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `collection_id` int(11) unsigned NOT NULL,
  `article_id` int(11) unsigned NOT NULL,
  `created_by` varchar(200)  NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
)

DROP TABLE IF EXISTS `article_history_category`;
CREATE TABLE `article_history_category` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `article_history_id` int(11) unsigned NOT NULL,
  `category_id` int(11) unsigned NOT NULL,
  `created_by` varchar(200)  NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
)

DROP TABLE IF EXISTS `article_category`;
CREATE TABLE `article_category` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(11) unsigned NOT NULL,
  `category_id` int(11) unsigned NOT NULL,
  `created_by` varchar(200)  NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
)

CREATE TABLE `article` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(200)  DEFAULT NULL,
  `content` longtext,
  `html_content` longtext,
  `content_type` int(11) NOT NULL,
  `user_id` int(11) unsigned NOT NULL,
  `parent_id` int(11) unsigned DEFAULT NULL,
  `desc_flag` int(11) unsigned  NOT NULL DEFAULT '0',
  `browse_count` int(11) unsigned NOT NULL DEFAULT '0',
  `comment_count` int(11) unsigned NOT NULL DEFAULT '0',
  `collect_count` int(11) unsigned NOT NULL DEFAULT '0',
  `like_count`  int(11) unsigned NOT NULL DEFAULT '0',
  `status` int(11) unsigned NOT NULL DEFAULT '0',
  `permission` int(11) unsigned NOT NULL DEFAULT '0',
  `modify_times` int(11) unsigned NOT NULL DEFAULT '0',
  `image_url` varchar(65535)  DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `last_comment_at` datetime DEFAULT NULL,
  `lastuser_id` int(11) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

