import { and, eq } from "drizzle-orm";
import { db } from "..";
import { feedFollows, feeds, users } from "../schema";

export async function createFeedFollow(userID: string, feedID: string) {
  await db
    .insert(feedFollows)
    .values({ userID: userID, feedID: feedID })
    .returning();

  const [feedFollowWithData] = await db
    .select({
      id: feedFollows.id,
      createdAt: feedFollows.createdAt,
      updatedAt: feedFollows.updatedAt,
      userName: users.name,
      feedName: feeds.name,
    })
    .from(feedFollows)
    .innerJoin(users, eq(users.id, feedFollows.userID))
    .innerJoin(feeds, eq(feeds.id, feedFollows.feedID));

  return feedFollowWithData;
}

export async function deleteFeedFollows() {
  await db.delete(feedFollows);
}

export async function getFeedFollowsForUser(userID: string) {
  return await db
    .select({
      id: feedFollows.id,
      createdAt: feedFollows.createdAt,
      updatedAt: feedFollows.updatedAt,
      feedName: feeds.name,
    })
    .from(feedFollows)
    .where(eq(feedFollows.userID, userID))
    .leftJoin(feeds, eq(feeds.id, feedFollows.feedID));
}

export async function deleteFeedFollow(feedID: string, userID: string) {
  await db
    .delete(feedFollows)
    .where(and(eq(feedFollows.feedID, feedID), eq(feedFollows.userID, userID)));
}
