<script lang="ts">
  import NoticeList from "./components/NoticeList.svelte";

  import logo from './assets/images/ajou-logo.png'
  import {Parse} from '../wailsjs/go/main/App'
  import {main} from '../wailsjs/go/models/models'

  // var notices: Array<main.Notice>;
    // let notice = new main.Notice()
    // let notices: Array<main.Notice> = [];

  let notices_promise = Parse("", 15);
</script>

{#await notices_promise then notices}
<main>
  <img alt="ajou logo" id="logo" src="{logo}">

  <NoticeList {notices} />
</main>
{/await}

<style>
  #logo {
    display: block;
    width: 300px;
    height: 282px;
    margin: auto;
    /* padding: 10% 0 0; */
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }
</style>
