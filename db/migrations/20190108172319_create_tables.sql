
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE articles (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  title varchar(32) DEFAULT NULL,
  body text DEFAULT NULL,
  tag_ids varchar(32)  DEFAULT NULL,
  created_at datetime DEFAULT NULL,
  updated_at datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE comments (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  title varchar(32) DEFAULT NULL,
  body text DEFAULT NULL,
  created_at datetime DEFAULT NULL,
  updated_at datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE tags (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(32) DEFAULT NULL,
  created_at datetime DEFAULT NULL,
  updated_at datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE articles;
DROP TABLE comments;
DROP TABLE tags;