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

CREATE TABLE "casereport" (
    id TEXT constraint casereport_pk primary key , 
    message TEXT,
    case_id int,
    case_desc TEXT,
    run_time datetime,
    status int,
    run_user TEXT,
    run_ip TEXT
);

CREATE TABLE "suitereport" (
    id INTEGER constraint casereport_pk primary key autoincrement, 
)


CREATE TABLE step (
  id INTEGER constraint step_pk primary key,
  description TEXT, 
  target_element_id INTEGER not null, 
  action_id INTEGER not null ,
  target_page int
 )


DELETE  FROM case_step;
DELETE  FROM cases;
DELETE FROM sqlite_sequence WHERE name = 'case_step';
DELETE FROM sqlite_sequence WHERE name = 'cases';



DELETE  FROM casereport;
DELETE FROM sqlite_sequence WHERE name = 'casereport';

# DROP TABLE casereport;


ALTER TABLE cases ADD COLUMN openId integer;