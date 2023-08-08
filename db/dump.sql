--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2 (Debian 14.2-1.pgdg110+1)
-- Dumped by pg_dump version 14.0



SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: gcalendar; Type: DATABASE; Schema: -; Owner: user
--

CREATE DATABASE gcalendar WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


ALTER DATABASE gcalendar OWNER TO "user";

\connect gcalendar

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: gcalendar; Type: SCHEMA; Schema: -; Owner: user
--

CREATE SCHEMA gcalendar;


ALTER SCHEMA gcalendar OWNER TO "user";

SET default_tablespace = '';

SET default_table_access_method = heap;


CREATE EXTENSION hstore WITH SCHEMA "gcalendar";


DROP TABLE IF EXISTS "notification_config";
CREATE TABLE "gcalendar"."notification_config" (
    "data" bytea NOT NULL,
    "calendar_uuid" character (128) NOT NULL,
    "channel_uuid" character(128) NOT NULL,
    "user_uuid" character (36) NOT NULL,

CONSTRAINT "notification_config_calendar_uuid" PRIMARY KEY ("calendar_uuid")
) WITH (oids = false);

CREATE UNIQUE INDEX "notification_config_channel_uuid" ON "gcalendar"."notification_config" USING btree ("channel_uuid");
CREATE INDEX "notification_config_user_uuid" ON "gcalendar"."notification_config" USING btree ("user_uuid");


DROP TABLE IF EXISTS "user_channel";
CREATE TABLE "gcalendar"."user_channel" (
    "data" bytea NOT NULL,
    "user_uuid" character varying(256) NOT NULL,
    "calendar_uuid" character varying(256) NOT NULL,
    "channel_uuid" character varying(256) NOT NULL,
    CONSTRAINT "user_channel_calendar_uuid" PRIMARY KEY ("calendar_uuid")
) WITH (oids = false);

CREATE INDEX "user_channel_user_uuid" ON "gcalendar"."user_channel" USING btree ("user_uuid");
CREATE INDEX "user_channel_channel_uuid" ON "gcalendar"."user_channel" USING btree ("channel_uuid");



DROP TABLE IF EXISTS "user_token";
CREATE TABLE "gcalendar"."user_token" (
    "data" bytea NOT NULL,
    "user_resource_uuid" character (128) NOT NULL,
    "user_uuid" character (36) NOT NULL,
    CONSTRAINT "user_token_user_uuid" PRIMARY KEY ("user_uuid")
) WITH (oids = false);

CREATE UNIQUE INDEX "user_token_user_resource_uuid" ON "gcalendar"."user_token" USING btree ("user_resource_uuid");


ALTER TABLE ONLY "gcalendar"."notification_config" ADD CONSTRAINT "notification_config_calendar_uuid_fkey" FOREIGN KEY (calendar_uuid) REFERENCES "gcalendar"."user_channel"(calendar_uuid) NOT DEFERRABLE;

ALTER TABLE ONLY "gcalendar"."user_channel" ADD CONSTRAINT "user_channel_user_uuid_fkey" FOREIGN KEY (user_uuid) REFERENCES "gcalendar"."user_token"(user_uuid) NOT DEFERRABLE;

-- 2022-04-05 20:00:29.193965+00

