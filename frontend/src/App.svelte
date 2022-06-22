<script lang="ts">
  import NoticeList from "./components/NoticeList.svelte";

  import logo from "./assets/images/ajou-logo.png";
  import { GetWeather, Parse } from "../wailsjs/go/main/App";
  // import {main} from '../wailsjs/go/models/models'

  // var notices: Array<main.Notice>;
  // let notice = new main.Notice()
  // let notices: Array<main.Notice> = [];

  let notices_promise = Parse("", 7);
  let weather_promise = GetWeather();

  let number_of_notice = 7;
  function updateNotice() {
    notices_promise = Parse("", number_of_notice);
  }
</script>

<main>
  <!-- <img alt="ajou logo" id="logo" src={logo} /> -->
  {#await weather_promise}
    <p>🌞 날씨 불러오는 중...</p>
  {:then weather}
  <img alt="weather" id="weather" src={weather.icon} />
  <p>현재 날씨 {weather.current_stat}, 최저 기온은 {weather.min_temp}, 최고 기온은 {weather.max_temp}<br/>
    낮 강수량은 {weather.rain_day}, 밤 강수량은 {weather.rain_night}</p>
  {:catch error}
    <p style="color:red">{error.message}</p>
  {/await}
  <p style="display:inline">공지 갯수:</p><input type=number min=1 max=50 style="text-align:center" name="notice" bind:value={number_of_notice} placeholder="공지 갯수" /><button on:click={ updateNotice }>공지 불러오기</button>
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
</style>
