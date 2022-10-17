import { getClimateData } from "./wikipedia";

async function main() {
    const city = process.argv[2];
    const climateData = await getClimateData(city).catch(e => {
        console.error(e);
        process.exit(1);
    });

    console.log(climateData);
}

main();