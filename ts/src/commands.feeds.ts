import { createFeed, getFeedsWithUser } from "./db/queries/feeds";
import { createFeedFollow } from "./db/queries/feed_follows";
import { User } from "./db/schema";
import { printFeed } from "./feeds";
import { getPostsForUser } from "./db/queries/posts";

export async function handlerAddFeed(
  cmdName: string,
  user: User,
  ...args: string[]
) {
  if (args.length < 2) {
    throw Error(`usage: ${cmdName} <feed_name> <url>`);
  }

  const feedName = args[0];
  const url = args[1];

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

export async function handlerBrowse(cmdName: string, ...args: string[]) {
  if (args.length > 1 || args.length < 0) {
    throw Error(`usage: ${cmdName} <limit> (optional)`);
  }

  const limit = Number(args[0]);
  const safeLimit = Number.isFinite(limit) ? limit : 2;
  const posts = await getPostsForUser(safeLimit);

  for (const post of posts) {
    console.log(`  - ${post.title}`);
  }
}
