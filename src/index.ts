import { getNumbeoData } from "./numbeo";
import { getWalkScoreData } from "./walkScore";
import { getWikipediaData } from "./wikipedia";
import { program } from "commander";

program
    .option("-w, --walkScore", "show WalkScore table")
    .option("-c, --climate", "show climate data table")
    .argument(
        "<city>",
        'city to find stats for (e.g, "Austin, TX" or "Amsterdam, Netherlands")',
    )
    .action(async (city, options) => {
        const wikipediaDataPromise = getWikipediaData(city).catch((e) => {
            console.error(`Could not get Wikipedia data: ${e}`);
            process.exit(1);
        });
        const walkScoreDataPromise = getWalkScoreData(city).catch((e) => {
            console.error(`Could not get WalkScore data: ${e}`);
        });
        const numbeoDataPromise = getNumbeoData(city).catch((e) => {
            console.error(`Could not get Numbeo data: ${e}`);
        });

        const [wikipediaData, walkScoreData, numbeoData] = await Promise.all([
            wikipediaDataPromise,
            walkScoreDataPromise,
            numbeoDataPromise,
        ]);

        const cityStats = {
            ...wikipediaData,
            walkScore: walkScoreData,
            costOfLiving: numbeoData,
        };

        console.log(JSON.stringify(cityStats, undefined, 2));

        if (options.climate) {
            console.log("Climate Data");
            console.table(cityStats.climateData);
        }
        if (options.walkScore) {
            console.log("WalkScore by Neighborhood");
            console.table(cityStats.walkScore?.byNeighborhood);
        }
    });

program.parse(process.argv);
