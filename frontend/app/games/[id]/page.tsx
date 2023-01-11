import MinesweeperCanvas from "../../../components/MinesweeperCanvas/MinesweeperCanvas";

interface CurrentGameProps {
  params: { id: string };
}

export default async function CurrentGame({ params }: CurrentGameProps) {
  const boardExample = [
    ["-", "2", "H", "H", "H"],
    ["-", "2", "H", "H", "H"],
    ["1", "3", "H", "H", "H"],
    ["H", "H", "H", "H", "H"],
    ["H", "H", "H", "H", "H"],
  ];

  return (
    <div>
      <h1>Current Game</h1>
      <p>{params.id}</p>
      <MinesweeperCanvas board={boardExample} />
    </div>
  );
}
