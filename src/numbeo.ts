import axios from "axios";
import { load } from "cheerio";
import { camelCase } from "change-case";

const numbeoService = axios.create({
    baseURL: "https://www.numbeo.com/",
});

export async function getNumbeoData(cityName: string) {
    const pageHtml = await scrapeHtml();
    const costOfLivingIndices = getCostOfLivingIndicies(pageHtml);

    return costOfLivingIndices.find((col) =>
        col.city?.toLowerCase().includes(cityName.toLowerCase()),
    );
}

async function scrapeHtml() {
    const { data } = await numbeoService.get(
        "/cost-of-living/rankings_current.jsp",
    );
    return data as string;
}

function getCostOfLivingIndicies(html: string) {
    const $ = load(html);
    const table = $("table#t2");
    const headers = $("th", table)
        .toArray()
        .map((header) => camelCase($(header).text()));
    const rows = $("tbody tr");
    const rowList = $(rows)
        .toArray()
        .map((row) => {
            const r: Record<string, any> = {};
            $("td", row).each((i, item) => {
                const header = headers[i];
                if (header !== "rank") {
                    const value = $(item).text();
                    r[header] = isNaN(+value) ? value : +value;
                }
            });
            return r;
        });
    return rowList;
}
