import { getWalkScoreData } from "./walkScore";
import { getClimateData } from "./wikipedia";

async function main() {
    const city = process.argv[2];
    const climateDataPromise = getClimateData(city).catch(e => {
        console.error(e);
        process.exit(1);
    });
    const walkScoreDataPromise = getWalkScoreData(city).catch(e => {
        console.error(e);
        process.exit(1);
    });

    const [climateData, walkScoreData] = await Promise.all([climateDataPromise, walkScoreDataPromise]);

    console.log(climateData);
    console.log(walkScoreData);
}

main();