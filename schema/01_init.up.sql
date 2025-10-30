CREATE TABLE students
(
    id serial not null unique,
    name varchar(255) not null,
    surname varchar(255) not null unique,
    group_name varchar(64) not null
);

CREATE TABLE certificates
(
    id serial not null unique,
    filename_hash varchar(255) not null unique
);

CREATE TABLE students_certificates
(
    id serial not null unique,
    student_id int references students(id) on delete cascade not null,
    certificates_id int references certificates(id) on delete cascade not null
);