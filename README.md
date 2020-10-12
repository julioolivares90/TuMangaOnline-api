# TuMangaOnline-api 

## **:package: Principales herramientas usadas**

- [x] fiber
- [x] go colly


## :rocket:  tumangaonline-api API enlace
https://tumangaonline-api.herokuapp.com/

### Start Server using go

```
go run App.go
```

### Request & Response Ejemplos de uso

### obtiene todos los mangas populares
https://tumangaonline-api.herokuapp.com/api/v1/manga/populares

```json
{
    "statusCode": 200,
    "data": [
        {
            "title": "Nishuume Cheat no Tensei Madoushi 〜Saikyou ga 1000-nengo ni Tensei",
            "score": "0.00",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/53154/nishuume-cheat-no-tensei-madoushi-saikyou-ga-1000-nengo-ni-tensei",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f5377efbcf2a.jpg"
        },
        {
            "title": "Koibito wo Netorare, Yuusha Party kara Tsuihou saretakedo, EX Skill [Kotei Dameiji] ni Mezamete Muteki no Sonzai ni. Saa, Fukushuu wo Hajimeyou.",
            "score": "0.00",
            "type": "MANGA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manga/52729/me-robaron-a-mi-novia-y-me-expulsaron-del-equipo-del-heroe-pero-desperte-la-habilidad-unica-dano-anulado-y-me-volvi-invencible-es-hora-de-conseguir-mi-venganza",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f3f6eb227f85.jpg"
        },
        {
            "title": "Shijou Saikyou Orc-san no Tanoshii Tanetsuke Harem Zukuri",
            "score": "8.06",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/46258/shijou-saikyou-orc-san-no-tanoshii-tanetsuke-harem-zukuri",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5d8ef12e50424.jpg"
        },
        {
            "title": "Dr. Stone",
            "score": "8.65",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/23741/dr-stone",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5d924d4309b18.jpg"
        },
        {
            "title": "Isekai de Slow Life wo (Ganbou)",
            "score": "10.00",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/47452/isekai-de-slow-life-wo-ganbou",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5e03c83021f8c.jpg"
        },
        {
            "title": "CUANDO LA CHICA MALVADA AMA",
            "score": "9.25",
            "type": "MANHWA",
            "demography": "Shoujo",
            "mangaUrl": "https://lectortmo.com/library/manhwa/53373/cuando-la-chica-malvada-ama",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f5fb88a121f1.jpg"
        },
        {
            "title": "The monster duchess and contract princess",
            "score": "9.69",
            "type": "MANHWA",
            "demography": "Shoujo",
            "mangaUrl": "https://lectortmo.com/library/manhwa/44038/la-duquesa-monstruosa-y-la-princesa-del-contrato",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5dbe02e90f5f0.jpg"
        },
        {
            "title": "Henkyou Gurashi no Maou, Tensei shite Saikyou no Majutsushi ni naru",
            "score": "9.00",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/52914/henkyou-gurashi-no-maou-tensei-shite-saikyou-no-majutsushi-ni-naru",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f493e4733287.jpg"
        },
        {
            "title": "Genjitsushugisha no Oukokukaizouki",
            "score": "8.38",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/28684/Genjitsushugisha-no-Oukok",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5d03083bc484e.jpg"
        },
        {
            "title": "Doctor Elise",
            "score": "8.91",
            "type": "MANHWA",
            "demography": "Shoujo",
            "mangaUrl": "https://lectortmo.com/library/manhwa/42389/doctora-elise",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5c606fabbaac4.jpg"
        },
        {
            "title": "Sobreviviendo Ante El Héroe",
            "score": "9.67",
            "type": "MANHWA",
            "demography": "Shoujo",
            "mangaUrl": "https://lectortmo.com/library/manhwa/42407/sobreviviendo-como-la-esposa-del-heroe",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5c3fcb1c0919c.jpg"
        },
        {
            "title": "Ankoku Kishi no Ore desu Ga Saikyou no Seikishi wo Mezashimasu",
            "score": "10.00",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/51126/ankoku-kishi-no-ore-desu-ga-saikyou-no-seikishi-wo-mezashimasu",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5eefdccf8042f.jpg"
        },
        {
            "title": "Rebirth After 80.000 Years Passed",
            "score": "8.25",
            "type": "MANHUA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manhua/47011/rebirth-after-80000-years-passed",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5ddad4d7b0eb2.jpg"
        },
        {
            "title": "Higyaku Danshi Fujisaki-kun",
            "score": "0.00",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/54103/higyaku-danshi-fujisaki-kun",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f82b5cbacc60.jpg"
        },
        {
            "title": "DUNGEON KURASHI NO MOTO YUUSHA",
            "score": "0.00",
            "type": "MANGA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manga/40904/dungeon-kurashi-no-moto-yuusha",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5ebb067d9f7e8.jpg"
        },
        {
            "title": "Isekai Seikatsu No Susume",
            "score": "10.00",
            "type": "MANGA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manga/49886/isekai-seikatsu-no-susume",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f558280540b2.jpg"
        },
        {
            "title": "El Jugador",
            "score": "7.58",
            "type": "MANHWA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/55/el-jugador",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f20d3e481d70.jpg"
        },
        {
            "title": "Martial Peak",
            "score": "9.44",
            "type": "MANHUA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhua/38921/martial-peak",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5eaf975a5fc80.jpg"
        },
        {
            "title": "Kagami No Mukou No Saihate Toshokan: Kou No Yuusha To Itsuwari No Maou",
            "score": "8.00",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/45568/kagami-no-mukou-no-saihate-toshokan-kou-no-yuusha-to-itsuwari-no-maou",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5e349e054f4f2.jpg"
        },
        {
            "title": "Overgeared",
            "score": "8.00",
            "type": "MANHWA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/48986/overgeared",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5e83a1879ce82.jpg"
        },
        {
            "title": "Bocchi Tenseiki",
            "score": "0.00",
            "type": "MANGA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manga/52981/bocchi-tenseiki",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f4c6c982d03a.jpg"
        },
        {
            "title": "Mi papá es demasiado fuerte",
            "score": "10.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/52126/mi-papa-es-demasiado-fuerte",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f348e21f0e45.jpg"
        },
        {
            "title": "Kimi wa Meido-sama",
            "score": "0.00",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/52295/kimi-wa-meido-sama",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f2b00abb3d05.jpg"
        },
        {
            "title": "Karakai Jouzu no (Moto) Takagi-san",
            "score": "9.20",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/29136/karakai-jouzu-no-moto-takagi-san",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5c7b0450287b7.jpg"
        },
        {
            "title": "LA LEY DEL INSO",
            "score": "9.67",
            "type": "MANHWA",
            "demography": "Shoujo",
            "mangaUrl": "https://lectortmo.com/library/manhwa/44379/la-ley-del-inso",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5ef924062a837.jpg"
        },
        {
            "title": "The Rise of the Unemployed Wise Man",
            "score": "0.00",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/51223/the-rise-of-the-unemployed-wise-man",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f498cbfcd20c.jpg"
        },
        {
            "title": "Elf-san wa Yaserarenai",
            "score": "8.00",
            "type": "MANGA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manga/22161/elf-san-wa-yaserarenai",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5b2f4ef557bc8.jpg"
        },
        {
            "title": "Yuujin Character wa Taihen desu ka?",
            "score": "0.00",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/47122/yuujin-character-wa-taihen-desu-ka",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5de4bd13bd696.jpg"
        },
        {
            "title": "LA AMANTE DEL CABALLERO DRAGÓN",
            "score": "0.00",
            "type": "MANGA",
            "demography": "Shoujo",
            "mangaUrl": "https://lectortmo.com/library/manga/45851/la-amante-del-caballero-dragon",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5d6a02ff84c43.jpg"
        },
        {
            "title": "Kanojo, Hitomishirimasu",
            "score": "10.00",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/51098/kanojo-hitomishirimasu",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5eee852476a68.jpg"
        }
    ]
}
```
### obtiene todos los mangas populares seinen 
https://tumangaonline-api.herokuapp.com/api/v1/manga/populares-seinen
```json
{
    "statusCode": 200,
    "data": [
        {
            "title": "Dueña de la Pensión",
            "score": "0.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/53758/duena-de-la-pension",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f7186dfd183c.jpg"
        },
        {
            "title": "Jimiko wa Igai ni Ero Katta",
            "score": "0.00",
            "type": "MANGA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manga/49040/jimiko-wa-igai-ni-erokatta",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5e87a536c19a4.jpg"
        },
        {
            "title": "Young Boss",
            "score": "10.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/49943/young-boss",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5eb71213bcb93.jpg"
        },
        {
            "title": "SISTERS AND FATHER",
            "score": "0.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/47560/sisters-and-father",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5e0ba6dcc40ae.jpg"
        },
        {
            "title": "Luvslave",
            "score": "10.00",
            "type": "MANGA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manga/44087/luvslave",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5cd105350800b.jpg"
        },
        {
            "title": "A KNOWING SISTER",
            "score": "0.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/53156/a-knowing-sister",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f53afdf115ed.jpg"
        },
        {
            "title": "The Girl Hiding in the Wall",
            "score": "0.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/49956/the-girl-hiding-in-the-wall",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5ec0c9a5e9b9d.jpg"
        },
        {
            "title": "Cowgirl´s Riding-Position Makes Me Cum",
            "score": "10.00",
            "type": "MANGA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manga/46302/cowgirls-riding-position-makes-me-cum",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5d94b98f26c01.jpg"
        },
        {
            "title": "Bullied Boy's Tongue Revenge",
            "score": "10.00",
            "type": "MANGA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manga/46118/bullied-boys-tongue-revenge",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5d84e2f18a859.jpg"
        },
        {
            "title": "Cram School Scandal",
            "score": "8.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/50813/cram-school-scandal",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f5e66e7894f2.jpg"
        },
        {
            "title": "Uq Holder",
            "score": "7.67",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/8777/uq-holder",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5e283e1726909.jpg"
        },
        {
            "title": "A Killer Woman",
            "score": "0.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/49207/a-killer-woman",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5e90fd0651625.jpg"
        },
        {
            "title": "Depredadores de la Estética",
            "score": "0.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/51526/depredadores-de-la-estetica",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f03fe5f3e512.jpg"
        },
        {
            "title": "Mala Sangre",
            "score": "0.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/50385/bad-blood",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5ecc89bc91261.jpg"
        },
        {
            "title": "Grabbed My Boobs For 16 Hours",
            "score": "10.00",
            "type": "MANGA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manga/46218/grabbed-my-boobs-for-16-hours",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5d8b7e4dd3c8b.jpg"
        },
        {
            "title": "Ejaculation",
            "score": "8.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/48710/ejaculation",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5e6eeca783872.jpg"
        },
        {
            "title": "El vastago",
            "score": "0.00",
            "type": "NOVELA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/novel/54074/el-vastago",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f80f6d734cbc.jpg"
        },
        {
            "title": "Tengo una Mansión en el mundo Post-apocalíptico",
            "score": "8.50",
            "type": "NOVELA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/novel/45123/tengo-una-mansion-en-el-mundo-post-apocaliptico",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5d2e05cfaec2c.jpg"
        },
        {
            "title": "La novia de mi amigo",
            "score": "8.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/49044/la-novia-de-mi-amigo",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5e87d57437924.jpg"
        },
        {
            "title": "SISTERS WIVE",
            "score": "9.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/41955/sisters-wive",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5c13d0d017c81.jpg"
        },
        {
            "title": "Necesito que me ayudes",
            "score": "9.00",
            "type": "MANHWA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/manhwa/48708/necesito-que-me-ayudes",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5e6eea57e8bb3.jpg"
        }
    ]
}
```
### obtener la informacion de un manga 
https://tumangaonline-api.herokuapp.com/api/v1/manga/info?mangaUrl=https://lectortmo.com/library/manga/23741/dr-stone
```json
{
    "statusCode": 200,
    "data": {
        "title": "Dr. Stone\n( 2017 )",
        "image": "https://otakuteca.com/images/books/cover/5d924d4309b18.jpg",
        "tipo": "MANGA",
        "score": "8.65",
        "demografia": "Shounen",
        "descripcion": "Senkuu y Taiju son dos amigos y estudiantes de preparatoria, el primero de ellos es un genio del club de química mientras que el otro es un grandullón musculoso e idiota. Después de 5 largos años, Taiju ha decidido declararle su amor a Yuzuriha, pero no todo siempre sale como nos gustaría... En la tierra ocurre un suceso que convierte a todo el mundo en piedra. ¿Cómo lograrán sobrevivir en este mundo post apocalíptico?",
        "estado": "Publicándose",
        "generos": [
            "Acción",
            "Aventura",
            "Apocalíptico",
            "Comedia",
            "Supervivencia"
        ],
        "capitulo": [
            {
                "Title": "Capítulo 169.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/614301"
            },
            {
                "Title": "Capítulo 168.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/610128"
            },
            {
                "Title": "Capítulo 167.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/605865"
            },
            {
                "Title": "Capítulo 166.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/600460"
            },
            {
                "Title": "Capítulo 165.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/597354"
            },
            {
                "Title": "Capítulo 164.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/593167"
            },
            {
                "Title": "Capítulo 163.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/588720"
            },
            {
                "Title": "Capítulo 162.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/583313"
            },
            {
                "Title": "Capítulo 161.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/575204"
            },
            {
                "Title": "Capítulo 160.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/569910"
            },
            {
                "Title": "Capítulo 159.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/560692"
            },
            {
                "Title": "Capítulo 158.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/555993"
            },
            {
                "Title": "Capítulo 157.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/552008"
            },
            {
                "Title": "Capítulo 156.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/546212"
            },
            {
                "Title": "Capítulo 155.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/542982"
            },
            {
                "Title": "Capítulo 154.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/538556"
            },
            {
                "Title": "Capítulo 153.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/534337"
            },
            {
                "Title": "Capítulo 152.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/530049"
            },
            {
                "Title": "Capítulo 151.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/525892"
            },
            {
                "Title": "Capítulo 150.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/521795"
            },
            {
                "Title": "Capítulo 149.00:  Z=149",
                "UrlLeer": "https://lectortmo.com/view_uploads/517215"
            },
            {
                "Title": "Capítulo 148.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/508632"
            },
            {
                "Title": "Capítulo 147.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/501182"
            },
            {
                "Title": "Capítulo 146.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/497515"
            },
            {
                "Title": "Capítulo 145.00:  Z=145",
                "UrlLeer": "https://lectortmo.com/view_uploads/493633"
            },
            {
                "Title": "Capítulo 144.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/490458"
            },
            {
                "Title": "Capítulo 143.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/487496"
            },
            {
                "Title": "Capítulo 142.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/485099"
            },
            {
                "Title": "Capítulo 141.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/482560"
            },
            {
                "Title": "Capítulo 140.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/479199"
            },
            {
                "Title": "Capítulo 139.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/477590"
            },
            {
                "Title": "Capítulo 138.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/475211"
            },
            {
                "Title": "Capítulo 137.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/472893"
            },
            {
                "Title": "Capítulo 136.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/470392"
            },
            {
                "Title": "Capítulo 135.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/468030"
            },
            {
                "Title": "Capítulo 134.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/458662"
            },
            {
                "Title": "Capítulo 133.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/455410"
            },
            {
                "Title": "Capítulo 132.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/453206"
            },
            {
                "Title": "Capítulo 131.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/451300"
            },
            {
                "Title": "Capítulo 130.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/446807"
            },
            {
                "Title": "Capítulo 129.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/444699"
            },
            {
                "Title": "Capítulo 128.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/442215"
            },
            {
                "Title": "Capítulo 127.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/440032"
            },
            {
                "Title": "Capítulo 126.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/437795"
            },
            {
                "Title": "Capítulo 125.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/435418"
            },
            {
                "Title": "Capítulo 124.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/433541"
            },
            {
                "Title": "Capítulo 123.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/431531"
            },
            {
                "Title": "Capítulo 122.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/428997"
            },
            {
                "Title": "Capítulo 121.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/426859"
            },
            {
                "Title": "Capítulo 120.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/425218"
            },
            {
                "Title": "Capítulo 119.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/423182"
            },
            {
                "Title": "Capítulo 118.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/421095"
            },
            {
                "Title": "Capítulo 117.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/419304"
            },
            {
                "Title": "Capítulo 116.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/414865"
            },
            {
                "Title": "Capítulo 115.00:  Un Segundo y un Grano",
                "UrlLeer": "https://lectortmo.com/view_uploads/412452"
            },
            {
                "Title": "Capítulo 114.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/410260"
            },
            {
                "Title": "Capítulo 113.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/408369"
            },
            {
                "Title": "Capítulo 112.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/406164"
            },
            {
                "Title": "Capítulo 111.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/404231"
            },
            {
                "Title": "Capítulo 110.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/402152"
            },
            {
                "Title": "Capítulo 109.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/400129"
            },
            {
                "Title": "Capítulo 108.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/398226"
            },
            {
                "Title": "Capítulo 107.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/396468"
            },
            {
                "Title": "Capítulo 106.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/394505"
            },
            {
                "Title": "Capítulo 105.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/392423"
            },
            {
                "Title": "Capítulo 104.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/390457"
            },
            {
                "Title": "Capítulo 103.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/386445"
            },
            {
                "Title": "Capítulo 102.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/384562"
            },
            {
                "Title": "Capítulo 101.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/382552"
            },
            {
                "Title": "Capítulo 100.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/380719"
            },
            {
                "Title": "Capítulo 99.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/376095"
            },
            {
                "Title": "Capítulo 98.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/374013"
            },
            {
                "Title": "Capítulo 97.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/371900"
            },
            {
                "Title": "Capítulo 96.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/369704"
            },
            {
                "Title": "Capítulo 95.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/367624"
            },
            {
                "Title": "Capítulo 94.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/365790"
            },
            {
                "Title": "Capítulo 93.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/363396"
            },
            {
                "Title": "Capítulo 92.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/361621"
            },
            {
                "Title": "Capítulo 91.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/359284"
            },
            {
                "Title": "Capítulo 90.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/357168"
            },
            {
                "Title": "Capítulo 89.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/350623"
            },
            {
                "Title": "Capítulo 88.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/347744"
            },
            {
                "Title": "Capítulo 87.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/346342"
            },
            {
                "Title": "Capítulo 86.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/344868"
            },
            {
                "Title": "Capítulo 85.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/343172"
            },
            {
                "Title": "Capítulo 84.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/341681"
            },
            {
                "Title": "Capítulo 83.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/339986"
            },
            {
                "Title": "Capítulo 82.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/338612"
            },
            {
                "Title": "Capítulo 81.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/336969"
            },
            {
                "Title": "Capítulo 80.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/335281"
            },
            {
                "Title": "Capítulo 79.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/333848"
            },
            {
                "Title": "Capítulo 78.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/332253"
            },
            {
                "Title": "Capítulo 77.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/330680"
            },
            {
                "Title": "Capítulo 76.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/329284"
            },
            {
                "Title": "Capítulo 75.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/327340"
            },
            {
                "Title": "Capítulo 74.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/325605"
            },
            {
                "Title": "Capítulo 73.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/324010"
            },
            {
                "Title": "Capítulo 72.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/322056"
            },
            {
                "Title": "Capítulo 71.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/320476"
            },
            {
                "Title": "Capítulo 70.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/318903"
            },
            {
                "Title": "Capítulo 69.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/315629"
            },
            {
                "Title": "Capítulo 68.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/314014"
            },
            {
                "Title": "Capítulo 67.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/312421"
            },
            {
                "Title": "Capítulo 66.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/310694"
            },
            {
                "Title": "Capítulo 65.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/309146"
            },
            {
                "Title": "Capítulo 64.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/307552"
            },
            {
                "Title": "Capítulo 63.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/305671"
            },
            {
                "Title": "Capítulo 62.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/303965"
            },
            {
                "Title": "Capítulo 61.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/302801"
            },
            {
                "Title": "Capítulo 60.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/301621"
            },
            {
                "Title": "Capítulo 59.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/300268"
            },
            {
                "Title": "Capítulo 58.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/298914"
            },
            {
                "Title": "Capítulo 57.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/297743"
            },
            {
                "Title": "Capítulo 56.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/295142"
            },
            {
                "Title": "Capítulo 55.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/293650"
            },
            {
                "Title": "Capítulo 54.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/292365"
            },
            {
                "Title": "Capítulo 53.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/291110"
            },
            {
                "Title": "Capítulo 52.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/289612"
            },
            {
                "Title": "Capítulo 51.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/288366"
            },
            {
                "Title": "Capítulo 50.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/287133"
            },
            {
                "Title": "Capítulo 49.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/285849"
            },
            {
                "Title": "Capítulo 48.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/284489"
            },
            {
                "Title": "Capítulo 47.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/282996"
            },
            {
                "Title": "Capítulo 46.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/281709"
            },
            {
                "Title": "Capítulo 45.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/280051"
            },
            {
                "Title": "Capítulo 44.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/278876"
            },
            {
                "Title": "Capítulo 43.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/276248"
            },
            {
                "Title": "Capítulo 42.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/274835"
            },
            {
                "Title": "Capítulo 41.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/271995"
            },
            {
                "Title": "Capítulo 40.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/270727"
            },
            {
                "Title": "Capítulo 39.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/268456"
            },
            {
                "Title": "Capítulo 38.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/267200"
            },
            {
                "Title": "Capítulo 37.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/265948"
            },
            {
                "Title": "Capítulo 36.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/264827"
            },
            {
                "Title": "Capítulo 35.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/263573"
            },
            {
                "Title": "Capítulo 34.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/262463"
            },
            {
                "Title": "Capítulo 33.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/261347"
            },
            {
                "Title": "Capítulo 32.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/260358"
            },
            {
                "Title": "Capítulo 31.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/259108"
            },
            {
                "Title": "Capítulo 30.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/257789"
            },
            {
                "Title": "Capítulo 29.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/256544"
            },
            {
                "Title": "Capítulo 28.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/255142"
            },
            {
                "Title": "Capítulo 27.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/253485"
            },
            {
                "Title": "Capítulo 26.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/252350"
            },
            {
                "Title": "Capítulo 25.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/251098"
            },
            {
                "Title": "Capítulo 24.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/249688"
            },
            {
                "Title": "Capítulo 23.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/248415"
            },
            {
                "Title": "Capítulo 22.50",
                "UrlLeer": "https://lectortmo.com/view_uploads/246054"
            },
            {
                "Title": "Capítulo 22.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/246141"
            },
            {
                "Title": "Capítulo 21.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/244261"
            },
            {
                "Title": "Capítulo 20.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/242959"
            },
            {
                "Title": "Capítulo 19.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/241778"
            },
            {
                "Title": "Capítulo 18.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/239758"
            },
            {
                "Title": "Capítulo 17.00",
                "UrlLeer": "https://lectortmo.com/view_uploads/238102"
            },
            {
                "Title": "Capítulo 16.00:  Kohaku",
                "UrlLeer": "https://lectortmo.com/view_uploads/236728"
            },
            {
                "Title": "Capítulo 15.00:  Los dos países del mundo de piedra",
                "UrlLeer": "https://lectortmo.com/view_uploads/235447"
            },
            {
                "Title": "Capítulo 14.00:  En lo que crees",
                "UrlLeer": "https://lectortmo.com/view_uploads/234401"
            },
            {
                "Title": "Capítulo 13.00:  El comienzo del mundo de piedra",
                "UrlLeer": "https://lectortmo.com/view_uploads/233284"
            },
            {
                "Title": "Capítulo 12.00:  Epílogo del Prólogo",
                "UrlLeer": "https://lectortmo.com/view_uploads/232165"
            },
            {
                "Title": "Capítulo 11.00:  El arma de la ciencia",
                "UrlLeer": "https://lectortmo.com/view_uploads/230929"
            },
            {
                "Title": "Capítulo 10.00:  La banda de la ciencia",
                "UrlLeer": "https://lectortmo.com/view_uploads/229892"
            },
            {
                "Title": "Capítulo 9.50:  Extra - ¡¡La cerámica: Intentémos hacerlo!!",
                "UrlLeer": "https://lectortmo.com/view_uploads/228698"
            },
            {
                "Title": "Capítulo 9.00:  Senku vs Tsukasa",
                "UrlLeer": "https://lectortmo.com/view_uploads/228552"
            },
            {
                "Title": "Capítulo 8.00:  Encender la señal",
                "UrlLeer": "https://lectortmo.com/view_uploads/226598"
            },
            {
                "Title": "Capítulo 7.00:  La aventura de la pólvora",
                "UrlLeer": "https://lectortmo.com/view_uploads/225353"
            },
            {
                "Title": "Capítulo 6.00:  Taiju vs Tsukasa",
                "UrlLeer": "https://lectortmo.com/view_uploads/223984"
            },
            {
                "Title": "Capítulo 5.00:  Yuzuriha",
                "UrlLeer": "https://lectortmo.com/view_uploads/222895"
            },
            {
                "Title": "Capítulo 4.00:  La almeja blanca pura",
                "UrlLeer": "https://lectortmo.com/view_uploads/221922"
            },
            {
                "Title": "Capítulo 3.00:  Rey del mundo de piedra",
                "UrlLeer": "https://lectortmo.com/view_uploads/220666"
            },
            {
                "Title": "Capítulo 2.00:  Fantasia vs Ciencia",
                "UrlLeer": "https://lectortmo.com/view_uploads/219474"
            },
            {
                "Title": "Capítulo 1.00:  Mundo de piedra",
                "UrlLeer": "https://lectortmo.com/view_uploads/218193"
            }
        ]
    }
}
```
### obtener las imagenes de un capitulo de un manga
https://tumangaonline-api.herokuapp.com/api/v1/manga/paginas?lectorTMO=https://lectortmo.com/view_uploads/569910
```json
{
    "statusCode": 200,
    "data": [
        "https://img1.tucomiconline.com/uploads/5f2701e221574/f64623e3.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/1e55867a.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/863bd767.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/2fd25756.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/7a465f81.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/ded9be6a.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/54739f43.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/a2b12cfa.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/1db24da4.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/d073729a.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/71042aa9.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/c22a10c7.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/8286bd16.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/930dfe36.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/cbebfd4a.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/611e4bb0.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/b6c20dc9.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/42671792.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/e5b57b1e.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/7558c109.jpg",
        "https://img1.tucomiconline.com/uploads/5f2701e221574/a110143b.jpg"
    ]
}
```
### realizar una busqueda de un manga
### listado de parametros 
- title
- _page
- orderItem
- orderDir
- filter_by
- Type
- demography
- status
- translation_status
- webcomic
- yonkoma
- amateur
- erotic

## enlace 
https://tumangaonline-api.herokuapp.com/api/v1/manga/library?title=naruto
```json
{
    "statusCode": 200,
    "data": [
        {
            "title": "Naruto",
            "score": "8.50",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/150/naruto",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5ed5b1dfc26c3.jpg"
        },
        {
            "title": "Boruto: Naruto Next Generations",
            "score": "6.92",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/15772/boruto",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f3ef9435be6c.jpg"
        },
        {
            "title": "Joke Box 7 Naruto Fanbook-Kakashi x Iruka",
            "score": "0.00",
            "type": "DOUJINSHI",
            "demography": "Shoujo",
            "mangaUrl": "https://lectortmo.com/library/doujinshi/40948/joke-box-7-naruto-fanbook-kakashi-x-iruka",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5ba4189b74e65.jpg"
        },
        {
            "title": "Naruto - Watashi no Kareshi wa Keikokubijin (Doujinshi)",
            "score": "0.00",
            "type": "DOUJINSHI",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/doujinshi/46640/naruto-watashi-no-kareshi-wa-keikokubijin-doujinshi",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5db834ffccd43.jpg"
        },
        {
            "title": "Naruto Gaiden: Nanadaime Hokage to Akairo no Hanatsuzuki",
            "score": "8.69",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/11344/naruto-gaiden-nanadaime-hokage-to-akairo-no-hanatsuzuki",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5edb2d0bc3bb2.jpg"
        },
        {
            "title": "Naruto Color",
            "score": "10.00",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/48176/naruto-color",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5e40d9a4efc46.jpg"
        },
        {
            "title": "Boruto ~Naruto The Movie~ Special Gaiden",
            "score": "8.51",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/12855/Boruto-Naruto-The-Movie-Specia",
            "mangaImagen": "https://otakuteca.com/images/books/cover/12855_TMOmanga013644.jpg"
        },
        {
            "title": "Naruto Especial: Boruto Road to B",
            "score": "8.35",
            "type": "ONE SHOT",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/one_shot/12697/Naruto-Especial-Boruto-Road-to",
            "mangaImagen": "https://otakuteca.com/images/books/cover/12697_TMOmanga043300.jpg"
        },
        {
            "title": "Naruto: The Path Lit by the Full Moon",
            "score": "8.29",
            "type": "ONE SHOT",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/one_shot/15552/naruto-gaiden-michita-tsuki-ga-terasu-michi",
            "mangaImagen": "https://otakuteca.com/images/books/cover/15552_TMOManga5719e0cca25b9.jpg"
        },
        {
            "title": "Naruto (Piloto)",
            "score": "7.99",
            "type": "ONE SHOT",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/one_shot/31871/Naruto-Piloto",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f46aa3cb4e9c.jpg"
        },
        {
            "title": "Naruto- festival de verano",
            "score": "0.00",
            "type": "DOUJINSHI",
            "demography": "Shoujo",
            "mangaUrl": "https://lectortmo.com/library/doujinshi/52147/naruto-festival-de-verano",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f23b7fbea104.jpg"
        },
        {
            "title": "Naruto: The Path Lit by the Full Moon",
            "score": "8.71",
            "type": "ONE SHOT",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/one_shot/12844/Naruto-ga-Hokage-ni-Natta-Hi",
            "mangaImagen": "https://otakuteca.com/images/books/cover/12844_TMOmanga024024.jpg"
        },
        {
            "title": "In a Different World with the Naruto System",
            "score": "0.00",
            "type": "NOVELA",
            "demography": "Seinen",
            "mangaUrl": "https://lectortmo.com/library/novel/41779/in-a-different-world-with-the-naruto-system",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5c027a1806966.jpg"
        },
        {
            "title": "Naruto parodia",
            "score": "5.46",
            "type": "MANGA",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/manga/12081/Naruto-parodia",
            "mangaImagen": "https://otakuteca.com/images/books/cover/12081_TMOmanga094331.jpg"
        },
        {
            "title": "NARUTO GAIDEN NUEVO CICLO",
            "score": "5.59",
            "type": "DOUJINSHI",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/doujinshi/19625/NARUTO-GAIDEN-NUEVO-CICLO",
            "mangaImagen": "https://otakuteca.com/images/books/cover/19625_TMOManga57c76213ecc34.jpg"
        },
        {
            "title": "Naruto ND",
            "score": "6.75",
            "type": "DOUJINSHI",
            "demography": "Shounen",
            "mangaUrl": "https://lectortmo.com/library/doujinshi/16018/Naruto-ND",
            "mangaImagen": "https://otakuteca.com/images/books/cover/16018_TMOManga5737cbea47cd2.jpg"
        },
        {
            "title": "Naruto-sazanka",
            "score": "0.00",
            "type": "DOUJINSHI",
            "demography": "Shoujo",
            "mangaUrl": "https://lectortmo.com/library/doujinshi/52754/naruto-sazanka",
            "mangaImagen": "https://otakuteca.com/images/books/cover/5f40eaa6178fb.jpg"
        }
    ]
}
```

