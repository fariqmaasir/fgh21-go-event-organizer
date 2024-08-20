insert into
    "users" (
        "email",
        "password",
        "username"
    )
VALUES (
        'admin@mail.com',
        '$argon2id$v=19$m=65536,t=3,p=4$nNZriJvwkgphr0dr7DLsNw$xRzlQ5noz7Dr3EnsetQRkxKHOmHG+sx2NpKZBxDKspc',
        'admin'
    );

insert into
    "profile" (
        "picture",
        "full_name",
        "birth_date",
        "gender",
        "phone_number",
        "profession",
        "nationality_id",
        "user_id"
    )
VALUES (
        'https://img.freepik.com/free-vector/business-user-cog_78370-7040.jpg',
        'Administrator',
        '2000-01-01 00:00:00.000 +0700',
        '0',
        '0812345678910',
        'Fullstack Developer',
        83,
        1
    );

insert into
    events (
        "image",
        "title",
        "date",
        "descriptions",
        "created_by"
    )
values (
        'https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1709863296-eJ48e4nUQ3TQv9YzBNOoxKUh6VyspQu8.jpg?width=1024&quality=90',
        'MOTOGP MANDALIKA 2024',
        '2024-09-27 10:00:00.000 +0700',
        'Event internasional MotoGP Mandalika akan dilaksanakan pada 27-29 September 2024. Perlombaan ini akan dilaksanakan di Pertamina Mandalika International Circuit, Lombok Tengah, Nusa Tenggara Barat (NTB).MotoGP ini sangat dinanti-nanti oleh masyarakat Indonesia, terutama yang berasal dari Pulau Lombok. Selain dapat menyaksikan MotoGP secara langsung, event ini juga akan memberikan keuntungan bagi warga sekitar dalam bidang pariwisata dan ekonomi.',
        1
    ),
    (
        'https://assets.techverse.asia/media/2024/06/24/screenshot_2024_06_24_141837.webp',
        'BRUNO MARS CONCERT',
        '2024-09-13 22:00:00.000 +0700',
        'Penyanyi Bruno Mars akan menyelenggarakan konser selama dua hari di Jakarta pada 13-14 September 2024 berlokasi di Jakarta International Stadium (JIS),Konser bertajuk Bruno Mars Live in Jakarta 2024 ini ditargetkan bisa menggaet total 100 ribu penonton selama dua hari pelaksanaan konsernya.',
        1
    ),
    (
        'https://konikepri.id/wp-content/uploads/2023/01/100000.jpeg',
        'PON XXI ACEH-SUMUT 2024',
        '2024-08-20 06:00:00.000 +0700',
        'Pekan Olahraga Nasional 2024 (biasa disingkat PON XXI/PON ACEH-SUMUT 2024) adalah penyelenggaraan ke-21 dari Pekan Olahraga Nasional (PON) yang merupakan ajang multi-olahraga nasional utama yang dijadwalkan berlangsung dari 8 hingga 20 September 2024 dengan Aceh dan Sumatera Utara sebagai tuan rumah.',
        1
    ),
    (
        'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQQC28j_FXLATnKHAPjTU8lpEh69bLOIHyAicMLuo6YhDB2RobPLNL2pnPKdXEDWXFHgPk&usqp=CAU',
        'The 90s Festival 2024',
        '2024-08-10 20:00:00.000 +0700',
        'Festival musik bergengsi, The 90s Festival 2024 bersiap digelar selama dua hari berturut pada 10 Agustus 2024 dan 11 Agustus 2024. Lokasinya di Gambir Expo Kemayoran, Jakarta.',
        1
    ),
    (
        'https://akcdn.detik.net.id/community/media/visual/2023/03/27/red-velvet.jpeg?w=1000',
        'Red Velvet 4Th Concert',
        '2024-09-07 20:00:00.000 +0700',
        'Red Velvet akan menggelar fancon di Indonesia pada 7 September 2024. Fancon bertajuk Happiness: My Dear, ReVe1uv itu akan diadakan di Beach City International Stadium.Jakarta jadi salah satu kota yang akan dikunjungi Red Velvet dalam lawatan mereka ke Asia edisi 2024. Tur fancon itu akan dimulai di Seoul pada 3-4 Agustus, kemudian ke Bangkok selama dua hari pada 17-18 Agustus 2024. Red Velvet kemudian menyambangi Indonesia pada 7 September, diikuti Manila pada 17 September, lalu berakhir di Macau pada 28 September.',
        1
    ),
    (
        'https://storage.jakarta-tourism.go.id/public/events/REqPzLRHe99fVy8hiS5Tiux2r88pVf-metaU25hcGluc3RhLmFwcF8zODEyNjQ2MzZfMTc4ODcxMTgyNzk5MjIyNDVfMzQyOTk0ODA2Mjk2MjExMjMxNV9uXzEwODAuanBn-.jpg',
        'Jakarta Desert Week',
        '2024-09-25 10:00:00.000 +0700',
        'Jakarta Dessert Week (JDW) merupakan rangkaian kolektif kegiatan dessert tahunan yang berkomitmen untuk meningkatkan kesadaran akan perkembangan industri makanan manis. JDW merupakan acara pertama yang memperkenalkan konsep restaurant-week di Indonesia, memberikan pengalaman unik bagi para pecinta makanan manis untuk menjelajahi berbagai hidangan lezat dari para ahli di bidangnya. Sebagai sebuah festival yang merayakan kekayaan dan keragaman dessert, JDW juga menjadi platform kreatif bagi para chef, pelaku industri, dan penggemar kuliner untuk berkumpul, berbagi, dan menikmati kreasi terbaik dari industri dessert di Jakarta.',
        1
    ),
    (
        'https://upload.wikimedia.org/wikipedia/commons/a/aa/Ngayogjazz_2012.jpg',
        'Ngayogjazz',
        '2024-11-16 20:00:00.000 +0700',
        'Ngayogjazz adalah festival musik jazz tahunan di Daerah Istimewa Yogyakarta yang diselenggarakan oleh para seniman dan komunitas musik jazz lokal Yogyakarta sejak tahun 2007.[1] Kegiatan tahunan ini terbuka bagi masyarakat dan gratis. Setiap tahunnya kegiatan ini diselenggarakan di desa-desa yang ada di Yogyakarta, hal ini bertujuan untuk melibatkan masyarakat dan memajukan pariwisata setempat serta mematahkan stigma bahwa jazz merupakan musik mahal yang dapat dinikmati segelintir orang saja.',
        1
    ),
    (
        'https://eventdaerah.kemenparekraf.go.id/storage/app/uploads/public/664/b3b/342/664b3b342d730842313501.png',
        'Wayang Jogja Night Carnival',
        '2024-10-07 20:00:00.000 +0700',
        'Wayang Jogja Night Carnival (WJNC) telah menjadi event unggulan Kota Yogyakarta yang diadakan untuk memperingati Hari Ulang Tahun Kota Yogyakarta pada 7 Oktober setiap tahunnya. WJNC merupakan karnaval jalanan (art on the street) dengan menggabungkan tema pewayangan yang telah dilaksanakan secara rutin sejak tahun 2016.  Penampilan karnaval ini melibatkan seni koreografi, busana, musik kontemporer, dan permainan lighting. Pada perkembangannya, konsep karnaval berubah menjadi street art dengan tidak meninggalkan unsur utama dari WJNC, yaitu Tugu Jogja, kendaraan hias, Wayang dan diadakan pada malam hari.',
        1
    ),
    (
        'https://eventdaerah.kemenparekraf.go.id/storage/app/uploads/public/664/c12/bae/664c12baecd5f940064969.png',
        'Jakarta Dessert Week',
        '2024-09-13 20:00:00.000 +0700',
        'Aceh merupakan salah satu daerah di Indonesia yang memiliki beragam alat musik perkusi. Berdasarkan sejarah lahir dan berkembangnya, alat musik perkusi di Aceh sering ditampilkan secara kolektif pada acara-acara sakral adat istiadat dan keagamaan. Aceh Perkusi merupakan agenda dua tahunan yang dilaksanakan untuk mengangkat keberagaman ini. Keberagaman perkusi tersebut juga turut dikemas dalam bentuk kolaborasi dengan perkusi dari provinsi lain serta mancanegara. Sehingga menjadikan event ini sangat unik, menarik, atraktif, dan berkarakter.',
        1
    ),
    (
        'https://eventdaerah.kemenparekraf.go.id/storage/app/uploads/public/64e/498/731/64e4987315159134586625.jpeg',
        'Festival Pesona Raja Ampat',
        '2024-10-18 20:00:00.000 +0700',
        'Kepulauan Raja Ampat merupakan jantung segitiga karang dunia “Heart of Coral Triangle“ dan sudah diakui sebagai salah satu kawasan yang memiliki terumbu karang terbaik di dunia. Selain itu, Raja Ampat juga memiliki hamparan padang lamun, hutan bakau, pantai berpasir, dan pantai tebing berbatu. Keberlangsungan ekosistem darat dan laut di Kepulauan Raja Ampat merupakan kesatuan proses yang saling berpengaruh sehingga menjadi sangat penting untuk dijaga kelestariannya. Untuk itu, Festival Pesona Raja Ampat 2023 menekankan konsep tematik terhadap harmoni dari pegunungan hingga perairan laut, yaitu “Miracle Raja Ampat, Embrace Culture Preserve Nature“.',
        1
    );