"use client";

import { useEffect, useState } from "react";
import getGame from "../../apis/minesweeper-backend/getGame";
import MinesweeperCanvas from "../MinesweeperCanvas/MinesweeperCanvas";

interface MinesweeperGameProps {
  id: string;
}

const MinesweeperGame = ({ id }: MinesweeperGameProps) => {
  // TODO: Add game model
  const [game, setGame] = useState<any>(null);
  const [gameOver, setGameOver] = useState(false);

  useEffect(() => {
    const fetchGame = async () => {
      const game = await getGame(id);
      setGame(game);

      if (game.state !== "in progress") setGameOver(true);
    };

    fetchGame();
  }, [id]);

  return (
    <>
      {gameOver ? (
        <div>{`Game Over: ${game.state}`}</div>
      ) : (
        game && <MinesweeperCanvas game={game} setGame={setGame} />
      )}
    </>
  );
};

export default MinesweeperGame;
