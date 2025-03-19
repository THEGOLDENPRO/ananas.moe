import { env } from "$env/dynamic/public";

import type { File } from "$lib/types";
import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch, params }) => {
    var file: File | null = null;
    
    const fileReq = await fetch(`${env.PUBLIC_API_URL}/file/${params.id}`);

    if (fileReq.ok) {
        file = await fileReq.json();
    }

    return { file: file };
};