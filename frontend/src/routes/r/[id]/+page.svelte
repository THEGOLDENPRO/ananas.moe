<script lang="ts">
    import { goto } from "$app/navigation";
    import type { PageData } from "./$types";

    let { data }: { data: PageData }  = $props();

    if (data.redirect) {
        window.location.href = data.redirect.url;
    } else {
        goto("/");
    }
</script>

<svelte:head>
    <title>{data.redirect?.title}</title>

    <meta content={data.redirect?.title} property="og:title">
    <meta content={data.redirect?.description} property="og:description">
    <meta content={data.redirect?.image} property='og:image'>
</svelte:head>
