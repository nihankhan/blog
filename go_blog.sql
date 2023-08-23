SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

-- Database Name: blog

CREATE DATABASE IF NOT EXISTS blog;

-- Table structure for table "articles"

USE blog;

CREATE TABLE IF NOT EXISTS blog.articles (
    id int(11) NOT NULL,
    image VARCHAR(255) DEFAULT NULL,
    slug VARCHAR(255) NOT NULL,
    title VARCHAR(70) NOT NULL,
    content TEXT NOT NULL,
    author int(11) NOT NULL,
    created_at DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Table structure for table: user

CREATE TABLE IF NOT EXISTS blog.users (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(64) NOT NULL,
    email VARCHAR(330) NOT NULL,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Indexes for table: articles

ALTER TABLE blog.articles
    ADD PRIMARY KEY(`id`),
    ADD UNIQUE KEY `slug` (`slug`),
    ADD KEY `author` (`author`);

-- Indexes for table: users

ALTER TABLE blog.users
    ADD UNIQUE KEY(`email`);

-- AUTO_INCREMENT for table: articles    

ALTER TABLE blog.articles
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

-- AUTO_INCREMENT for table: users    

ALTER TABLE blog.users
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
