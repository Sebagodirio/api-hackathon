INSERT INTO hackathons (country,state,hackathon_year) VALUES ('United States','Chicago','2022');
INSERT INTO hackathons (country,state,hackathon_year) VALUES ('United States','Miami','2022');
INSERT INTO hackathons (country,state,hackathon_year) VALUES ('United States','New York','2022');

INSERT INTO developments (name,description,score, hackathon_id) VALUES ('test_name1','lorem_ipsum',200,1);
INSERT INTO developments (name,description,score, hackathon_id) VALUES ('test_name2','lorem_ipsum',150,1);
INSERT INTO developments (name,description,score, hackathon_id) VALUES ('test_name3','lorem_ipsum',230,1);
INSERT INTO developments (name,description,score, hackathon_id) VALUES ('test_name4','lorem_ipsum',18,1);
INSERT INTO developments (name,description,score, hackathon_id) VALUES ('test_name5','lorem_ipsum',25,1);
INSERT INTO developments (name,description,score, hackathon_id) VALUES ('test_name6','lorem_ipsum',125,1);
INSERT INTO developments (name,description,score, hackathon_id) VALUES ('test_name7','lorem_ipsum',340,1);
INSERT INTO developments (name,description,score, hackathon_id) VALUES ('test_name8','lorem_ipsum',200,1);
INSERT INTO developments (name,description,score, hackathon_id) VALUES ('test_name9','lorem_ipsum',220,1);

INSERT INTO developments (name,description,score, hackathon_id) VALUES ('dummy_name1','lorem_ipsum',200,2);
INSERT INTO developments (name,description,score, hackathon_id) VALUES ('dummy_name2','lorem_ipsum',150,2);
INSERT INTO developments (name,description,score, hackathon_id) VALUES ('dummy_name3','lorem_ipsum',230,2);
INSERT INTO developments (name,description,score, hackathon_id) VALUES ('dummy_name4','lorem_ipsum',18,2);
INSERT INTO developments (name,description,score, hackathon_id) VALUES ('dummy_name5','lorem_ipsum',25,2);
INSERT INTO developments (name,description,score, hackathon_id) VALUES ('dummy_name6','lorem_ipsum',125,2);
INSERT INTO developments (name,description,score, hackathon_id) VALUES ('dummy_name7','lorem_ipsum',340,2);
INSERT INTO developments (name,description,score, hackathon_id) VALUES ('dummy_name8','lorem_ipsum',200,2);
INSERT INTO developments (name,description,score, hackathon_id) VALUES ('dummy_name9','lorem_ipsum',220,2);

INSERT INTO developers (first_name,last_name,email,development_id) VALUES('Jhon','Smith','jhonsmith@gmail.com',1);
INSERT INTO developers (first_name,last_name,email,development_id) VALUES('Jhon','William','jhonsmith@gmail.com',1);
INSERT INTO developers (first_name,last_name,email,development_id) VALUES('Karen','Gomez','jhonsmith@gmail.com',1);
INSERT INTO developers (first_name,last_name,email,development_id) VALUES('Jean','Deep','jhonsmith@gmail.com',2);
INSERT INTO developers (first_name,last_name,email,development_id) VALUES('Jhon','Shal','jhonsmith@gmail.com',2);
INSERT INTO developers (first_name,last_name,email,development_id) VALUES('Jenny','Sean','jhonsmith@gmail.com',2);
INSERT INTO developers (first_name,last_name,email,development_id) VALUES('Jhon','Cork','jhonsmith@gmail.com',3);
INSERT INTO developers (first_name,last_name,email,development_id) VALUES('Marta','Alisson','jhonsmith@gmail.com',3);
INSERT INTO developers (first_name,last_name,email,development_id) VALUES('Jhon','Smith','jhonsmith@gmail.com',3);
INSERT INTO developers (first_name,last_name,email,development_id) VALUES('Abigail','Bored','jhonsmith@gmail.com',4);
INSERT INTO developers (first_name,last_name,email,development_id) VALUES('Jhon','Thise','jhonsmith@gmail.com',4);
INSERT INTO developers (first_name,last_name,email,development_id) VALUES('Loren','Godirio','jhonsmith@gmail.com',4);
INSERT INTO developers (first_name,last_name,email,development_id) VALUES('Jhon','Jhonsson','jhonsmith@gmail.com',10);
INSERT INTO developers (first_name,last_name,email,development_id) VALUES('Jhon','Clark','jhonsmith@gmail.com',10);
INSERT INTO developers (first_name,last_name,email,development_id) VALUES('Nicolas','William','jhonsmith@gmail.com',10);
INSERT INTO developers (first_name,last_name,email,development_id) VALUES('Jhon','Clark','jhonsmith@gmail.com',11);
INSERT INTO developers (first_name,last_name,email,development_id) VALUES('Camile','Sean','jhonsmith@gmail.com',11);

INSERT INTO users (email,password) VALUES ('admin@admin.com','admin');