import { desc, eq } from "drizzle-orm";
import { db } from "..";
import { feedFollows, feeds, posts } from "../schema";

export async function createPost(
  title: string,
  url: string,
  publishedAt: Date,
  feedID: string,
  description: string,
) {
  const [result] = await db
    .insert(posts)
    .values({
      title: title,
      url: url,
      publishedAt: publishedAt,
      feedID: feedID,
      description: description,
    })
    .returning();
  return result;
}

export async function getPostsForUser(userID: string, limit: number) {
  return await db
    .select({
      id: posts.id,
      createdAt: posts.createdAt,
      updatedAt: posts.updatedAt,
      title: posts.title,
      url: posts.url,
      description: posts.description,
      publishedAt: posts.publishedAt,
      feedID: posts.feedID,
      feedName: feeds.name,
    })
    .from(posts)
    .innerJoin(feedFollows, eq(posts.feedID, feedFollows.feedID))
    .innerJoin(feeds, eq(posts.feedID, feeds.id))
    .where(eq(feedFollows.userID, userID))
    .orderBy(desc(posts.publishedAt))
    .limit(limit);
}
