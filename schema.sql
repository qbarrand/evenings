create table topic (
    id integer primary key,
    name text,
    short_name text unique,

    created_at date,
    updated_at date,
    deleted_at date
);

create table session (
    id integer primary key not null,
    start date,
    end date,
    topic_id int references topic,

    created_at date,
    updated_at date,
    deleted_at date
);

insert into topic(name, short_name) values("Tests psychotechniques", "PSY");
insert into topic(name, short_name) values("Golang", "GO");
insert into topic(name, short_name) values("Java", "JAVA");
insert into topic(name, short_name) values("Algorithmique", "ALGO");