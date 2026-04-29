import { eq, sql } from "drizzle-orm";
import { db } from "..";
import { feeds, users, Feed } from "../schema";

export async function createFeed(name: string, url: string, user_id: string) {
  const [result] = await db
    .insert(feeds)
    .values({ name: name, url: url, userID: user_id })
    .returning();
  return result;
}

export async function deleteFeeds() {
  await db.delete(feeds);
}

export async function getFeeds() {
  return await db.select().from(feeds);
}

export async function getFeedsWithUser() {
  return await db
    .select()
    .from(feeds)
    .innerJoin(users, eq(feeds.userID, users.id));
}

export async function getFeed(url: string) {
  const [feed] = await db.select().from(feeds).where(eq(feeds.url, url));
  return feed;
}

export async function markFeedFetched(feedID: string) {
  const now = new Date();
  await db
    .update(feeds)
    .set({ updatedAt: now, lastFetchedAt: now })
    .where(eq(feeds.id, feedID));
}

export async function getNextFeedToFetch() {
  const [feed] = await db.execute<Feed>(
    sql`SELECT * FROM feeds ORDER BY last_fetched_at NULLS FIRST LIMIT 1`,
  );
  return feed;
}
