import { exit } from "node:process";
import { CommandsRegistry, registerCommand, runCommand } from "./commands";
import { handlerReset } from "./command_reset";
import { handlerLogin, handlerRegister } from "./commands_users";

async function main() {
  const cmds: CommandsRegistry = {};
  registerCommand(cmds, "login", handlerLogin);
  registerCommand(cmds, "register", handlerRegister);
  registerCommand(cmds, "reset", handlerReset);

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
