CREATE DATABASE agreastd DEFAULT CHARACTER SET 'utf8mb4';
USE test;
CREATE TABLE agrists (
  id  INT AUTO_INCREMENT NOT NULL,
  username char(30),
  PRIMARY KEY (id)
);

INSERT INTO agrists (username) VALUES ("yamadaryunosuke");
INSERT INTO agrists (username) VALUES ("yamadaryunosuke1");
INSERT INTO agrists (username) VALUES ("test");

CREATE TABLE farmlands(
  id  INT AUTO_INCREMENT NOT NULL,
  farmlandname char(30),
  PRIMARY KEY (id)
);

INSERT INTO farmlands (farmlandname) VALUES ("農耕地A");
INSERT INTO farmlands (farmlandname) VALUES ("農耕地B");
INSERT INTO farmlands (farmlandname) VALUES ("農耕地C");

CREATE TABLE possessions(
  id int AUTO_INCREMENT NOT NULL,
  agrist_id int,
  farmland_id int,
  PRIMARY KEY (id,agrist_id,farmland_id),
  FOREIGN KEY(agrist_id) REFERENCES agrists(id),
  FOREIGN KEY(farmland_id) REFERENCES farmlands(id)
);

INSERT INTO possessions (agrist_id,farmland_id) VALUES (1,1);
INSERT INTO possessions (agrist_id,farmland_id)VALUES (1,2);

CREATE TABLE temphums(
  id  INT AUTO_INCREMENT NOT NULL,
  farmland_id int,
  day date,
  temp float,
  humi float,
  PRIMARY KEY (id)
);

INSERT INTO temphums(farmland_id,day,temp,humi) VALUES (1,'2023/01/17',20.3,40);
INSERT INTO temphums(farmland_id,day,temp,humi) VALUES (2,'2023/01/17',25,72);

CREATE TABLE crops(
  id INT AUTO_INCREMENT NOT NULL,
  cropsname char(30),
  PRIMARY KEY (id)
);

INSERT INTO crops (cropsname) VALUES ("ピーマン");

CREATE TABLE plantings(
  id  int AUTO_INCREMENT NOT NULL,
  farmland_id int,
  crops_id int,
  startday date,
  PRIMARY KEY(id,farmland_id,crops_id),
  FOREIGN KEY(crops_id) REFERENCES crops(id),
  FOREIGN KEY(farmland_id) REFERENCES farmlands(id)
);
INSERT INTO plantings (farmland_id,startday,crops_id) VALUES ( 1,'2022/12/28',1 );
INSERT INTO plantings (farmland_id,startday,crops_id) VALUES ( 2,'2022/12/27',1 );

CREATE TABLE machines(
  id  INT AUTO_INCREMENT NOT NULL,
  machine_num char(10),
  farmland_id int,
  PRIMARY KEY(id),
  FOREIGN KEY(farmland_id)REFERENCES farmlands(id)
);
INSERT INTO machines (machine_num,farmland_id)VALUES("123455678",1);
INSERT INTO machines (machine_num,farmland_id)VALUES("000000000",1);

CREATE TABLE photos(
  id int AUTO_INCREMENT NOT NULL,
  identifer int,
  amount int,
  url char(100),
  PRIMARY KEY (id)
);
INSERT INTO photos (identifer,amount,url) VALUES (1,1,"ピーマン.jpg");
INSERT INTO photos (identifer,amount,url) VALUES (1,2,"ピーマン2.jpg");
INSERT INTO photos (identifer,amount,url) VALUES (1,2,"ピーマン3.jpg");
INSERT INTO photos (identifer,amount,url) VALUES (2,1,"シシトウ-min.jpg");

CREATE TABLE harvests(
  id int AUTO_INCREMENT NOT NULL,
  machine_id int,
  amount int,
  photo_id int,
  day date,
  PRIMARY KEY (id,machine_id,photo_id),
  FOREIGN KEY(photo_id) REFERENCES photos(id)
);
INSERT INTO harvests (machine_id,photo_id,amount,day) VALUES (1,1,10,"2023/01/17");
INSERT INTO harvests (machine_id,photo_id,amount,day) VALUES (1,2,10,"2023/01/18");
INSERT INTO harvests (machine_id,photo_id,amount,day) VALUES (1,3,10,"2023/01/18");

CREATE TABLE unharvests(
  id int AUTO_INCREMENT NOT NULL,
  machine_id int,
  amount int,
  photo_id int,
  day date,
  PRIMARY KEY (id,machine_id,photo_id),
  FOREIGN KEY(machine_id) REFERENCES machines(id),
  FOREIGN KEY(photo_id) REFERENCES photos(id)
);
INSERT INTO unharvests (machine_id,photo_id,amount,day) VALUES (1,2,10,"2023/01/17");

