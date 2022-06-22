<script lang="ts">
  import NoticeList from "./components/NoticeList.svelte";

  import logo from "./assets/images/ajou-logo.png";
  import { GetWeather, Parse } from "../wailsjs/go/main/App";
  // import {main} from '../wailsjs/go/models/models'

  // var notices: Array<main.Notice>;
  // let notice = new main.Notice()
  // let notices: Array<main.Notice> = [];

  let notices_promise = Parse("", 15);
  let weather_promise = GetWeather();
</script>

<main>
  <img alt="ajou logo" id="logo" src={logo} />
  {#await weather_promise}
    <p>🌞 Loading weather...</p>
  {:then weather}
  <img alt="weather" id="weather" src={weather.icon} />
  <p>날씨: 현재 {weather.current_stat}, 최저 기온은 {weather.min_temp}, 최고 기온은 {weather.max_temp}</p>
  {:catch error}
    <p style="color:red">{error.message}</p>
  {/await}
  {#await notices_promise then notices}
    <NoticeList {notices} />
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
