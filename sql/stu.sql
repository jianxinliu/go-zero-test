create table `student`
(
    `id`   int(10)     not null auto_increment,
    `name` varchar(20) not null,
    `age`  int,
    `city` varchar(100),
    primary key (id)
) engine = innoDB comment = 'student'