<script>
    export let callback;
    export let name;

</script>

<!-- [{notice.id}]: {notice.title} ({notice.date})<br />{notice.category} -->
<li>
    <a class="category" href="/" on:click|preventDefault={callback}>{name}</a>
</li>

<style>
    a {
        color: inherit;
        transition: 0.2s;
        -webkit-transition: 0.2s;
    }

    a:hover {
        color: hotPink;
    }

    @media (-webkit-min-device-pixel-ratio: 0) {
        a {
            background-color: white;
            background-image: -webkit-linear-gradient(
                left,
                hotPink 0%,
                orange 50%,
                transparent 50%
            );
            background-position: 100% 0;
            background-size: 200% 200%;
            color: transparent;
            transition: 0.1s 0.2s;
            background-clip: text;
            -webkit-transition: 0.1s 0.2s;
            -webkit-background-clip: text;
        }
        a:hover {
            background-position: 0 0;
            color: transparent;
            transition: 0.4s 0;
        }
    }

    .category {
        font-size: 1.25rem;
        font-weight: bold;
        text-decoration: none;
    }

    li {
        display: inline;
    }
</style>
