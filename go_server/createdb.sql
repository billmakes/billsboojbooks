-- SQL script for creating the database tables in SQLite.
-- Use it by piping into sqlite3 with a new DB name:
--
--  $ sqlite3 somedb.db < createdb.sql
--
-- Eli Bendersky [http://eli.thegreenplace.net]
-- This code is in the public domain.


-- Foreign key constraints have to be enabled explicitly in SQLite.
PRAGMA foreign_keys = ON;

create table Book (
    bookID integer primary key,
    title text,
    author text,
    year text,
    read integer,
    readDate date
);

create table Tag (
    tagID integer primary key,
    name text unique
);

-- Linking table for the many-to-many relationship between Tag and Book
create table BookTag (
    bookID integer,
    tagID integer,

    foreign key(bookID) references Book(bookID),
    foreign key(tagID) references Tag(tagID)
);
