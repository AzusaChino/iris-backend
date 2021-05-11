create table article (
    id integer not null constraint article_pk primary key autoincrement,
    title text,
    remark text,
    pic text,
    content text
);