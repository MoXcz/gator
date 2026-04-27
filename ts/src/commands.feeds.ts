import { readConfig } from "./config";
import { createFeed, getFeedsWithUser } from "./db/queries/feed";
import { createFeedFollow } from "./db/queries/feed_follows";
import { getUser } from "./db/queries/users";
import { printFeed } from "./feeds";

export async function handlerAddFeed(cmdName: string, ...args: string[]) {
  if (args.length < 2) {
    throw Error(`usage: ${cmdName} <feed_name> <url>`);
  }

  const feedName = args[0];
  const url = args[1];
  const config = readConfig();
  const user = await getUser(config.currentUserName);

  const feed = await createFeed(feedName, url, user.id);
  await createFeedFollow(user.id, feed.id);

  console.log(`Feed for ${feed.name} created and followed`);
  printFeed(feed, user);
}

export async function handlerFeeds(cmdName: string, ...args: string[]) {
  if (args.length > 0) {
    throw Error(`usage: ${cmdName}`);
  }

  const result = await getFeedsWithUser();

  for (const r of result) {
    console.log(`Name: ${r.feeds.name}`);
    console.log(`URL:  ${r.feeds.url}`);
    console.log(`User: ${r.users.name}\n`);
  }
}
