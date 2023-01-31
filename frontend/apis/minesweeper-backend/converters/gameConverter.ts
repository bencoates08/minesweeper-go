import { Game } from "../../../models";

type APIGame = {
  board: string[][];
  cells_remaining: number;
  id: string;
  name: string;
  state: string;
};

export const gameConverter = () => {
  return {
    toAPI: (game: Game): APIGame => {
      return {
        board: game.board,
        cells_remaining: game.cellsRemaining,
        id: game.id,
        name: game.name,
        state: game.state,
      };
    },

    fromAPI(game: APIGame): Game {
      return {
        board: game.board,
        cellsRemaining: game.cells_remaining,
        id: game.id,
        name: game.name,
        state: game.state,
      };
    },
  };
};
