<script>
    import PeopleItem from "./PeopleItem.svelte";

    export let peoples;
</script>

<ul>
    {#each peoples as people (people.id)}
        <PeopleItem {people} />
        <br />
    {/each}
</ul>

<style>
    ul {
        display: table;
        margin-left: auto;
        margin-right: auto;
        list-style-type: decimal;
    }
</style>
