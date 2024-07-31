create database todo_list;

drop table if exists task;
create table if not exists task(
    id serial auto_increment primary key,
    title text,
    sub_title text,
    due_date timestamp,
    completed_at timestamp,
    status varchar(128),

    created_at timestamp default current_timestamp,
    updated_at timestamp on update current_timestamp
);
