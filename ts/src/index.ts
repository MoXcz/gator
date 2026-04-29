import { exit } from "node:process";
import {
  CommandHandler,
  CommandsRegistry,
  registerCommand,
  runCommand,
  UserCommandHandler,
} from "./commands";
import { handlerReset } from "./command_reset";
import { handlerLogin, handlerRegister } from "./commands_users";
import { handlerAgg } from "./command_agg";
import { handlerAddFeed, handlerFeeds } from "./commands.feeds";
import {
  handleFollow,
  handleFollowing,
  handleUnfollow,
} from "./command_follow";
import { readConfig } from "./config";
import { getUser } from "./db/queries/users";

async function main() {
  const cmds: CommandsRegistry = {};
  registerCommand(cmds, "login", handlerLogin);
  registerCommand(cmds, "register", handlerRegister);
  registerCommand(cmds, "reset", handlerReset);
  registerCommand(cmds, "agg", handlerAgg);
  registerCommand(cmds, "addfeed", middlewareLoggedIn(handlerAddFeed));
  registerCommand(cmds, "feeds", handlerFeeds);
  registerCommand(cmds, "follow", middlewareLoggedIn(handleFollow));
  registerCommand(cmds, "following", middlewareLoggedIn(handleFollowing));
  registerCommand(cmds, "unfollow", middlewareLoggedIn(handleUnfollow));

  if (process.argv.length < 3) {
    console.log("invalid number of arguments");
    exit(1);
  }

  const cmdName = process.argv[2];
  const args = process.argv.slice(3);
  await runCommand(cmds, cmdName, ...args);
  process.exit(0);
}

main().catch((err) => {
  console.error(err.message);
  exit(1);
});

function middlewareLoggedIn(handler: UserCommandHandler): CommandHandler {
  return async function (cmdName: string, ...args: string[]) {
    const config = readConfig();
    const username = config.currentUserName;
    if (!username) {
      throw new Error(`User not set`);
    }

    const user = await getUser(username);
    if (!user) {
      throw new Error(`User ${username} not found`);
    }
    await handler(cmdName, user, ...args);
  };
}
