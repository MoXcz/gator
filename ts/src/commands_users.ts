import { readConfig, setUser } from "./config";
import { createUser, getUser, getUsers } from "./db/queries/users";

export async function handlerRegister(cmdName: string, ...args: string[]) {
  if (args.length < 1) {
    throw Error(`usage: ${cmdName} <name>`);
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

export async function handlerLogin(cmdName: string, ...args: string[]) {
  if (args.length < 1) {
    throw Error(`usage: ${cmdName} <name>`);
  }
  const username = args[0];
  const user = await getUser(username);
  if (!user) {
    throw Error("user does not exist");
  }

  setUser(username);
  console.log(`User ${username} has been set`);
}
