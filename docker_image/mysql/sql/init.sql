CREATE DATABASE agreast-db DEFAULT CHARACTER SET 'utf8mb4';
USE agreast-db;
CREATE TABLE agrist (
  id  INT AUTO_INCREMENT NOT NULL,
  name char(30),
  possession_id int,
  PRIMARY KEY (id)
);

INSERT INTO agrist (name,possession_id) VALUES ("yamadaryunosuke",1);

CREATE TABLE farmland(
  id  INT AUTO_INCREMENT NOT NULL,
  name char(30),
  PRIMARY KEY (id)
);

INSERT INTO farmland (name) VALUES ("農耕地A");
INSERT INTO farmland (name) VALUES ("農耕地B");

CREATE TABLE possession(
  id int AUTO_INCREMENT NOT NULL,
  agrist_id int,
  farmland_id int,
  PRIMARY KEY (id,agrist_id,farmland_id)
);

INSERT INTO possession (agrist_id,farmland_id) VALUES (1,1);
INSERT INTO possession (agrist_id,farmland_id)VALUES (1,2);

CREATE TABLE temphum(
  id  INT AUTO_INCREMENT NOT NULL,
  farmland_id int,
  day date,
  temp float,
  humi float,
  PRIMARY KEY (id)
);

INSERT INTO temphum(farmland_id,day,temp,humi) VALUES (1,'2023/01/17',20.3,40);
INSERT INTO temphum(farmland_id,day,temp,humi) VALUES (2,'2023/01/17',25,72);

CREATE TABLE crops(
  id INT AUTO_INCREMENT NOT NULL,
  name char(30),
  PRIMARY KEY (id)
);

INSERT INTO crops (name) VALUES ("ピーマン");

CREATE TABLE planting(
  id  int AUTO_INCREMENT NOT NULL,
  farmland_id int,
  crops_id int,
  startday date,
  PRIMARY KEY(id)
);
INSERT INTO planting (farmland_id,startday,crops_id) VALUES ( 1,'2022/12/28',1 );
INSERT INTO planting (farmland_id,startday,crops_id) VALUES ( 2,'2022/12/27',1 );

CREATE TABLE machine(
  id  INT AUTO_INCREMENT NOT NULL,
  machine_num char(10),
  PRIMARY KEY(id)
);
INSERT INTO machine (machine_num)VALUES("123455678");

CREATE TABLE photo(
  id int AUTO_INCREMENT NOT NULL,
  num int,
  url char(100),
  PRIMARY KEY (id)
);
INSERT INTO photo (identifier,num,url) VALUES (1,9,"****//");

CREATE TABLE harvest(
  id int AUTO_INCREMENT NOT NULL,
  machine_id int,
  amount int,
  photo_id int,
  day date,
  PRIMARY KEY (id,machine_id,photo_id),
  FOREIGN KEY(machine_id) REFERENCES machine(id),
  FOREIGN KEY(photo_id) REFERENCES photo(id)
);
INSERT INTO harvest (machine_id,photo_id,amount,day) VALUES (1,1,10,"2023/01/17");

CREATE TABLE unharvest(
  id int AUTO_INCREMENT NOT NULL,
  machine_id int,
  amount int,
  photo_id int,
  day date,
  PRIMARY KEY (id,machine_id,photo_id),
  FOREIGN KEY(machine_id) REFERENCES machine(id),
  FOREIGN KEY(photo_id) REFERENCES photo(id)
);
INSERT INTO harvest (machine_id,photo_id,amount,day) VALUES (1,1,10,"2023/01/17");

