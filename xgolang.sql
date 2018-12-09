use xgolang;

create table if not exists user (
    user_id int(11) primary key auto_increment,
    username varchar(150) not null,
    password varchar(150) not null,
    age int(11)
);