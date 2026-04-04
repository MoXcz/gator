import { deleteUsers } from "./db/queries/users";

export async function handlerReset(_: string, ...args: string[]) {
  if (args.length > 0) {
    throw Error("expected no arguments");
  }

  try {
    await deleteUsers();
  } catch {
    throw Error("could not delete users from table");
  }
  console.log("users table deleted");
}
