# City Stats

This CLI tool aggregates population, climate, cost of living, and walk score data from Wikipedia, Walk Score, and Numbeo for a given city and provides it as a JSON object.

## Example

```
> yarn start -c -w "Austin, TX"
{
  "city": "Austin, Texas",
  "climateData": {
    "Average precipitation days (≥ 0.01 in)": {
      "Jan": 7.6,
      "Feb": 7.7,
      "Mar": 8.9,
      "Apr": 7.1,
      "May": 8.9,
      "Jun": 7.4,
      "Jul": 4.9,
      "Aug": 4.8,
      "Sep": 7.1,
      "Oct": 7,
      "Nov": 6.9,
      "Dec": 7.5,
      "Year": 85.8
    },
    "Average snowy days (≥ 0.1 in)": {
      "Jan": 0.2,
      "Feb": 0.3,
      "Mar": 0,
      "Apr": 0,
      "May": 0,
      "Jun": 0,
      "Jul": 0,
      "Aug": 0,
      "Sep": 0,
      "Oct": 0,
      "Nov": 0,
      "Dec": 0.1,
      "Year": 0.6
    },
    "Average relative humidity (%)": {
      "Jan": 67.2,
      "Feb": 66,
      "Mar": 64.2,
      "Apr": 66.4,
      "May": 71.4,
      "Jun": 69.5,
      "Jul": 65.1,
      "Aug": 63.8,
      "Sep": 68.4,
      "Oct": 67.1,
      "Nov": 68.7,
      "Dec": 67.6,
      "Year": 67.1
    },
    "Mean monthly sunshine hours": {
      "Jan": 163.8,
      "Feb": 169.3,
      "Mar": 205.9,
      "Apr": 205.8,
      "May": 227.1,
      "Jun": 285.5,
      "Jul": 317.2,
      "Aug": 297.9,
      "Sep": 233.8,
      "Oct": 215.6,
      "Nov": 168.3,
      "Dec": 153.5,
      "Year": 2643.7
    },
    "Percent possible sunshine": {
      "Jan": 51,
      "Feb": 54,
      "Mar": 55,
      "Apr": 53,
      "May": 54,
      "Jun": 68,
      "Jul": 74,
      "Aug": 73,
      "Sep": 63,
      "Oct": 61,
      "Nov": 53,
      "Dec": 48,
      "Year": 60
    },
    "Average ultraviolet index": {
      "Jan": 4,
      "Feb": 6,
      "Mar": 8,
      "Apr": 9,
      "May": 10,
      "Jun": 11,
      "Jul": 11,
      "Aug": 10,
      "Sep": 9,
      "Oct": 7,
      "Nov": 5,
      "Dec": 4,
      "Year": 8
    },
    "Record high °F": {
      "Jan": 90,
      "Feb": 99,
      "Mar": 98,
      "Apr": 99,
      "May": 104,
      "Jun": 109,
      "Jul": 109,
      "Aug": 112,
      "Sep": 112,
      "Oct": 100,
      "Nov": 91,
      "Dec": 90,
      "Year": 112
    },
    "Mean maximum °F": {
      "Jan": 80.1,
      "Feb": 84.2,
      "Mar": 87.7,
      "Apr": 91.8,
      "May": 95.5,
      "Jun": 99.5,
      "Jul": 102.3,
      "Aug": 103.9,
      "Sep": 99.9,
      "Oct": 93.7,
      "Nov": 85.3,
      "Dec": 80.5,
      "Year": 105.3
    },
    "Average high °F": {
      "Jan": 62.5,
      "Feb": 66.5,
      "Mar": 73.3,
      "Apr": 80.3,
      "May": 86.9,
      "Jun": 93.2,
      "Jul": 96.6,
      "Aug": 97.8,
      "Sep": 91.4,
      "Oct": 82.5,
      "Nov": 71.5,
      "Dec": 63.9,
      "Year": 80.5
    },
    "Average low °F": {
      "Jan": 41.8,
      "Feb": 45.8,
      "Mar": 52.2,
      "Apr": 58.9,
      "May": 66.8,
      "Jun": 72.9,
      "Jul": 75,
      "Aug": 75.1,
      "Sep": 70.1,
      "Oct": 60.8,
      "Nov": 50.5,
      "Dec": 43.4,
      "Year": 59.4
    },
    "Mean minimum °F": {
      "Jan": 27.1,
      "Feb": 30.3,
      "Mar": 34.8,
      "Apr": 42.8,
      "May": 53.4,
      "Jun": 65,
      "Jul": 70.1,
      "Aug": 69.3,
      "Sep": 58.5,
      "Oct": 43.7,
      "Nov": 33.8,
      "Dec": 28.6,
      "Year": 24.2
    },
    "Record low °F": {
      "Jan": -2,
      "Feb": -1,
      "Mar": 18,
      "Apr": 30,
      "May": 40,
      "Jun": 51,
      "Jul": 57,
      "Aug": 58,
      "Sep": 41,
      "Oct": 30,
      "Nov": 20,
      "Dec": 4,
      "Year": -2
    },
    "Average precipitation inches": {
      "Jan": 2.64,
      "Feb": 1.89,
      "Mar": 2.88,
      "Apr": 2.42,
      "May": 5.04,
      "Jun": 3.68,
      "Jul": 1.96,
      "Aug": 2.74,
      "Sep": 3.45,
      "Oct": 3.91,
      "Nov": 2.92,
      "Dec": 2.72,
      "Year": 36.25
    },
    "Average snowfall inches": {
      "Jan": 0,
      "Feb": 0.2,
      "Mar": 0,
      "Apr": 0,
      "May": 0,
      "Jun": 0,
      "Jul": 0,
      "Aug": 0,
      "Sep": 0,
      "Oct": 0,
      "Nov": 0,
      "Dec": 0,
      "Year": 0.2
    },
    "Average dew point °F": {
      "Jan": 36.1,
      "Feb": 39.6,
      "Mar": 46.2,
      "Apr": 55,
      "May": 63.3,
      "Jun": 68.2,
      "Jul": 68.9,
      "Aug": 68.4,
      "Sep": 65.5,
      "Oct": 56.5,
      "Nov": 47.7,
      "Dec": 39.4,
      "Year": 54.6
    }
  },
  "population": 964177,
  "area": "326.51 sq mi",
  "elevation": "289–1,450 ft",
  "walkScore": {
    "average": {
      "walk": 42,
      "transit": 35,
      "bike": 54
    },
    "byNeighborhood": [
      {
        "name": "Downtown",
        "walkScore": 92,
        "transitScore": 68,
        "bikeScore": 90,
        "population": 7412
      },
      {
        "name": "West University",
        "walkScore": 92,
        "transitScore": 65,
        "bikeScore": 94,
        "population": 15358
      },
      {
        "name": "University of Texas-Austin",
        "walkScore": 89,
        "transitScore": 72,
        "bikeScore": 93,
        "population": 8338
      },
      {
        "name": "North University",
        "walkScore": 84,
        "transitScore": 66,
        "bikeScore": 92,
        "population": 4747
      },
      {
        "name": "East Cesar Chavez",
        "walkScore": 84,
        "transitScore": 49,
        "bikeScore": 93,
        "population": 3278
      },
      {
        "name": "Central East Austin",
        "walkScore": 84,
        "transitScore": 53,
        "bikeScore": 85,
        "population": 4372
      },
      {
        "name": "Old West Austin",
        "walkScore": 83,
        "transitScore": 49,
        "bikeScore": 85,
        "population": 4424
      },
      {
        "name": "Bouldin Creek",
        "walkScore": 82,
        "transitScore": 54,
        "bikeScore": 78,
        "population": 5487
      },
      {
        "name": "Holly",
        "walkScore": 82,
        "transitScore": 49,
        "bikeScore": 94,
        "population": 3990
      },
      {
        "name": "Hancock",
        "walkScore": 79,
        "transitScore": 60,
        "bikeScore": 90,
        "population": 5001
      },
      {
        "name": "Hyde Park",
        "walkScore": 78,
        "transitScore": 56,
        "bikeScore": 94,
        "population": 6547
      },
      {
        "name": "Chestnut",
        "walkScore": 78,
        "transitScore": 53,
        "bikeScore": 90,
        "population": 1860
      },
      {
        "name": "Triangle State",
        "walkScore": 76,
        "transitScore": 58,
        "bikeScore": 87,
        "population": 1706
      },
      {
        "name": "Zilker",
        "walkScore": 75,
        "transitScore": 47,
        "bikeScore": 82,
        "population": 5772
      },
      {
        "name": "Upper Boggy Creek",
        "walkScore": 73,
        "transitScore": 54,
        "bikeScore": 85,
        "population": 5193
      },
      {
        "name": "Dawson",
        "walkScore": 71,
        "transitScore": 49,
        "bikeScore": 67,
        "population": 2634
      },
      {
        "name": "Brentwood",
        "walkScore": 70,
        "transitScore": 54,
        "bikeScore": 83,
        "population": 7409
      },
      {
        "name": "Crestview",
        "walkScore": 70,
        "transitScore": 52,
        "bikeScore": 87,
        "population": 3987
      },
      {
        "name": "Rosedale",
        "walkScore": 70,
        "transitScore": 45,
        "bikeScore": 81,
        "population": 3351
      },
      {
        "name": "Rosewood",
        "walkScore": 70,
        "transitScore": 50,
        "bikeScore": 86,
        "population": 4193
      },
      {
        "name": "South River City",
        "walkScore": 69,
        "transitScore": 52,
        "bikeScore": 75,
        "population": 6490
      },
      {
        "name": "Highland",
        "walkScore": 69,
        "transitScore": 56,
        "bikeScore": 86,
        "population": 4057
      },
      {
        "name": "North Loop",
        "walkScore": 67,
        "transitScore": 53,
        "bikeScore": 93,
        "population": 4169
      },
      {
        "name": "Wooten",
        "walkScore": 64,
        "transitScore": 49,
        "bikeScore": 74,
        "population": 5345
      },
      {
        "name": "Govalle",
        "walkScore": 63,
        "transitScore": 48,
        "bikeScore": 89,
        "population": 4102
      },
      {
        "name": "Riverside",
        "walkScore": 63,
        "transitScore": 52,
        "bikeScore": 66,
        "population": 12008
      },
      {
        "name": "North Shoal Creek",
        "walkScore": 59,
        "transitScore": 47,
        "bikeScore": 80,
        "population": 3485
      },
      {
        "name": "St. Johns",
        "walkScore": 58,
        "transitScore": 48,
        "bikeScore": 59,
        "population": 9478
      },
      {
        "name": "Windsor Road",
        "walkScore": 57,
        "transitScore": 43,
        "bikeScore": 79,
        "population": 2978
      },
      {
        "name": "Westgate",
        "walkScore": 57,
        "transitScore": 46,
        "bikeScore": 62,
        "population": 3883
      },
      {
        "name": "Old Enfield",
        "walkScore": 57,
        "transitScore": 39,
        "bikeScore": 79,
        "population": 1180
      },
      {
        "name": "West Congress",
        "walkScore": 56,
        "transitScore": 49,
        "bikeScore": 58,
        "population": 2623
      },
      {
        "name": "North Lamar",
        "walkScore": 56,
        "transitScore": 44,
        "bikeScore": 49,
        "population": 6502
      },
      {
        "name": "Sweetbriar",
        "walkScore": 55,
        "transitScore": 51,
        "bikeScore": 58,
        "population": 6068
      },
      {
        "name": "Windsor Park",
        "walkScore": 55,
        "transitScore": 45,
        "bikeScore": 66,
        "population": 14662
      },
      {
        "name": "St. Edwards",
        "walkScore": 55,
        "transitScore": 47,
        "bikeScore": 60,
        "population": 5290
      },
      {
        "name": "South Lamar",
        "walkScore": 54,
        "transitScore": 48,
        "bikeScore": 66,
        "population": 8136
      },
      {
        "name": "North Austin",
        "walkScore": 54,
        "transitScore": 45,
        "bikeScore": 61,
        "population": 27911
      },
      {
        "name": "Allandale",
        "walkScore": 53,
        "transitScore": 45,
        "bikeScore": 74,
        "population": 9130
      },
      {
        "name": "Garrison Park",
        "walkScore": 53,
        "transitScore": 42,
        "bikeScore": 65,
        "population": 10723
      },
      {
        "name": "RMMA",
        "walkScore": 53,
        "transitScore": 47,
        "bikeScore": 77,
        "population": 2488
      },
      {
        "name": "Gateway",
        "walkScore": 53,
        "transitScore": 37,
        "bikeScore": 57,
        "population": 1121
      },
      {
        "name": "Galindo",
        "walkScore": 51,
        "transitScore": 46,
        "bikeScore": 57,
        "population": 3511
      },
      {
        "name": "South Manchaca",
        "walkScore": 51,
        "transitScore": 45,
        "bikeScore": 60,
        "population": 6465
      },
      {
        "name": "Georgian Acres",
        "walkScore": 51,
        "transitScore": 48,
        "bikeScore": 54,
        "population": 8610
      },
      {
        "name": "Coronado Hills",
        "walkScore": 50,
        "transitScore": 44,
        "bikeScore": 46,
        "population": 3598
      },
      {
        "name": "MLK",
        "walkScore": 49,
        "transitScore": 47,
        "bikeScore": 69,
        "population": 4863
      },
      {
        "name": "West Austin",
        "walkScore": 45,
        "transitScore": 32,
        "bikeScore": 60,
        "population": 10514
      },
      {
        "name": "North Burnet",
        "walkScore": 45,
        "transitScore": 40,
        "bikeScore": 59,
        "population": 4931
      },
      {
        "name": "Montropolis",
        "walkScore": 44,
        "transitScore": 46,
        "bikeScore": 50,
        "population": 10521
      },
      {
        "name": "Heritage Hills",
        "walkScore": 43,
        "transitScore": 39,
        "bikeScore": 55,
        "population": 6089
      },
      {
        "name": "Windsor Hills",
        "walkScore": 42,
        "transitScore": 31,
        "bikeScore": 41,
        "population": 7102
      },
      {
        "name": "Parker Lane",
        "walkScore": 42,
        "transitScore": 42,
        "bikeScore": 57,
        "population": 9599
      },
      {
        "name": "Franklin Park",
        "walkScore": 40,
        "transitScore": 41,
        "bikeScore": 41,
        "population": 16574
      },
      {
        "name": "Barton Hills",
        "walkScore": 36,
        "transitScore": 34,
        "bikeScore": 49,
        "population": 8019
      },
      {
        "name": "East Congress",
        "walkScore": 35,
        "transitScore": 45,
        "bikeScore": 45,
        "population": 3069
      },
      {
        "name": "University Hills",
        "walkScore": 35,
        "transitScore": 41,
        "bikeScore": 34,
        "population": 4662
      },
      {
        "name": "Pleasant Valley",
        "walkScore": 35,
        "transitScore": 53,
        "bikeScore": 57,
        "population": 12676
      },
      {
        "name": "McKinney",
        "walkScore": 32,
        "transitScore": 36,
        "bikeScore": 32,
        "population": 4517
      },
      {
        "name": "Pecan Springs Springdale",
        "walkScore": 30,
        "transitScore": 40,
        "bikeScore": 40,
        "population": 4875
      },
      {
        "name": "MLK-183",
        "walkScore": 29,
        "transitScore": 41,
        "bikeScore": 50,
        "population": 7799
      },
      {
        "name": "Northwest Hills - Far West",
        "walkScore": 25,
        "transitScore": 18,
        "bikeScore": 21,
        "population": 4545
      },
      {
        "name": "Johnston Terrace",
        "walkScore": 24,
        "transitScore": 39,
        "bikeScore": 60,
        "population": 1850
      },
      {
        "name": "East Oak Hill",
        "walkScore": 23,
        "transitScore": 21,
        "bikeScore": 45,
        "population": 13354
      },
      {
        "name": "Village at Western Oaks",
        "walkScore": 22,
        "transitScore": 17,
        "bikeScore": 50,
        "population": 6567
      },
      {
        "name": "Circle C Ranch",
        "walkScore": 17,
        "transitScore": 8,
        "bikeScore": 39,
        "population": 7391
      },
      {
        "name": "West Oak Hill",
        "walkScore": 14,
        "transitScore": 18,
        "bikeScore": 27,
        "population": 16911
      },
      {
        "name": "Southeast Austin",
        "walkScore": 12,
        "transitScore": 29,
        "bikeScore": 30,
        "population": 2322
      }
    ]
  },
  "costOfLiving": {
    "city": "Austin, TX, United States",
    "costOfLivingIndex": 68.13,
    "rentIndex": 55.73,
    "costOfLivingPlusRentIndex": 62.06,
    "groceriesIndex": 68.19,
    "restaurantPriceIndex": 71.88,
    "localPurchasingPowerIndex": 154.74
  }
}
Climate Data
┌────────────────────────────────────────┬───────┬───────┬───────┬───────┬───────┬───────┬───────┬───────┬───────┬───────┬───────┬───────┬────────┐
│                (index)                 │  Jan  │  Feb  │  Mar  │  Apr  │  May  │  Jun  │  Jul  │  Aug  │  Sep  │  Oct  │  Nov  │  Dec  │  Year  │
├────────────────────────────────────────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼────────┤
│ Average precipitation days (≥ 0.01 in) │  7.6  │  7.7  │  8.9  │  7.1  │  8.9  │  7.4  │  4.9  │  4.8  │  7.1  │   7   │  6.9  │  7.5  │  85.8  │
│     Average snowy days (≥ 0.1 in)      │  0.2  │  0.3  │   0   │   0   │   0   │   0   │   0   │   0   │   0   │   0   │   0   │  0.1  │  0.6   │
│     Average relative humidity (%)      │ 67.2  │  66   │ 64.2  │ 66.4  │ 71.4  │ 69.5  │ 65.1  │ 63.8  │ 68.4  │ 67.1  │ 68.7  │ 67.6  │  67.1  │
│      Mean monthly sunshine hours       │ 163.8 │ 169.3 │ 205.9 │ 205.8 │ 227.1 │ 285.5 │ 317.2 │ 297.9 │ 233.8 │ 215.6 │ 168.3 │ 153.5 │ 2643.7 │
│       Percent possible sunshine        │  51   │  54   │  55   │  53   │  54   │  68   │  74   │  73   │  63   │  61   │  53   │  48   │   60   │
│       Average ultraviolet index        │   4   │   6   │   8   │   9   │  10   │  11   │  11   │  10   │   9   │   7   │   5   │   4   │   8    │
│             Record high °F             │  90   │  99   │  98   │  99   │  104  │  109  │  109  │  112  │  112  │  100  │  91   │  90   │  112   │
│            Mean maximum °F             │ 80.1  │ 84.2  │ 87.7  │ 91.8  │ 95.5  │ 99.5  │ 102.3 │ 103.9 │ 99.9  │ 93.7  │ 85.3  │ 80.5  │ 105.3  │
│            Average high °F             │ 62.5  │ 66.5  │ 73.3  │ 80.3  │ 86.9  │ 93.2  │ 96.6  │ 97.8  │ 91.4  │ 82.5  │ 71.5  │ 63.9  │  80.5  │
│             Average low °F             │ 41.8  │ 45.8  │ 52.2  │ 58.9  │ 66.8  │ 72.9  │  75   │ 75.1  │ 70.1  │ 60.8  │ 50.5  │ 43.4  │  59.4  │
│            Mean minimum °F             │ 27.1  │ 30.3  │ 34.8  │ 42.8  │ 53.4  │  65   │ 70.1  │ 69.3  │ 58.5  │ 43.7  │ 33.8  │ 28.6  │  24.2  │
│             Record low °F              │  -2   │  -1   │  18   │  30   │  40   │  51   │  57   │  58   │  41   │  30   │  20   │   4   │   -2   │
│      Average precipitation inches      │ 2.64  │ 1.89  │ 2.88  │ 2.42  │ 5.04  │ 3.68  │ 1.96  │ 2.74  │ 3.45  │ 3.91  │ 2.92  │ 2.72  │ 36.25  │
│        Average snowfall inches         │   0   │  0.2  │   0   │   0   │   0   │   0   │   0   │   0   │   0   │   0   │   0   │   0   │  0.2   │
│          Average dew point °F          │ 36.1  │ 39.6  │ 46.2  │  55   │ 63.3  │ 68.2  │ 68.9  │ 68.4  │ 65.5  │ 56.5  │ 47.7  │ 39.4  │  54.6  │
└────────────────────────────────────────┴───────┴───────┴───────┴───────┴───────┴───────┴───────┴───────┴───────┴───────┴───────┴───────┴────────┘
WalkScore by Neighborhood
┌─────────┬──────────────────────────────┬───────────┬──────────────┬───────────┬────────────┐
│ (index) │             name             │ walkScore │ transitScore │ bikeScore │ population │
├─────────┼──────────────────────────────┼───────────┼──────────────┼───────────┼────────────┤
│    0    │          'Downtown'          │    92     │      68      │    90     │    7412    │
│    1    │      'West University'       │    92     │      65      │    94     │   15358    │
│    2    │ 'University of Texas-Austin' │    89     │      72      │    93     │    8338    │
│    3    │      'North University'      │    84     │      66      │    92     │    4747    │
│    4    │     'East Cesar Chavez'      │    84     │      49      │    93     │    3278    │
│    5    │    'Central East Austin'     │    84     │      53      │    85     │    4372    │
│    6    │      'Old West Austin'       │    83     │      49      │    85     │    4424    │
│    7    │       'Bouldin Creek'        │    82     │      54      │    78     │    5487    │
│    8    │           'Holly'            │    82     │      49      │    94     │    3990    │
│    9    │          'Hancock'           │    79     │      60      │    90     │    5001    │
│   10    │         'Hyde Park'          │    78     │      56      │    94     │    6547    │
│   11    │          'Chestnut'          │    78     │      53      │    90     │    1860    │
│   12    │       'Triangle State'       │    76     │      58      │    87     │    1706    │
│   13    │           'Zilker'           │    75     │      47      │    82     │    5772    │
│   14    │     'Upper Boggy Creek'      │    73     │      54      │    85     │    5193    │
│   15    │           'Dawson'           │    71     │      49      │    67     │    2634    │
│   16    │         'Brentwood'          │    70     │      54      │    83     │    7409    │
│   17    │         'Crestview'          │    70     │      52      │    87     │    3987    │
│   18    │          'Rosedale'          │    70     │      45      │    81     │    3351    │
│   19    │          'Rosewood'          │    70     │      50      │    86     │    4193    │
│   20    │      'South River City'      │    69     │      52      │    75     │    6490    │
│   21    │          'Highland'          │    69     │      56      │    86     │    4057    │
│   22    │         'North Loop'         │    67     │      53      │    93     │    4169    │
│   23    │           'Wooten'           │    64     │      49      │    74     │    5345    │
│   24    │          'Govalle'           │    63     │      48      │    89     │    4102    │
│   25    │         'Riverside'          │    63     │      52      │    66     │   12008    │
│   26    │     'North Shoal Creek'      │    59     │      47      │    80     │    3485    │
│   27    │         'St. Johns'          │    58     │      48      │    59     │    9478    │
│   28    │        'Windsor Road'        │    57     │      43      │    79     │    2978    │
│   29    │          'Westgate'          │    57     │      46      │    62     │    3883    │
│   30    │        'Old Enfield'         │    57     │      39      │    79     │    1180    │
│   31    │       'West Congress'        │    56     │      49      │    58     │    2623    │
│   32    │        'North Lamar'         │    56     │      44      │    49     │    6502    │
│   33    │         'Sweetbriar'         │    55     │      51      │    58     │    6068    │
│   34    │        'Windsor Park'        │    55     │      45      │    66     │   14662    │
│   35    │        'St. Edwards'         │    55     │      47      │    60     │    5290    │
│   36    │        'South Lamar'         │    54     │      48      │    66     │    8136    │
│   37    │        'North Austin'        │    54     │      45      │    61     │   27911    │
│   38    │         'Allandale'          │    53     │      45      │    74     │    9130    │
│   39    │       'Garrison Park'        │    53     │      42      │    65     │   10723    │
│   40    │            'RMMA'            │    53     │      47      │    77     │    2488    │
│   41    │          'Gateway'           │    53     │      37      │    57     │    1121    │
│   42    │          'Galindo'           │    51     │      46      │    57     │    3511    │
│   43    │       'South Manchaca'       │    51     │      45      │    60     │    6465    │
│   44    │       'Georgian Acres'       │    51     │      48      │    54     │    8610    │
│   45    │       'Coronado Hills'       │    50     │      44      │    46     │    3598    │
│   46    │            'MLK'             │    49     │      47      │    69     │    4863    │
│   47    │        'West Austin'         │    45     │      32      │    60     │   10514    │
│   48    │        'North Burnet'        │    45     │      40      │    59     │    4931    │
│   49    │        'Montropolis'         │    44     │      46      │    50     │   10521    │
│   50    │       'Heritage Hills'       │    43     │      39      │    55     │    6089    │
│   51    │       'Windsor Hills'        │    42     │      31      │    41     │    7102    │
│   52    │        'Parker Lane'         │    42     │      42      │    57     │    9599    │
│   53    │       'Franklin Park'        │    40     │      41      │    41     │   16574    │
│   54    │        'Barton Hills'        │    36     │      34      │    49     │    8019    │
│   55    │       'East Congress'        │    35     │      45      │    45     │    3069    │
│   56    │      'University Hills'      │    35     │      41      │    34     │    4662    │
│   57    │      'Pleasant Valley'       │    35     │      53      │    57     │   12676    │
│   58    │          'McKinney'          │    32     │      36      │    32     │    4517    │
│   59    │  'Pecan Springs Springdale'  │    30     │      40      │    40     │    4875    │
│   60    │          'MLK-183'           │    29     │      41      │    50     │    7799    │
│   61    │ 'Northwest Hills - Far West' │    25     │      18      │    21     │    4545    │
│   62    │      'Johnston Terrace'      │    24     │      39      │    60     │    1850    │
│   63    │       'East Oak Hill'        │    23     │      21      │    45     │   13354    │
│   64    │  'Village at Western Oaks'   │    22     │      17      │    50     │    6567    │
│   65    │       'Circle C Ranch'       │    17     │      8       │    39     │    7391    │
│   66    │       'West Oak Hill'        │    14     │      18      │    27     │   16911    │
│   67    │      'Southeast Austin'      │    12     │      29      │    30     │    2322    │
└─────────┴──────────────────────────────┴───────────┴──────────────┴───────────┴────────────┘
```
