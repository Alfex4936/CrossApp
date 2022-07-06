<script>
    export let people;
    // console.log(people);

    import { BrowserOpenURL } from "../../wailsjs/runtime";
</script>

<li>
    <p
        style="font-weight: bold"
        on:click={() =>
            BrowserOpenURL("mailto:" + people.email + "?subject=안녕하세요.")}
    >
        {people.korNm} - {people.deptNm}
    </p>
    이메일: {people.email}, 번호: {"031-219-" + people.telNo}
</li>

<style>
    a {
        color: inherit;
        font-weight: bold;
        text-decoration: none;
    }
</style>
