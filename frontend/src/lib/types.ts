interface Album {
    id: number;
    image: string;
    name: string;
}

interface Artist {
    id: number;
    name: string;
    image: string;
}

interface ExternalIds {
    spotify: string[];
    appleMusic: string[];
}

interface Track {
    albums: Album[];
    artists: Artist[];
    durationMs: number;
    explicit: boolean;
    externalIds: ExternalIds;
    id: number;
    name: string;
}

export interface TrackItem {
    position: number;
    streams: number;
    playedMs: number;
    indicator: any | undefined;
    track: Track;
}

export interface StatsFM {
    items: TrackItem[];
}

export interface Project {
    title: string;
    preview: string;
    url: string;
}

export interface Cloud {
    image: string;
}

export interface Redirect {
    id: string;
    title: string;
    description: string;
    image: string;
    url: string;
}

export interface File {
    id: string;
    file: string;
}