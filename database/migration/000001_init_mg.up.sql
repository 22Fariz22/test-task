CREATE table persons
(
  user_id UUID PRIMARY KEY,
  first_name  VARCHAR(30) NOT NULL CHECK,
  surname VARCHAR(30) NOT NULL CHECK,
  patronymic VARCHAR(30),
  age INTEGER,
  gender VARCHAR(10),
  nationality VARCHAR(30)
);