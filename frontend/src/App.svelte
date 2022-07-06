<script lang="ts">
    import FullCalendar, { type CalendarOptions } from "svelte-fullcalendar";
    import daygridPlugin from "@fullcalendar/daygrid";

    import PeopleList from "./components/PeopleList.svelte";
    import NoticeList from "./components/NoticeList.svelte";
    import CategoryItem from "./components/category/CategoryItem.svelte";
    import CategoryList, {
        categories,
    } from "./components/category/CategoryList.svelte";

    import { Tabs, Tab, TabList, TabPanel } from "svelte-tabs";

    import { ajou_events } from "./components/EventList.svelte";
    import { GetPeople, GetWeather, Parse } from "../wailsjs/go/main/App";
    // TODO: animated tabs

    let options: CalendarOptions = {
        // dateClick: (event) => alert('date click! ' + event.dateStr),
        events: ajou_events,
        buttonText: {
            today: "오늘",
        },
        initialView: "dayGridMonth",
        locale: "ko",
        plugins: [daygridPlugin],
    };

    let weather_promise = GetWeather();
    let people_promise = GetPeople("아주");
    let keyword_of_people = "";

    let notices_promise = Parse("", 7);
    let number_of_notice = 7;
    let last_cate = 0;
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

    function updateNotice() {
        notices_promise = Parse(
            make_link(last_cate, number_of_notice, keyword_of_notice),
            number_of_notice
        );
    }

    function updatePeople() {
        people_promise = GetPeople(keyword_of_people);
    }
</script>

<svelte:head>
    <!-- <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@fortawesome/fontawesome-free@6.1.1/css/fontawesome.min.css" integrity="sha384-zIaWifL2YFF1qaDiAo0JFgsmasocJ/rqu7LKYH8CoBEXqGbb9eO+Xi3s6fQhgFWM" crossorigin="anonymous"> -->
    <meta content="text/html;charset=utf-8" http-equiv="Content-Type" />
    <!-- <link
        rel="stylesheet"
        href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css"
    /> -->
</svelte:head>

<main>
    <Tabs>
        <TabList>
            <Tab>공지 사항</Tab>
            <Tab>학사 일정</Tab>
            <Tab>인물 검색</Tab>
        </TabList>

        <!-- 공지사항 -->
        <TabPanel>
            <CategoryList>
                {#each categories as cate, i}
                    <CategoryItem
                        callback={() => {
                            last_cate = cate.id;
                            updateNotice();
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

        <!-- 인물검색 -->
        <TabPanel>
            <input
                type="text"
                minlength="1"
                maxlength="15"
                style="text-align:center"
                name="keyword"
                bind:value={keyword_of_people}
                on:change={updatePeople}
                placeholder="아주"
            />

            {#await people_promise}
                <p>💌 인물 불러오는 중...</p>
            {:then peoples}
                <PeopleList {peoples} />
            {:catch _}
                <p>인물 불러오는 중 에러 발생!</p>
            {/await}
            <!-- <iframe
                class="media"
                src="https://mportal.ajou.ac.kr/system/phone/phone.do"
                title="아주 BB"
                width="90%"
                height="700px"
                frameborder="0"
            /> -->
            <!-- {@html bb_iframe} -->
        </TabPanel>
        <!-- 인물검색 -->
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
        /* color: rgb(55, 55, 55); */
    }

    :global(.fc-h-event) {
        border: 1px solid black;
        background-color: black;
    }

    /* 일요일 헤더 */
    :global(th.fc-col-header-cell.fc-day.fc-day-sun > div) {
        background-color: red;
    }

    /* 일요일 numbers */
    :global(.fc-day-sun > div > div.fc-daygrid-day-top > a) {
        color: red;
        background-color: white;
    }

    /* 토요일 헤더 */
    :global(th.fc-col-header-cell.fc-day.fc-day-sat > div) {
        background-color: blue;
    }

    /* 일요일 numbers */
    :global(.fc-day-sat > div > div.fc-daygrid-day-top > a) {
        color: blue;
        background-color: white;
    }

    /* :global(.fc-today-button:after) {
        content: "오늘";
    } */
</style>
