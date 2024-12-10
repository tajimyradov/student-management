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
    group_id   int references groups (id),
    type_id    int references types (id),
    unique (lesson_id, teacher_id, student_id, group_id, type_id)
);

create table types
(
    id   serial primary key,
    name varchar(128) not null
);

create table timetables
(
    id serial primary key ,
    weekday         int not null,
    group_id        integer references groups (id),
    lesson_id       integer references lessons (id),
    time_id         integer references times (id),
    auditory_id     integer references auditories (id),

    alt_lesson_id   integer references lessons (id),
    alt_auditory_id integer references auditories (id),
    type_id         int references types (id),
    unique (weekday, group_id, lesson_id, time_id, auditory_id, alt_auditory_id, alt_lesson_id, type_id)
);

create table absences
(
    id         serial primary key,
    student_id int references students (id),
    group_id   integer references groups (id),
    lesson_id  integer references lessons (id),
    time_id    integer references times (id),
    teacher_id int references teachers (id),
    type_id    int references types (id),
    date       date default now(),
    status int default 0,
    is_absent bool
);

CREATE MATERIALIZED VIEW timetable_view AS
select t.weekday,
       t.group_id,
       g.name as group_name,
       t.lesson_id,
       l.name as lesson_name,
       t.time_id,
       ti.start_time as start_time,
       ti.end_time as end_time,
       t.auditory_id,
       a.name as auditory_name,
       t.alt_lesson_id,
       ll.name as alt_lesson_name,
       t.alt_auditory_id,
       aa.name as alt_auditory_name,
       t.type_id,
       ty.name as type_name
from timetables as t
         join groups as g on g.id = t.group_id
         join lessons as l on l.id = t.lesson_id
         join times as ti on ti.id = t.time_id
         join types as ty on ty.id = t.type_id
         left join auditories as a on a.id = t.auditory_id
         left join lessons as ll on ll.id = t.alt_lesson_id
         left join auditories as aa on aa.id = t.alt_auditory_id
order by t.weekday, t.time_id;