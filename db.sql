drop table if exists public.core_data_stamp cascade;
create table public.core_data_stamp
(
    status integer default 1,
    created_date timestamp without time zone,
    created_by varchar(50),
    updated_date timestamp without time zone,
    updated_by varchar(50),
    restored_date timestamp without time zone,
    restored_by varchar(50),
    deleted_date timestamp without time zone,
    deleted_by varchar(50)
);

drop table if exists public.auth_role cascade;
create table public.auth_role
(
    id serial primary key,
    name varchar(100) unique
) inherits (public.core_data_stamp);

drop table if exists public.auth_user cascade;
create table public.auth_user
(
    id serial primary key,
    username varchar(100) unique,
    email varchar(100) unique,
    password varchar(255),
    role_id integer references public.auth_role (id) not null
) inherits (public.core_data_stamp);

/* TRIGGER */
CREATE OR REPLACE FUNCTION data_stamp() RETURNS TRIGGER AS $data_stamp$
BEGIN
    IF (TG_OP = 'INSERT') THEN
    IF(NEW.created_by IS NULL) THEN
        NEW.created_by := CURRENT_USER;
    END IF;
    NEW.created_date := NOW();
    END IF;

    IF (TG_OP = 'UPDATE') THEN
    CASE NEW.status
        WHEN 1,2 THEN
            IF (OLD.status = 3) THEN
                IF (NEW.restored_by IS NULL) THEN
                    NEW.restored_by := CURRENT_USER;
                END IF;
                NEW.restored_date := NOW();
            ELSE
                IF (NEW.updated_by IS NULL) THEN
                    NEW.updated_by := CURRENT_USER;
                END IF;
                NEW.updated_date := NOW();
            END IF;
        WHEN 3 THEN
            IF (NEW.deleted_by IS NULL) THEN
                NEW.deleted_by := CURRENT_USER;
            END IF;
            NEW.deleted_date := NOW();
    END CASE;
    END IF;
    RETURN NEW;
END;
$data_stamp$ LANGUAGE plpgsql;

DO $$
DECLARE
    t record;
BEGIN
    FOR t IN 
        SELECT * FROM information_schema.columns
        WHERE TRUE 
        AND column_name = 'created_date'
        AND table_schema = 'public'
    LOOP
        EXECUTE format('CREATE TRIGGER data_stamp
                        BEFORE INSERT OR UPDATE ON %I.%I
                        FOR EACH ROW EXECUTE PROCEDURE data_stamp()',
                        t.table_schema, t.table_name);
    END LOOP;
END;
$$ LANGUAGE plpgsql;



/* DUMP DATA */
/* auth_role */
insert into public.auth_role(name) values ('superuser'), 
('admin');

/* auth_user */
insert into public.auth_user(username, email, password, role_id) values('root', 'root@prabowo.com', '12345', 1);

