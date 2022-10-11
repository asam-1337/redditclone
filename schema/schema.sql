CREATE TABLE IF NOT EXISTS users
(
    id serial not null unique,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE IF NOT EXISTS posts
(
    id serial not null unique primary key ,
    post_type varchar(255) not null,
    title text not null,
    post_category varchar(255) not null,
    post_text text,
    url varchar(255),
    score int not null default 1,
    views int not null default 1,
    upvote_percentage int not null default 100,
    user_id int references users (id) on delete cascade not null,
    created varchar(255) not null
);

CREATE TABLE IF NOT EXISTS comments
(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    post_id int references posts(id) on delete cascade not null,
    body varchar(255) not null,
    created varchar(255) not null
);

CREATE TABLE IF NOT EXISTS votes
(
    user_id int references users(id) on delete cascade not null,
    post_id int references posts(id) on delete cascade not null,
    vote int not null default 1,
    primary key (user_id, post_id)
);
