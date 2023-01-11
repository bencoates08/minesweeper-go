"use client";

import MinesweeperCanvas from "../MinesweeperCanvas/MinesweeperCanvas";

interface MinesweeperGameProps {
  // TODO: replace with actual game model
  game: any;
}

const MinesweeperGame = ({ game }: MinesweeperGameProps) => {
  return <MinesweeperCanvas board={game.board} />;
};

export default MinesweeperGame;
