--
-- Base de datos: `email`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `aliases`
--
CREATE TABLE IF NOT EXISTS `aliases` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `local` VARCHAR(255) NOT NULL DEFAULT '',
  `remoto` TEXT DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------
--
-- Estructura de tabla para la tabla `users`
--
CREATE TABLE IF NOT EXISTS `users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `userid` INT NOT NULL DEFAULT 65000,
  `login` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `email` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `password` VARCHAR(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `maildir` MEDIUMTEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `identificacion` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `grupo` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `dominio` INT NOT NULL DEFAULT 0,
  `quota` BIGINT NOT NULL DEFAULT 30000000,
  `policyd_messagequota` INT UNSIGNED NOT NULL DEFAULT 100,
  `policyd_messagetally` INT NOT NULL DEFAULT 0,
  `policyd_timestamp` INT DEFAULT NULL,
  `smtpok` TINYINT(1) NOT NULL DEFAULT 1,
  `imapok` TINYINT(1) NOT NULL DEFAULT 1,
  `pop3ok` TINYINT(1) NOT NULL DEFAULT 1,
  `active` TINYINT(1) NOT NULL DEFAULT 1,
  `created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `userid` (`userid`),
  KEY `email` (`email`(250)),
  KEY `password` (`password`),
  KEY `dominio` (`dominio`),
  KEY `login` (`login`(250))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------
--
-- Estructura de tabla para la tabla `transport`
--
CREATE TABLE IF NOT EXISTS `transport` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `domain` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `transport` VARCHAR(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  UNIQUE KEY `domain_2` (`domain`),
  KEY `domain` (`domain`),
  KEY `id` (`id`),
  KEY `transport` (`transport`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

------
-- Creamos un nuevo usuario
create user 'web'@'localhost' identified by 'password';
grant select, insert, update, delete, create, references, alter on email.* to 'web'@'localhost';
flush privileges;

-- mysql -D email -u web -p