import { GameSettings } from "../../../models";

type APISettings = {
  name: string;
  height: Number;
  width: Number;
  bombs: Number;
};

export const settingsConverter = () => {
  return {
    toAPI: (settings: GameSettings): APISettings => {
      return {
        name: settings.name,
        height: Number(settings.height),
        width: Number(settings.width),
        bombs: Number(settings.bombs),
      };
    },

    fromAPI(settings: APISettings): GameSettings {
      return {
        name: settings.name,
        height: settings.height.valueOf(),
        width: settings.width.valueOf(),
        bombs: settings.bombs.valueOf(),
      };
    },
  };
};
