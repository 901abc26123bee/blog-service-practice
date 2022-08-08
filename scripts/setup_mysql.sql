/*
    blog_article     blog_article
          |               |
          |               |
          blog_article_tag

docker run --name blog-mysql -e MYSQL_ROOT_PASSWORD=123456 -p 3310:3306 mysql/mysql-server:8.0
*/


CREATE DATABASE
IF
	NOT EXISTS blog_service DEFAULT CHARACTER
	SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

USE blog_service;

CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT 'tag',

  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'creation_time',
  `created_by` varchar(100) DEFAULT '' COMMENT 'creator',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'modification_time',
  `modified_by` varchar(100) DEFAULT '' COMMENT 'modifier',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'deletion_time',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '0:not deleted, 1:deleted',

  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '0:disabled, 1:enable',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='blog_tag';

CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT 'article_title',
  `desc` varchar(255) DEFAULT '' COMMENT 'article_description',
  `cover_image_url` varchar(255) DEFAULT '' COMMENT 'cover_image_url',
  `content` longtext COMMENT 'article_content',

  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'creation_time',
  `created_by` varchar(100) DEFAULT '' COMMENT 'creator',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'modification_time',
  `modified_by` varchar(100) DEFAULT '' COMMENT 'modifier',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'deletion_time',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '0:not deleted, 1:deleted',

  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '0:disabled, 1:enable',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='blog_article';

CREATE TABLE `blog_article_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NOT NULL COMMENT 'article_id',
  `tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'tag_id',

  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'creation_time',
  `created_by` varchar(100) DEFAULT '' COMMENT 'creator',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'modification_time',
  `modified_by` varchar(100) DEFAULT '' COMMENT 'modifier',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'deletion_time',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '0:not deleted, 1:deleted',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='blog_article_tag';

-- 創建一個新的數據表，用於存儲簽發的認證信息
CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `app_key` varchar(20) DEFAULT '' COMMENT 'Key',
  `app_secret` varchar(50) DEFAULT '' COMMENT 'Secret',
  
  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'creation_time',
  `created_by` varchar(100) DEFAULT '' COMMENT 'creator',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'modification_time',
  `modified_by` varchar(100) DEFAULT '' COMMENT 'modifier',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'deletion_time',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '0:not deleted, 1:deleted',

  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='blog_authorization';

INSERT INTO `blog_service`.`blog_auth`(`id`, `app_key`, `app_secret`, `created_on`, `created_by`, `modified_on`, `modified_by`, `deleted_on`, `is_del`) VALUES (1, 'tdlemon', 'qwe123hahaha', 0, 'tdlemon', 0, '', 0, 0);
