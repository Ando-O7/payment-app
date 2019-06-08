DROP TABLE IF EXISTS items;

CREATE TABLE item (
    id integer AUTO_INCREMENT,
    name varchar(255),
    description varchar(255),
    amount integer,
    primary key(id)
);

INSERT INTO items(
    name,
    description,
    amount
)
VALUES(
    'toy',
    'test toy',
    3000
);

INSERT INTO item(
    name,
    description,
    amount
)
VALUES(
    'game',
    'test game',
    7000
);