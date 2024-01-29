-- Create the database if it doesn't exist
CREATE DATABASE IF NOT EXISTS hb_student_tracker1;
USE hb_student_tracker1;

-- Drop tables in the correct order to avoid foreign key constraint issues
DROP TABLE IF EXISTS `reviews`;
DROP TABLE IF EXISTS `course_student`;
DROP TABLE IF EXISTS `course`;
DROP TABLE IF EXISTS `instructor`;
DROP TABLE IF EXISTS `instructor_detail`;
DROP TABLE IF EXISTS `student`;
DROP TABLE IF EXISTS `reviews`;

-- Create student table
CREATE TABLE student (
                         id BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
                         first_name VARCHAR(255) DEFAULT NULL,
                         last_name VARCHAR(255) DEFAULT NULL,
                         email VARCHAR(255) DEFAULT NULL
) AUTO_INCREMENT = 1 DEFAULT CHARSET=UTF8;

-- Insert sample data into student table
INSERT INTO student (first_name, last_name, email) VALUES
                                                       ('Alice', 'Johnson', 'alice.johnson@email.com'),
                                                       ('Bob', 'Brown', 'bob.brown@email.com'),
                                                       ('Carol', 'Davis', 'carol.davis@email.com'),
                                                       ('David', 'Evans', 'david.evans@email.com');

-- Create instructor_detail table
CREATE TABLE `instructor_detail` (
                                     `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
                                     `youtube_channel` VARCHAR(128) DEFAULT NULL,
                                     `hobby` VARCHAR(45) DEFAULT NULL,
                                     PRIMARY KEY (`id`)
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `instructor_detail` (youtube_channel, hobby) VALUES
                                                             ('channelA', 'Gardening'),
                                                             ('channelB', 'Cooking'),
                                                             ('channelC', 'Photography'),
                                                             ('channelD', 'Hiking');

-- Create instructor table
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

INSERT INTO `instructor` (first_name, last_name, email, instructor_detail_id) VALUES
                                                                                  ('Eva', 'Fisher', 'eva.fisher@email.com', 1),
                                                                                  ('Frank', 'Green', 'frank.green@email.com', 2),
                                                                                  ('Grace', 'Hall', 'grace.hall@email.com', 3),
                                                                                  ('Harry', 'Irwin', 'harry.irwin@email.com', 4);


-- Create course table
CREATE TABLE `course` (
                          `id` BIGINT NOT NULL AUTO_INCREMENT,
                          `title` varchar(128) DEFAULT NULL,
                          `instructor_id` BIGINT UNSIGNED DEFAULT NULL,
                          PRIMARY KEY (`id`),
                          UNIQUE KEY `TITLE_UNIQUE` (`title`),
                          KEY `FK_INSTRUCTOR_idx` (`instructor_id`),
                          CONSTRAINT `FK_INSTRUCTOR` FOREIGN KEY (`instructor_id`) REFERENCES `instructor` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

INSERT INTO `course` (title, instructor_id) VALUES
                                                ('Mathematics 101', 1),
                                                ('Physics for Beginners', 2),
                                                ('Introduction to Programming', 3),
                                                ('World History', 4);

-- Create course_student table
CREATE TABLE `course_student` (
                                  `course_id` BIGINT  NOT NULL,
                                  `student_id` BIGINT  NOT NULL,
                                  PRIMARY KEY (`course_id`, `student_id`),
                                  KEY `FK_STUDENT_idx` (`student_id`),
                                  CONSTRAINT `FK_COURSE` FOREIGN KEY (`course_id`) REFERENCES `course` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
                                  CONSTRAINT `FK_STUDENT` FOREIGN KEY (`student_id`) REFERENCES `student` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) DEFAULT CHARSET=UTF8;

INSERT INTO `course_student` (course_id, student_id) VALUES
                                                         (10, 1),
                                                         (11, 2),
                                                         (12, 3),
                                                         (13, 4);

CREATE TABLE `reviews` (
                           `id` BIGINT NOT NULL AUTO_INCREMENT,
                           `course_id` BIGINT NOT NULL,
                           `comment` TEXT,
                           `rating` TINYINT CHECK (rating >= 0 AND rating <= 5),
                           PRIMARY KEY (`id`),
                           FOREIGN KEY (`course_id`) REFERENCES `course` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) DEFAULT CHARSET=UTF8;
INSERT INTO `reviews` (course_id, comment, rating) VALUES
                                                       (10, 'Great introduction to mathematics.', 5),
                                                       (11, 'Very informative and well-paced physics course.', 4),
                                                       (12, 'Excellent course for beginners in programming.', 5),
                                                       (13, 'Good historical overview but needed more depth.', 3);

