--
-- PostgreSQL database dump
--

-- Dumped from database version 13.3
-- Dumped by pg_dump version 13.3

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: users; Type: TABLE; Schema: public; Owner: HaNzO
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(50),
    phone_number character varying(10),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    verified boolean,
    logged_in boolean,
    deleted_at timestamp without time zone
);


ALTER TABLE public.users OWNER TO "HaNzO";

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: HaNzO
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO "HaNzO";

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: HaNzO
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: verification; Type: TABLE; Schema: public; Owner: HaNzO
--

CREATE TABLE public.verification (
    id integer NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    otp_verified boolean,
    user_id integer,
    otp character varying(6)
);


ALTER TABLE public.verification OWNER TO "HaNzO";

--
-- Name: verification_id_seq; Type: SEQUENCE; Schema: public; Owner: HaNzO
--

CREATE SEQUENCE public.verification_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.verification_id_seq OWNER TO "HaNzO";

--
-- Name: verification_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: HaNzO
--

ALTER SEQUENCE public.verification_id_seq OWNED BY public.verification.id;


--
-- Name: verifications; Type: TABLE; Schema: public; Owner: HaNzO
--

CREATE TABLE public.verifications (
    id integer NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    otp_verified boolean,
    user_id integer,
    otp character varying(6)
);


ALTER TABLE public.verifications OWNER TO "HaNzO";

--
-- Name: verifications_id_seq; Type: SEQUENCE; Schema: public; Owner: HaNzO
--

CREATE SEQUENCE public.verifications_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.verifications_id_seq OWNER TO "HaNzO";

--
-- Name: verifications_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: HaNzO
--

ALTER SEQUENCE public.verifications_id_seq OWNED BY public.verifications.id;


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: HaNzO
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: verification id; Type: DEFAULT; Schema: public; Owner: HaNzO
--

ALTER TABLE ONLY public.verification ALTER COLUMN id SET DEFAULT nextval('public.verification_id_seq'::regclass);


--
-- Name: verifications id; Type: DEFAULT; Schema: public; Owner: HaNzO
--

ALTER TABLE ONLY public.verifications ALTER COLUMN id SET DEFAULT nextval('public.verifications_id_seq'::regclass);


--
-- Name: users phone_number_unique; Type: CONSTRAINT; Schema: public; Owner: HaNzO
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT phone_number_unique UNIQUE (phone_number);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: HaNzO
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: verification verification_pkey; Type: CONSTRAINT; Schema: public; Owner: HaNzO
--

ALTER TABLE ONLY public.verification
    ADD CONSTRAINT verification_pkey PRIMARY KEY (id);


--
-- Name: verifications verifications_pkey; Type: CONSTRAINT; Schema: public; Owner: HaNzO
--

ALTER TABLE ONLY public.verifications
    ADD CONSTRAINT verifications_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

