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

    printCityStats(cityStats);

    function printCityStats(stats: typeof cityStats) {
        console.log(
            JSON.stringify(
                {
                    city: stats.city,
                    population: stats.population,
                    elevation: stats.elevation,
                    area: stats.area,
                    averageWalkScore: stats.walkScore?.average,
                },
                undefined,
                2,
            ),
        );

        console.log("Climate Data");
        console.table(stats.climateData);

        if (stats.walkScore) {
            console.log("WalkScore by Neighborhood");
            console.table(stats.walkScore.byNeighborhood);
        }
    }
}

main();
