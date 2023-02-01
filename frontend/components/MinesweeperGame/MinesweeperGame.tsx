"use client";

import { useEffect, useState } from "react";
import getGame from "../../apis/minesweeper-backend/getGame";
import { Game } from "../../models";
import MinesweeperCanvas from "../MinesweeperCanvas/MinesweeperCanvas";
import styles from "./MinesweeperGame.module.scss";

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
    <div className={styles.gameWrapper}>
      {gameOver && <h1 className={styles.gameOver}>{`you ${game.state}`}</h1>}
      <div className={styles.gameInfo}>
        <p>{`${game.cellsRemaining} TO GO`}</p>
      </div>
      <MinesweeperCanvas game={game} setGame={setGame} />
    </div>
  );
};

export default MinesweeperGame;
