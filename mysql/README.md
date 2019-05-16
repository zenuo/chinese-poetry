# MySQL

## 创建数据库

```sql
CREATE DATABASE `chinese-poetry` CHARACTER SET 'utf8mb4';
```

## 诗表

```sql
CREATE TABLE `poet` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `dynasty`varchar(11) DEFAULT NULL,
  `author` text DEFAULT NULL,
  `paragraph` text DEFAULT NULL,
  `strains` text DEFAULT NULL,
  `title` text DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

## 词表

```sql
CREATE TABLE `ci` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `dynasty`varchar(11) DEFAULT NULL,
  `author` text DEFAULT NULL,
  `paragraph` text DEFAULT NULL,
  `rhythmic` text DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```