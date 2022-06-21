<!-- TodoList.svelte -->
<script>
    import NoticeItem from "./NoticeItem.svelte";

    export let notices;
</script>

<ul>
    {#each notices as notice (notice.id)}
        <NoticeItem {notice} />
        <br/>
    {/each}
</ul>

<style>
    ul {
        display: table;
        margin-left: auto;
        margin-right: auto;
    }
</style>
