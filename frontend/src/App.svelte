<script lang="ts">
    // import logo from "./assets/images/ajou-logo.png";

    // import {main} from '../wailsjs/go/models/models'
    // var notices: Array<main.Notice>;
    // let notice = new main.Notice()
    // let notices: Array<main.Notice> = [];

    import FullCalendar, { type CalendarOptions } from "svelte-fullcalendar";
    import daygridPlugin from "@fullcalendar/daygrid";
    let options: CalendarOptions = {
        // dateClick: (event) => alert('date click! ' + event.dateStr),
        events: [
            { title: "하계방학 시작, 2학기 휴학·복학 신청 시작", date: "2022-06-30" },
            { title: "2학기 1차 복학 신청", start: "2022-06-30", end: "2022-07-19" },
            { title: "2학기 등록", start: "2022-06-30", end: "2022-07-19" },
            { title: "2학기 예비 수강신청", start: "2022-07-21", end: "2022-07-23" },
            { title: "2학기 수강신청", start: "2022-08-08", end: "2022-08-12" },
            { title: "광복절", date: "2022-08-15" },
            { title: "2학기 개강", date: "2022-09-01" },
            { title: "추석 연휴", start: "2022-09-09", end: "2022-09-12" },
            { title: "개천절", date: "2022-10-03" },
            { title: "한글날", date: "2022-10-09" },
            { title: "2학기 중간시험", start: "2022-10-20", end: "2022-10-26" },
        ],
        initialView: "dayGridMonth",
        locale: "ko",
        plugins: [daygridPlugin],
    };

    import NoticeList from "./components/NoticeList.svelte";
    import { Tabs, Tab, TabList, TabPanel } from "svelte-tabs";
    import CategoryItem from "./components/category/CategoryItem.svelte";
    import CategoryList, {
        categories,
    } from "./components/category/CategoryList.svelte";
    import { GetWeather, Parse } from "../wailsjs/go/main/App";

    let weather_promise = GetWeather();
    let notices_promise = Parse("", 7);
    let number_of_notice = 7;
    let keyword_of_notice = "";

    // let weather_updater: ReturnType<typeof setTimeout>
    setTimeout(() => {
        weather_promise = GetWeather();
    }, 1000 * Math.floor(Math.random() * (610 - 590 + 1) + 590));

    const ajou_link = "https://www.ajou.ac.kr/kr/ajou/notice.do";

    function make_link(cateId = 0, nums = 7, search = ""): string {
        let basic = ajou_link + "?mode=list";
        basic += "&article.offset=0&articleLimit=" + nums;
        basic += "&srSearchKey=&srSearchVal=" + encodeURI(search);

        if (cateId > 0) {
            basic += "&srCategoryId=" + cateId;
        }
        return basic;
    }

    function updateNotice(cate_id = 0) {
        notices_promise = Parse(
            make_link(cate_id, number_of_notice, keyword_of_notice),
            number_of_notice
        );
    }
</script>

<svelte:head>
    <!-- <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@fortawesome/fontawesome-free@6.1.1/css/fontawesome.min.css" integrity="sha384-zIaWifL2YFF1qaDiAo0JFgsmasocJ/rqu7LKYH8CoBEXqGbb9eO+Xi3s6fQhgFWM" crossorigin="anonymous"> -->
    <link
        rel="stylesheet"
        href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css"
    />
</svelte:head>

<main>
    <Tabs>
        <TabList>
            <Tab>공지사항</Tab>
            <Tab>학사일정</Tab>
            <Tab>Three</Tab>
        </TabList>

        <!-- 공지사항 -->
        <TabPanel>
            <CategoryList>
                {#each categories as cate, i}
                    <CategoryItem
                        callback={() => {
                            // last_menu = cate.id;
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
                    최저 기온은 {weather.min_temp}, 최고 기온은 {weather.max_temp}<br
                    />
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
                on:input={updateNotice}
                placeholder="공지 갯수"
            />

            <input
                type="text"
                minlength="1"
                maxlength="15"
                style="text-align:center"
                name="keyword"
                bind:value={keyword_of_notice}
                on:change={updateNotice}
                placeholder="공지 검색"
            />

            <!-- <button on:click={() => (notices_promise = Parse(make_link(last_menu, number_of_notice), number_of_notice))}>공지 불러오기</button> -->
            <!-- <button on:click={() => (notices_promise = Parse(make_link(last_menu, number_of_notice), number_of_notice))}>공지 검색</button> -->

            {#await notices_promise}
                <p>💌 공지 불러오는 중...</p>
            {:then notices}
                <NoticeList {notices} />
            {:catch _}
                <p>공지 불러오는 중 에러 발생!</p>
            {/await}
        </TabPanel>
        <!-- 공지사항 -->

        <!-- 학사일정 -->
        <TabPanel>
            <FullCalendar {options} />
        </TabPanel>
        <!-- 학사일정 -->

        <!-- ㅊ -->
        <TabPanel>
            <h2>Panel Three</h2>
        </TabPanel>
        <!-- ㅊ -->
    </Tabs>
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
    :global(.svelte-tabs__tab-list) {
        border-bottom: none !important;
    }

    :global(.svelte-tabs li.svelte-tabs__selected) {
        border-bottom: 2px solid white;
        color: white;
        font-weight: bold;
        outline: none !important;
    }
    :global(.svelte-tabs li.svelte-tabs__tab) {
        /* background: blue; */
    }

    :global(.fc-h-event) {
        border: 1px solid black;
        background-color: black;
    }

    /* :global(.fc-today-button:after) {
        content: "오늘";
    } */
</style>
