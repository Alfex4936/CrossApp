<script>
    export let notice;

    let placeholder = "할 일을 입력해주세요";

    
</script>

<!-- [{notice.id}]: {notice.title} ({notice.date})<br />{notice.category} -->
<li>
    {notice.title}<br/>({notice.category}) {notice.writer} - {notice.date}
</li>

<style lang="scss">
</style>
