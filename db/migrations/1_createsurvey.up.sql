begin;


create sequence survey_id_seq
start with 1
increment by 1
no minvalue
no maxvalue
cache 1;

create sequence question_id_seq
start with 1
increment by 1
no minvalue
no maxvalue
cache 1;

create sequence answer_id_seq
start with 1
increment by 1
no minvalue
no maxvalue
cache 1;

create table survey (
id int primary key default nextval('survey_id_seq'),
desc text not null,
created date default now()
);

create type question_type as enum ('freeform');

create table question (
id int primary key default nextval('question_id_seq'),
survey_id int not null references survey(id),
prompt text not null,
type question_type default 'freeform',
created date default now()
);

create table answer (
id int primary key default nextval('answer_id_seq'),
question_id int not null references question(id),
value text not null,
user_info varchar(255),
created date default now()
);

end;
