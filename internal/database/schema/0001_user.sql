-- +goose up
-- +goose statementbegin
create type staff_status as enum (
	'Active',
	'On Assignment',
	'Leave of Absence',
	'Inactive'
);
create type org_type as enum (
	'Division',
	'Bureau',
	'Directorate',
	'Branch'
);
-- recursively organized hierarchical list
create table if not exists organizations (
	id char(40) primary key default concat('org_', gen_random_uuid()),
	name text not null,
	type org_type not null,	
	parent_id char(40), -- self references organization(id), 
	unique(name, parent_id)
);

create table if not exists users (
	id char(41) primary key default concat('user_', gen_random_uuid()),
	uid text not null,
	name text not null,
	job_title text not null,
	office text not null,
	organizations text not null references organizations(id),
	status staff_status not null,
	telephone text,
	unique(uid)
);
-- +goose statementend
