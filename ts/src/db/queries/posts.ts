import { db } from "..";
import { posts } from "../schema";

export async function createPost(
  title: string,
  url: string,
  publishedAt: Date,
  feedID: string,
) {
  const [result] = await db
    .insert(posts)
    .values({
      title: title,
      url: url,
      publishedAt: publishedAt,
      feedID: feedID,
    })
    .returning();
  return result;
}

export async function getPostsForUser(limit: number) {
  return await db.select().from(posts).orderBy(posts.publishedAt).limit(limit);
}
