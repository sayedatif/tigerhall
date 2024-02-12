CREATE USER 'tigerhall'@'%' IDENTIFIED BY 'secret';
GRANT ALL PRIVILEGES ON tigerhall.* TO 'tigerhall'@'%';
FLUSH PRIVILEGES;