import { Game } from "../../models";
import { gameConverter } from "./converters";

const getGame = async (gameID: string): Promise<Game> => {
  // TODO: think about replacing with config file
  try {
    const response = await fetch(
      `${process.env.NEXT_PUBLIC_HOST}/games/${gameID}`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
        cache: "no-store",
      }
    );

    if (!response.ok || response.status !== 200) {
      return Promise.reject("Error getting game");
    }

    return response.json().then((data) => gameConverter().fromAPI(data));
  } catch (error) {
    throw error;
  }
};

export default getGame;
