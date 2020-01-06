# xo tutorial
```
docker run --name test-mysql --rm -d -e MYSQL_ROOT_PASSWORD=my-pw -p 3306:3306 mysql:8.0.0
```

# using db
http://dimitrik.free.fr/mysqltechday_2013_10/world.sql より
```
CREATE DATABASE IF NOT EXISTS `test-xo-db`;
USE `test-xo-db`

DROP TABLE IF EXISTS `City`;
CREATE TABLE `City` (
  `ID` int(11) NOT NULL auto_increment,
  `Name` char(35) NOT NULL default '',
  `CountryCode` char(3) NOT NULL default '',
  `District` char(20) NOT NULL default '',
  `Population` int(11) NOT NULL default '0',
  PRIMARY KEY  (`ID`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;
```

# using xo
```
# xo 'mysql://root:my-pw@0.0.0.0/test-xo-db'
# ls
city.xo.go  README.md  xo_db.xo.go
```
