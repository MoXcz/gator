import { XMLParser } from "fast-xml-parser";
import { Feed, User } from "./db/schema";

type RSSFeed = {
  channel: {
    title: string;
    link: string;
    description: string;
    item: RSSItem[];
  };
};

type RSSItem = {
  title: string;
  link: string;
  description: string;
  pubDate: string;
};

export async function fetchFeed(feedURL: string) {
  try {
    const res = await fetch(feedURL, {
      headers: {
        "User-Agent": "gator",
      },
    });
    const xmlData = await res.text();

    const parser = new XMLParser();
    const parsedData = parser.parse(xmlData);
    const channel = parsedData.rss.channel || parsedData.channel;
    if (!channel) {
      throw Error(`invalid return value from ${feedURL}`);
    }

    const feed: RSSFeed = {
      channel: {
        title: channel.title,
        link: channel.link,
        description: channel.title,
        item: [],
      },
    };

    if (
      !feed.channel.title ||
      !feed.channel.link ||
      !feed.channel.description
    ) {
      throw Error("missing channel fields");
    }

    if (Array.isArray(channel.item)) {
      for (let i = 0; i < channel.item.length; i++) {
        const item: RSSItem = channel.item[i];
        if (!item.title || !item.description || !item.link || !item.pubDate) {
          continue;
        }

        feed.channel.item.push(item);
      }
    }

    return feed;
  } catch (e) {
    console.log(e);
  }
}

export function printFeed(feed: Feed, user: User) {
  console.log(user);
  console.log(feed);
}
