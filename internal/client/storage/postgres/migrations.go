package postgres

const createUsersTableQuery = `
CREATE TABLE IF NOT EXISTS  users
(
    id       uuid not null constraint users_pk primary key,
    name     varchar,
	access_token varchar,
	refresh_token varchar
);
`
const createCardsTableQuery = `
CREATE TABLE IF NOT EXISTS cards
(
    id uuid not null constraint cards_pk primary key,
    name varchar,
    card_holder_name varchar,
    number varchar,
    bank varchar,
    exp_month varchar,
    exp_year varchar,
    security_code varchar,
    meta jsonb
);
`
const createCredsTableQuery = `
CREATE TABLE IF NOT EXISTS creds
(
    id uuid not null constraint creds_pk primary key,
    name varchar,
    login varchar,
    password varchar,
    Meta jsonb
);
`
const createNotesTableQuery = `
CREATE TABLE IF NOT EXISTS notes
(
    id uuid not null constraint notes_pk primary key,
    name varchar,
    text varchar,
    Meta jsonb
);
`
const createFilesTableQuery = `
CREATE TABLE IF NOT EXISTS files
(
    id uuid not null constraint files_pk primary key,
    name varchar,
    fileName varchar,
    Meta jsonb
);
`
