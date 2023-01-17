import { Game, GameSettings } from "../../models";
import { gameConverter, settingsConverter } from "./converters";

const newGame = async (gameSettings: GameSettings): Promise<Game> => {
  const settings = settingsConverter().toAPI(gameSettings);

  // TODO: think about replacing with config file
  const req = new Request(`${process.env.NEXT_PUBLIC_HOST}/api/games`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(settings),
  });

  try {
    const response = await fetch(req);

    if (!response.ok || response.status !== 201) {
      return Promise.reject("Error creating new game");
    }

    return response.json().then((data) => gameConverter().fromAPI(data));
  } catch (error) {
    throw error;
  }
};

export default newGame;
