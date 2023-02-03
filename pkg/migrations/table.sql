CREATE TABLE `users` (
  `Id` varchar(255) PRIMARY KEY,
  `username` varchar(255),
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL
);

CREATE TABLE `manga` (
  `Id` bigint PRIMARY KEY NOT NULL,
  `Title` varchar(255) NOT NULL,
  `Thumb` varchar(255) NOT NULL,
  `Author` varchar(255) NOT NULL,
  `Publisher` varchar(255) NOT NULL,
  `YearPublished` varchar(255) NOT NULL,
  `Status` ENUM ('Publish', 'Draft') NOT NULL,
  `Genre` ENUM ('Manga', 'Manhwa', 'Manhua') NOT NULL DEFAULT "manga",
  `CreateBy` varchar(255)
);

CREATE TABLE `category` (
  `Id` varchar(255) PRIMARY KEY NOT NULL,
  `Name` varchar(255)
);

CREATE TABLE `category_manga` (
  `manga_id` varchar(255),
  `category_id` varchar(255)
);

CREATE TABLE `chapters` (
  `Id` varchar(255) PRIMARY KEY,
  `manga_id` varchar(255),
  `number` int,
  `publish_by` varchar(255),
  `content` varchar(255)
);
