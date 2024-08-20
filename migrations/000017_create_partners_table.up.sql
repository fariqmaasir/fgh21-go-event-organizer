create table partners (
    id serial primary key,
    image varchar(255),
    name varchar(100)
);

insert into
    partners ("image", "name")
values (
        'https://upload.wikimedia.org/wikipedia/commons/thumb/c/cd/Logo_Pesona_Indonesia_%28Kementerian_Pariwisata%29.png/1200px-Logo_Pesona_Indonesia_%28Kementerian_Pariwisata%29.png',
        'Wondeful Indonesia'
    ),
    (
        'https://bemubaya.com/assets/img/kementrian/senbud/Senbud.png',
        'Kementrian Seni Budaya'
    ),
    (
        'https://upload.wikimedia.org/wikipedia/commons/a/a2/Seal_of_the_Ministry_of_Tourism_and_Creative_Economy_of_the_Republic_of_Indonesia_%28Indonesian_version%29.svg',
        'Kemenparekraf/baparekraf'
    ),
    (
        'https://png.pngtree.com/png-clipart/20190614/original/pngtree-gudang-garam-logo-png-image_3718300.jpg',
        'Gudang Garam'
    ),
    (
        'https://upload.wikimedia.org/wikipedia/commons/thumb/e/eb/Logo_ovo_purple.svg/768px-Logo_ovo_purple.svg.png',
        'Ovo'
    ),
    (
        'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTVyXg71v_1IF35hyeGaNf7KiyGGp5YufsIYQ&s',
        'Fazztrack'
    );