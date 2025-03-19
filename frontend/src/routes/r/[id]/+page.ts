import { env } from "$env/dynamic/public";

import type { Redirect } from "$lib/types";
import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch, params }) => {
    var redirect: Redirect | null = null;
    
    const redirectReq = await fetch(`${env.PUBLIC_API_URL}/r/${params.id}`);

    if (redirectReq.ok) {
        redirect = await redirectReq.json();
    }

    return { redirect: redirect };
};