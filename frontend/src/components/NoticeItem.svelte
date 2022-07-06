<script>
    export let notice;

    import { BrowserOpenURL } from "../../wailsjs/runtime";
</script>

<li>
    <!-- <a target="_blank" rel="noopener" href={notice.link}>{notice.title}</a><br
    />({notice.category}) {notice.writer} - {notice.date} -->

    <p style="font-weight:bold" on:click={() => BrowserOpenURL(notice.link)}>
        {notice.title}
    </p>
    ({notice.category}) {notice.writer} - {notice.date}
</li>

<style>
    a {
        color: inherit;
        font-weight: bold;
        text-decoration: none;
    }
</style>
