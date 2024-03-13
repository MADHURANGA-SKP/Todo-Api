CREATE TABLE "users" (
  "user_id" bigserial PRIMARY KEY,
  "first_name" varchar ,
  "last_name" varchar ,
  "user_name" varchar ,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL
);

CREATE TABLE "todos" (
  "todo_id" bigserial PRIMARY KEY,
  "user_id" int NOT NULL,
  "title" varchar ,
  "time" varchar ,
  "date" varchar ,
  "completed" varchar 
);

ALTER TABLE "todos" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
