import { getWalkScoreData } from "./walkScore";
import { getWikipediaData } from "./wikipedia";

async function main() {
    const city = process.argv[2];
    const wikipediaDataPromise = getWikipediaData(city).catch((e) => {
        console.error(`Could not get Wikipedia data: ${e}`);
        process.exit(1);
    });
    const walkScoreDataPromise = getWalkScoreData(city).catch((e) => {
        console.error(`Could not get WalkScore data: ${e}`);
    });

    const [wikipediaData, walkScoreData] = await Promise.all([
        wikipediaDataPromise,
        walkScoreDataPromise,
    ]);

    const cityStats = {
        ...wikipediaData,
        walkScore: walkScoreData,
    };
    console.log(JSON.stringify(cityStats, undefined, 2));
}

main();
