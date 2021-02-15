-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `events`
(
    `id`            varchar(255) NOT NULL,
    `title`         varchar(67)  NOT NULL,
    `date`          datetime DEFAULT NULL,
    `duration`      int          NOT NULL,
    `description`   varchar(127),
    `owner_id`      int          NOT NULL,
    `notify_before` int          NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS events;

