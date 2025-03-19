import { env } from "$env/dynamic/public";

import type { Project, StatsFM, TrackItem } from "$lib/types";
import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch }) => {
    var spotifyStats: TrackItem[] | null = null;
    var projects: Project[] | null = null;

    const topTracks = await fetch("https://api.stats.fm/api/v1/users/r3tr0ananas/top/tracks?range=weeks");
    const projectsReq = await fetch(`${env.PUBLIC_API_URL}/projects`);
    // const cloudsReq = await fetch(`${env.PUBLIC_API_URL}/clouds`);

    if (topTracks.ok) {
        const tracks: StatsFM = await topTracks.json();
        spotifyStats = tracks.items.slice(0, 40);
    }

    if (projectsReq.ok) {
        projects = await projectsReq.json();
    }

    return { spotify: spotifyStats, projects }
};