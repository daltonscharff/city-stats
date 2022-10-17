import { getWalkScoreData } from "./walkScore";
import { getWikipediaData } from "./wikipedia";

async function main() {
    const city = process.argv[2];
    const wikipediaDataPromise = getWikipediaData(city).catch(e => {
        console.error(`Error getting Wikipedia data: ${e}`);
        process.exit(1);
    });
    const walkScoreDataPromise = getWalkScoreData(city).catch(e => {
        console.error(`Error getting WalkScore data: ${e}`);
        process.exit(1);
    });

    const [wikipediaData,
        walkScoreData
    ] = await Promise.all([wikipediaDataPromise,
        walkScoreDataPromise
    ]);

    console.log(wikipediaData);
    console.log(walkScoreData);
}

main();