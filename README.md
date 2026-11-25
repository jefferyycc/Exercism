# Exercism — Go

My solutions to the [Exercism Go track](https://exercism.org/tracks/go), kept as a public learning log while I build fluency in Go.

## How this repo works

- I solve each exercise locally against the track's test suite, then submit it on exercism.org.
- Submissions are synced here automatically by Exercism's GitHub Syncer, so the commit history doubles as a timeline of my practice.
- Each submission is saved as an *iteration*: `solutions/go/<exercise>/<iteration>/`. When I revisit an exercise to refactor, a new iteration folder appears next to the old one.
- Only my solution files live here; the instructions and tests for each exercise are on exercism.org.

<!-- PROGRESS:START -->
**Progress:** 35 exercises completed (13 concept · 22 practice) — last updated 2026-11-25.

## Concept exercises

Guided exercises that each introduce one Go language feature, in track order.

| Exercise | Concepts |
|---|---|
| [Gopher's Gorgeous Lasagna](solutions/go/lasagna/1/) | `basics` |
| [Annalyn's Infiltration](solutions/go/annalyns-infiltration/1/) | `booleans` |
| [Party Robot](solutions/go/party-robot/1/) | `string-formatting`, `packages` |
| [Weather Forecast](solutions/go/weather-forecast/1/) | `comments` |
| [Blackjack](solutions/go/blackjack/1/) | `conditionals-switch` |
| [Cars Assemble](solutions/go/cars-assemble/1/) | `arithmetic-operators`, `numbers` |
| [Vehicle Purchase](solutions/go/vehicle-purchase/1/) | `comparison`, `conditionals-if` |
| [Lasagna Master](solutions/go/lasagna-master/1/) | `functions` |
| [Need For Speed](solutions/go/need-for-speed/1/) | `structs` |
| [Bird Watcher](solutions/go/bird-watcher/1/) | `for-loops` |
| [Card Tricks](solutions/go/card-tricks/1/) | `slices`, `variadic-functions` |
| [Welcome To Tech Palace!](solutions/go/welcome-to-tech-palace/1/) | `strings`, `strings-package` |
| [Animal Magic](solutions/go/animal-magic/1/) | `randomness` |

## Practice exercises

Open-ended problems that apply those concepts, sorted by Exercism's difficulty rating (1–10).

| Exercise | Difficulty | Practices |
|---|---|---|
| [Hello World](solutions/go/hello-world/1/) | 1 | — |
| [Raindrops](solutions/go/raindrops/1/) | 1 | `basics`, `booleans` |
| [Two Fer](solutions/go/two-fer/1/) | 1 | `basics` |
| [Collatz Conjecture](solutions/go/collatz-conjecture/1/) | 2 | `conditionals-if`, `errors`, `for-loops` |
| [Darts](solutions/go/darts/1/) | 2 | `arithmetic-operators`, `floating-point-numbers` |
| [Difference of Squares](solutions/go/difference-of-squares/1/) | 2 | `arithmetic-operators`, `numbers` |
| [Hamming](solutions/go/hamming/1/) | 2 | `basics`, `booleans`, `comparison` |
| [Leap](solutions/go/leap/1/) | 2 | `basics` |
| [Luhn](solutions/go/luhn/1/) | 2 | `arithmetic-operators`, `for-loops` |
| [Anagram](solutions/go/anagram/1/) | 3 | `slices` |
| [Clock](solutions/go/clock/1/) | 3 | `string-formatting` |
| [ETL](solutions/go/etl/1/) | 3 | `for-loops`, `slices` |
| [Gigasecond](solutions/go/gigasecond/1/) | 3 | `packages` |
| [ISBN Verifier](solutions/go/isbn-verifier/1/) | 3 | `arithmetic-operators` |
| [List Ops](solutions/go/list-ops/1/) | 3 | `first-class-functions`, `slices` |
| [Phone Number](solutions/go/phone-number/1/) | 3 | `string-formatting` |
| [Protein Translation](solutions/go/protein-translation/1/) | 3 | `slices` |
| [Sum of Multiples](solutions/go/sum-of-multiples/1/) | 3 | `variadic-functions` |
| [Baffling Birthdays](solutions/go/baffling-birthdays/1/) | 4 | `randomness` |
| [Ledger](solutions/go/ledger/1/) | 5 | `string-formatting` |
| [Linked List](solutions/go/linked-list/1/) | 5 | `variadic-functions` |
| [POV](solutions/go/pov/1/) | 6 | `variadic-functions` |
<!-- PROGRESS:END -->

## Running a solution

The test files aren't stored here. To run one, download the exercise with the [Exercism CLI](https://exercism.org/docs/using/solving-exercises/working-locally), drop in my solution file, and run:

```bash
exercism download --track=go --exercise=<exercise>
go test ./...
```

## About me

I'm Jeffery Yu, a senior software engineer in San Francisco working on backend, platform, and cloud infrastructure. More at [jefferyycc.github.io](https://jefferyycc.github.io).

> If you're working through the Go track yourself, try each exercise before reading a solution here. That's where the learning is.
