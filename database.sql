create table faculties
(
    id   serial primary key,
    name varchar(128) not null,
    code varchar(128) default ''
);

create table departments
(
    id         serial primary key,
    name       varchar(128) not null,
    code       varchar(128) default '',
    faculty_id integer references faculties (id)
);

create table teachers
(
    id            serial primary key,
    first_name    varchar(128) not null,
    last_name     varchar(128) not null,
    code          varchar(128) default '',
    gender        boolean,
    username      varchar(128) not null,
    password      varchar(128) not null,
    image         varchar(128),
    department_id integer references departments (id)
);

create table professions
(
    id            serial primary key,
    name          varchar(128) not null,
    code          varchar(128) default '',
    department_id integer references departments (id)
);

create table groups
(
    id            serial primary key,
    name          varchar(128) not null,
    code          varchar(128) default '',
    year          smallint     not null,
    profession_id integer references professions (id)
);

create table students
(
    id         serial primary key,
    first_name varchar(128) not null,
    last_name  varchar(128) not null,
    code       varchar(128) default '',
    gender     boolean,
    birth_date DATE         not null,
    image      varchar(128),
    group_id   integer references groups (id)
);

create table times
(
    id         serial primary key,
    start_time time not null,
    end_time   time not null
);

create table lessons
(
    id   serial primary key,
    name varchar(128) not null,
    code varchar(128) default ''
);

create table auditories
(
    id   serial primary key,
    name varchar(128) not null
);

create table lesson_teacher_student_bindings
(
    lesson_id  int references lessons (id),
    teacher_id int references teachers (id),
    student_id int references students (id),
    group_id int references groups(id),
    unique (lesson_id, teacher_id, student_id,group_id)
);

create table timetables
(
    weekday     int not null,
    group_id    integer references groups (id),
    lesson_id   integer references lessons (id),
    time_id     integer references times (id),
    auditory_id integer references auditories (id),
    teacher_id integer references teachers(id),

    alt_lesson_id   integer references lessons (id),
    alt_auditory_id integer references auditories (id),
    alt_teacher_id integer references teachers(id),
    unique (weekday,group_id,lesson_id,time_id,auditory_id,alt_auditory_id,alt_lesson_id)
);

