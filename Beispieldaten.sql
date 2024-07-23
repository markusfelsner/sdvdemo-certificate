DROP DATABASE IF EXISTS sdvdemodatabase;

CREATE DATABASE sdvdemodatabase;
USE sdvdemodatabase;

DROP TABLE IF EXISTS sdvdemotabelle;

CREATE TABLE sdvdemotabelle (
    ID int,  
    Name varchar(255), 
    Age varchar(255) 
    );

INSERT INTO sdvdemotabelle (ID,Name,Age)
VALUES
  ("1","Charles Gray","44"),
  ("2","Chava Porter","55"),
  ("3","Tyrone Ramsey","66"),
  ("4","Wylie Dodson","77"),
  ("5","Flynn Stuart","88");

  ALTER TABLE
    `sdvdemotabelle` ADD PRIMARY KEY(ID);

ALTER TABLE
   sdvdemotabelle MODIFY ID BIGINT(20) NOT NULL AUTO_INCREMENT,
    AUTO_INCREMENT = 4;