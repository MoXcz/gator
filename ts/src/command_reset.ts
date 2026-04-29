import { deleteFeeds } from "./db/queries/feeds";
import { deleteFeedFollows } from "./db/queries/feed_follows";
import { deleteUsers } from "./db/queries/users";

export async function handlerReset(_: string, ...args: string[]) {
  if (args.length > 0) {
    throw Error("expected no arguments");
  }

  try {
    await deleteUsers();
    await deleteFeeds();
    await deleteFeedFollows();
  } catch {
    throw Error("could not delete all rows from tables");
  }
  console.log("all tables reset!");
}
