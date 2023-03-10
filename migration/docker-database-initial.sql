create table customers(
   id serial primary key,
   name varchar,
   cpf varchar
);

INSERT INTO customers(name, cpf) VALUES
('johnny silva', '474.390.678-19'),
('patricia braga', '148.888.388-48');