<script context="module">
    export let object = {
        number_of_notice: 7,
    };
</script>

<script lang="ts">
    // import logo from "./assets/images/ajou-logo.png";

    // import {main} from '../wailsjs/go/models/models'

    // var notices: Array<main.Notice>;
    // let notice = new main.Notice()
    // let notices: Array<main.Notice> = [];

    import NoticeList from "./components/NoticeList.svelte";

    import { GetWeather, Parse } from "../wailsjs/go/main/App";

    let weather_promise = GetWeather();
    let notices_promise = Parse("", 7);
    let number_of_notice = 7;
    let last_menu = 1;

    // let weather_updater: ReturnType<typeof setTimeout>
    setTimeout(() => {
        weather_promise = GetWeather();
    }, 1000 * Math.floor(Math.random() * (610 - 590 + 1) + 590));

    const ajou_link = "https://www.ajou.ac.kr/kr/ajou/notice.do";

    function updateNotice(menu = 1) {
        let category = "";

        switch (menu) {
            case 2: {
                category =
                    ajou_link +
                    "?mode=list&srCategoryId=" +
                    1 +
                    "&srSearchKey=&srSearchVal=&article.offset=0&articleLimit=" +
                    number_of_notice;
                break;
            }
            case 3: {
                category =
                    ajou_link +
                    "?mode=list&srCategoryId=" +
                    2 +
                    "&srSearchKey=&srSearchVal=&article.offset=0&articleLimit=" +
                    number_of_notice;
                break;
            }
            case 4: {
                category =
                    ajou_link +
                    "?mode=list&srCategoryId=" +
                    3 +
                    "&srSearchKey=&srSearchVal=&article.offset=0&articleLimit=" +
                    number_of_notice;
                break;
            }
            case 5: {
                category =
                    ajou_link +
                    "?mode=list&srCategoryId=" +
                    4 +
                    "&srSearchKey=&srSearchVal=&article.offset=0&articleLimit=" +
                    number_of_notice;
                break;
            }
            case 6: {
                category =
                    ajou_link +
                    "?mode=list&srCategoryId=" +
                    5 +
                    "&srSearchKey=&srSearchVal=&article.offset=0&articleLimit=" +
                    number_of_notice;
                break;
            }
            case 7: {
                category =
                    ajou_link +
                    "?mode=list&srCategoryId=" +
                    6 +
                    "&srSearchKey=&srSearchVal=&article.offset=0&articleLimit=" +
                    number_of_notice;
                break;
            }
            case 8: {
                category =
                    ajou_link +
                    "?mode=list&srCategoryId=" +
                    7 +
                    "&srSearchKey=&srSearchVal=&article.offset=0&articleLimit=" +
                    number_of_notice;
                break;
            }
            case 9: {
                category =
                    ajou_link +
                    "?mode=list&srCategoryId=" +
                    8 +
                    "&srSearchKey=&srSearchVal=&article.offset=0&articleLimit=" +
                    number_of_notice;
                break;
            }
            case 10: {
                category =
                    ajou_link +
                    "?mode=list&srCategoryId=" +
                    166 +
                    "&srSearchKey=&srSearchVal=&article.offset=0&articleLimit=" +
                    number_of_notice;
                break;
            }
            case 11: {
                category =
                    ajou_link +
                    "?mode=list&srCategoryId=" +
                    167 +
                    "&srSearchKey=&srSearchVal=&article.offset=0&articleLimit=" +
                    number_of_notice;
                break;
            }
            case 12: {
                category =
                    ajou_link +
                    "?mode=list&srCategoryId=" +
                    168 +
                    "&srSearchKey=&srSearchVal=&article.offset=0&articleLimit=" +
                    number_of_notice;
                break;
            }
            default: {
                category = "";
                break;
            }
        }

        notices_promise = Parse(category, number_of_notice);
    }
</script>

<main>
    <ul id="menu">
        <li>
            <a
                class="category"
                href="/"
                on:click|preventDefault={() => {
                    last_menu = 1;
                    updateNotice(1);
                }}>전체 공지</a
            >
        </li>
        |
        <li>
            <a
                class="category"
                href="/"
                on:click|preventDefault={() => {
                    last_menu = 2;
                    updateNotice(2);
                }}>학사</a
            >
        </li>
        |
        <li>
            <a
                class="category"
                href="/"
                on:click|preventDefault={() => {
                    last_menu = 3;
                    updateNotice(3);
                }}>비교과</a
            >
        </li>
        |
        <li>
            <a
                class="category"
                href="/"
                on:click|preventDefault={() => {
                    last_menu = 4;
                    updateNotice(4);
                }}>장학</a
            >
        </li>
        |
        <li>
            <a
                class="category"
                href="/"
                on:click|preventDefault={() => {
                    last_menu = 5;
                    updateNotice(5);
                }}>학술</a
            >
        </li>
        |
        <li>
            <a
                class="category"
                href="/"
                on:click|preventDefault={() => {
                    last_menu = 6;
                    updateNotice(6);
                }}>입학</a
            >
        </li>
        |
        <li>
            <a
                class="category"
                href="/"
                on:click|preventDefault={() => {
                    last_menu = 7;
                    updateNotice(7);
                }}>취업</a
            >
        </li>
        |
        <li>
            <a
                class="category"
                href="/"
                on:click|preventDefault={() => {
                    last_menu = 8;
                    updateNotice(8);
                }}>사무</a
            >
        </li>
        |
        <li>
            <a
                class="category"
                href="/"
                on:click|preventDefault={() => {
                    last_menu = 9;
                    updateNotice(9);
                }}>기타</a
            >
        </li>
        |
        <li>
            <a
                class="category"
                href="/"
                on:click|preventDefault={() => {
                    last_menu = 10;
                    updateNotice(10);
                }}>행사</a
            >
        </li>
        |
        <li>
            <a
                class="category"
                href="/"
                on:click|preventDefault={() => {
                    last_menu = 11;
                    updateNotice(11);
                }}>파란학기제</a
            >
        </li>
        |
        <li>
            <a
                class="category"
                href="/"
                on:click|preventDefault={() => {
                    last_menu = 12;
                    updateNotice(12);
                }}>학사일정</a
            >
        </li>
    </ul>

    <!-- <img alt="ajou logo" id="logo" src={logo} /> -->
    {#await weather_promise}
        <p>🌞 날씨 불러오는 중...</p>
    {:then weather}
        <img
            on:click={() => (weather_promise = GetWeather())}
            title="눌러서 새로고침 하세요"
            alt="weather"
            id="weather"
            src={weather.icon}
        />
        <p>
            아주대 현재 날씨 {weather.current_temp} ({weather.current_stat}),
            최저 기온은 {weather.min_temp}, 최고 기온은 {weather.max_temp}<br />
            낮 강수량은 {weather.rain_day}, 밤 강수량은 {weather.rain_night}<br
            />
            내일 낮 강수량은 {weather.tmrw_rain_day}, 밤 감수량은 {weather.tmrw_rain_night}
        </p>
    {:catch error}
        <p style="color:red">{error.message}</p>
    {/await}

    <p style="display:inline">공지 갯수:</p>
    <input
        type="number"
        min="1"
        max="50"
        style="text-align:center"
        name="notice"
        bind:value={number_of_notice}
        placeholder="공지 갯수"
    />
    <button on:click={() => updateNotice(last_menu)}>공지 불러오기</button>

    {#await notices_promise}
        <p>💌 공지 불러오는 중...</p>
    {:then notices}
        <NoticeList {notices} />
    {:catch _}
        <p>공지 불러오는 중 에러 발생!</p>
    {/await}
</main>

<style>
    #logo {
        display: block;
        width: 200px;
        height: 182px;
        margin: auto;
        /* padding: 10% 0 0; */
        background-position: center;
        background-repeat: no-repeat;
        background-size: 100% 100%;
        background-origin: content-box;
    }

    #weather {
        width: 10%;
    }

    ul#menu li {
        display: inline;
    }

    a {
        color: inherit;
        -webkit-transition: 0.2s;
    }

    a:hover {
        color: hotPink;
    }

    @media (-webkit-min-device-pixel-ratio: 0) {
        a {
            background-color: white;
            background-image: -webkit-linear-gradient(
                left,
                hotPink 0%,
                orange 50%,
                transparent 50%
            );
            background-position: 100% 0;
            background-size: 200% 200%;
            color: transparent;
            -webkit-transition: 0.1s 0.2s;
            -webkit-background-clip: text;
        }
        a:hover {
            background-position: 0 0;
            color: transparent;
            transition: 0.4s 0;
        }
    }

    ul#menu,
    .category {
        font-size: 1.25rem;
        font-weight: bold;
        text-decoration: none;
    }
</style>
