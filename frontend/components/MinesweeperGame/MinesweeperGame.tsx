"use client";

import { useEffect, useState } from "react";
import getGame from "../../apis/minesweeper-backend/getGame";
import { Game } from "../../models";
import MinesweeperCanvas from "../MinesweeperCanvas/MinesweeperCanvas";

interface MinesweeperGameProps {
  id: string;
}

const MinesweeperGame = ({ id }: MinesweeperGameProps) => {
  const [game, setGame] = useState<Game | null>(null);
  const [gameOver, setGameOver] = useState(false);

  useEffect(() => {
    const fetchGame = async () => {
      const game = await getGame(id);
      setGame(game);
    };

    fetchGame();
  }, [id]);

  useEffect(() => {
    setGameOver(game?.state !== "in progress");
  }, [game?.state]);

  if (!game) return <div>Loading...</div>;

  return (
    <>
      <h1>{gameOver ? `Game Over: ${game.state}` : "Minesweeper"}</h1>
      <MinesweeperCanvas
        game={game}
        setGame={setGame}
        setGameOver={setGameOver}
      />
      <p>{`Cells Rermaining: ${game.cellsRemaining}`}</p>
    </>
  );
};

export default MinesweeperGame;
