const reveal = async (
  gameID: string,
  row: number,
  col: number
): Promise<any> => {
  const position = { row, col };

  // TODO: think about replacing with config file
  const req = new Request(
    `${process.env.NEXT_PUBLIC_HOST}/api/games/${gameID}/reveal`,
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(position),
    }
  );

  try {
    const response = await fetch(req);

    if (!response.ok || response.status !== 200) {
      return Promise.reject("Error requesting reveal");
    }

    // ! TODO: add game model and converter
    return response.json();
  } catch (error) {
    throw error;
  }
};

export default reveal;
