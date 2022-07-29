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
INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (5, 'NFT-5', '2022-05-03 14:11:08.000000', '5', 'bla bla 5');
INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (6, 'NFT-6', '2022-05-08 06:01:08.000000', '7', 'bla bla 6');
INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (7, 'NFT-7', '2022-05-09 19:10:08.000000', '47', 'bla bla 7');
INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (8, 'NFT-8', '2022-05-16 23:10:08.000000', '54', 'bla bla 8');
INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (9, 'NFT-9', '2022-05-19 15:10:08.000000', '0.56', 'bla bla 9');
INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (10, 'NFT-10', '2022-04-07 12:10:08.000000', '0.012', 'bla bla 10');
INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (11, 'NFT-11', '2022-04-07 11:10:08.000000', '2.89', 'bla bla 11');
INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (12, 'NFT-12', '2022-04-06 10:10:08.000000', '7.16', 'bla bla 12');
INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (13, 'NFT-13', '2022-04-05 09:10:08.000000', '11', 'bla bla 13');
INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (14, 'NFT-14', '2022-04-04 08:10:08.000000', '34', 'bla bla 14');
INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (15, 'NFT-15', '2022-03-03 07:10:08.000000', '18', 'bla bla 15');
INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (16, 'NFT-16', '2022-03-08 06:10:08.000000', '22', 'bla bla 16');
INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (17, 'NFT-17', '2022-03-01 05:10:08.000000', '1000', 'bla bla 17');
INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (18, 'NFT-18', '2022-03-22 04:10:08.000000', '187', 'bla bla 18');
INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (19, 'NFT-19', '2022-02-14 17:10:08.000000', '1.789', 'bla bla 19');
INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (20, 'NFT-20', '2022-02-18 14:10:08.000000', '90', 'bla bla 20');
INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (21, 'NFT-21', '2022-02-08 04:10:08.000000', '99', 'bla bla 21');
INSERT INTO public."Products" (id, name, created_at, price, description) VALUES (22, 'NFT-22', '2022-01-10 02:54:08.000000', '2', 'bla bla 22');
