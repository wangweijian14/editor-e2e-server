-- SQLite
select * from sqlite_master where type = "table";

ALTER TABLE page ADD COLUMN url TEXT;


ALTER TABLE case_step ADD COLUMN case_id integet;

ALTER TABLE case_step ADD COLUMN score integet;

ALTER TABLE cases ADD COLUMN description TEXT;


ALTER TABLE cases ADD COLUMN exec_count INTEGER;

ALTER TABLE cases ADD COLUMN pass_count INTEGER;

ALTER TABLE cases ADD COLUMN skip_count INTEGER;

ALTER TABLE cases ADD COLUMN failed_count INTEGER;

ALTER TABLE cases ADD COLUMN his_id INTEGER;

ALTER TABLE cases DROP COLUMN cs_id;

CREATE TABLE "cases" (
 id INTEGER constraint cases_pk primary key autoincrement, 
 exec_count int, 
 pass_count int,
 skip_count INTEGER,
 failed_count INTEGER,
 his_id TEXT,
 description TEXT 
 );

DROP TABLE cases;