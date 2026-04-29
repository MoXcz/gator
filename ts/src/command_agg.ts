import { fetchFeed, scrapeFeeds } from "./feeds";

export async function handlerAgg(cmdName: string, ...args: string[]) {
  if (args.length < 1) {
    throw Error(`usage: ${cmdName} <time_between_reqs>`);
  }

  const time = args[0];
  const duration = parseDuration(time);

  console.log(`Collecting feeds every ${time}`);

  scrapeFeeds().catch((e) => {
    console.log(`Error: ${e}`);
  });

  const interval = setInterval(() => {
    scrapeFeeds().catch((e) => {
      console.log(`Error: ${e}`);
    });
  }, duration);

  await new Promise<void>((resolve) => {
    process.on("SIGINT", () => {
      console.log("Shutting down feed aggregator...");
      clearInterval(interval);
      resolve();
    });
  });
}

function parseDuration(durationStr: string): number {
  const regex = /^(\d+)(ms|s|m|h)$/;
  const match = durationStr.match(regex);

  if (!match) {
    throw new Error(`Invalid duration ${durationStr}`);
  }

  const value = Number(match[1]);
  const unit = match[2];

  switch (unit) {
    case "ms":
      return value;
    case "s":
      return value * 1000;
    case "m":
      return value * 60 * 1000;
    case "h":
      return value * 60 * 60 * 1000;
    default:
      throw new Error(`Unsupported unit: ${unit}`);
  }
}
