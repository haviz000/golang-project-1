-- Database: book_project_1

-- DROP DATABASE IF EXISTS book_project_1;

CREATE DATABASE book_project_1
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_Indonesia.1252'
    LC_CTYPE = 'English_Indonesia.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

-- Table: public.books

-- DROP TABLE IF EXISTS public.books;

CREATE TABLE IF NOT EXISTS public.books
(
    book_id bigint NOT NULL DEFAULT 'nextval('books_book_id_seq'::regclass)',
    book_name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    author character varying(50) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT books_pkey PRIMARY KEY (book_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.books
    OWNER to postgres;