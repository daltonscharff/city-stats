import axios from "axios";
import { load } from "cheerio";

const walkScoreService = axios.create({
    baseURL: "https://www.walkscore.com"
});


export async function getWalkScoreData(cityName: string) {
    const cityPath = await getPath(cityName);
    const pageHtml = await scrapeHtml(cityPath);

    const scores = getScores(pageHtml);
    const neighborhoods = getNeighborhoods(pageHtml);

    return { scores, neighborhoods };
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

function getScores(html: string) {
    const $ = load(html);
    const patternList = [/walk\/score\/(\d+).svg$/, /transit\/score\/(\d+).svg$/, /bike\/score\/(\d+).svg$/];
    const imageList = patternList.map((pattern) => $('img').filter((_, element) => pattern.test($(element).attr("src") ?? "")).toArray()[0]);
    const [walkScore, transitScore, bikeScore] = imageList.map((image, i) => image.attribs.src.match(patternList[i])?.[1]);

    return {
        walk: walkScore,
        transit: transitScore,
        bike: bikeScore
    };
}

function getNeighborhoods(html: string) {
    const $ = load(html);
    const neighborhoodRows = $('#hoods-list-table tbody tr');

    type Neighborhood = { name: string, walkScore: number, transitScore: number, bikeScore: number, population: number; };

    const neighborhoods: Neighborhood[] = [];

    neighborhoodRows.each((_, row) => {
        const neighborhood = {
            name: $(".name", row).text(),
            walkScore: parseInt($(".walkscore", row).text(), 10),
            transitScore: parseInt($(".transitscore", row).text(), 10),
            bikeScore: parseInt($(".bikescore", row).text(), 10),
            population: parseInt($(".population", row).text().replace(/,/g, ''), 10)
        };
        neighborhoods.push(neighborhood);
    });

    return neighborhoods;
}