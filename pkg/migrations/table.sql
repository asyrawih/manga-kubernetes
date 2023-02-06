CREATE TABLE `users` (
  `Id` varchar(255) PRIMARY KEY,
  `username` varchar(30) UNIQUE NOT NULL,
  `name` varchar(30) NOT NULL,
  `email` varchar(60) UNIQUE NOT NULL,
  `password` varchar(60) NOT NULL
);

CREATE TABLE `manga` (
  `id` bigint PRIMARY KEY NOT NULL,
  `title` varchar(125) NOT NULL,
  `thumb` varchar(100) NOT NULL,
  `author` varchar(30) NOT NULL,
  `publisher` varchar(30) NOT NULL,
  `year_published` varchar(4) NOT NULL,
  `status` ENUM ('Publish', 'Draft') NOT NULL,
  `genre` ENUM ('Manga', 'Manhwa', 'Manhua') NOT NULL DEFAULT "manga",
  `create_by` varchar(255)
);

CREATE INDEX `username_index` ON `users` (`username`);

CREATE UNIQUE INDEX `users_index_1` ON `users` (`Id`);

CREATE INDEX `create_by_status_index` ON `manga` (`create_by`, `status`);

CREATE UNIQUE INDEX `manga_index_1` ON `manga` (`id`);
