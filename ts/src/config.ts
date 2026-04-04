import fs from "fs";

export type Config = {
  dbUrl: string;
  currentUserName: string;
};

const filename = "./.gatorconfig.json";

export function setUser(userName: string) {
  // TODO: use to get ~/.config location os.homedir
  if (!fs.existsSync(filename)) {
    console.log("File doesn't exist");
    return;
  }

  const fileContents = fs.readFileSync(filename, { encoding: "utf-8" });
  const rawConfig = JSON.parse(fileContents);
  rawConfig.current_user_name = userName;

  const configWrite = JSON.stringify(rawConfig);

  fs.writeFileSync(filename, configWrite);
}

export function readConfig(): Config {
  const fileContents = fs.readFileSync(filename, { encoding: "utf-8" });
  const rawConfig = JSON.parse(fileContents);
  const config: Config = {
    dbUrl: rawConfig.db_url,
    currentUserName: rawConfig.current_user_name,
  };

  return config;
}
