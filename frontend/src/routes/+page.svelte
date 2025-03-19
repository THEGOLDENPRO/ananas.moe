<script lang="ts">
    import type { PageData } from "./$types";

    let { data }: { data: PageData }  = $props();
</script>

<svelte:head>
    <title>Home</title>

    <meta content="Ananas • Home" property="og:title">
    <meta content="My website" property="og:description">
    <meta content="ananas.moe" property="og:site_name">
    <meta name="theme-color" content="#ffcb00">
    <meta content="https://ananas.moe/me.webp" property='og:image'>

    <meta name="twitter:title" content="Ananas • Home">
    <meta name="twitter:description" content="My website">
    <meta name="twitter:card" content="summary_large_image">
    <meta name="twitter:image:src" content="https://ananas.moe/me.webp">
</svelte:head>

<div class="flex flex-col">
    <div class="flex justify-center items-center h-screen flex-col md:flex-row gap-10 md:gap-52">
        <div class="block md:hidden text-center">
            <img src="./me.webp" class="rounded-lg" alt="profile" width="300">
        </div>

        <div class="max-w-[50rem] text-center cursor-default select-none">
            <h1 class="text-2xl font-bold">Hi, I'm
                <a class="underline decoration-amber-300 decoration-[3px] underline-offset-4" href="https://en.wiktionary.org/wiki/Ananas#German">Ananas</a>.
            </h1>

            <span class="text-lg p-1">
                <p>I'm a software developer from 🇩🇪 Germany.</p>
            </span>

            <div class="flex justify-center items-center flex-row gap-10">
                <a href="https://github.com/r3tr0ananas" class="font-extrabold hover:scale-125 transition duration-300" aria-label="Link to my GitHub">
                    <img src="/github.svg" alt="Github logo" class="w-[34px] mx-auto">
                    <h1>GitHub</h1>
                </a>

                <a href="https://codeberg.org/bananas" class="font-extrabold hover:scale-125 transition duration-300" aria-label="Link to my Codeberg">
                    <img src="/codeberg.svg" alt="Codeberg logo" class="w-[34px] mx-auto">

                    <h1>Codeberg</h1>
                </a>
            </div>
        </div>

        <div class="hidden md:block text-center">
            <img src="./me.webp" class="rounded-lg" alt="profile" width="300">
        </div>
    </div>

    <div class="divider"></div>

    <div class="flex justify-center items-center">
        {#if data.spotify}
            <div class="flex flex-col w-full">
                <div class="mockup-browser border border-base-300 m-4">
                    <div class="mockup-browser-toolbar">
                      <div class="input">music</div>
                    </div>

                    <ul class="list bg-base-100 rounded-box shadow-md overflow-y-auto h-80">
                        <li class="p-4 pb-2 text-xs opacity-60 tracking-wide">Most listened tracks this month</li>
                        {#each data.spotify as trackItem}
                            <a class="list-row flex justify-between items-center" href="https://open.spotify.com/track/{trackItem.track.externalIds.spotify[0]}">
                                <div class="text-2xl opacity-40">{trackItem.position}</div>
    
                                <img class="skeleton size-10 rounded-box" loading="lazy" alt={trackItem.track.albums[0]?.name} src={trackItem.track.albums[0]?.image}/>
                                <div>
                                    <h1>{trackItem.track.name}</h1>
                                    <h2 class="text-xs uppercase font-semibold opacity-60">{trackItem.track.artists.map(artist => artist.name).join(" & ")}</h2>
                                </div>
    
                                <div class="ml-auto text-center">
                                    <h1>{Math.round(trackItem.playedMs / 1000 / 60)} minutes</h1>
                                    <h2 class="text-xs uppercase font-semibold opacity-60">{trackItem.streams} Streams</h2>
                                </div>
                            </a>
                        {/each}
                    </ul>
                </div>
            </div>
        {:else}
            <div class="flex justify-center items-center bg-base-100 rounded-box shadow-md h-80 w-[500px]">
                <h1>Couldn't pull stats.fm data :sad:</h1>
            </div>
        {/if}
    </div>

    <div class="divider"></div>

    <div class="flex justify-center items-center">
        {#if data.projects}
            <div class="flex flex-col">
                <div class="mockup-browser border border-base-300 m-4">
                    <div class="mockup-browser-toolbar">
                      <div class="input">projects</div>
                    </div>

                    <div class="grid xl:grid-cols-2 shadow-md py-4 px-8 gap-4">
                        {#each data.projects as project}
                            <a class="flex flex-col justify-center items-center" href={project.url}>
                                <img src={project.preview} loading="lazy" alt={project.title} class="rounded-box shadow-md">
                                <h1 class="font-bold text-xl">{project.title}</h1>
                            </a>
                        {/each}
                    </div>
                </div>
            </div>
        {:else}
            <div class="flex justify-center items-center bg-base-100 rounded-box shadow-md h-80 w-[500px]">
                <h1>Couldn't pull projects data</h1>
            </div>
        {/if}
    </div>
</div>