-- +goose Up
-- SQL in this section is executed when the migration is applied.

begin;

create sequence survey_api_survey_id_seq
start with 1
increment by 1
no minvalue
no maxvalue
cache 1;

create sequence survey_api_question_id_seq
start with 1
increment by 1
no minvalue
no maxvalue
cache 1;

create sequence survey_api_answer_id_seq
start with 1
increment by 1
no minvalue
no maxvalue
cache 1;

create table survey_api_survey (
id int primary key default nextval('survey_api_survey_id_seq'),
description text not null,
created date default now()
);

create type survey_api_question_type as enum ('freeform');

create table survey_api_question (
id int primary key default nextval('survey_api_question_id_seq'),
survey_id int not null references survey_api_survey(id),
prompt text not null,
question_type survey_api_question_type default 'freeform',
created date default now()
);

create table survey_api_answer (
id int primary key default nextval('survey_api_answer_id_seq'),
question_id int not null references survey_api_question(id),
value text not null,
user_info varchar(255),
created date default now()
);

end;


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

begin;

drop sequence survey_api_survey_id_seq;
drop sequence survey_api_question_id_seq;
drop sequence survey_api_answer_id_seq;
drop table survey_api_survey;
drop type survey_api_question_type;
drop table survey_api_question;
drop table survey_api_answer;

end;
