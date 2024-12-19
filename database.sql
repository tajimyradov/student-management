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
    username   varchar(128) not null,
    password   varchar(128) not null,
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
    id              serial primary key,
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
    status     int  default 0,
    note       text
);

create MATERIALIZED VIEW absences_view as
select a.id,

       d.id          as department_id,
       d.name        as department_name,
       f.id          as faculty_id,
       f.name        as faculty_name,


       g.id          as group_id,
       g.name        as group_name,
       l.id          as lesson_id,

       l.name        as lesson_name,
       t.id          as time_id,
       t.start_time,
       t.end_time,
       t2.id         as teacher_id,
       t2.first_name as teacher_first_name,
       t2.last_name  as teacher_last_name,
       s.id          as student_id,
       s.first_name  as student_first_name,
       s.last_name   as student_last_name,
       g.year        as student_year,
       t4.first_name as faculty_dean_first_name,
       t4.last_name  as faculty_dean_last_name,
       t5.first_name as department_lead_first_name,
       t5.last_name  as department_lead_last_name,
       t3.id         as type_id,
       t3.name       as type_name,
       a.date,
       a.status,
       a.note

from absences as a
         join groups g on a.group_id = g.id
         join lessons l on l.id = a.lesson_id
         join students s on s.id = a.student_id
         join times t on a.time_id = t.id
         join teachers t2 on a.teacher_id = t2.id
         join types t3 on a.type_id = t3.id
         join professions p on g.profession_id = p.id
         join departments d on p.department_id = d.id
         join faculties f on d.faculty_id = f.id
         join deans d2 on f.id = d2.faculty_id
         join teachers as t4 on d2.teacher_id = t4.id
         join department_leads dl on d.id = dl.department_id
         join teachers as t5 on dl.teacher_id = t5.id
order by a.date;


CREATE MATERIALIZED VIEW timetable_view AS
select t.weekday,
       t.group_id,
       g.name        as group_name,
       t.lesson_id,
       l.name        as lesson_name,
       t.time_id,
       ti.start_time as start_time,
       ti.end_time   as end_time,
       t.auditory_id,
       a.name        as auditory_name,
       t.alt_lesson_id,
       ll.name       as alt_lesson_name,
       t.alt_auditory_id,
       aa.name       as alt_auditory_name,
       t.type_id,
       ty.name       as type_name
from timetables as t
         join groups as g on g.id = t.group_id
         join lessons as l on l.id = t.lesson_id
         join times as ti on ti.id = t.time_id
         join types as ty on ty.id = t.type_id
         left join auditories as a on a.id = t.auditory_id
         left join lessons as ll on ll.id = t.alt_lesson_id
         left join auditories as aa on aa.id = t.alt_auditory_id
order by t.weekday, t.time_id;


create table deans
(
    faculty_id int references faculties (id),
    teacher_id int references teachers (id)
);

create table department_leads
(
    department_id int references departments (id),
    teacher_id    int references teachers (id)
);