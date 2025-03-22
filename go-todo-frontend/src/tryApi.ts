import { fetchMe } from "./api";

const tryFetchMe = async () => {
    try {
        const posts = await fetchMe();
        console.log("Fetched posts:", posts);
    } catch (error) {
        console.error(error)
    }
}

tryFetchMe();