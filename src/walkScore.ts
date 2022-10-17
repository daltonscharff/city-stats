import axios from "axios";
import { load } from "cheerio";

const walkScoreService = axios.create({
    baseURL: "https://www.walkscore.com"
});


export async function getScores(cityName: string) {
    const cityPath = await getPath(cityName);
    const pageHtml = await scrapeHtml(cityPath);

    const $ = load(pageHtml);
    const patternList = [/walk\/score\/(\d+).svg$/, /transit\/score\/(\d+).svg$/, /bike\/score\/(\d+).svg$/];
    const imageList = patternList.map((pattern) => $('img').filter((_, element) => pattern.test($(element).attr("src") ?? "")).toArray()[0]);
    const [walkScore, transitScore, bikeScore] = imageList.map((image, i) => image.attribs.src.match(patternList[i])?.[1]);

    return {
        walkScore,
        transitScore,
        bikeScore
    };
}

async function getPath(cityName: string) {
    type CityPathQueryResult = {
        query: string,
        entities: boolean,
        suggestions: { path: string, name: string; }[];
    };
    const { data } = await walkScoreService.get<CityPathQueryResult>("/auth/search_suggest", {
        params: {
            query: cityName,
            skip_entities: 0
        }
    });

    if (data.suggestions.length === 0) throw new Error("city not found");

    return data.suggestions[0].path;
}

async function scrapeHtml(path: string) {
    const { data } = await walkScoreService.get(path);
    return data as string;
}