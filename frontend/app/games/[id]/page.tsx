interface CurrentGameProps {
  params: { id: string };
}

export default async function CurrentGame({ params }: CurrentGameProps) {
  return (
    <div>
      <h1>Current Game</h1>
      <p>{params.id}</p>
    </div>
  );
}
