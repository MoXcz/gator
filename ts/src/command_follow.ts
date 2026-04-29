import { getFeed } from "./db/queries/feeds";
import {
  createFeedFollow,
  deleteFeedFollow,
  getFeedFollowsForUser,
} from "./db/queries/feed_follows";
import { User } from "./db/schema";

export async function handleFollow(
  cmdName: string,
  user: User,
  ...args: string[]
) {
  if (args.length < 1) {
    throw Error(`usage: ${cmdName} <url>`);
  }

  const url = args[0];
  const feed = await getFeed(url);

  const feedFollow = await createFeedFollow(user.id, feed.id);

  console.log(`Feed ${feedFollow.feedName} followed by ${feedFollow.userName}`);
}

export async function handleFollowing(
  cmdName: string,
  user: User,
  ...args: string[]
) {
  if (args.length > 0) {
    throw Error(`usage: ${cmdName}`);
  }

  const feedFollows = await getFeedFollowsForUser(user.id);
  console.log(`Feeds followed by ${user.name}`);
  for (const fd of feedFollows) {
    console.log(`  - ${fd.feedName}`);
  }
}

export async function handleUnfollow(
  cmdName: string,
  user: User,
  ...args: string[]
) {
  if (args.length < 1) {
    throw Error(`usage: ${cmdName} <url>`);
  }

  const url = args[0];
  const feed = await getFeed(url);
  await deleteFeedFollow(feed.id, user.id);

  console.log(`Feed ${feed.name} removed for ${user.name}`);
}
