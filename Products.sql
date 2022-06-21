create table "Products"
(
    id          serial
        constraint products_pk
            primary key,
    name        text,
    created_at  timestamp,
    price       text,
    description text
);

alter table "Products"
    owner to "grpcSimple";

create unique index products_id_uindex
    on "Products" (id);

INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (3, 'NFT-3', '2022-06-15 16:20:00.000000', '76', 'bla bla 3');
INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (1, 'nft-1', '2022-06-04 15:09:46.000000', '4', 'bla bla 1');
INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (4, 'NFT-4', '2022-06-10 05:00:00.000000', '0.94', 'bla bla 4');
INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (2, 'NFT-2', '2022-06-08 04:10:08.000000', '547', 'bla bla 2');
