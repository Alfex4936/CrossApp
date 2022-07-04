<script>
    export let people;
    // console.log(people);
</script>

<li>
    <a
        target="_blank"
        rel="noopener"
        href={"mailto:" + people.email + "?subject=안녕하세요."}
        >{people.korNm} - {people.deptNm}</a
    ><br />이메일: {people.email}, 번호: {"031-219-" + people.telNo}
</li>

<style>
    a {
        color: inherit;
        font-weight: bold;
        text-decoration: none;
    }
</style>
