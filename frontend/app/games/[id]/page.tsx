import getGameByID from "../../../apis/minesweeper-backend/getGame";
import MinesweeperGame from "../../../components/MinesweeperGame/MinesweeperGame";

interface CurrentGameProps {
  params: { id: string };
}

export default async function CurrentGame({ params }: CurrentGameProps) {
  const game = await getGameByID(params.id);

  return (
    <div>
      <MinesweeperGame game={game} />
    </div>
  );
}
