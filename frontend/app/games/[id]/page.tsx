import MinesweeperGame from "../../../components/MinesweeperGame/MinesweeperGame";

interface CurrentGameProps {
  params: { id: string };
}

export default async function CurrentGame({ params }: CurrentGameProps) {
  return (
    <div>
      <MinesweeperGame id={params.id} />
    </div>
  );
}
