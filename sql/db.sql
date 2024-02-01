--
-- PostgreSQL database dump
--

-- Dumped from database version 16.1 (Debian 16.1-1.pgdg120+1)
-- Dumped by pg_dump version 16.1 (Debian 16.1-1.pgdg120+1)

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
-- Name: account; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA account;


ALTER SCHEMA account OWNER TO postgres;

--
-- Name: transaction; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA transaction;


ALTER SCHEMA transaction OWNER TO postgres;

--
-- Name: wallet; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA wallet;


ALTER SCHEMA wallet OWNER TO postgres;

--
-- Name: add_transaction(bigint, double precision, character varying, bigint); Type: FUNCTION; Schema: transaction; Owner: postgres
--

CREATE FUNCTION transaction.add_transaction(ptransaction_source_wallet_id bigint, ptransaction_amount double precision, ptransaction_currency character varying, ptransaction_target_wallet_id bigint) RETURNS bigint
    LANGUAGE plpgsql
    AS $$
declare
    vTransaction_id bigint;

begin
        vTransaction_id = nextval('id_seq'::regclass);
insert into "transaction".list (transaction_id, transaction_source_wallet_id, transaction_amount, transaction_currency, transaction_target_wallet_id, transaction_status) values (vTransaction_id, ptransaction_source_wallet_id,ptransaction_amount, ptransaction_currency, ptransaction_target_wallet_id,'created');

    return vTransaction_id;
end
$$;


ALTER FUNCTION transaction.add_transaction(ptransaction_source_wallet_id bigint, ptransaction_amount double precision, ptransaction_currency character varying, ptransaction_target_wallet_id bigint) OWNER TO postgres;

--
-- Name: id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.id_seq OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: list; Type: TABLE; Schema: account; Owner: postgres
--

CREATE TABLE account.list (
    account_id bigint DEFAULT nextval('public.id_seq'::regclass) NOT NULL,
    account_number bigint NOT NULL,
    account_currency character varying NOT NULL,
    account_amount double precision DEFAULT 0 NOT NULL,
    account_wallet_id bigint NOT NULL,
    lock smallint DEFAULT 0 NOT NULL
);


ALTER TABLE account.list OWNER TO postgres;

--
-- Name: list; Type: TABLE; Schema: transaction; Owner: postgres
--

CREATE TABLE transaction.list (
    transaction_id bigint DEFAULT nextval('public.id_seq'::regclass) NOT NULL,
    transaction_source_wallet_id bigint NOT NULL,
    transaction_amount double precision NOT NULL,
    transaction_currency character varying NOT NULL,
    transaction_target_wallet_id bigint NOT NULL,
    transaction_status character varying NOT NULL
);


ALTER TABLE transaction.list OWNER TO postgres;

--
-- Name: list; Type: TABLE; Schema: wallet; Owner: postgres
--

CREATE TABLE wallet.list (
    wallet_id bigint DEFAULT nextval('public.id_seq'::regclass) NOT NULL,
    wallet_number bigint NOT NULL
);


ALTER TABLE wallet.list OWNER TO postgres;

--
-- Name: list list_pkey; Type: CONSTRAINT; Schema: account; Owner: postgres
--

ALTER TABLE ONLY account.list
    ADD CONSTRAINT list_pkey PRIMARY KEY (account_id);


--
-- Name: list list_pkey; Type: CONSTRAINT; Schema: transaction; Owner: postgres
--

ALTER TABLE ONLY transaction.list
    ADD CONSTRAINT list_pkey PRIMARY KEY (transaction_id);


--
-- Name: list list_pkey; Type: CONSTRAINT; Schema: wallet; Owner: postgres
--

ALTER TABLE ONLY wallet.list
    ADD CONSTRAINT list_pkey PRIMARY KEY (wallet_id);


--
-- Name: list fk_wallet; Type: FK CONSTRAINT; Schema: account; Owner: postgres
--

ALTER TABLE ONLY account.list
    ADD CONSTRAINT fk_wallet FOREIGN KEY (account_wallet_id) REFERENCES wallet.list(wallet_id);


--
-- Name: list fk_wallet; Type: FK CONSTRAINT; Schema: transaction; Owner: postgres
--

ALTER TABLE ONLY transaction.list
    ADD CONSTRAINT fk_wallet FOREIGN KEY (transaction_source_wallet_id) REFERENCES wallet.list(wallet_id);


--
-- PostgreSQL database dump complete
--