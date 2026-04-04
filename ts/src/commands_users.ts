import { setUser } from "./config";
import { createUser, getUser } from "./db/queries/users";

export async function handlerRegister(_: string, ...args: string[]) {
  if (args.length < 1) {
    throw Error("expected username argument");
  }

  const username = args[0];
  const user = await getUser(username);
  if (user) {
    throw Error(`User ${user.name} already exists`);
  }

  const newUser = await createUser(username);
  setUser(newUser.name);
  console.log(`User ${newUser.name} created`);
}

export async function handlerLogin(_: string, ...args: string[]) {
  if (args.length < 1) {
    throw Error("expected username argument");
  }
  const username = args[0];
  const user = await getUser(username);
  if (!user) {
    throw Error("user does not exist");
  }

  setUser(username);
  console.log(`User ${username} has been set`);
}
