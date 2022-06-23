<script>
    export let notice;

</script>

<!-- [{notice.id}]: {notice.title} ({notice.date})<br />{notice.category} -->
<li>
    <a target="_blank" rel="noopener" href="{notice.link}">{notice.title}</a><br/>({notice.category}) {notice.writer} - {notice.date}
</li>

<style>
    a { color: inherit; } 
</style>
