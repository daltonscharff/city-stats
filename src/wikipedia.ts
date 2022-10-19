import axios from "axios";
import { Cheerio, Element, load } from "cheerio";

const wikipediaService = axios.create({
    baseURL: "https://en.wikipedia.org/w/api.php",
});

export async function getWikipediaData(cityName: string) {
    const cityPage = await getCityPageData(cityName);
    const pageHtml = await scrapeHtml(cityPage.pageid);
    const climateData = parseClimateTable(pageHtml);
    const { population, area, elevation } = parseInfobox(pageHtml);

    return {
        city: cityPage.title,
        climateData,
        population,
        area,
        elevation,
    };
}

async function getCityPageData(cityName: string) {
    type WikipediaQueryResult = {
        batchcomplete: string;
        continue: {
            gsroffset: number;
            continue: string;
        };
        query?: {
            pages: Record<
                string,
                {
                    pageid: string;
                    ns: number;
                    title: string;
                    index: number;
                }
            >;
        };
    };
    const { data: wikipediaQueryResult } =
        await wikipediaService.get<WikipediaQueryResult>("/", {
            params: {
                action: "query",
                gsrlimit: 1,
                gsrsearch: cityName,
                format: "json",
                generator: "search",
            },
        });
    if (!wikipediaQueryResult.query) {
        throw new Error("city not found");
    }

    const pageId = Object.keys(wikipediaQueryResult.query.pages)[0];
    return wikipediaQueryResult.query.pages[pageId];
}

async function scrapeHtml(pageId: string) {
    type WikipediaPageResult = {
        parse: {
            title: string;
            pageid: number;
            text: {
                "*": string;
            };
        };
    };
    const { data: wikipediaPageResult } =
        await wikipediaService.get<WikipediaPageResult>("/", {
            params: {
                action: "parse",
                pageid: pageId,
                format: "json",
            },
        });
    const pageHtml = wikipediaPageResult.parse.text["*"];
    return pageHtml;
}

function parseClimateTable(html: string) {
    const $ = load(html);
    const climateTableElement = $("table.wikitable").filter((_, element) =>
        /climate data/i.test($("tbody tr th", element).text()),
    );
    const climateTableRows = $("tr", climateTableElement).filter((i) => i >= 1);
    const climateTableRowList: string[][] = [];
    climateTableRows.map((_, row) => {
        const rowList = $(row)
            .map((_, element) => {
                return $("th", element)
                    .map((_, e) => {
                        return $(e).text().trim();
                    })
                    .toArray()
                    .concat(
                        $("td", element)
                            .map((_, e) => {
                                return $(e).text().trim();
                            })
                            .toArray(),
                    );
            })
            .toArray();
        climateTableRowList.push(rowList);
    });

    const climateTableObject: Record<string, any> = {};
    climateTableRowList
        .filter((row, _, array) => row.length === array[0].length)
        .forEach((row, i, array) => {
            if (i === 0) return;
            const obj: Record<string, string> = {};
            array[0]
                .filter((_, i) => i > 0)
                .forEach((element, i) => (obj[element] = row[i + 1]));
            climateTableObject[row[0]] = obj;
        });

    function clean(climateTable: typeof climateTableObject) {
        Object.keys(climateTable).forEach((key) => {
            const keyPattern = /\((?:(\w|°))+\)$/;
            if (!key.match(keyPattern)) return;

            const cleanedKey = key.replace(keyPattern, "").trim();
            const valuePattern = /\(.+\)$/;
            climateTable[cleanedKey] = {};
            Object.keys(climateTable[key]).forEach((k) => {
                climateTable[cleanedKey][k] = climateTable[key][k]
                    .replace(valuePattern, "")
                    .trim();
            });
            delete climateTable[key];
        });

        Object.keys(climateTable).forEach((key) => {
            Object.keys(climateTable[key]).forEach((k) => {
                climateTable[key][k] = parseFloat(
                    climateTable[key][k].replace("−", "-"),
                );
            });
        });

        return climateTable;
    }

    const cleanedClimateTable = clean(climateTableObject);
    return cleanedClimateTable;
}

function parseInfobox(html: string) {
    function getPopulation(infobox: Cheerio<Element>) {
        const populationElement = $("tr.mergedtoprow", infobox)
            .filter((_, row) => /population/i.test($(row).text()))
            .first();
        const population = parseInt(
            $(".infobox-data", populationElement.next())
                .text()
                .replace(/,/g, ""),
        );
        return population;
    }

    function getArea(infobox: Cheerio<Element>) {
        const areaElement = $("tr.mergedtoprow", infobox)
            .filter((_, row) => /area/i.test($(row).text()))
            .first();
        const area = $(".infobox-data", areaElement.next()).text();
        return area;
    }

    function getElevation(infobox: Cheerio<Element>) {
        const elevationElement = $("tr.mergedtoprow", infobox)
            .filter((_, row) => /elevation/i.test($(row).text()))
            .first();
        const elevation = $(".infobox-data", elevationElement).text();
        return elevation;
    }

    const $ = load(html);
    const infoboxElement = $(".infobox");
    return {
        population: getPopulation(infoboxElement),
        area: getArea(infoboxElement),
        elevation: getElevation(infoboxElement),
    };
}
