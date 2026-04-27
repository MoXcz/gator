import { fetchFeed } from "./feeds";

export async function handlerAgg(cmdName: string, ...args: string[]) {
  if (args.length < 1) {
    throw Error(`usage: ${cmdName} <url>`);
  }

  const url = args[0];
  const feed = await fetchFeed(url);
  console.log(feed);
}
