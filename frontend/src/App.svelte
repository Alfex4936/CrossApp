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
    import CategoryItem from "./components/category/CategoryItem.svelte";
    import CategoryList, {
        categories,
    } from "./components/category/CategoryList.svelte";

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

    function make_link(cateId = 1, nums = 7): string {
        return (
            ajou_link +
            "?mode=list&srCategoryId=" +
            cateId +
            "&srSearchKey=&srSearchVal=&article.offset=0&articleLimit=" +
            nums
        );
    }

    function updateNotice(menu = 1) {
        let category = "";

        switch (menu) {
            case 2: {
                category = make_link(1, number_of_notice);
                break;
            }
            case 3: {
                category = make_link(2, number_of_notice);
                break;
            }
            case 4: {
                category = make_link(3, number_of_notice);
                break;
            }
            case 5: {
                category = make_link(4, number_of_notice);
                break;
            }
            case 6: {
                category = make_link(5, number_of_notice);
                break;
            }
            case 7: {
                category = make_link(6, number_of_notice);
                break;
            }
            case 8: {
                category = make_link(7, number_of_notice);
                break;
            }
            case 9: {
                category = make_link(8, number_of_notice);
                break;
            }
            case 10: {
                category = make_link(166, number_of_notice);
                break;
            }
            case 11: {
                category = make_link(167, number_of_notice);
                break;
            }
            case 12: {
                category = make_link(168, number_of_notice);
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
    <CategoryList>
        {#each categories as cate, i}
            <CategoryItem
                callback={() => {
                    last_menu = cate.id;
                    updateNotice(cate.id);
                }}
                name={cate.name}
            />
            {#if i != 11}
            &nbsp;|&nbsp;
            {/if}
        {/each}
    </CategoryList>

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
    /* #logo {
        display: block;
        width: 200px;
        height: 182px;
        margin: auto;
        padding: 10% 0 0;
        background-position: center;
        background-repeat: no-repeat;
        background-size: 100% 100%;
        background-origin: content-box;
    } */

    #weather {
        width: 10%;
    }

</style>
