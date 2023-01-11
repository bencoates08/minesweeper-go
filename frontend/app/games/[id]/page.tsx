import getGameByID from "../../../apis/minesweeper-backend/getGame";
import MinesweeperGame from "../../../components/MinesweeperGame/MinesweeperGame";

interface CurrentGameProps {
  params: { id: string };
}

const getGame = async (id: string) => {
  return getGameByID(id);
};

export default async function CurrentGame({ params }: CurrentGameProps) {
  const game = await getGame(params.id);

  return (
    <div>
      <MinesweeperGame game={game} />
    </div>
  );
}
