CREATE TABLE `tbl_user` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT null default '',
    `email` varchar(50) NOT null default '',
    `pwd` varchar(50) NOT null default '',
    `role` int not null default 0,
    `active` int not null default 0 COMMENT '1-active; 2-inactive',
    `disable_url` varchar(255) not null default '',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by` int not null default 0,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `updated_by` int not null default 0,
    PRIMARY KEY (`id`),
    UNIQUE KEY `tbl_user_un` (`email`),
    KEY `tbl_user_email_IDX` (`email`,`pwd`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8MB4 COMMENT='user info';

CREATE TABLE `tbl_role` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT null default '',
    `active` int not null default 0 COMMENT '1-active; 2-inactive',
    `urls` varchar(255) not null default '',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by` int not null default 0,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `updated_by` int not null default 0,
    PRIMARY KEY (`id`),
    KEY `name_idx` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8MB4 COMMENT='role info';

CREATE TABLE `tbl_url` (
   `id` int unsigned NOT NULL AUTO_INCREMENT,
   `url` varchar(255) NOT null default '',
   `active` int not null default 0 COMMENT '1-active; 2-inactive',
   `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
   `created_by` int not null default 0,
   `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   `updated_by` int not null default 0,
   PRIMARY KEY (`id`),
   KEY `url_idx` (`url`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8MB4 COMMENT='url info';