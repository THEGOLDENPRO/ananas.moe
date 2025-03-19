<script lang="ts">
    import type { PageData } from "./$types";

    import { env } from "$env/dynamic/public";
    import { goto } from "$app/navigation";

    import { Download, File } from "@lucide/svelte";

    let { data }: { data: PageData } = $props();

    if (!data.file) {
        goto("/");
    }
</script>

<div class="flex flex-col justify-center items-center h-screen gap-4">
    <File class="size-24"/>

    <h1 class="text-2xl font-bold">{data.file!.file}</h1>

    <a class="btn text-center" href="{env.PUBLIC_API_URL}/file/{data.file!.id}/download">
        <Download />
        Download
    </a>
</div>