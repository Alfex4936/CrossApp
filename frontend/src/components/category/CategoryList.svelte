<script context="module">
    export let categories = [
        {
            id: 0,
            name: "전체 공지",
        },
        {
            id: 1,
            name: "학사",
        },
        {
            id: 2,
            name: "비교과",
        },
        {
            id: 3,
            name: "장학",
        },
        {
            id: 4,
            name: "학술",
        },
        {
            id: 5,
            name: "입학",
        },
        {
            id: 6,
            name: "취업",
        },
        {
            id: 7,
            name: "사무",
        },
        {
            id: 8,
            name: "기타",
        },
        {
            id: 166,
            name: "행사",
        },
        {
            id: 167,
            name: "파란학기제",
        },
        {
            id: 168,
            name: "학사일정",
        },
    ]
</script>

<ul id="menu">
    <slot />
</ul>

<style>
    ul {
        display: table;
        margin-left: auto;
        margin-right: auto;
    }

    #menu {
        font-size: 1.25rem;
        font-weight: bold;
        text-decoration: none;
    }
</style>
