CREATE DATABASE IF NOT EXISTS hb_student_tracker1;

USE hb_student_tracker1;

DROP TABLE IF EXISTS student;
DROP TABLE IF EXISTS `instructor`;
DROP TABLE IF EXISTS `instructor_detail`;
DROP TABLE IF EXISTS `course`;
DROP TABLE IF EXISTS `course_student`;
DROP TABLE IF EXISTS `reviews`;

CREATE TABLE student (
                         id BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
                         first_name VARCHAR(255) DEFAULT NULL,
                         last_name VARCHAR(255) DEFAULT NULL,
                         email VARCHAR(255) DEFAULT NULL
) AUTO_INCREMENT = 1 DEFAULT CHARSET=UTF8;

INSERT INTO student (first_name, last_name, email) VALUES
                                                       ('Ivan', 'Hornik', 'ih@gmail.com'),
                                                       ('Jane', 'Smith', 'jane.smith@email.com'),
                                                       ('Jim', 'Beam', 'jim.beam@email.com');

ALTER TABLE `course` DROP FOREIGN KEY `FK_INSTRUCTOR`;
DROP TABLE IF EXISTS `instructor`;
DROP TABLE IF EXISTS `instructor_detail`;

CREATE TABLE `instructor_detail` (
                                     `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
                                     `youtube_channel` VARCHAR(128) DEFAULT NULL,
                                     `hobby` VARCHAR(45) DEFAULT NULL,
                                     PRIMARY KEY (`id`)
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE TABLE `instructor` (
                              `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
                              `first_name` VARCHAR(45) DEFAULT NULL,
                              `last_name` VARCHAR(45) DEFAULT NULL,
                              `email` VARCHAR(45) DEFAULT NULL,
                              `instructor_detail_id` BIGINT UNSIGNED DEFAULT NULL,
                              PRIMARY KEY (`id`),
                              KEY `FK_DETAIL_idx` (`instructor_detail_id`),
                              CONSTRAINT `FK_DETAIL` FOREIGN KEY (`instructor_detail_id`) REFERENCES `instructor_detail` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `course`;

CREATE TABLE `course` (
                          `id` BIGINT NOT NULL AUTO_INCREMENT,
                          `title` varchar(128) DEFAULT NULL,
                          `instructor_id` BIGINT UNSIGNED DEFAULT NULL,
                          PRIMARY KEY (`id`),
                          UNIQUE KEY `TITLE_UNIQUE` (`title`),
                          KEY `FK_INSTRUCTOR_idx` (`instructor_id`),
                          CONSTRAINT `FK_INSTRUCTOR` FOREIGN KEY (`instructor_id`) REFERENCES `instructor` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;
DROP TABLE IF EXISTS `course_student`;

CREATE TABLE `course_student` (
                                  `course_id` BIGINT  NOT NULL,
                                  `student_id` BIGINT  NOT NULL,

                                  PRIMARY KEY (`course_id`, `student_id`),

                                  KEY `FK_STUDENT_idx` (`student_id`),

                                  CONSTRAINT `FK_COURSE` FOREIGN KEY (`course_id`) REFERENCES `course` (`id`)
                                      ON DELETE NO ACTION ON UPDATE NO ACTION,

                                  CONSTRAINT `FK_STUDENT` FOREIGN KEY (`student_id`) REFERENCES `student` (`id`)
                                      ON DELETE NO ACTION ON UPDATE NO ACTION
) DEFAULT CHARSET=UTF8;





