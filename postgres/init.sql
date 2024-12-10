create table IF NOT EXISTS public.products
(
    id       uuid    not null primary key,
    name     text    not null,
    upc      text    not null,
    price    decimal not null,
    quantity integer not null
);

alter table public.products
    owner to postgres;

INSERT INTO public.products (id, name, upc, price, quantity)
VALUES ('0193a998-3c53-70cc-8e81-e16d8f9472f3', 'Blaze Innovative Scale', '038899974595', 66.65, 57);
INSERT INTO public.products (id, name, upc, price, quantity)
VALUES ('0193a998-408f-7e59-9bce-61b293733f68', 'Modular Wireless Fitness Tracker', '059573664981', 441.91, 66);
INSERT INTO public.products (id, name, upc, price, quantity)
VALUES ('0193a998-43cb-7587-a018-a6a5f9a9197c', 'Blue Fridge Boost', '004011552917', 401.55, 57);
INSERT INTO public.products (id, name, upc, price, quantity)
VALUES ('0193a998-4643-71ea-8be6-e8c259284ccb', 'Prime Compact Phone', '041184580461', 641.62, 44);
INSERT INTO public.products (id, name, upc, price, quantity)
VALUES ('0193a998-48cb-7568-b974-5aae66a83f61', 'Modular Voice-Controlled Fan', '073274785658', 333.09, 78);
