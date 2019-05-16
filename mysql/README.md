# MySQL

## 创建数据库

```sql
CREATE DATABASE `chinese-poetry` CHARACTER SET 'utf8mb4';
```

## 宋诗表

```sql
CREATE TABLE `poet` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `dynasty`varchar(11) DEFAULT NULL,
  `author` varchar(31) DEFAULT NULL,
  `paragraph` varchar(1023) DEFAULT NULL,
  `strains` varchar(1023) DEFAULT NULL,
  `title` varchar(511) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```