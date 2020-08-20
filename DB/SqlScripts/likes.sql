CREATE table likes (
    id SERIAL NOT NULL,
    self_id integer NOT NULL,
    lover_id integer NOT NULL,
    hasLove boolean  NOT NULL,
    PRIMARY KEY (id),
    foreign key (self_id) references public.users(id),
    foreign key (lover_id) references public.users(id)
);