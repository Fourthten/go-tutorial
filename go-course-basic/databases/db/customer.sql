CREATE TABLE IF NOT EXISTS `customer` (
  `id` varchar(100) NOT NULL,
  `name` varchar(100) NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DELETE FROM `customer`;

ALTER TABLE `customer`
  ADD COLUMN email varchar(100),
  ADD COLUMN balance int DEFAULT 0,
  ADD COLUMN rating DOUBLE DEFAULT 0.0,
  ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  ADD COLUMN birth_date DATE,
  ADD COLUMN married BOOLEAN DEFAULT false;

DESC `customer`;

INSERT INTO `customer` (`id`, `name`, `email`, `balance`, `rating`, `birth_date`, `married`)
VALUES ('setia', 'Setia', 'setia@mail.com', 10000, 80.5, '1999-01-01', true),
      ('ajung', 'Ajung', 'ajung@mail.com', 100000, 90.0, '1998-01-01', true),
      ('tia', 'Tia', 'tia@mail.com', 50000, 75.0, '1998-02-01', true);

CREATE TABLE IF NOT EXISTS `user` (
  `username` VARCHAR(100) NOT NULL,
  `password` VARCHAR(100) NOT NULL,
  PRIMARY KEY(username)
) ENGINE=InnoDB;

SELECT * FROM `user`;

INSERT INTO `user` (`username`, `password`) VALUES ('admin', 'admin');

CREATE TABLE `comments` (
  id INT NOT NULL AUTO_INCREMENT,
  email VARCHAR(100) NOT NULL,
  comment TEXT,
  PRIMARY KEY (id)
) ENGINE=InnoDB;

DESC `comments`;

SELECT * FROM `comments`;