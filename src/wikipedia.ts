import axios from "axios";
import { load } from "cheerio";

const wikipediaService = axios.create({
    baseURL: "https://en.wikipedia.org/w/api.php"
});

export async function getClimateData(cityName: string) {
    const pageId = await getCityPageId(cityName);
    const pageHtml = await scrapeHtml(pageId);
    const climateTable = parseClimateTable(pageHtml);

    return climateTable;
}

async function getCityPageId(cityName: string) {
    type WikipediaQueryResult = {
        batchcomplete: string,
        continue: {
            "gsroffset": number;
            "continue": string;
        },
        query?: {
            "pages": Record<string, {
                "pageid": number,
                "ns": number,
                "title": string,
                "index": number;
            }>;
        };
    };
    const { data: wikipediaQueryResult } = await wikipediaService.get<WikipediaQueryResult>("/", {
        params: {
            action: "query",
            gsrlimit: 1,
            gsrsearch: cityName,
            format: "json",
            generator: "search"
        }
    });
    if (!wikipediaQueryResult.query) {
        throw new Error("city not found");
    }

    const pageId = Object.keys(wikipediaQueryResult.query.pages)[0];
    return pageId;
}

async function scrapeHtml(pageId: string) {
    type WikipediaPageResult = {
        parse: {
            title: string,
            pageid: number,
            text: {
                "*": string;
            };
        };
    };
    const { data: wikipediaPageResult } = await wikipediaService.get<WikipediaPageResult>("/", {
        params: {
            action: "parse",
            pageid: pageId,
            format: "json",
        }
    });
    const pageHtml = wikipediaPageResult.parse.text["*"];
    return pageHtml;
}

function parseClimateTable(html: string) {
    const $ = load(html);
    const climateTableElement = $("table.wikitable").filter((_, element) => /climate data/i.test($("tbody tr th", element).text()));
    const climateTableRows = $("tr", climateTableElement).filter((i) => i >= 1);
    const climateTableRowList: string[][] = [];
    climateTableRows.map((_, row) => {
        const rowList = $(row).map((_, element) => {
            return $("th", element).map((_, e) => {
                return $(e).text().trim();
            }).toArray().concat($("td", element).map((_, e) => {
                return $(e).text().trim();
            }).toArray());
        }).toArray();
        climateTableRowList.push(rowList);
    });

    const climateTableObject: Record<string, any> = {};
    climateTableRowList.filter((row, _, array) => row.length === array[0].length).forEach((row, i, array) => {
        if (i === 0) return;
        const obj: Record<string, string> = {};
        array[0].filter((_, i) => i > 0).forEach((element, i) => obj[element] = row[i + 1]);
        climateTableObject[row[0]] = obj;
    });

    return climateTableObject;
}